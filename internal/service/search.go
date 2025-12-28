package service

import (
	"context"
	"newsApp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchArticles(
	ctx context.Context,
	db *mongo.Database,
	query string,
) ([]models.Article, error) {

	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	return findArticles(ctx, db, filter)
}
