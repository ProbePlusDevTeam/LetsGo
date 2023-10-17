package database

import (
     "os"
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo/options"
)


var client *mongo.Client

// ConnectMongoDB connects to the MongoDB database and returns a MongoDB client instance.
func ConnectMongoDB() (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

    // Connect to MongoDB
    c, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }

    // Check if the connection is successful
    err = c.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }

    return c, nil
}

func GetCollection(CollectionName string) *mongo.Collection {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    dbname := os.Getenv("DB_NAME")
    client,_ := ConnectMongoDB()
	var collection *mongo.Collection = client.Database(dbname).Collection(CollectionName)
	return collection

}