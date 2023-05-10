package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"

	"github.com/elastic/go-elasticsearch/v8"
)

type blanks struct {
	commonEs.EsClient[commonEsModel.Blank]
}

func (bl blanks) SearchBlank(index string, req commonEsModel.BlankReq) ([]commonEs.Result[commonEsModel.Blank], error) {
	body := &bytes.Buffer{}
	// blank 查找
	var elems []elm
	terms := make(map[string][]interface{})
	if req.Type != "" {
		terms["type"] = []interface{}{req.Type}
	}
	terms["tags"] = make([]interface{}, len(req.Tags))
	for i, v := range req.Tags {
		terms["tags"][i] = v
	}
	elems = append(elems, BuildTerms(terms)...)
	elems = append(elems, BuildMatch("all_text", req.Text))
	elems = append(elems, BuildTimeRange("update_at", req.UpdatedAtGte, req.UpdatedAtLt))
	elems = append(elems, BuildTimeRange("create_at", req.CreatedAtGte, req.CreatedAtLt))

	var shouldElems []elm
	shouldTerms := make(map[string][]interface{})

	if len(req.MatterIDs) > 0 {
		shouldTerms["matter_ids"] = make([]interface{}, len(req.MatterIDs))
		for i, v := range req.MatterIDs {
			shouldTerms["matter_ids"][i] = v
		}
		shouldElems = append(shouldElems, BuildOrTerms("matter_ids", req.MatterIDs))
	}

	elems = append(elems, BuildBool(nil, shouldElems))

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

	result, err := bl.SearchDoc(index, body)
	if err != nil {
		panic(err)
	}

	return result.Hits.Hits, nil
}

func (bl blanks) CreateDocWithID(index string, id string, req commonEsModel.Blank) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return bl.CreateDocByID(index, id, bytes.NewBuffer(body))
}

func newBlanks(cli *elasticsearch.Client) *blanks {
	return &blanks{commonEs.EsClient[commonEsModel.Blank]{Client: cli}}
}
