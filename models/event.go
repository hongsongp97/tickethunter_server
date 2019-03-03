package models

type Event struct {
	Id           uint   `bson:"_id,omitempty" json:"id,omitempty"`
	Title        string `bson:"title,omitempty" json:"title,omitempty"`
	Description  string `bson:"description" json:"description"`
	AvatarUrl    string `bson:"avatar_url" json:"avatar_url"`
	Location     string `bson:"location" json:"location"`
	Category     string `bson:"category" json:"category"`
	ImageUrl     string `bson:"image_url" json:"image_url"`
	OwnerId      uint   `bson:"owner_id" json:"owner_id"`
	Time         string `bson:"time,omitempty" json:"time,omitempty"`
	TicketNumber string `bson:"ticket_number" json:"ticket_number"`
	TicketPrice  string `bson:"ticket_price" json:"ticket_price"`
}