package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"plusone/backend/types"
)

func GetByID(id string) (*types.User, bool, error) {
	var user *types.User

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
