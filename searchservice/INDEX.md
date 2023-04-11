[toc]

### 创建索引
```json
PUT /resource_userID
{
  "mappings": {
    "properties": {
      "id": {
        "type": "long"
      },
      "uuid": {
        "type": "keyword"
      },
      "pid": {
        "type": "long"
      },
      "user_id": {
        "type": "long"
      },
      "dir": {
        "type": "boolean"
      },
      "ext": {
        "type": "keyword"
      },
      "name": {
        "type": "text",
        "analyzer": "ik_max_word"
      },
      "size": {
        "type": "long"
      },
      "privacy": {
        "type": "boolean"
      },
      "path": {
        "type": "text"
      },
      "times": {
        "type": "long"
      },
      "create_at": {
        "type": "date"
      },
      "update_at": {
        "type": "date"
      },
      "deleted_at": {
        "type": "date"
      },
      "visit_time": {
        "type": "date"
      },
      "sha256": {
        "type": "keyword"
      }
    }
  }
}
```

```json
PUT /blank_userID
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
        "analyzer": "ik_max_word",
        "copy_to": ["all_text^2"]
      },
      "content": {
        "type": "text",
        "analyzer": "ik_max_word",
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
```

### 查询语句
```json
get /resource/_search
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
```