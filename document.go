package wison

import (
	"fmt"
)

type Type int8

const (
	UNKNOWN Type = iota
	NULL
	STRING
	BLOB
	BOOL
	INT
	UINT
	FLOAT64
	FLOAT32
	UUID
	MAP
	ARRAY
	DATA
	DOC
)

// Node represents any abstract wison node.
type Node []byte

// Type returns the wison data type.
func (d Node) Type() Type {
	if len(d) == 0 {
		return UNKNOWN
	}

	return d[0]
}

// Len returns the size of the node to skip to the next node (re-slice) and includes
// the type byte.
func (d Node) Len() (int, error) {
	switch d.Type() {
	case NULL:
		return 1, nil
	case STRING:
		fallthrough
	case BLOB:
		len, read, err := readUvarint32(d[1:])
		return len + read + 1, err
	case BOOL:
		return 1 + 1, nil
	case FLOAT64:
		return 8 + 1, nil
	case FLOAT32:
		return 4 + 1, nil
	case UUID:
		return 16 + 1, nil
	case UINT:
		_, read, err := readUvarint32(d[1:])
		return read + 1, err
	case INT:
		_, read, err := readVarint32(d[1:])
		return read + 1, err
	case UNKNOWN:
		fallthrough
	default:
		return 0, fmt.Errorf("not a node")
	}
}

// Document represents an in-memory document.
type Document []byte

// Type returns the wison data type.
func (d Document) Type() Type {
	return Node(d).Type()
}

// Root returns nil or the root node.
func (d Document) Root() Node {
	if d[1] >= NULL && d[1] < DOC {
		return Node(d[1:])
	}

	return nil
}
