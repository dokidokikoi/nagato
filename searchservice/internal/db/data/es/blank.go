package es

import (
	"bytes"
	"encoding/json"
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
	elems = append(elems, BuildBool(nil, []elm{BuildMatch("all_text", req.Text)}))
	elems = append(elems, BuildRange("update_at", req.UpdatedAtGte, req.UpdatedAtLt))
	elems = append(elems, BuildRange("create_at", req.CreatedAtGte, req.CreatedAtLt))

	// 嵌套对象查找
	var nestedElems []elm
	nestedTerms := make(map[string][]interface{})
	nestedTerms["matters.ext"] = []interface{}{req.Ext}
	nestedTerms["matters.dir"] = []interface{}{req.Dir}
	nestedTerms["matters.privacy"] = []interface{}{req.Privacy}
	nestedTerms["matters.sha256"] = []interface{}{req.Sha256}
	nestedElems = append(nestedElems, BuildTerms(nestedTerms)...)
	nestedElems = append(nestedElems, BuildMatch("matters.al_text", req.Text))

	nestedElems = append(nestedElems, BuildRange("matters.size", req.TimesGte, req.TimesLt))
	nestedElems = append(nestedElems, BuildRange("matters.update_at", req.UpdatedAtGte, req.UpdatedAtLt))
	nestedElems = append(nestedElems, BuildRange("matters.create_at", req.CreatedAtGte, req.CreatedAtLt))

	query := BulidQuery(
		BuildBool(
			nil,
			[]elm{
				BuildBool(elems, nil),
				BuildNested(req.Nested, BulidQuery(BuildBool(nestedElems, nil), nil)),
			},
		),
		BuildHighLight(req.Highlight),
	)

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
	return bl.CreateDocByID(index, id, bytes.NewBuffer(body))
}

func newBlanks(cli *elasticsearch.Client) *blanks {
	return &blanks{commonEs.EsClient[commonEsModel.Blank]{Client: cli}}
}

var blankIndex = `
{
	"mappings": {
	  "properties": {
		"id": {
		  "type": "long"
		},
		"type": {
		  "type": "keyword"
		},
		"title": {
		  "type": "text",
		  "analyzer": "ik_smart",
		  "copy_to": ["all_text^2"]
		},
		"content": {
		  "type": "text",
		  "analyzer": "ik_smart",
		  "copy_to": ["all_text"]
		},
		"tags": {
		  "type": "keyword"
		},
		"matter_ids": {
		  "type": "long"
		},
		"updated_at": {
		  "type": "date"
		},
		"created_at": {
		  "type": "date"
		},
		"all_text": {
		  "type": "text"
		}
	  }
	}
  }
`
