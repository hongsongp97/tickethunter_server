package dao

import (
	"context"
	"errors"
	"fmt"
	"log"

	. "github.com/hongsongp97/tickethunter_server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao struct {
	Dao        *Dao
	Collection *mongo.Collection
}

const (
	USER_COLLECTION = "user"
)

func (userDao *UserDao) Init(dao *Dao) {
	log.SetFlags(log.Lshortfile)

	userDao.Dao = dao
	userDao.Collection = userDao.Dao.DatabaseHandle.Collection(USER_COLLECTION)
}

// Find list of users
func (userDao *UserDao) FindAll() ([]User, error) {
	var users []User
	ctx := context.Background()

	cur, err := userDao.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer cur.Close(context.Background())

	for cur.Next(ctx) {
		user := &User{}
		if err := cur.Decode(user); err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, *user)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

// Find a user by its id
func (userDao *UserDao) FindById(id string) (User, error) {
	var user User
	ctx := context.Background()

	cur, err := userDao.Collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		return user, err
	}

	defer cur.Close(context.Background())

	if cur.Next(ctx) {
		if err := cur.Decode(&user); err != nil {
			log.Println(err)
			return user, err
		}
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

// Insert a user into database
func (userDao *UserDao) Insert(user User) error {
	res, err := userDao.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err)
		return err
	}
	id := res.InsertedID
	fmt.Printf("id: %s\n", id)
	return err
}

// Delete an existing user
func (userDao *UserDao) Delete(userId string) error {
	ctx := context.Background()
	res, err := userDao.Collection.DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		log.Println(err)
	}
	if (*res).DeletedCount == 0 {
		err = errors.New("No matched record!")
	}
	return err
}

// Update an existing user
func (userDao *UserDao) Update(user User) error {
	ctx := context.Background()

	data := bson.D{{"$set", user}}

	log.Println(bson.M{"_id": user.Id})
	res, err := userDao.Collection.UpdateOne(ctx, bson.M{"_id": user.Id}, data)
	if err != nil {
		log.Println(err)
	}
	if (*res).MatchedCount == 0 {
		err = errors.New("No matched record!")
	}
	return err
}
