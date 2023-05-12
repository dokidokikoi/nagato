package matter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"nagato/apiservice/internal/model"
	"nagato/apiservice/rpc/client"
	"nagato/common/es"
	commonEsModel "nagato/common/es/model"
	"strings"
)

const INDEX_RESOURCE = "resource"

func (s matterSrv) CreateResource(ctx context.Context, example *model.Matter) error {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(example)
	if err != nil {
		return err
	}

	err = client.SearchSrv.CreateDocByID(ctx, INDEX_RESOURCE, fmt.Sprintf("%d", example.ID), body.String())
	return err
}

func (s matterSrv) GetResourceMate(ctx context.Context, name string, version int) (*model.Matter, error) {
	resBytes, err := client.SearchSrv.GetDoc(ctx, INDEX_RESOURCE, fmt.Sprintf("%s_%d", name, version))
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
	searchBytes, err := client.SearchSrv.SearchDoc(ctx, INDEX_RESOURCE, doc)
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
	searchBytes, err := client.SearchSrv.SearchDoc(ctx, INDEX_RESOURCE, doc)
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

func (r matterSrv) Search(ctx context.Context, userID uint, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Resource, int64, error) {
	res, err := r.store.Matters().SearchResource(userID, resourceReq)
	if err != nil {
		return nil, 0, err
	}

	resource := make([]commonEsModel.Resource, len(res))
	for i := range res {
		resource[i] = res[i].Source
	}

	return resource, 0, nil
}

func (r matterSrv) CreateIndices(userID uint, indexReq string) error {
	return r.store.Matters().CreateIndices(userID, indexReq)
}

func (r matterSrv) CreateDocWithID(userID uint, id string, req commonEsModel.Resource) error {
	return r.store.Matters().CreateDocWithID(userID, id, req)
}

func (r matterSrv) UpdateDoc(userID uint, id string, req commonEsModel.Resource) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return r.store.Matters().UpdateDoc(userID, id, bytes.NewBuffer(data))
}

func (r matterSrv) DelDoc(userID uint, id string) error {
	return r.store.Matters().DelDoc(userID, id)
}
