package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "newhope/app/struct"
)

//insert
func InsertAccount(document interface{}) error {
	err := collMap["account"].InsertOne(document)
	return err
}

//condition query
func QueryConditionAccount2json(filter interface{}) (string, error) {
	cursor, err := collMap["account"].Find(filter)
	if err != nil {
		return "[]", err
	}
	if err := cursor.Err(); err != nil {
		return "[]", err
	}
	var all = make([]interface{}, 0)
	for cursor.Next(context.Background()) {
		var account _struct.AccountInfo
		err = cursor.Decode(&account)
		if err == nil {
			all = append(all, &account)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//update
func UpdateAccount(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["account"].UpdateOne(filter, update, updateOpts)
	return err
}

//delete
func DeleteAccount(filter interface{}) error {
	err := collMap["account"].DeleteOne(filter)
	return err
}

func DeleteManyAccount(filter interface{}) error {
	err := collMap["account"].DeleteMany(filter)
	return err
}
