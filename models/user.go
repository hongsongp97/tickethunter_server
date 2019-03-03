package models

type User struct {
	Id           string `bson:"_id,omitempty" json:"id"`
	FirstName    string `bson:"first_name" json:"first_name"`
	LastName     string `bson:"last_name" json:"last_name"`
	UserName     string `bson:"user_name" json:"user_name"`
	EmailAddress string `bson:"email" json:"email"`
	Phone        string `bson:"phone" json:"phone"`
	Address      string `bson:"address" json:"address"`
	AvatarURL    string `bson:"avatar_url" json:"avatar_url"`
	Description  string `bson:"description" json:"description"`
}
