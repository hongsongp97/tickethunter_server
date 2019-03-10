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

func (eventDao *EventDao) FindByJoinedUser(id string) ([]UserByEachEvent, error) {
	var (
		userByEachEvent []UserByEachEvent
		joinedUsers     []string
		event           Event
	)
	ctx := context.Background()

	cur, err := eventDao.Collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		log.Print("Event dao - FindByJoinedUser: " + err.Error())
		return userByEachEvent, err
	}
	defer cur.Close(context.Background())

	if cur.Next(ctx) {
		if err := cur.Decode(&event); err != nil {
			log.Print("Event dao - FindByJoinedUser: " + err.Error())
			return userByEachEvent, err
		}
	}

	if err := cur.Err(); err != nil {
		log.Print("Event dao - FindByJoinedUser: " + err.Error())
		return userByEachEvent, err
	}

	joinedUsers = event.JoinedUsers
	for _, joinedUser := range joinedUsers {
		var id UserByEachEvent
		log.Println("joinedUser: " + (string)(joinedUser))
		id.Id = (string)(joinedUser)
		userByEachEvent = append(userByEachEvent, id)
	}

	return userByEachEvent, err
}

func (eventDao *EventDao) FindByFollowedUser(id string) ([]UserByEachEvent, error) {
	var (
		userByEachEvent []UserByEachEvent
		followedUsers   []string
		event           Event
	)
	ctx := context.Background()

	cur, err := eventDao.Collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		log.Print("Event dao - FindByFollowedUser: " + err.Error())
		return userByEachEvent, err
	}
	defer cur.Close(context.Background())

	if cur.Next(ctx) {
		if err := cur.Decode(&event); err != nil {
			log.Print("Event dao - FindByFollowedUser: " + err.Error())
			return userByEachEvent, err
		}
	}

	if err := cur.Err(); err != nil {
		log.Print("Event dao - FindByFollowedUser: " + err.Error())
		return userByEachEvent, err
	}

	followedUsers = event.FollowedUsers
	for _, followedUser := range followedUsers {
		var id UserByEachEvent
		log.Println("joinedUser: " + (string)(followedUser))
		id.Id = (string)(followedUser)
		userByEachEvent = append(userByEachEvent, id)
	}

	return userByEachEvent, err
}

func (eventDao *EventDao) FindByCategoryId(key string) ([]Event, error) {
	var (
		events []Event
		event  Event
	)

	ctx := context.Background()

	cur, err := eventDao.Collection.Find(ctx, bson.D{{"category", key}})

	if err != nil {
		log.Print("Event dao - FindByCategoryId: " + err.Error())
		return events, err
	}
	defer cur.Close(context.Background())

	for cur.Next(ctx) {
		if err := cur.Decode(&event); err != nil {
			log.Print("Event dao - FindByCategoryId: " + err.Error())
			return events, err
		}
		events = append(events, event)
	}

	if err := cur.Err(); err != nil {
		log.Print("Event dao - FindByCategoryId: " + err.Error())
		return events, err
	}

	return events, err
}

// Insert a event into database
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

// Delete an existing event
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

// Update an existing event
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
