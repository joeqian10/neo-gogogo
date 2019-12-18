package io

import (
	"bytes"
	"encoding/binary"
	"io"
)

// BinaryReader is a convenient wrapper around a io.Reader and err object.
// Used to simplify error handling when reading into a struct with many fields.
type BinaryReader struct {
	r   io.Reader
	Err error
}

// NewBinaryReaderFromIO makes a BinaryReader from io.Reader.
func NewBinaryReaderFromIO(ior io.Reader) *BinaryReader {
	return &BinaryReader{r: ior}
}

// NewBinaryReaderFromBuf makes a BinaryReader from byte buffer.
func NewBinaryReaderFromBuf(b []byte) *BinaryReader {
	r := bytes.NewReader(b)
	return NewBinaryReaderFromIO(r)
}

// ReadLE reads from the underlying io.Reader
// into the interface v in little-endian format.
func (r *BinaryReader) ReadLE(v interface{}) {
	if r.Err != nil {
		return
	}
	r.Err = binary.Read(r.r, binary.LittleEndian, v)
}

// ReadBE reads from the underlying io.Reader
// into the interface v in big-endian format.
func (r *BinaryReader) ReadBE(v interface{}) {
	if r.Err != nil {
		return
	}
	r.Err = binary.Read(r.r, binary.BigEndian, v)
}

// ReadUnit64Bytes reads from the underlying io.Reader
// into the interface v in little-endian format
func (br *BinaryReader) ReadUnit64Bytes() []byte {
	b := make([]byte, 8)
	br.ReadLE(b)
	if br.Err != nil {
		return nil
	}
	return b;
}

// ReadVarUint reads a variable-length-encoded integer from the
// underlying reader.
func (r *BinaryReader) ReadVarUint() uint64 {
	if r.Err != nil {
		return 0
	}

	var b uint8
	r.Err = binary.Read(r.r, binary.LittleEndian, &b)

	if b == 0xfd {
		var v uint16
		r.Err = binary.Read(r.r, binary.LittleEndian, &v)
		return uint64(v)
	}
	if b == 0xfe {
		var v uint32
		r.Err = binary.Read(r.r, binary.LittleEndian, &v)
		return uint64(v)
	}
	if b == 0xff {
		var v uint64
		r.Err = binary.Read(r.r, binary.LittleEndian, &v)
		return v
	}

	return uint64(b)
}

// ReadBytes reads the next set of bytes from the underlying reader.
// ReadVarUInt() is used to determine how large that slice is
func (r *BinaryReader) ReadBytes() []byte {
	n := r.ReadVarUint()
	b := make([]byte, n)
	r.ReadLE(b)
	return b
}

// ReadString calls ReadBytes and casts the results as a string.
func (r *BinaryReader) ReadString() string {
	b := r.ReadBytes()
	return string(b)
}
