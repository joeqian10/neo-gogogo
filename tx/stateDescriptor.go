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
func (s *StateDescriptor) Deserialize(r *io.BinaryReader) {
	r.ReadLE(&s.Type)

	s.Key = r.ReadVarBytes()
	s.Value = r.ReadVarBytes()
	s.Field = r.ReadVarString()
}

// Serialize implements Serializable interface.
func (s *StateDescriptor) Serialize(w *io.BinaryWriter) {
	w.WriteLE(s.Type)
	w.WriteVarBytes(s.Key)
	w.WriteVarBytes(s.Value)
	w.WriteVarString(s.Field)
}
