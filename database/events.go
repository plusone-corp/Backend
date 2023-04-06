package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetLatestEvent(userId primitive.ObjectID) (*types.Event, bool, error) {

	var post *types.Event

	filter := bson.D{{
		Key:   "author",
		Value: userId,
	}}

	opt := options.FindOne().SetSort(bson.D{{"createdAt", -1}})
	err := EventsCollection.FindOne(Context, filter, opt).Decode(&post)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return post, true, nil
}

func GetAllEvent(userId primitive.ObjectID) (*[]types.Event, error) {
	var events []types.Event

	filter := bson.D{{
		"author",
		userId,
	}}

	opt := options.Find().SetSort(bson.D{{"createdAt", 1}})
	cursor, err := PostCollection.Find(Context, filter, opt)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(Context, &events); err != nil {
		return nil, err
	}

	return &events, nil
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
