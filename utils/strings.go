package utils

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
