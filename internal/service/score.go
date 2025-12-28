package service

import (
	"context"
	"newsApp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetByScore(
	ctx context.Context,
	db *mongo.Database,
	minScore float64,
) ([]models.Article, error) {

	filter := bson.M{
		"relevance_score": bson.M{"$gte": minScore},
	}

	return findArticles(ctx, db, filter)
}
