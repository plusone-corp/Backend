package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"plusone/backend/types"
)

func GetPostID(id primitive.ObjectID) (*types.Post, bool, error) {
	var user *types.User

	filter := bson.D{{
		Key: "posts",
		Value: bson.D{{
			"$elemMatch", bson.D{{
				"id", id,
			}},
		}},
	}}

	err := UserCollection.FindOne(Context, filter).Decode(&user)
	if err == mongo.ErrNoDocuments || user == nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	var post *types.Post

	for _, v := range user.Posts {
		if v.Id == id {
			post = &v
		}
	}

	return post, true, nil
}

func CreatePost(userId primitive.ObjectID, post types.Post) (*types.Post, bool, error) {
	user, found, err := GetByID(userId)
	if !found && err == nil {
		return nil, false, nil
	} else if !found && err != nil {
		return nil, false, err
	}

	post.Author = user.ID
	_, err = UserCollection.UpdateOne(Context, bson.D{{"_id", user.ID}}, bson.D{{"$push", bson.D{{"posts", post}}}})
	if err != nil {
		return nil, false, err
	}
	return &post, true, nil
}

func GetLastestPost(userId primitive.ObjectID) (*types.Post, bool, error) {
	user, found, err := GetByID(userId)
	if !found && err == nil {
		return nil, false, nil
	} else if !found && err != nil {
		return nil, false, err
	}

	post := _getMaxValue(user.Posts)

	return post, true, nil
}

func _getMaxValue(arr []types.Post) *types.Post {
	max := arr[0]
	for _, v := range arr {
		if v.CreatedAt.Unix() > max.CreatedAt.Unix() {
			max = v
		}
	}

	return &max
}
