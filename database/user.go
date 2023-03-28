package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"plusone/backend/types"
)

func GetByID(id primitive.ObjectID) (*types.User, bool, error) {
	var user *types.User

	log.Println(id.String())

	err := UserCollection.FindOne(Context, bson.D{{"_id", id}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func GetByEmail(email string) (*types.User, bool, error) {
	var user *types.User

	err := UserCollection.FindOne(Context, bson.D{{"email", email}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func GetByUsername(username string) (*types.User, bool, error) {
	var user *types.User

	err := UserCollection.FindOne(Context, bson.D{{"username", username}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func CreateUser(user types.User) (bool, error) {
	_, err := UserCollection.InsertOne(Context, user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetManyUserID(ids []primitive.ObjectID) (*[]types.UserSensored, bool, error) {
	var users []types.UserSensored
	cursor, err := UserCollection.Find(Context, bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return nil, false, err
	}

	if err = cursor.All(Context, &users); err != nil {
		return nil, false, err
	}
	return &users, true, nil
}
