package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/zeebo/xxh3"
)

const (
	bufferSize = 32 * 1024
)

var bufPool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, bufferSize)
		return &b
	},
}

func createFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var hash string

	hash, err = hashXXH3(file)

	if err != nil {
		return "", fmt.Errorf("hashing failed: %w", err)
	}

	return hash, nil

}

func hashXXH3(r io.Reader) (string, error) {
	hasher := xxh3.New()
	buffPtr := bufPool.Get().(*[]byte)
	buff := *buffPtr
	defer bufPool.Put(buffPtr)

	if _, err := io.CopyBuffer(hasher, r, buff); err != nil {
		return "", fmt.Errorf("XXH3 hashing failed: %w", err)
	}

	hash := hasher.Sum128()
	return fmt.Sprintf("%016x%016x", hash.Hi, hash.Lo), nil
}
