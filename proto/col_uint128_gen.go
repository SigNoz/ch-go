// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"github.com/go-faster/errors"
)

// ColUInt128 represents UInt128 column.
type ColUInt128 []UInt128

// Compile-time assertions for ColUInt128.
var (
	_ Input  = ColUInt128{}
	_ Result = (*ColUInt128)(nil)
	_ Column = (*ColUInt128)(nil)
)

// Type returns ColumnType of UInt128.
func (ColUInt128) Type() ColumnType {
	return ColumnTypeUInt128
}

// Rows returns count of rows in column.
func (c ColUInt128) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt128) Reset() {
	*c = (*c)[:0]
}

// NewArrUInt128 returns new Array(UInt128).
func NewArrUInt128() *ColArr {
	return &ColArr{
		Data: new(ColUInt128),
	}
}

// AppendUInt128 appends slice of UInt128 to Array(UInt128).
func (c *ColArr) AppendUInt128(data []UInt128) {
	d := c.Data.(*ColUInt128)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes UInt128 rows to *Buffer.
func (c ColUInt128) EncodeColumn(b *Buffer) {
	const size = 128 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binPutUInt128(
			b.Buf[offset:offset+size],
			v,
		)
		offset += size
	}
}

// DecodeColumn decodes UInt128 rows from *Reader.
func (c *ColUInt128) DecodeColumn(r *Reader, rows int) error {
	const size = 128 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			binUInt128(data[i:i+size]),
		)
	}
	*c = v
	return nil
}
