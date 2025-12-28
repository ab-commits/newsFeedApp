package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"newsApp/internal/db"
	"newsApp/internal/models"
)

func main() {
	mongoDB, err := db.ConnectMongo("mongodb://localhost:27017", "news_db")
	if err != nil {
		log.Fatal(err)
	}

	collection := mongoDB.Collection("articles")
	files, err := filepath.Glob("./data/*.json")
	if err != nil {
		fmt.Println("error occured", err.Error())
		return
	}
	ctx := context.Background()
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var articles []models.Article
		if err := json.Unmarshal(data, &articles); err != nil {
			log.Fatal("JSON parse error:", err)
		}

		docs := make([]interface{}, 0, len(articles))
		for _, article := range articles {
			docs = append(docs, article)
		}

		if len(docs) > 0 {
			res, err := collection.InsertMany(ctx, docs)
			if err != nil {
				log.Fatal("InsertMany failed:", err)
			}
			log.Printf("Inserted %d documents from %s\n", len(res.InsertedIDs), file)
		}
	}
}
