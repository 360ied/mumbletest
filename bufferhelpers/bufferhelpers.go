package bufferhelpers

import (
	"bytes"
	"encoding/binary"
)

// a wrapper over WriteByte
func WriteBEUint8(buf *bytes.Buffer, i uint8) {
	buf.WriteByte(i)
}

func WriteBEUint16(buf *bytes.Buffer, i uint16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	buf.Write(b)
}
func WriteBEUint32(buf *bytes.Buffer, i uint32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	buf.Write(b)
}
func WriteBEUint64(buf *bytes.Buffer, i uint64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	buf.Write(b)
}
