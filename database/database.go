package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	connectTimeout           = 5
)

type Resource struct {
	DB *mongo.Database
}

func (r *Resource) Close() {
	logrus.Warning("Closing all db connections")
}

func CreateResource() (*Resource, error) {
	dbName := "db"
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		logrus.Errorf("Failed to create client: %v", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Errorf("Failed to connect to server: %v", err)
		return nil, err
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Errorf("Failed to ping cluster: %v", err)
		return nil, err
	}

	logrus.Info("Database connected ..")
	return &Resource{DB: client.Database(dbName)}, nil
}