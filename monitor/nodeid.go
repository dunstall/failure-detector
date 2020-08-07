package monitor

import (
	"encoding/binary"
	"errors"
)

var (
	ErrInvalidID = errors.New("invalid node ID")
)

type NodeID uint32

func (id *NodeID) Encode() []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(*id))
	return b
}

func (id *NodeID) Decode(b []byte) error {
	if len(b) < 4 {
		return ErrInvalidID
	}
	*id = NodeID(binary.BigEndian.Uint32(b[:4]))
	return nil
}
