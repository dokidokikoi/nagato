package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

func TestEs(t *testing.T) {
	cert, _ := ioutil.ReadFile("http_ca.crt")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "ADR*piFezssmbUhhN8*S",
		CACert:   cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	body := &bytes.Buffer{}
	var elems []elm
	terms := map[string][]interface{}{"type": {"playlist"}, "tags": {"golang", "elasticsearch"}}
	elems = append(elems, BuildTerms(terms)...)
	elems = append(elems, BuildBool(nil, BuildMatchs(map[string]string{"title": "资源管理 golang", "content": "资源管理 golang"})))

	b, _ := json.Marshal(BulidQuery(BuildBool(elems, nil), BuildHighLight([]string{"title"}), 0, 0, nil))
	fmt.Println(string(b))

	if err := json.NewEncoder(body).Encode(BulidQuery(BuildBool(elems, nil), BuildHighLight([]string{"title"}), 0, 0, nil)); err != nil {
		panic(err)
	}

	client := commonEs.EsClient[commonEsModel.Blank]{Client: es}
	resp, err := client.SearchDoc("blank", body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}

func TestCreateIndex(t *testing.T) {
	cert, _ := ioutil.ReadFile("http_ca.crt")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "ADR*piFezssmbUhhN8*S",
		CACert:   cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	response, _ := es.Indices.Create("blank", es.Indices.Create.WithBody(strings.NewReader(`
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
	`)))

	t.Log(response)
}
