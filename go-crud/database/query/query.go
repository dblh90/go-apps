package query

import (
	"github.com/dblH90/go-crud/config"
	"github.com/dblH90/go-crud/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoAppDB struct {
	App config.GoAppTools
	DB  mongo.Client
}

func NewGoAppDB(app config.GoAppTools, db mongo.Client) database.DBRepo {
	return &GoAppDB{
		App: app,
		DB:  db,
	}
}

func (g *GoAppDB) InsertUser() {
	return
}
