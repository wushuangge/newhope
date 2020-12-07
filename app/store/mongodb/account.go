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
func QueryConditionAccount(filter interface{}) (_struct.AccountInfo, error) {
	var account _struct.AccountInfo
	cursor, err := collMap["account"].Find(filter)
	if err != nil {
		return account, err
	}
	if err := cursor.Err(); err != nil {
		return account, err
	}
	for cursor.Next(context.Background()) {
		err = cursor.Decode(&account)
		if err == nil {
			break
		}
	}
	cursor.Close(context.Background())
	return account, nil
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
