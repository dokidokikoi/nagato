package tools

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
)

// 计算数据散列值并返回
func CalculateHash(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func GetHashFromHeader(h http.Header) string {
	digest := h.Get("digest")
	if len(digest) < 9 {
		return ""
	}

	if digest[:8] != "SHA-256=" {
		return ""
	}

	return digest[8:]
}

func GetSizeFromHeader(h http.Header) uint {
	size, _ := strconv.ParseInt(h.Get("content-length"), 0, 32)

	return uint(size)
}

// 每个分片的大小,size/DATA_SHARDS 再向上取整
func CalculatePerShard(a, b uint) uint {
	return (a + b - 1) / b
}
