package compress

import (
	"fmt"
	"io"

	"github.com/go-faster/city"
	"github.com/go-faster/errors"
	"github.com/pierrec/lz4/v4"
)

// Reader decodes compressed blocks.
type Reader struct {
	reader io.Reader
	data   []byte
	pos    int64
	raw    []byte
	header []byte
}

func formatU128(v city.U128) string {
	return fmt.Sprintf("0x%x%x", v.Low, v.High)
}

// readBlock reads next compressed data into raw and decompresses into data.
func (r *Reader) readBlock() error {
	r.pos = 0

	_ = r.header[headerSize-1]
	if _, err := io.ReadFull(r.reader, r.header); err != nil {
		return errors.Wrap(err, "header")
	}

	var (
		rawSize  = int(bin.Uint32(r.header[hRawSize:])) - rawSizeOffset
		dataSize = int(bin.Uint32(r.header[hDataSize:]))
	)
	if dataSize < 0 || dataSize > maxDataSize {
		return errors.Errorf("data size should be %d < %d < %d", 0, dataSize, maxDataSize)
	}
	if rawSize < 0 || rawSize > maxBlockSize {
		return errors.Errorf("raw size should be %d < %d < %d", 0, rawSize, maxBlockSize)
	}

	r.data = append(r.data[:0], make([]byte, dataSize)...)
	r.raw = append(r.raw[:0], r.header...)
	r.raw = append(r.raw, make([]byte, rawSize)...)
	_ = r.raw[:rawSize+headerSize-1]

	if _, err := io.ReadFull(r.reader, r.raw[headerSize:]); err != nil {
		return errors.Wrap(err, "read raw")
	}
	hGot := city.U128{
		Low:  bin.Uint64(r.raw[0:8]),
		High: bin.Uint64(r.raw[8:16]),
	}
	h := city.CH128(r.raw[hMethod:])
	if hGot != h {
		return errors.Errorf("data corrupted: hash mismatch: %v (actual) != %v (got in header)",
			formatU128(h), formatU128(hGot),
		)
	}
	switch m := Method(r.header[hMethod]); m {
	case LZ4:
		n, err := lz4.UncompressBlock(r.raw[headerSize:], r.data)
		if err != nil {
			return errors.Wrap(err, "uncompress")
		}
		if n != dataSize {
			return errors.Errorf("unexpected uncompressed data size: %d (actual) != %d (got in header)",
				n, dataSize,
			)
		}
	default:
		return errors.Errorf("compression 0x%02x not implemented", m)
	}

	return nil
}

// Read implements io.Reader.
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.pos >= int64(len(r.data)) {
		if err := r.readBlock(); err != nil {
			return 0, errors.Wrap(err, "read next block")
		}
	}
	n = copy(p, r.data[r.pos:])
	r.pos += int64(n)
	return n, nil
}

// NewReader returns new *Reader from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		reader: r,
		header: make([]byte, headerSize),
	}
}