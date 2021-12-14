// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"github.com/go-faster/errors"
)

// ColUInt16 represents UInt16 column.
type ColUInt16 []uint16

// Compile-time assertions for ColUInt16.
var (
	_ Input  = ColUInt16{}
	_ Result = (*ColUInt16)(nil)
	_ Column = (*ColUInt16)(nil)
)

// Type returns ColumnType of UInt16.
func (ColUInt16) Type() ColumnType {
	return ColumnTypeUInt16
}

// Rows returns count of rows in column.
func (c ColUInt16) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt16) Reset() {
	*c = (*c)[:0]
}

// NewArrUInt16 returns new Array(UInt16).
func NewArrUInt16() *ColArr {
	return &ColArr{
		Data: new(ColUInt16),
	}
}

// AppendUInt16 appends slice of uint16 to Array(UInt16).
func (c *ColArr) AppendUInt16(data []uint16) {
	d := c.Data.(*ColUInt16)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes UInt16 rows to *Buffer.
func (c ColUInt16) EncodeColumn(b *Buffer) {
	const size = 16 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		bin.PutUint16(
			b.Buf[offset:offset+size],
			v,
		)
		offset += size
	}
}

// DecodeColumn decodes UInt16 rows from *Reader.
func (c *ColUInt16) DecodeColumn(r *Reader, rows int) error {
	const size = 16 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			bin.Uint16(data[i:i+size]),
		)
	}
	*c = v
	return nil
}
