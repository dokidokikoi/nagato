package es

import "time"

type elm map[string]interface{}

func BulidQuery(esBool, highlight elm, from, size int, selectField []string) elm {
	res := elm{
		"from": from,
		"sort": []interface{}{
			elm{"_score": "desc"},
			elm{"_doc": "asc"},
		},
	}
	if esBool != nil {
		res["query"] = esBool
	}
	if highlight != nil {
		res["highlight"] = highlight
	}
	if size > 0 {
		res["size"] = size
	}

	if len(selectField) > 0 {
		res["_source"] = selectField
	}
	return res
}

func BuildBool(must, should []elm) elm {
	mustElms := BuildMust(must)
	shouldElms := BuildShould(should)
	if mustElms != nil && shouldElms != nil {
		return elm{
			"bool": elm{
				"must":   mustElms["must"],
				"should": mustElms["should"],
			},
		}
	}
	if mustElms != nil {
		return elm{
			"bool": elm{
				"must": mustElms["must"],
			},
		}
	}
	if shouldElms != nil {
		return elm{
			"bool": elm{
				"should": mustElms["should"],
			},
		}
	}
	return nil
}

func BuildMust(element []elm) elm {
	var res []elm
	for _, e := range element {
		if e != nil {
			res = append(res, e)
		}
	}
	if len(res) <= 0 {
		return nil
	}
	return elm{
		"must": res,
	}
}

func BuildShould(element []elm) elm {
	var res []elm
	for _, e := range element {
		if e != nil {
			res = append(res, e)
		}
	}
	if len(res) <= 0 {
		return nil
	}
	return elm{
		"should": res,
	}
}

func BuildTerm(field, value string) elm {
	return elm{
		"term": elm{
			field: elm{
				"value": value,
			},
		},
	}
}

func BuildTerms(m map[string][]interface{}) []elm {
	var res []elm
	for k, v := range m {
		for _, val := range v {
			res = append(res, elm{
				"term": elm{
					k: elm{
						"value": val,
					},
				},
			})
		}
	}

	return res
}

func BuildOrTerms(field string, m interface{}) elm {

	return elm{
		"terms": elm{
			field: m,
		},
	}
}

func BuildMatchPhrasePrefix(field, value string) elm {
	return elm{
		"match_phrase_prefix": elm{
			field: elm{
				"query":          value,
				"max_expansions": 50,
				"slop":           50,
			},
		},
	}
}

func BuildMatchPhrasePrefixs(m map[string]string) []elm {
	var res []elm
	for k, v := range m {
		res = append(res, elm{
			"match_phrase_prefix": elm{
				k: elm{
					"query":          v,
					"max_expansions": 50,
					"slop":           50,
				},
			},
		})
	}

	return res
}

func BuildMatch(field, value string) elm {
	if value == "" {
		return elm{
			"match_all": elm{},
		}
	}
	return elm{
		"match": elm{
			field: elm{
				"query":          value,
				"max_expansions": 50,
			},
		},
	}
}

func BuildMatchs(m map[string]string) []elm {
	var res []elm
	for k, v := range m {
		res = append(res, elm{
			"match": elm{
				k: elm{
					"query":          v,
					"max_expansions": 50,
				},
			},
		})
	}

	return res
}

func BuildRange(field string, gte *uint, lt *uint) elm {
	e := make(elm)
	var tmp *uint = nil

	if gte != tmp {
		e["gte"] = gte
	}
	if lt != tmp {
		e["lt"] = lt
	}

	if len(e) < 1 {
		return nil
	}
	return elm{
		"range": elm{
			field: e,
		},
	}
}

func BuildTimeRange(field string, gte time.Time, lt time.Time) elm {
	e := make(elm)

	if !gte.IsZero() {
		e["gte"] = gte
	}
	if !lt.IsZero() {
		e["lt"] = lt
	}

	if len(e) < 1 {
		return nil
	}
	return elm{
		"range": elm{
			field: e,
		},
	}
}

func BuildHighLight(fields []string) elm {
	elms := make(elm, 0)
	for _, v := range fields {
		elms[v] = elm{
			"require_field_match": false,
		}
	}
	return elm{
		"fields":              elms,
		"pre_tags":            []string{"<em>"},
		"post_tags":           []string{"</em>"},
		"number_of_fragments": 5,
		"fragment_size":       150,
		"no_match_size":       150,
		"order":               "score",
	}
}

func BuildNested(path string, query elm) elm {
	return elm{
		"path":  path,
		"query": query,
	}
}

/*
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "type": {
              "value": "playlist"
            }
          }
        },
        {
          "term": {
            "tags": {
              "value": "golang"
            }
          }
        },
        {
          "term": {
            "tags": {
              "value": "elasticsearch"
            }
          }
        },
        {
          "bool": {
            "should": [
              {
                "match_phrase_prefix": {
                  "title": {
                    "query": "资源管理 golang",
                    "max_expansions": 50,
                    "slop": 50
                  }
                }
              },
              {
                "match_phrase_prefix": {
                  "content": {
                    "query": "资源管理 golang",
                    "max_expansions": 50,
                    "slop": 50
                  }
                }
              }
            ]
          }
        },
        {
          "range": {
            "created_at": {
              "gte": "2022-01-01T00:00:00",   // 大于等于指定时间
              "lt": "2022-02-01T00:00:00"     // 小于指定时间
            }
          }
        }
      ]
    }
  },
  "highlight" : {
		"fields" : {
			"name" : { "number_of_fragments" : 0 }
		}
	},
	"size" : 25,
	"sort" : [ { "_score" : "desc" }, { "_doc" : "asc" } ]
}
*/
