package es

var BlankIndex = `
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
		"all_text": {
		  "type": "text",
		  "analyzer": "ik_max_word"
		}
	  }
	}
  }
`

var ResourceIndex = `
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
		}
	  }
	}
  }
`
