package handlers

import (
	"context"
    "net/http"
    "github.com/gin-gonic/gin"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/ProbePlusDevTeam/LetsGo/config"
    "github.com/ProbePlusDevTeam/LetsGo/api/models"
)

func CreateUSer(detail *gin.Context){
	var newuser models.User
	client,_ := database.ConnectMongoDB()
	collection := client.Database("Sample_GO_DB").Collection("Users")
	if error := detail.BindJSON(&newuser); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	_, err := collection.InsertOne(context.Background(), newuser)
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
        return 
    }
	detail.JSON(http.StatusCreated,newuser)
}

func GetUsers(c *gin.Context){  
	client,_ := database.ConnectMongoDB()
	collection := client.Database("Sample_GO_DB").Collection("Users")
	filter :=bson.M{}
	cursor,_ := collection.Find(context.Background(), filter)
	var users []models.User
	for cursor.Next(context.Background()) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
            return
        }
        users = append(users, user)
    }
    c.JSON(http.StatusOK, users) //send an HTTP response in JSON format the HTTP status code & data to be serialized and sent as JSON
}

func UpdateUser(detail *gin.Context) {
	id := detail.Param("id")
	client,_ := database.ConnectMongoDB()
	var updateuser models.User
	if error := detail.BindJSON(&updateuser); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	collection := client.Database("Sample_GO_DB").Collection("Users")
	filter := bson.M{"id": id}
	update := bson.M{"$set": updateuser}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(false)
	result := collection.FindOneAndUpdate(context.Background(), filter,update,options)
	if result.Err() == mongo.ErrNoDocuments {
        detail.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
        return
	}
	detail.JSON(http.StatusOK, updateuser)
}
