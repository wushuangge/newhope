package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "newhope/app/struct"
)

//insert
func InsertRecord(document interface{}) error {
	err := collMap["record"].InsertOne(document)
	return err
}

//query all
func QueryAllRecord() (string, error) {
	cursor, err := collMap["record"].Find(bson.D{})
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
func QueryConditionRecord(filter interface{}) (string, error) {
	cursor, err := collMap["record"].Find(filter)
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
func UpdateRecord(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["record"].UpdateOne(filter, update, updateOpts)
	return err
}

//delete
func DeleteRecord(filter interface{}) error {
	err := collMap["record"].DeleteOne(filter)
	return err
}
