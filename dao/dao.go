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

func (dao *DAO) Connect() {
	client, err := mongo.Connect(context.TODO(), dao.Server)
	dao.Client = client

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = dao.Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func (dao *DAO) Disconnect() {
	err := dao.Client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func (dao *DAO) GetDatabase() {

	dao.DatabaseHandle = dao.Client.Database(dao.Database)

	fmt.Printf("Connected to \"%s\" database.\n", dao.Database)
}
