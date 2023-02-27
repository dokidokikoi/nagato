package tools

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
)

// 计算数据散列值并返回
func CalculateHash(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
