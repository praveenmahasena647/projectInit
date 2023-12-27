package dbs

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Uri          string
	DBName       string
	DBCollection *mongo.Collection
	connection   *mongo.Client
)

func Connect() error {
	var err error
	var clientOptions = options.Client().ApplyURI(Uri)
	connection, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	DBCollection = connection.Database(DBName).Collection("")
	return nil
}
