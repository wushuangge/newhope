package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "newhope/app/struct"
)

//insert
func InsertTerms(document interface{}) error {
	err := collMap["terms"].InsertOne(document)
	return err
}

//query all
func QueryAllTerms() (string, error) {
	cursor, err := collMap["terms"].Find(bson.D{})
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var entryInfo _struct.TermsInfo
		err = cursor.Decode(&entryInfo)
		if err == nil {
			all = append(all, &entryInfo)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//condition query
func QueryConditionTerms(filter interface{}) (string, error) {
	cursor, err := collMap["terms"].Find(filter)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var entryInfo _struct.TermsInfo
		err = cursor.Decode(&entryInfo)
		if err == nil {
			all = append(all, &entryInfo)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//update
func UpdateTerms(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["terms"].UpdateOne(filter, update, updateOpts)
	return err
}

//delete
func DeleteTerms(filter interface{}) error {
	err := collMap["terms"].DeleteOne(filter)
	return err
}
