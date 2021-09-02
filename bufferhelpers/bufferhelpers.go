package bufferhelpers

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"mumbletest/bufferpool"
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

// a wrapper over ReadByte
func ReadBEUint8(r *bufio.Reader) (uint8, error) {
	return r.ReadByte()
}

func ReadBEUint16(r *bufio.Reader) (uint16, error) {
	b := make([]byte, 2)
	if _, err := io.ReadFull(r, b); err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(b), nil
}

func ReadBEUint32(r *bufio.Reader) (uint32, error) {
	b := make([]byte, 4)
	if _, err := io.ReadFull(r, b); err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(b), nil
}

func ReadBEUint64(r *bufio.Reader) (uint64, error) {
	b := make([]byte, 8)
	if _, err := io.ReadFull(r, b); err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint64(b), nil
}

// Reads exactly len bytes from r into a bytes.Buffer pulled from bufferpool.GetBuffer
func ReadLenIntoBuf(r io.Reader, len int64) (*bytes.Buffer, error) {
	buf := bufferpool.GetBuffer()

	n, err := buf.ReadFrom(io.LimitReader(r, len))
	if err != nil {
		return nil, err
	}
	if n != len {
		return nil, fmt.Errorf("ReadLenIntoBuf: bytes read (%d) != required len (%d)", n, len)
	}

	return buf, nil
}

func BufferIntoReader(buf *bytes.Buffer) *bytes.Reader {
	return bytes.NewReader(buf.Bytes())
}
