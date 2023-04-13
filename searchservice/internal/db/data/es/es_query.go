package es

type elm map[string]interface{}

func BulidQuery(esBool, highlight elm) elm {
	return elm{
		"query":     esBool,
		"highlight": highlight,
		"size":      25,
		"sort": []interface{}{
			elm{"_score": "desc"},
			elm{"_doc": "asc"},
		},
	}
}

func BuildBool(must, should []elm) elm {
	return elm{
		"bool": elm{
			"must":   must,
			"should": should,
		},
	}
}

func BulidMust(element []elm) elm {
	return elm{
		"must": element,
	}
}

func BulidShould(element []elm) elm {
	return elm{
		"should": element,
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
	return elm{
		"match": elm{
			field: elm{
				"query":          value,
				"max_expansions": 50,
				"slop":           50,
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

func BuildRange(field string, gte interface{}, lt interface{}) elm {
	e := make(elm)
	if gte != nil {
		e["gte"] = gte
	}
	if lt != nil {
		e["lt"] = lt
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
