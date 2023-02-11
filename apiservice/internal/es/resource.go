package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	commonEs "nagato/common/es"
	"strings"
)

const INDEX_RESOURCE = "metadata"

type Resource struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
	Size    int64  `json:"size"`
	Hash    string `json:"hash"`
}

type resourceEs struct {
	cli commonEs.EsClient[Resource]
}

func (r *resourceEs) CreateLastestResource(name, hash string, size int64) error {
	latest, err := r.SearchLatestVersion(name)
	if err != nil {
		return err
	}
	latestVersion := 1
	if latest != nil {
		latestVersion = latest.Version + 1
	}

	body := &bytes.Buffer{}
	err = json.NewEncoder(body).Encode(&Resource{
		Name:    name,
		Version: latestVersion,
		Size:    size,
		Hash:    hash,
	})
	if err != nil {
		return err
	}
	err = r.cli.CreateDocByID(INDEX_RESOURCE, fmt.Sprintf("%s_%d", name, latestVersion), body)
	return err
}

func (r *resourceEs) SearchLatestVersion(name string) (*Resource, error) {
	doc := fmt.Sprintf(`{"query":{"match":{"name":"%s"}},"sort": [{"version":{"order":"desc"}}],"from":0,"size":1}`, name)
	searchRes, err := r.cli.SearchDoc(INDEX_RESOURCE, strings.NewReader(doc))
	if err != nil {
		return nil, err
	}

	list := searchRes.GetSourceList()
	if list == nil {
		return nil, nil
	}

	return list[0], nil
}

func (r *resourceEs) GetResourceMate(name string, version int) (*Resource, error) {
	res, err := r.cli.GetDoc(INDEX_RESOURCE, fmt.Sprintf("%s_%d", name, version))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *resourceEs) SearchResourceAllVersion(name string, from, size int) ([]*Resource, error) {
	doc := fmt.Sprintf(`{"query":?,"sort": [{"version":{"order":"desc"}}],"from":%d,"size":%d}`, from, size)
	if strings.Trim(name, " ") != "" {
		doc = strings.Replace(doc, "?", fmt.Sprintf(`{"match_phrase_prefix":{"name":"%s"}}`, name), 1)
	} else {
		doc = strings.Replace(doc, "?", `{"match_all":{}}`, 1)
	}
	searchRes, err := r.cli.SearchDoc(INDEX_RESOURCE, strings.NewReader(doc))
	if err != nil {
		return nil, err
	}

	list := searchRes.GetSourceList()
	if list == nil {
		return nil, nil
	}

	return list, nil
}

func newResource(cli *esStore) *resourceEs {
	return &resourceEs{
		cli: commonEs.EsClient[Resource]{
			Client: cli.cli,
		},
	}
}
