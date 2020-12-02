package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (coll *Coll) InsertOne(document interface{}) error {
	_, err := coll.collection.InsertOne(context.Background(), document)
	return err
}

func (coll *Coll) Find(filter interface{}) (*mongo.Cursor, error) {
	cursor, err := coll.collection.Find(context.Background(), filter)
	return cursor, err
}

func (coll *Coll) DeleteOne(filter interface{}) error {
	_, err := coll.collection.DeleteOne(context.Background(), filter)
	return err
}

func (coll *Coll) DeleteMany(filter interface{}) error {
	_, err := coll.collection.DeleteMany(context.Background(), filter)
	return err
}

func (coll *Coll) UpdateOne(filter interface{}, update interface{}, opts *options.UpdateOptions) error {
	_, err := coll.collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}

func (coll *Coll) CountDocuments(filter interface{}) (int64, error) {
	return coll.collection.CountDocuments(context.Background(), filter)
}

func (coll *Coll) FindPaging(ctx context.Context, filter interface{}, findoptions *options.FindOptions) (*mongo.Cursor, error) {
	cursor, err := coll.collection.Find(ctx, filter, findoptions)
	return cursor, err
}
