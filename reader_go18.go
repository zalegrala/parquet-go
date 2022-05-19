package parquet

import (
	"io"
	"reflect"
)

type GenericReader[T any] struct {
	base Reader
	read readFunc[T]
}

func NewGenericReader[T any](input io.ReaderAt, options ...ReaderOption) *GenericReader[T] {
	c, err := NewReaderConfig(options...)
	if err != nil {
		panic(err)
	}

	if c.Schema == nil {
		var model T
		c.Schema = schemaOf(reflect.TypeOf(model))
	}

	f, err := openFile(input)
	if err != nil {
		panic(err)
	}

	r := &GenericReader[T]{
		base: Reader{
			file: reader{
				schema:   c.Schema,
				rowGroup: fileRowGroupOf(f),
			},
		},
	}

	if !nodesAreEqual(c.Schema, f.schema) {
		r.base.file.rowGroup = convertRowGroupTo(r.base.file.rowGroup, c.Schema)
	}

	r.base.read.init(r.base.file.schema, r.base.file.rowGroup)
	r.read = readFuncOf[T](r.base.file.schema)
	return r
}

func NewGenericRowGroupReader[T any](rowGroup RowGroup, options ...ReaderOption) *GenericReader[T] {
	c, err := NewReaderConfig(options...)
	if err != nil {
		panic(err)
	}

	if c.Schema == nil {
		var model T
		c.Schema = schemaOf(reflect.TypeOf(model))
	}

	r := &GenericReader[T]{
		base: Reader{
			file: reader{
				schema:   c.Schema,
				rowGroup: rowGroup,
			},
		},
	}

	if !nodesAreEqual(c.Schema, rowGroup.Schema()) {
		r.base.file.rowGroup = convertRowGroupTo(r.base.file.rowGroup, c.Schema)
	}

	r.base.read.init(r.base.file.schema, r.base.file.rowGroup)
	r.read = readFuncOf[T](r.base.file.schema)
	return r
}

func (r *GenericReader[T]) Reset() {
	r.base.Reset()
}

func (r *GenericReader[T]) Read(rows []T) (int, error) {
	return r.read(r, rows)
}

func (r *GenericReader[T]) ReadRows(rows []Row) (int, error) {
	return r.base.ReadRows(rows)
}

func (r *GenericReader[T]) Schema() *Schema {
	return r.base.Schema()
}

func (r *GenericReader[T]) NumRows() int64 {
	return r.base.NumRows()
}

func (r *GenericReader[T]) SeekToRow(rowIndex int64) error {
	return r.base.SeekToRow(rowIndex)
}

func (r *GenericReader[T]) Close() error {
	return r.base.Close()
}

func (r *GenericReader[T]) readRows(rows []T) (int, error) {
	if cap(r.base.rowbuf) < len(rows) {
		r.base.rowbuf = make([]Row, len(rows))
	} else {
		r.base.rowbuf = r.base.rowbuf[:len(rows)]
	}

	n, err := r.base.ReadRows(r.base.rowbuf)
	if n > 0 {
		schema := r.base.Schema()

		for i, row := range r.base.rowbuf[:n] {
			if err := schema.Reconstruct(&rows[i], row); err != nil {
				return i, err
			}
		}
	}
	return n, err
}

var (
	_ Rows                = (*GenericReader[any])(nil)
	_ RowReaderWithSchema = (*Reader)(nil)

	_ Rows                = (*GenericReader[struct{}])(nil)
	_ RowReaderWithSchema = (*GenericReader[struct{}])(nil)

	_ Rows                = (*GenericReader[map[struct{}]struct{}])(nil)
	_ RowReaderWithSchema = (*GenericReader[map[struct{}]struct{}])(nil)
)

type readFunc[T any] func(*GenericReader[T], []T) (int, error)

func readFuncOf[T any](schema *Schema) readFunc[T] {
	var model T
	switch t := reflect.TypeOf(model); t.Kind() {
	case reflect.Interface, reflect.Map:
		return (*GenericReader[T]).readRows
	case reflect.Struct:
		return readFuncOfStruct[T](t, schema)
	default:
		panic("cannot create reader for values of type " + t.String())
	}
}

func readFuncOfStruct[T any](t reflect.Type, schema *Schema) readFunc[T] {
	return func(w *GenericReader[T], rows []T) (int, error) {
		return w.readRows(rows)
	}
}