package es

import (
	"nagato/common/es"
)

func CreateLastestResource(name, hash string, size int64) error {
	return es.AddVersion(name, hash, size)
}

func GetResourceMate(name string, version int) (*es.Metadata, error) {
	return es.GetMetadata(name, version)
}

func SearchResourceLatestVersion(name string) (*es.Metadata, error) {
	return es.SearchLatestVersion(name)
}

func SearchResourceAllVersion(name string, from, size int) ([]es.Metadata, error) {
	return es.SearchAllVersion(name, from, size)
}
