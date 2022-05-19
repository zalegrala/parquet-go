package parquet_test

import (
	"bytes"
	"errors"
	"io"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/segmentio/parquet-go"
)

func TestGenericReader(t *testing.T) {
	testGenericReader[booleanColumn](t)
	testGenericReader[int32Column](t)
	testGenericReader[int64Column](t)
	testGenericReader[int96Column](t)
	testGenericReader[floatColumn](t)
	testGenericReader[doubleColumn](t)
	testGenericReader[byteArrayColumn](t)
	testGenericReader[fixedLenByteArrayColumn](t)
	testGenericReader[stringColumn](t)
	testGenericReader[indexedStringColumn](t)
	testGenericReader[uuidColumn](t)
	testGenericReader[decimalColumn](t)
	testGenericReader[addressBook](t)
	testGenericReader[contact](t)
	testGenericReader[listColumn2](t)
	testGenericReader[listColumn1](t)
	testGenericReader[listColumn0](t)
	testGenericReader[nestedListColumn1](t)
	testGenericReader[nestedListColumn](t)
}

func testGenericReader[Row any](t *testing.T) {
	var prng = rand.New(rand.NewSource(2))
	var model Row
	t.Run(reflect.TypeOf(model).Name(), func(t *testing.T) {
		f := func(rows []Row) bool {
			if len(rows) == 0 {
				return true // TODO: fix support for parquet files with zero rows
			}
			buffer := new(bytes.Buffer)
			writer := parquet.NewGenericWriter[Row](buffer)
			_, err := writer.Write(rows)
			if err != nil {
				t.Error(err)
				return false
			}
			if err := writer.Close(); err != nil {
				t.Error(err)
				return false
			}
			reader := parquet.NewGenericReader[Row](bytes.NewReader(buffer.Bytes()))
			result := make([]Row, len(rows))
			n, err := reader.Read(result)
			if err != nil && !errors.Is(err, io.EOF) {
				t.Error(err)
				return false
			}
			if n < len(rows) {
				t.Errorf("not enough values were read: want=%d got=%d", len(rows), n)
				return false
			}
			if !reflect.DeepEqual(rows, result) {
				t.Errorf("rows mismatch:\nwant: %+v\ngot:  %+v", rows, result)
				return false
			}
			return true
		}
		if err := quick.Check(f, &quick.Config{Rand: prng}); err != nil {
			t.Error(err)
		}
	})
}

func BenchmarkGenericReader(b *testing.B) {
	data := make([]byte, 16*benchmarkReaderNumRows)
	prng := rand.New(rand.NewSource(0))
	prng.Read(data)

	values := make([]uuidColumn, benchmarkReaderNumRows)
	for i := range values {
		j := (i + 0) * 16
		k := (i + 1) * 16
		copy(values[i].Value[:], data[j:k])
	}

	rowbuf := make([]uuidColumn, benchmarkReaderRowsPerStep)
	buffer := parquet.NewGenericBuffer[uuidColumn]()
	buffer.Write(values)

	b.Run("go1.17", func(b *testing.B) {
		reader := parquet.NewRowGroupReader(buffer)
		benchmarkRowsPerSecond(b, func() int {
			for i := range rowbuf {
				if err := reader.Read(&rowbuf[i]); err != nil {
					if err != io.EOF {
						b.Fatal(err)
					} else {
						reader.Reset()
					}
				}
			}
			return len(rowbuf)
		})
	})

	b.Run("go1.18", func(b *testing.B) {
		reader := parquet.NewGenericRowGroupReader[uuidColumn](buffer)
		benchmarkRowsPerSecond(b, func() int {
			n, err := reader.Read(rowbuf)
			if err != nil {
				if err != io.EOF {
					b.Fatal(err)
				} else {
					reader.Reset()
				}
			}
			return n
		})
	})
}