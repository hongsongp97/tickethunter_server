package dao

import (
	"context"
	"errors"
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
	log.SetFlags(log.Lshortfile)

	userDAO.Collection = userDAO.Dao.DatabaseHandle.Collection(COLLECTION)
}

// Find list of users
func (userDAO *UserDAO) FindAll() ([]User, error) {
	var users []User
	ctx := context.Background()

	cur, err := userDAO.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cur.Close(context.Background())

	for cur.Next(ctx) {
		user := &User{}
		if err := cur.Decode(user); err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, *user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}

// Find a user by its id
func (userDAO *UserDAO) FindById(id string) (User, error) {
	var user User
	ctx := context.Background()

	cur, err := userDAO.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	defer cur.Close(context.Background())

	if cur.Next(ctx) {
		if err := cur.Decode(&user); err != nil {
			log.Fatal(err)
			return user, err
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return user, err
	}

	return user, nil
}

// Insert a user into database
func (userDAO *UserDAO) Insert(user User) error {
	res, err := userDAO.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	id := res.InsertedID
	fmt.Printf("id: %s\n", id)
	return err
}

// Delete an existing user
func (userDAO *UserDAO) Delete(userId string) error {
	ctx := context.Background()
	res, err := userDAO.Collection.DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		log.Fatal(err)
	}
	if (*res).DeletedCount == 0 {
		err = errors.New("No matched record!")
	}
	return err
}

// Update an existing user
// func (ud *UserDAO) Update(user User) error {
// err := db.C(COLLECTION).UpdateId(user.ID, &user)
// return err
// }
