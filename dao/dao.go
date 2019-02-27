package dao

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type DAO struct {
	Server         string
	Database       string
	Client         *mongo.Client
	DatabaseHandle *mongo.Database
}

func (dao *DAO) ConnectToDB() {
	var err error

	dao.Client, err = mongo.Connect(context.TODO(), dao.Server)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = dao.Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get *DatabaseHandle
	dao.DatabaseHandle = dao.Client.Database(dao.Database)

	fmt.Printf("Connected to \"%s\" database.\n", dao.Database)
}

func (dao *DAO) Disconnect() {
	err := dao.Client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
