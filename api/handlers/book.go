package handlers

import (
	"context"
    "net/http"
	"log"
    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/ProbePlusDevTeam/LetsGo/config"
    "github.com/ProbePlusDevTeam/LetsGo/api/models"
)

/*Function to get all the books details */
func GetBooks(c *gin.Context){  //'c' represents context of incoming HTTP request containing information about request(headers,query parameters & response writer)
	client,_ := database.ConnectMongoDB()
	collection := client.Database("Sample_GO_DB").Collection("Books")
	filter :=bson.M{}
	cursor,_ := collection.Find(context.Background(), filter)
	var books []models.Book
	for cursor.Next(context.Background()) {
        var book models.Book
        if err := cursor.Decode(&book); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
            return
        }
        books = append(books, book)
    }
    c.JSON(http.StatusOK, books) //send an HTTP response in JSON format the HTTP status code & data to be serialized and sent as JSON
}

/*Function to get a book detail*/
func GetBook(c *gin.Context){
	var book models.Book
	id:= c.Param("id")
	client,_ := database.ConnectMongoDB()
	collection := client.Database("Sample_GO_DB").Collection("Books")
	filter := bson.M{"id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&book)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
        return
    }
	c.JSON(http.StatusOK, book)
}

/*function to create a Book detail */
func PostBook(detail *gin.Context){
	client,_ := database.ConnectMongoDB()
	collection := client.Database("Sample_GO_DB").Collection("Books")
	var newbook models.Book
	if error := detail.BindJSON(&newbook); error!=nil{     //If the JSON data in the request body is successfully parsed & matches the structure of newbook variable, this function call will return nil (indicating no error).
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	_, err := collection.InsertOne(context.Background(), newbook)
    if err != nil {
        log.Fatalf("Failed to insert book: %v", err)
        return 
    }
	detail.JSON(http.StatusCreated,newbook)
}



func UpdateBook(detail *gin.Context) {
	id := detail.Param("id")
	client,_ := database.ConnectMongoDB()
	var updatebook models.Book
	if error := detail.BindJSON(&updatebook); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	collection := client.Database("Sample_GO_DB").Collection("Books")
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatebook}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(false)
	result := collection.FindOneAndUpdate(context.Background(), filter,update,options)
	if result.Err() == mongo.ErrNoDocuments {
        detail.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
        return
	}
	detail.JSON(http.StatusOK, updatebook)
}


/*Function to delete a book detail */
func DelBook(detail *gin.Context) {
    id := detail.Param("id")
	client,_ := database.ConnectMongoDB()
	collection := client.Database("Sample_GO_DB").Collection("Books")
	filter := bson.M{"id": id}
	result,_ := collection.DeleteOne(context.Background(), filter)
    
    if result.DeletedCount == 1 {
        detail.JSON(http.StatusOK, gin.H{"message": "Details successfully deleted"})
    } else {
        detail.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
    }
}
    