package service

import (
	"context"
	"newsApp/internal/models"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRelevantArticles(
	ctx context.Context,
	db *mongo.Database,
	llm *models.LLMResult,
	userLat, userLng float64,
) ([]models.Article, error) {

	collection := db.Collection("articles")
	filter := bson.M{}

	switch llm.Intent {

	case "source":
		// Example: "New York Times"
		if len(llm.Entities) > 0 {
			filter["source_name"] = bson.M{"$in": llm.Entities}
		}

	case "category":
		// Example: "technology"
		if len(llm.Concepts) > 0 {
			filter["category"] = bson.M{"$in": llm.Concepts}
		}

	case "score":
		filter["relevance_score"] = bson.M{"$gte": 0.7}

	case "search":
		// text search on title + description
		filter["$or"] = []bson.M{
			{"title": bson.M{"$regex": llm.RawQuery, "$options": "i"}},
			{"description": bson.M{"$regex": llm.RawQuery, "$options": "i"}},
		}

	case "nearby":
		if userLat != 0 && userLng != 0 {
			filter["latitude"] = bson.M{
				"$gte": userLat - 0.5,
				"$lte": userLat + 0.5,
			}
			filter["longitude"] = bson.M{
				"$gte": userLng - 0.5,
				"$lte": userLng + 0.5,
			}
		}
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var articles []models.Article
	if err := cur.All(ctx, &articles); err != nil {
		return nil, err
	}

	// Sort by relevance
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].RelevanceScore > articles[j].RelevanceScore
	})

	return articles, nil
}
