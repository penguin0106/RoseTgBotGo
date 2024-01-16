package database

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

type MongoDB struct {
	Client *mongo.Client
	// Добавьте другие поля по необходимости
}

func NewMongoDB(uri string) (*MongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		Client: client,
	}, nil
}

func (db *MongoDB) Close() {
	err := db.Client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
