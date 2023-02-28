package stream

import (
	"io"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	servers := []string{
		":10151",
		":10152",
		":10150",
		":10153",
	}

	uuids := []string{
		"5AB1E4B2-FC87-4463-8FE8-3A1371AC9543",
		"2090FB4C-FD29-40A6-A3A7-544E957F731E",
		"E91553A6-D87F-4B5D-99FC-5D9189139B4A",
		"939BC421-5F04-4DFF-85B9-134988673C13",
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		getStream, err := NewRSResumableGetStream(servers, uuids, 240628)
		if err != nil {
			panic(err)
		}
		io.Copy(w, getStream)
	})

	http.ListenAndServe(":8888", nil)
}
