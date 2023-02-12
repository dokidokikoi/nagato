package mongo

import "go.mongodb.org/mongo-driver/mongo"

type logs struct {
	db *mongo.Database
}

func newLogs(d *Store) *logs {
	return &logs{db: d.db}
}
