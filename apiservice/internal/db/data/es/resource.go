package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"

	"github.com/elastic/go-elasticsearch/v8"
)

type resources struct {
	commonEs.EsClient[commonEsModel.Resource]
}

func (r resources) Search(index string, req commonEsModel.ResourceReq) ([]commonEs.Result[commonEsModel.Resource], error) {
	body := &bytes.Buffer{}

	var elems []elm
	terms := make(map[string][]interface{})
	if req.Ext != "" {
		terms["ext"] = []interface{}{req.Ext}
	}
	if req.Dir != nil {
		terms["dir"] = []interface{}{req.Dir}
	}
	if req.Sha256 != "" {
		terms["sha256"] = []interface{}{req.Sha256}
	}
	terms["privacy"] = []interface{}{req.Privacy}

	elems = append(elems, BuildTerms(terms)...)
	elems = append(elems, BuildMatch("all_text", req.Text))

	elems = append(elems, BuildRange("times", req.TimesGte, req.TimesLt))
	elems = append(elems, BuildTimeRange("update_at", req.UpdatedAtGte, req.UpdatedAtLt))
	elems = append(elems, BuildTimeRange("create_at", req.CreatedAtGte, req.CreatedAtLt))

	query := BulidQuery(
		BuildBool(elems, nil),
		BuildHighLight(req.Highlight),
		req.Page,
		req.PageSize,
		req.Select,
	)

	res, _ := json.Marshal(query)
	fmt.Println(string(res))

	if err := json.NewEncoder(body).Encode(query); err != nil {
		panic(err)
	}

	result, err := r.SearchDoc(index, body)
	if err != nil {
		panic(err)
	}

	return result.Hits.Hits, nil
}

func (r resources) CreateDocWithID(index string, id string, req commonEsModel.Resource) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return r.CreateDocByID(index, id, bytes.NewBuffer(body))
}

func newResources(cli *elasticsearch.Client) *resources {
	return &resources{commonEs.EsClient[commonEsModel.Resource]{Client: cli}}
}
