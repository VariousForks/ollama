package blob

import (
	"io"
	"os"

	"github.com/ollama/ollama/server/internal/chunks"
)

type Chunker struct {
	size int64
	f    *os.File

	completed []chunks.Chunk
}

func (cw *Chunker) Put(d Digest, size int64, r io.Reader) (int, error) {
	panic("TODO")
}

func (cw *Chunker) Complete() bool {
	panic("TODO")
}

func (cw *Chunker) Close() error {
	return cw.f.Close()
}

// completed returns a single chunk that covers the entire size.
func completed(size int64) []chunks.Chunk {
	return []chunks.Chunk{{End: size - 1}}
}

func (c *DiskCache) Chunked(d Digest, size int64) (*Chunker, error) {
	name := c.GetFile(d)
	info, err := os.Stat(name)
	if err == nil && info.Size() == size {
		return &Chunker{completed: completed(size)}, nil
	}
}
