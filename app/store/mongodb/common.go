package mongodb

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB : class of Mongo
type MongoDB struct {
	username string
	password string
	ip       string
	port     string
	client   *mongo.Client
}

// Coll : class of collection
type Coll struct {
	collection *mongo.Collection
}

// DB : class of database
type DB struct {
	name     string
	database *mongo.Database
}

func Interface2json(i interface{}) string {
	jsons, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(jsons)
}

func Interfaces2json(i interface{}) string {
	var str = ""
	for _, one := range i.([]interface{}) {
		str += "," + Interface2json(one)
	}
	if len(str) == 0 {
		return "[]"
	}
	return "[" + str[1:] + "]"
}
