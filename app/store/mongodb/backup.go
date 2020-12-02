package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "newhope/app/struct"
)

//insert
func InsertBackup(document interface{}) error {
	err := collMap["backup"].InsertOne(document)
	return err
}

//query all
func QueryAllBackup() (string, error) {
	cursor, err := collMap["backup"].Find(bson.D{})
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var entryInfo _struct.EntryInfo
		err = cursor.Decode(&entryInfo)
		if err == nil {
			all = append(all, &entryInfo)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//condition query
func QueryConditionBackup(filter interface{}) (string, error) {
	cursor, err := collMap["backup"].Find(filter)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var entryInfo _struct.EntryInfo
		err = cursor.Decode(&entryInfo)
		if err == nil {
			all = append(all, &entryInfo)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//update
func UpdateBackup(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["backup"].UpdateOne(filter, update, updateOpts)
	return err
}
