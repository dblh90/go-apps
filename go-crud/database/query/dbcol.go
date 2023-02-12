package query

import "go.mongodb.org/mongo-driver/mongo"

func User(db mongo.Client, collection string) mongo.Collection {
	var userCol = db.Database("go-crud").Collection(collection)
	return *userCol
}
