package utils

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

func StringToObjectIDs(ids []string) ([]primitive.ObjectID, error) {
	newIds := []primitive.ObjectID{}
	for _, v := range ids {
		objId, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return newIds, errors.New("failed to convert string")
		}
		newIds = append(newIds, objId)
	}
	return newIds, nil
}

func StringToObjectId(id string) (*primitive.ObjectID, error) {
	parts := strings.Split(id, `"`)
	splittedId := parts[1]
	objId, err := primitive.ObjectIDFromHex(splittedId)
	if err != nil {
		return nil, errors.New("failed to convert string")
	}

	return &objId, nil
}
