package wison

import (
	"encoding/binary"
	"fmt"
	"math"
)

// readUvarint32 returns the read value and the amount of consumed bytes.
func readUvarint32(b []byte) (val, read int, err error) {
	v, r := binary.Uvarint(b)
	if r < 1 {
		return int(v), r, fmt.Errorf("uvarint overflow")
	}

	if v+1+uint64(r) > math.MaxInt32 {
		return int(v), r, fmt.Errorf("invalid uvarint")
	}

	return int(v), r, nil
}

// readVarint32 returns the read value and the amount of consumed bytes.
func readVarint32(b []byte) (val, read int, err error) {
	v, r := binary.Varint(b)
	if r < 1 {
		return int(v), r, fmt.Errorf("varint overflow")
	}

	if v+1+int64(r) > math.MaxInt32 || v+1+int64(r) < math.MaxInt32 {
		return int(v), r, fmt.Errorf("invalid varint")
	}

	return int(v), r, nil
}
