package utils

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"plusone/backend/types"
	"strconv"
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

func StringArrToGeoJSON(arr string) (*types.GeoJSON, error) {
	parts := strings.Split(arr, ",")
	// Convert the string slice to a float64 slice
	var coords []float64
	for _, part := range parts {
		f, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return nil, err
		}
		coords = append(coords, f)
	}

	// Create a new GeoJSONPoint object using the converted coordinates
	geoPoint := types.GeoJSON{
		Type:        "Point",
		Coordinates: coords,
	}

	return &geoPoint, nil
}
