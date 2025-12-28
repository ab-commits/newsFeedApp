package service

import (
	"context"
	"newsApp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBySource(
	ctx context.Context,
	db *mongo.Database,
	source string,
) ([]models.Article, error) {

	filter := bson.M{
		"source_name": source,
	}

	return findArticles(ctx, db, filter)
}
