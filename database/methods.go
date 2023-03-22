package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetByID(id string) (*User, bool, error) {
	var user *User

	err := UserCollection.FindOne(Context, bson.D{{"_id", id}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func GetByEmail(email string) (*User, bool, error) {
	var user *User

	err := UserCollection.FindOne(Context, bson.D{{"email", email}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func GetByUsername(username string) (*User, bool, error) {
	var user *User

	err := UserCollection.FindOne(Context, bson.D{{"username", username}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func CreateUser(user User) (bool, error) {
	_, err := UserCollection.InsertOne(Context, user)
	if err != nil {
		return false, err
	}

	return true, nil
}
