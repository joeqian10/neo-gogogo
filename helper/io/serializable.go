package io

import (
	"bytes"
	"io"
)

// Serializable defines the binary encoding/decoding interface. Errors are
// returned via BinaryReader/BinaryWriter Err field. These functions must have safe
// behavior when passed BinaryReader/BinaryWriter with Err already set. Invocations
// to these functions tend to be nested, with this mechanism only the top-level
// caller should handle the error once and all the other code should just not
// panic in presence of error.
type Serializable interface {
	Deserialize(*BinaryReader)
	Serialize(*BinaryWriter)
}

func ToArray(se Serializable) ([]byte, error) {
	buffer := &bytes.Buffer{}
	writer := NewBinaryWriterFromIO(io.Writer(buffer))
	se.Serialize(writer)
	return buffer.Bytes(), writer.Err
}

func AsSerializable(se Serializable, data []byte) error {
	buffer := bytes.NewBuffer(data)
	reader := NewBinaryReaderFromIO(io.Reader(buffer))
	se.Deserialize(reader)
	return reader.Err
}
