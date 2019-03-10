package model

type Comment struct {
	Id        string `bson:"_id,omitempty" json:"id,omitempty"`
	OwnerId   string `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	EventId   string `bson:"event_id,omitempty" json:"event_id,omitempty"`
	Content   string `bson:"content,omitempty" json:"content,omitempty"`
	Timestamp string `bson:"time_stamp" json:"time_stamp"`
}
