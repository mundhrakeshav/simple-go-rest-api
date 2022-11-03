package db

import (
	"context"
	"errors"
	logger "mundhrakeshav/go-http/pkg/log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBStoreInterface interface {
	GetDb() *mongo.Database
	GetClient() (*mongo.Client, error)
	Coll(name string, opts ...*options.CollectionOptions) *mongo.Collection
	Disconnect() error
}
type DBStoreImpl struct {
	user_db *mongo.Database
	client  *mongo.Client
}

var (
	DBStore DBStoreImpl
	ctx     context.Context
)

func init() {
	godotenv.Load(".env")
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		AuthSource: "admin",
		Username:   os.Getenv("MONGO_USERNAME"),
		Password:   os.Getenv("MONGO_PASSWORD"),
	})

	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		logger.Log.Fatal("error while connecting with mongo")
	}

	err = mongoclient.Ping(ctx, readpref.Primary())

	if err != nil {
		logger.Log.Fatal("error while trying to ping mongo")
	}
	user_db := mongoclient.Database("userdb")
	DBStore = DBStoreImpl{
		client:  mongoclient,
		user_db: user_db,
	}
	logger.Log.Info("DB Initialized")
}

func (md *DBStoreImpl) GetUserDb() *mongo.Database {
	return md.user_db
}

func (md *DBStoreImpl) GetClient() (*mongo.Client, error) {
	if md.client != nil {
		return md.client, nil
	}
	return nil, errors.New("client is missing (nil) in Mongo Data Store")
}

func (md *DBStoreImpl) GetUsersCollection() *mongo.Collection {
	return md.user_db.Collection("users")
}

func (md *DBStoreImpl) Disconnect() error {
	err := md.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}
