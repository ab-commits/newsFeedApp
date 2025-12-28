package service

import (
	"context"
	"newsApp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findArticles(
	ctx context.Context,
	db *mongo.Database,
	filter bson.M,
) ([]models.Article, error) {

	cur, err := db.Collection("articles").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var articles []models.Article
	if err := cur.All(ctx, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}
