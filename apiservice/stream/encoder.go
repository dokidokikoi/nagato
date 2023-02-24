package stream

import (
	"crypto/sha256"
	"hash"
	"io"

	"github.com/klauspost/reedsolomon"
)

const (
	BLOCK_PER_SHARD = 8000
	BLOCK_SIZE      = BLOCK_PER_SHARD * DATA_SHARDS
)

type encoder struct {
	writers []io.Writer
	hashs   []hash.Hash
	enc     reedsolomon.Encoder
	cache   []byte
}

func (e *encoder) Write(p []byte) (n int, err error) {
	length := len(p)
	current := 0
	for length != 0 {
		next := BLOCK_SIZE - len(e.cache)
		if next > length {
			next = length
		}
		e.cache = append(e.cache, p[current:current+next]...)

		if len(e.cache) == BLOCK_SIZE {
			e.Flush()
		}
		current += next
		length -= next
	}
	return len(p), nil
}

func (e *encoder) Flush() {
	if len(e.cache) == 0 {
		return
	}

	shards, _ := e.enc.Split(e.cache)
	e.enc.Encode(shards)
	for i := range shards {
		e.writers[i].Write(shards[i])
		e.hashs[i].Write(shards[i])
	}
	e.cache = []byte{}
}

func NewEncoder(writers []io.Writer) *encoder {
	enc, _ := reedsolomon.New(DATA_SHARDS, PARITY_SHARDS)
	hashs := make([]hash.Hash, len(writers))
	for i, _ := range hashs {
		hashs[i] = sha256.New()
	}
	return &encoder{writers, hashs, enc, nil}
}
