package tx

import (
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// StateType represents the type of StateDescriptor.
type StateType uint8

// Valid DescStateType constants.
const (
	Account   StateType = 0x40
	Validator StateType = 0x48
)

// StateDescriptor ..
type StateDescriptor struct {
	Type  StateType
	Key   []byte
	Value []byte
	Field string
}

// Deserialize implements Serializable interface.
func (s *StateDescriptor) Deserialize(r *io.BinReader) {
	r.ReadLE(&s.Type)

	s.Key = r.ReadBytes()
	s.Value = r.ReadBytes()
	s.Field = r.ReadString()
}

// Serialize implements Serializable interface.
func (s *StateDescriptor) Serialize(w *io.BinWriter) {
	w.WriteLE(s.Type)
	w.WriteBytes(s.Key)
	w.WriteBytes(s.Value)
	w.WriteString(s.Field)
}
