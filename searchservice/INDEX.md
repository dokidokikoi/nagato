[toc]

### 创建索引
```json
PUT /resource
{
  "mappings": {
    "properties": {
      "id": {
        "type": "long"
      },
      "uuid": {
        "type": "keyword"
      },
      "puuid": {
        "type": "keyword"
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
        "analyzer": "ik_max_word",
        "copy_to": ["all_text"]
      },
      "size": {
        "type": "long"
      },
      "privacy": {
        "type": "boolean"
      },
      "path": {
        "type": "text",
        "analyzer": "ik_max_word",
        "copy_to": ["all_text"]
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
      },
      "all_text": {
        "type": "text",
        "analyzer": "ik_max_word"
      },
      "children": {
        "type": "nested",
        "properties": {
          "id": {
            "type": "long"
          },
          "uuid": {
            "type": "keyword"
          },
          "puuid": {
            "type": "keyword"
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
            "analyzer": "ik_max_word",
            "copy_to": ["all_text"]
          },
          "size": {
            "type": "long"
          },
          "privacy": {
            "type": "boolean"
          },
          "path": {
            "type": "text",
            "analyzer": "ik_max_word",
            "copy_to": ["all_text"]
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
          },
          "children": {
            "type": "object"
          },
          "all_text": {
            "type": "text",
            "analyzer": "ik_max_word"
          }
        }
      }
    }
  }
}
```

```json
PUT /blank
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
        "copy_to": ["all_text^10"]
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
      "matters": {
        "type": "nested",
        "properties": {
          "id": {
            "type": "long"
          },
          "uuid": {
            "type": "keyword"
          },
          "puuid": {
            "type": "keyword"
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
            "analyzer": "ik_max_word",
            "copy_to": ["all_text"]
          },
          "size": {
            "type": "long"
          },
          "privacy": {
            "type": "boolean"
          },
          "path": {
            "type": "text",
            "analyzer": "ik_max_word",
            "copy_to": ["all_text"]
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
          },
          "children": {
            "type": "object"
          },
          "all_text": {
            "type": "text",
            "analyzer": "ik_max_word"
          }
        }
      },
      "all_text": {
        "type": "text",
        "analyzer": "ik_max_word"
      }
    }
  }
}
```

### 查询语句
```json
get /blank/_search
{
  "query": {
    "bool": {
      "should": [
        {
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
        {
          "nested": {
            "path": "matters",
            "query": {
              "bool": {
                "must": [
                  {
                    "term": {
                      "ext": {
                        "value": ".zip"
                      }
                    }
                  },
                  {
                    "term": {
                      "dir": {
                        "value": false
                      }
                    }
                  },
                  {
                    "term": {
                      "privacy": {
                        "value": false
                      }
                    }
                  },
                  {
                    "match": {
                      "all_text": {
                        "query": "乔碧萝"
                      }
                    }
                  },
                  {
                    "term": {
                      "sha256": {
                        "value": "fsfsddfsf"
                      }
                    }
                  },
                  {
                    "range": {
                      "size": {
                        "gte": "23565",   // 大于等于指定大小
                        "lt": "786543"     // 小于指定大小
                      }
                    }
                  }
                ]
              }
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

```json
get /resource/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "bool": {
            "must": [
              {
                "term": {
                  "ext": {
                    "value": ".zip"
                  }
                }
              },
              {
                "term": {
                  "dir": {
                    "value": false
                  }
                }
              },
              {
                "term": {
                  "privacy": {
                    "value": false
                  }
                }
              },
              {
                "match": {
                  "all_text": {
                    "query": "乔碧萝"
                  }
                }
              },
              {
                "term": {
                  "sha256": {
                    "value": "fsfsddfsf"
                  }
                }
              },
              {
                "range": {
                  "size": {
                    "gte": "23565",   // 大于等于指定大小
                    "lt": "786543"     // 小于指定大小
                  }
                }
              }
            ]
          }
        },
        {
          "nested": {
            "path": "children",
            "query": {
              "bool": {
                "must": [
                  {
                    "term": {
                      "ext": {
                        "value": ".zip"
                      }
                    }
                  },
                  {
                    "term": {
                      "dir": {
                        "value": false
                      }
                    }
                  },
                  {
                    "term": {
                      "privacy": {
                        "value": false
                      }
                    }
                  },
                  {
                    "match": {
                      "all_text": {
                        "query": "乔碧萝"
                      }
                    }
                  },
                  {
                    "term": {
                      "sha256": {
                        "value": "fsfsddfsf"
                      }
                    }
                  },
                  {
                    "range": {
                      "size": {
                        "gte": "23565",   // 大于等于指定大小
                        "lt": "786543"     // 小于指定大小
                      }
                    }
                  }
                ]
              }
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