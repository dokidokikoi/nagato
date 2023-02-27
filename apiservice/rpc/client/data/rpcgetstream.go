package data

type RpcGetStream struct {
	client    DataRecv
	cache     []byte
	cacheSize int
}

func (r *RpcGetStream) Read(p []byte) (n int, err error) {
	if r.cacheSize <= 0 {
		err = r.getData()
		if err != nil {
			return 0, err
		}
	}

	length := len(p)
	if r.cacheSize < length {
		length = r.cacheSize
	}
	r.cacheSize -= length
	copy(p, r.cache[:length])
	r.cache = r.cache[length:]

	return length, nil
}

func (r *RpcGetStream) getData() error {
	req, err := r.client.Recv()
	if err != nil {
		return err
	}

	r.cache = req.GetData()
	r.cacheSize = len(r.cache)
	return nil
}

func NewRpcGetStream(client DataRecv) *RpcGetStream {
	return &RpcGetStream{client, nil, 0}
}
