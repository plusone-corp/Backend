package database

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"plusone/backend/types"
)

func GetPostID(id primitive.ObjectID) (*types.ResPost, bool, error) {
	var query *types.Post

	filter := bson.D{{
		Key:   "_id",
		Value: id,
	}}

	err := PostCollection.FindOne(Context, filter).Decode(&query)
	if err == mongo.ErrNoDocuments || query == nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	post, found, err := validateUser(*query)
	if !found && err == nil {
		return nil, false, errors.New(fmt.Sprintf("User with ID %v not found", query.Id))
	} else if !found && err != nil {
		return nil, false, err
	}

	return post, true, nil
}

func CreatePost(post types.Post) (*types.Post, bool, error) {
	_, err := PostCollection.InsertOne(Context, post)
	if err != nil {
		return nil, false, err
	}
	_, err = EventsCollection.UpdateOne(Context, bson.D{{"_id", post.Event}}, bson.D{{"$push", bson.D{{"posts", post.Id}}}})
	return &post, true, nil
}

func GetLastestPost(userId primitive.ObjectID) (*types.Post, bool, error) {

	var posts []*types.Post

	filter := bson.D{{
		Key:   "author",
		Value: userId,
	}}

	opt := options.Find().SetSort(bson.D{{"createdAt", 1}})
	cursor, err := PostCollection.Find(Context, filter, opt)
	if err != nil {
		return nil, false, err
	}

	if err = cursor.All(Context, &posts); err != nil {
		return nil, false, err
	}

	return posts[len(posts)-1], true, nil
}

func GetManyPostsID(ids []primitive.ObjectID) (*[]types.Post, bool, error) {
	var posts []types.Post
	cursor, err := PostCollection.Find(Context, bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return nil, false, err
	}

	if err = cursor.All(Context, &posts); err != nil {
		return nil, false, err
	}
	return &posts, true, nil
}

func GetAllPost(userId primitive.ObjectID) (*[]types.Post, bool, error) {
	var posts []types.Post

	filter := bson.D{{
		"author",
		userId,
	}}

	opt := options.Find().SetSort(bson.D{{"createdAt", 1}})
	cursor, err := PostCollection.Find(Context, filter, opt)
	if err != nil {
		return nil, false, err
	}

	if err = cursor.All(Context, &posts); err != nil {
		return nil, false, err
	}

	return &posts, true, nil
}

func validateUser(res types.Post) (*types.ResPost, bool, error) {
	user, found, err := GetUserByID(res.Author)
	if !found && err == nil {
		return nil, false, nil
	} else if !found && err != nil {
		return nil, false, err
	}

	post := &types.ResPost{
		Id:          res.Id,
		Description: res.Description,
		Event:       res.Event,
		Title:       res.Title,
		Author:      SensoreUser(*user),
		Image:       res.Image,
		Reactions:   res.Reactions,
		Comments:    res.Comments,
		CreatedAt:   res.CreatedAt,
	}

	return post, true, nil
}
