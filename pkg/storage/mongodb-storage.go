package storage

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"journal/models"
	"log"
	"os"
	"time"
)

type MongoDBStorage struct {
	DB *mongo.Collection
}

// NewMongoDBStorage initializes the MongoDB database and returns a storage collection instance
func NewMongoDBStorage(databaseName string, collectionName string) (*MongoDBStorage, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to create mongo client: %w", err)
	}

	coll := client.Database(databaseName).Collection(collectionName)

	return &MongoDBStorage{
		DB: coll,
	}, nil
}

func (s *MongoDBStorage) CreateEntry(entry models.Entry) error {
	_, err := s.DB.InsertOne(context.Background(), entry)
	return err
}

func (s *MongoDBStorage) LoadEntries() ([]models.Entry, error) {

	cursor, err := s.DB.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var entries []models.Entry
	for cursor.Next(context.Background()) {
		var entry models.Entry
		if err := cursor.Decode(&entry); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func (s *MongoDBStorage) GetEntry(id string) (models.Entry, error) {
	var entry models.Entry
	err := s.DB.FindOne(context.Background(), bson.M{"id": id}).Decode(&entry)
	if err != nil {
		return entry, err
	}
	return entry, nil
}

func (s *MongoDBStorage) UpdateEntry(entry models.Entry) error {
	_, err := s.DB.UpdateOne(
		context.Background(),
		bson.M{"id": entry.ID},
		bson.M{"$set": bson.M{
			"id":      entry.ID,
			"title":   entry.Title,
			"content": entry.Content,
			"created": entry.Created,
			"updated": entry.Updated,
		}},
	)
	return err
}

func (s *MongoDBStorage) DeleteEntry(id string) error {
	_, err := s.DB.DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

// SaveEntries saves the journal entries to the SQLite database (insert or update)
func (s *MongoDBStorage) SaveEntries(entries []models.Entry) error {
	for _, entry := range entries {
		filter := bson.M{"id": entry.ID}
		update := bson.M{
			"$set": bson.M{
				"id":      entry.ID,
				"title":   entry.Title,
				"content": entry.Content,
				"created": entry.Created,
				"updated": entry.Updated,
			},
		}
		opts := options.Update().SetUpsert(true)

		_, err := s.DB.UpdateOne(context.Background(), filter, update, opts)
		if err != nil {
			return err
		}
	}

	return nil
}
