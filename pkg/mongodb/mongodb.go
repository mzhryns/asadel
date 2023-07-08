package mongodb

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
}

func Connect(dbHost, dbPort, dbUsername, dbPass string, ctx context.Context) (*mongo.Client, error) {
	db := &mongodb{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUsername: dbUsername,
		DBPassword: dbPass,
	}

	return connect(db, ctx)
}

func connect(param *mongodb, ctx context.Context) (*mongo.Client, error) {
	uri := buildURI(param)
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func buildURI(param *mongodb) string {
	username := url.QueryEscape(param.DBUsername)
	password := url.QueryEscape(param.DBPassword)

	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		username,
		password,
		param.DBHost,
		param.DBPort,
	)
}
