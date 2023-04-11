package service

import (
	"nagato/searchservice/internal/db"
	"nagato/searchservice/internal/db/data"
)

type IService interface {
}

type srv struct {
	store db.Store
}

func NewSrv() IService {

	return &srv{store: data.GetDataCenter()}
}
