// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import "github.com/go-faster/errors"

// ColumnInt64 represents Int64 column.
type ColumnInt64 []int64

// Compile-time assertions for ColumnInt64.
var (
	_ Input  = ColumnInt64{}
	_ Result = (*ColumnInt64)(nil)
)

// Type returns ColumnType of Int64.
func (ColumnInt64) Type() ColumnType {
	return ColumnTypeInt64
}

// Rows returns count of rows in column.
func (c ColumnInt64) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColumnInt64) Reset() {
	*c = (*c)[:0]
}

// EncodeColumn encodes Int64 rows to *Buffer.
func (c ColumnInt64) EncodeColumn(b *Buffer) {
	for _, v := range c {
		b.PutInt64(v)
	}
}

// DecodeColumn decodes Int64 rows from *Reader.
func (c *ColumnInt64) DecodeColumn(r *Reader, rows int) error {
	const size = 64 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			int64(bin.Uint64(data[i:i+size])),
		)
	}
	*c = v
	return nil
}