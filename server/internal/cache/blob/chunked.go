package blob

import (
	"crypto/sha256"
	"io"
	"os"
	"sync"

	"github.com/ollama/ollama/server/internal/chunks"
)

type Chunker struct {
	cache *DiskCache
	size  int64
	f     *os.File // nil means pre-validated

	mu        sync.Mutex
	completed []chunks.Chunk
}

func (cw *Chunker) Complete() bool {
	panic("TODO")
}

// Put puts a chunk of data into the chunked file. The chunk must not overlap
// with any previously put chunks. The Digest is the digest of the data in the
// chunk, not the whole file.
func (cw *Chunker) Put(c chunks.Chunk, d Digest, r io.Reader) (int, error) {
	if cw.f == nil {
		return 0, os.ErrInvalid
	}
	w := &checkWriter{
		d:      d,
		offset: c.Start,
		size:   c.Size(),
		h:      sha256.New(),
		f:      cw.f,
	}
	_ = w
	panic("TODO")
}

// Close closes the chunked file. It must be called after all calls to Put.
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

	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return nil, err
	}

	return &Chunker{size: size, f: f}, nil
}
