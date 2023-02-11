<p align="center"><img width="180" src="./logo.png" alt="logo"></p>
<h1 align="center">nagato</h1>
<h4 align="center">Data Integration Thought Entity</h4>



```json
PUT /resource
{
  "mappings": {
    "properties": {
      "uuid": {
        "type": "keyword"
      },
      "puuid": {
        "type": "keyword"
      },
      "user_uuid": {
        "type": "keyword"
      },
      "username": {
        "type": "text",
        "analyzer": "ik_max_word"
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
      "tags": {
        "type": "keyword"
      },
      "priveiws": {
        "type": "keyword",
        "index": false
      },
      "size": {
        "type": "long"
      },
      "privacy": {
        "type": "boolean"
      },
      "path": {
        "type": "text",
        "analyzer": "ik_max_word"
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