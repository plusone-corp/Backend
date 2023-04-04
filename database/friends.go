package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"plusone/backend/types"
)

func GetFriends(userId primitive.ObjectID) (*[]types.UserFiltered, error) {
	var users []types.UserFiltered
	cursor, err := UserCollection.Find(Context, bson.D{{"friends", bson.D{{"$elemMatch", bson.D{{"$eq", userId}}}}}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(Context)

	if err = cursor.All(Context, &users); err != nil { // pass pointer to users slice
		return nil, err
	}

	return &users, nil
}

func AddFriend(userId primitive.ObjectID, targetId string) error {
	friendId, err := primitive.ObjectIDFromHex(targetId)
	if err != nil {
		return err
	}

	userUpdate := mongo.NewUpdateOneModel()
	userUpdate.SetFilter(bson.D{{"_id", userId}})
	userUpdate.SetUpdate(bson.D{{"$push", bson.D{{"friends", friendId}}}})

	friendUpdate := mongo.NewUpdateOneModel()
	friendUpdate.SetFilter(bson.D{{"_id", friendId}})
	friendUpdate.SetUpdate(bson.D{{"$push", bson.D{{"friends", userId}}}})

	writeModel := make([]mongo.WriteModel, 0)
	writeModel = append(writeModel, userUpdate)
	writeModel = append(writeModel, friendUpdate)
	_, err = UserCollection.BulkWrite(Context, writeModel)
	if err != nil {
		return err
	}

	return nil
}

func GetAllFriends(userId primitive.ObjectID) (*[]types.UserFiltered, error) {
	var friends []types.UserFiltered

	filter := bson.D{{
		"author",
		userId,
	}}

	opt := options.Find().SetSort(bson.D{{"createdAt", 1}})
	cursor, err := PostCollection.Find(Context, filter, opt)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(Context, &friends); err != nil {
		return nil, err
	}

	return &friends, nil
}
