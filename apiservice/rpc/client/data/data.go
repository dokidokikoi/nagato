package data

import pb "nagato/common/proto/data"

type DataRecv interface {
	Recv() (Data, error)
}

type Data interface {
	GetData() []byte
}

type TempFileDataRecv struct {
	cli pb.Data_GetTempFileClient
}

func (t TempFileDataRecv) Recv() (Data, error) {
	return t.cli.Recv()
}

type FileDataRecv struct {
	cli pb.Data_GetMatterClient
}

func (t FileDataRecv) Recv() (Data, error) {
	return t.cli.Recv()
}
