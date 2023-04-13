package es

import (
	"bytes"
	"encoding/json"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"

	"github.com/elastic/go-elasticsearch/v8"
)

type resources struct {
	commonEs.EsClient[commonEsModel.Resource]
}

func (r resources) Search(index string, req commonEsModel.ResourceReq) ([]commonEs.Result[commonEsModel.Resource], error) {
	body := &bytes.Buffer{}
	// blank 查找
	var elems []elm
	terms := make(map[string][]interface{})

	terms["matters.ext"] = []interface{}{req.Ext}
	terms["matters.dir"] = []interface{}{req.Dir}
	terms["matters.privacy"] = []interface{}{req.Privacy}
	terms["matters.sha256"] = []interface{}{req.Sha256}
	elems = append(elems, BuildTerms(terms)...)
	elems = append(elems, BuildMatch("matters.al_text", req.Text))

	elems = append(elems, BuildRange("matters.size", req.TimesGte, req.TimesLt))
	elems = append(elems, BuildRange("matters.update_at", req.UpdatedAtGte, req.UpdatedAtLt))
	elems = append(elems, BuildRange("matters.create_at", req.CreatedAtGte, req.CreatedAtLt))

	// 嵌套对象查找
	var nestedElems []elm
	nestedTerms := make(map[string][]interface{})

	nestedTerms["children.ext"] = []interface{}{req.Ext}
	nestedTerms["children.dir"] = []interface{}{req.Dir}
	nestedTerms["children.privacy"] = []interface{}{req.Privacy}
	nestedTerms["children.sha256"] = []interface{}{req.Sha256}
	nestedElems = append(nestedElems, BuildTerms(nestedTerms)...)
	nestedElems = append(nestedElems, BuildMatch("children.al_text", req.Text))

	nestedElems = append(nestedElems, BuildRange("children.size", req.TimesGte, req.TimesLt))
	nestedElems = append(nestedElems, BuildRange("children.update_at", req.UpdatedAtGte, req.UpdatedAtLt))
	nestedElems = append(nestedElems, BuildRange("children.create_at", req.CreatedAtGte, req.CreatedAtLt))

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
	return r.CreateDocByID(index, id, bytes.NewBuffer(body))
}

func newResources(cli *elasticsearch.Client) *resources {
	return &resources{commonEs.EsClient[commonEsModel.Resource]{Client: cli}}
}
