package bufferpool

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

// Get a blank buffer
func GetBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

// Return the buffer, you do not have to reset it before putting it back
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
