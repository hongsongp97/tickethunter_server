package model

type Event struct {
	Id            string   `bson:"_id,omitempty" json:"id,omitempty"`
	Title         string   `bson:"title" json:"title"`
	Description   string   `bson:"description" json:"description"`
	AvatarUrl     string   `bson:"avatar_url" json:"avatar_url"`
	Location      string   `bson:"location" json:"location"`
	Email         string   `bson:"email" json:"email"`
	Category      string   `bson:"category" json:"category"`
	ImageUrl      []string `bson:"image_url" json:"image_url"`
	OwnerId       string   `bson:"own_id" json:"own_id"`
	Time          string   `bson:"time" json:"time"`
	TicketNumber  uint     `bson:"ticket_number" json:"ticket_number"`
	TicketPrice   uint     `bson:"ticket_price" json:"ticket_price"`
	Phone         string   `bson:"phone" json:"phone"`
	JoinedUsers   []string `bson:"joined_users" json:"joined_users"`
	FollowedUsers []string `bson:"followed_users" json:"followed_users"`
}

type UserByEachEvent struct {
	Id string `bson:"_id,omitempty" json:"id,omitempty"`
}
