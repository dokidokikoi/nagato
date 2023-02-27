package tools

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

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

func GetOffsetFromHeader(h http.Header) uint {
	byteRange := h.Get("range")
	if len(byteRange) < 7 {
		return 0
	}
	if byteRange[:6] != "bytes=" {
		return 0
	}
	bytePos := strings.Split(byteRange[6:], "-")
	offset, _ := strconv.ParseUint(bytePos[0], 0, 32)
	return uint(offset)
}

func ParseRangeFromHeader(h http.Header) (start, end int64, err error) {
	rangeHeader := h.Get("range")
	if rangeHeader == "" {
		return 0, 0, errors.New("no range header")
	}

	bytesRange := strings.Split(rangeHeader, "=")
	if len(bytesRange) < 2 {
		return 0, 0, errors.New("no range header")
	}
	if strings.Trim(bytesRange[0], " ") != "bytes" {
		return 0, 0, errors.New("only support bytes range")
	}
	bytePos := strings.Split(bytesRange[1], "-")
	if len(bytePos) == 1 {
		start, err = strconv.ParseInt(bytePos[0], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		end = -1
	} else if len(bytePos) == 2 {
		start, err = strconv.ParseInt(bytePos[0], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		end, err = strconv.ParseInt(bytePos[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}
	} else {
		return 0, 0, errors.New("invalid range header format")
	}

	return
}
