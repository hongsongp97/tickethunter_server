package dao

import (
	"context"
	"fmt" // . "github.com/hongsongp97/tickethunter_server/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
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
// func (userDAO *UserDAO) FindAll() ([]User, error) {
// 	var users []User

// 	cur, err := userDAO.Collection.Find(context.Background(), bson.D{})

// 	return users, err
// }

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
