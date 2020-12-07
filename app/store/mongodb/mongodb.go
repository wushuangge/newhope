package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"newhope/config"
)

var collMap map[string]*Coll

func (db *DB) GetCollByName(collName string) *Coll {
	coll := new(Coll)
	coll.collection = db.database.Collection(collName)
	return coll
}

func (db *DB) DeleteColl(coll string) error {
	return db.GetCollByName(coll).collection.Drop(context.Background())
}

func (mongoDB *MongoDB) SetLinkParameters(username string, password string, ip string, port string) {
	mongoDB.username = username
	mongoDB.password = password
	mongoDB.ip = ip
	mongoDB.port = port
}

func (mongoDB *MongoDB) getDBUri() string {
	return "mongodb://" + mongoDB.username + ":" + mongoDB.password + "@" + mongoDB.ip + ":" + mongoDB.port
}

func (mongoDB *MongoDB) GetDBByName(name string) *DB {
	db := new(DB)
	db.name = name
	db.database = mongoDB.client.Database(name)
	return db
}

func ConnMongoDB() *MongoDB {
	mongoDB := new(MongoDB)
	mongoDB.SetLinkParameters(config.GetMongoUserName(), config.GetMongoPassword(), config.GetMongoIp(), config.GetMongoPort())
	mongoDB.client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoDB.getDBUri()))
	return mongoDB
}

func (mongoDB *MongoDB) Ping() error {
	return mongoDB.client.Ping(context.Background(), readpref.Primary())
}

func InitMongoDB() error {
	mongoDB := ConnMongoDB()
	err := mongoDB.Ping()
	if err != nil {
		return err
	}
	task := mongoDB.GetDBByName("newhope")
	collMap = make(map[string]*Coll)
	collMap["record"] = task.GetCollByName("record")
	collMap["account"] = task.GetCollByName("account")
	collMap["backup"] = task.GetCollByName("backup")
	collMap["terms"] = task.GetCollByName("terms")
	return nil
}

//创建索引
func createIndex(key string, coll *Coll) error {
	idx := mongo.IndexModel{
		Keys:    bsonx.Doc{{key, bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
	_, err := coll.collection.Indexes().CreateOne(context.Background(), idx)
	return err
}
