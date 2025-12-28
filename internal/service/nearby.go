package service

import (
	"context"
	"newsApp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNearby(
	ctx context.Context,
	db *mongo.Database,
	lat, lng float64,
) ([]models.Article, error) {

	filter := bson.M{
		"latitude":  bson.M{"$gte": lat - 0.5, "$lte": lat + 0.5},
		"longitude": bson.M{"$gte": lng - 0.5, "$lte": lng + 0.5},
	}

	return findArticles(ctx, db, filter)
}
