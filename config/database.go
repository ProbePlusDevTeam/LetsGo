package database

import (
    "fmt"
    "log"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/ProbePlusDevTeam/LetsGo/api/models"
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


// func ConnectMongoDB() *mongo.Client {
//     Mongo_URL := "mongodb://127.0.0.1:27017"
//     client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))
    
//     if err != nil {
//     log.Fatal(err)
//     }
    
//     ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
//     err = client.Connect(ctx)
//     defer cancel()
    
//     if err != nil {
//     log.Fatal(err)
//     }
    
//     fmt.Println("Connected to mongoDB")
//     return client
//    }






func InsertBook(newBook *models.Book) error {
    fmt.Println("Success1")
    db := client.Database("Sample_GO_DB")
    fmt.Println("Success2")
    collection := db.Collection("Books")
    fmt.Println("Success3")
    // collection := client.Database("Sample_GO_DB").Collection("Books")
    _, err := collection.InsertOne(context.Background(), newBook)
    fmt.Println("Success")
    if err != nil {
        log.Fatalf("Failed to insert book: %v", err)
        return err
    }
    return nil
}