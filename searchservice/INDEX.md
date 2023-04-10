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