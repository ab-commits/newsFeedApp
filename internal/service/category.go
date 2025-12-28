package service

import (
	"context"
	"newsApp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetByCategory(
	ctx context.Context,
	db *mongo.Database,
	category string,
) ([]models.Article, error) {

	filter := bson.M{
		"category": bson.M{"$in": []string{category}},
	}

	return findArticles(ctx, db, filter)
}
