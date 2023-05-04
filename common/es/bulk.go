package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var (
	batch int = 255
)

type Bulk[T any] struct {
	Index string
	Data  []T
}

func (b Bulk[T]) BuildMeta(getID func(t T) interface{}) [][]byte {
	count := len(b.Data)
	var buf bytes.Buffer
	var res [][]byte
	for i, a := range b.Data {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%v" } }%s`, getID(a), "\n"))

		data, err := json.Marshal(a)
		if err != nil {
			log.Fatalf("Cannot encode meta %d: %s, index: %s", getID(a), err, b.Index)
		}
		data = append(data, "\n"...) // <-- Comment out to trigger failure for batch

		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)

		if i > 0 && i%batch == 0 || i == count-1 {
			res = append(res, buf.Bytes())
		}
		buf.Reset()
	}

	return res
}
