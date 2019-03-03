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

type EventDao struct {
	Dao        *Dao
	Collection *mongo.Collection
}

const (
	EVENT_COLLECTION = "event"
)

func (eventDao *EventDao) Init(dao *Dao) {
	log.SetFlags(log.Lshortfile)

	eventDao.Dao = dao
	eventDao.Collection = eventDao.Dao.DatabaseHandle.Collection(EVENT_COLLECTION)
}

// Find list of events
func (eventDao *EventDao) FindAll() ([]Event, error) {
	var events []Event
	ctx := context.Background()

	cur, err := eventDao.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer cur.Close(context.Background())

	for cur.Next(ctx) {
		event := &Event{}
		if err := cur.Decode(event); err != nil {
			log.Println(err)
			return nil, err
		}
		events = append(events, *event)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return events, nil
}

// Find a user by its id
func (eventDao *EventDao) FindById(id string) (Event, error) {
	var event Event
	ctx := context.Background()

	cur, err := eventDao.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return event, err
	}

	defer cur.Close(context.Background())

	if cur.Next(ctx) {
		if err := cur.Decode(&event); err != nil {
			log.Println(err)
			return event, err
		}
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return event, err
	}

	return event, nil
}

// Insert a user into database
func (eventDao *EventDao) Insert(event Event) error {
	res, err := eventDao.Collection.InsertOne(context.Background(), event)
	if err != nil {
		log.Println(err)
		return err
	}
	id := res.InsertedID
	fmt.Printf("id: %s\n", id)
	return err
}

// Delete an existing user
func (eventDao *EventDao) Delete(eventId string) error {
	ctx := context.Background()
	res, err := eventDao.Collection.DeleteOne(ctx, bson.M{"_id": eventId})
	if err != nil {
		log.Println(err)
	}
	if (*res).DeletedCount == 0 {
		err = errors.New("No matched record!")
	}
	return err
}

// Update an existing user
func (eventDao *EventDao) Update(event Event) error {
	ctx := context.Background()

	data := bson.D{{"$set", event}}

	log.Println(bson.M{"_id": event.Id})
	res, err := eventDao.Collection.UpdateOne(ctx, bson.M{"_id": event.Id}, data)
	if err != nil {
		log.Println(err)
	}
	if (*res).MatchedCount == 0 {
		err = errors.New("No matched record!")
	}
	return err
}
