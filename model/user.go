package model

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	FirstName    string        `bson:"name" json:"first_name"`
	LastName     string        `bson:"name" json:"last_name"`
	UserName     string        `bson:"name" json:"user_name"`
	Password     string        `bson:"name" json:"password"`
	EmailAddress string        `bson:"name" json:"email"`
	Phone        string        `bson:"name" json:"phone"`
	Address      string        `bson:"name" json:"address"`
	AvatarURL    string        `bson:"cover_image" json:"avatar_url"`
	Description  string        `bson:"description" json:"description"`
}
