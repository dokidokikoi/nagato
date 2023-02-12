package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Store struct {
	db *mongo.Database
}

func (d *Store) Logs() *logs {
	return newLogs(d)
}
