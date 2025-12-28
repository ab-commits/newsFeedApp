package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"newsApp/internal/db"
	"newsApp/internal/models"

	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println(len(files))
	for _, file := range files {
		data, _ := ioutil.ReadFile(file)
		fmt.Println(string(data))
		var articles []models.Article
		json.Unmarshal(data, &articles)

		for _, article := range articles {
			_, err := collection.InsertOne(ctx, bson.M{
				"id":               article.ID,
				"title":            article.Title,
				"description":      article.Description,
				"url":              article.URL,
				"publication_date": article.PublicationDate,
				"source_name":      article.SourceName,
				"category":         article.Category,
				"relevance_score":  article.RelevanceScore,
				"latitude":         article.Latitude,
				"longitude":        article.Longitude,
			})
			if err != nil {
				log.Println("Insert error:", err)
			}
		}
		fmt.Println("Loaded file:", file)
	}
	fmt.Println("All files loaded")
}
