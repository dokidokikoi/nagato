package matter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"nagato/apiservice/internal/model"
	"nagato/apiservice/rpc/server"
	"nagato/common/es"
	"strings"
)

const INDEX_RESOURCE = "resource"

func (s matterSrv) CreateResource(ctx context.Context, example *model.Matter) error {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(example)
	if err != nil {
		return err
	}

	err = server.SearchSrv.CreateDocByID(ctx, INDEX_RESOURCE, fmt.Sprintf("%d", example.ID), body.String())
	return err
}

func (s matterSrv) GetResourceMate(ctx context.Context, name string, version int) (*model.Matter, error) {
	resBytes, err := server.SearchSrv.GetDoc(ctx, INDEX_RESOURCE, fmt.Sprintf("%s_%d", name, version))
	if err != nil {
		return nil, err
	}

	res := &model.Matter{}
	json.Unmarshal(resBytes, res)
	return res, nil
}

func (s matterSrv) SearchResourceAllVersion(ctx context.Context, name string, from, size int) ([]*model.Matter, error) {
	doc := fmt.Sprintf(`{"query":?,"sort": [{"version":{"order":"desc"}}],"from":%d,"size":%d}`, from, size)
	if strings.Trim(name, " ") != "" {
		doc = strings.Replace(doc, "?", fmt.Sprintf(`{"match_phrase_prefix":{"name":"%s"}}`, name), 1)
	} else {
		doc = strings.Replace(doc, "?", `{"match_all":{}}`, 1)
	}
	searchBytes, err := server.SearchSrv.SearchDoc(ctx, INDEX_RESOURCE, doc)
	if err != nil {
		return nil, err
	}

	searchRes := &es.SearchResult[model.Matter]{}
	json.Unmarshal(searchBytes, searchRes)
	list := searchRes.GetSourceList()
	if list == nil {
		return nil, nil
	}

	return list, nil
}

func (s matterSrv) SearchLatestVersion(ctx context.Context, name string) (*model.Matter, error) {
	doc := fmt.Sprintf(`{"query":{"match":{"name":"%s"}},"sort": [{"version":{"order":"desc"}}],"from":0,"size":1}`, name)
	searchBytes, err := server.SearchSrv.SearchDoc(ctx, INDEX_RESOURCE, doc)
	if err != nil {
		return nil, err
	}

	searchRes := &es.SearchResult[model.Matter]{}
	json.Unmarshal(searchBytes, searchRes)
	list := searchRes.GetSourceList()
	if list == nil {
		return nil, nil
	}

	return list[0], nil
}
