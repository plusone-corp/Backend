package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"plusone/backend/types"
)

func GetEventID(id primitive.ObjectID) (*types.ResEvent, bool, error) {
	var res *types.Event

	err := EventsCollection.FindOne(Context, bson.D{{"_id", id}}).Decode(&res)
	if err == mongo.ErrNoDocuments || res == nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	event, found, err := validateEvent(res)
	if !found && err == nil {
		return nil, false, nil
	} else if !found && err != nil {
		return nil, false, err
	}

	return event, true, nil
}

func CreateEvent(event types.Event) (*types.Event, error) {
	_, err := EventsCollection.InsertOne(Context, event)
	if err != nil {
		return nil, err
	}

	_, err = UserCollection.UpdateOne(Context, bson.D{{"_id", event.Author}}, bson.D{{"$push", bson.D{{"events", event.Id}}}})

	return &event, nil
}

func SensoreUser(userData types.User) types.UserFiltered {
	user := types.UserFiltered{Username: userData.Username, Avatar: userData.Avatar, DisplayName: userData.DisplayName, Description: userData.Description, Level: userData.Level}
	return user
}

func validateEvent(res *types.Event) (*types.ResEvent, bool, error) {
	author, found, err := GetUserByID(res.Author)
	if !found && err == nil {
		return nil, false, nil
	} else if !found && err != nil {
		return nil, false, err
	}

	inviters, found, err := GetManyUserID(res.Invites)
	if !found && err != nil {
		return nil, false, err
	}

	posts, found, err := GetManyPostsID(res.Posts)
	if !found && err != nil {
		return nil, false, err
	}

	post := &types.ResEvent{
		Id:          res.Id,
		Description: res.Description,
		Posts:       *posts,
		Title:       res.Title,
		Author:      SensoreUser(*author),
		Invites:     *inviters,
		Reactions:   res.Reactions,
		Comments:    res.Comments,
		CreatedAt:   res.CreatedAt,
	}

	return post, true, nil
}

func GetManyEventsID(ids []primitive.ObjectID) (*[]types.Event, bool, error) {
	var posts []types.Event
	cursor, err := EventsCollection.Find(Context, bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return nil, false, err
	}

	if err = cursor.All(Context, &posts); err != nil {
		return nil, false, err
	}
	return &posts, true, nil
}
