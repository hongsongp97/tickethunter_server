package dao

import (
	"context"
	"fmt"
	"log"

	. "github.com/hongsongp97/tickethunter_server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDAO struct {
	Dao        *DAO
	Collection *mongo.Collection
}

const (
	COLLECTION = "user"
)

func (userDAO *UserDAO) Init() {
	userDAO.Collection = userDAO.Dao.DatabaseHandle.Collection(COLLECTION)
}

// Find list of users
func (userDAO *UserDAO) FindAll() ([]User, error) {
	var users []User
	ctx := context.Background()

	cur, err := userDAO.Collection.Find(ctx, bson.D{})
	if err != nil {
		// log.Fatal("fsnfjksdf")
		return users, err
	}

	defer cur.Close(context.Background())

	for cur.Next(ctx) {
		elem := &bson.D{}
		if err := cur.Decode(elem); err != nil {
			log.Fatal(err)
			return users, err
		}

		log.Println(elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return users, err
	}

	return users, err
}

// Find a user by its id
// func (ud *UserDAO) FindById(id string) (User, error) {
// var user User
// err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
// return user, err
// }

// Insert a user into database
func (userDAO *UserDAO) Insert() error {
	res, err := userDAO.Collection.InsertOne(context.Background(), bson.M{"name": "Son"})
	if err != nil {
		return err
	}
	id := res.InsertedID
	fmt.Printf("id: %s", id)
	return err
}

// Delete an existing user
// func (ud *UserDAO) Delete(user User) error {
// err := db.C(COLLECTION).Remove(&user)
// return err
// }

// Update an existing user
// func (ud *UserDAO) Update(user User) error {
// err := db.C(COLLECTION).UpdateId(user.ID, &user)
// return err
// }
