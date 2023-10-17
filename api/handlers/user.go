package handlers

import (
	"fmt"
	"strconv"
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

var UserCollection *mongo.Collection = database.GetCollection("Users")


func CreateUSer(detail *gin.Context){
	var newuser models.User
	if error := detail.BindJSON(&newuser); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	_, err := UserCollection.InsertOne(context.Background(), newuser)
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
        return 
    }
	detail.JSON(http.StatusCreated,newuser)
}

func GetUsers(c *gin.Context){  
	filter :=bson.M{}
	cursor,_ := UserCollection.Find(context.Background(), filter)
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

func GetUser(c *gin.Context){
	var newuser models.User
	id,_ := strconv.Atoi(c.Param("id"))
	filter := bson.M{"id": id}
	err := UserCollection.FindOne(context.Background(), filter).Decode(&newuser)
	fmt.Println(filter)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
        return
    }
	c.JSON(http.StatusOK, newuser)
}

func UpdateUser(detail *gin.Context) {
	id,_ := strconv.Atoi(detail.Param("id"))
	var updateuser models.User
	if error := detail.BindJSON(&updateuser); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	filter := bson.M{"id": id}
	update := bson.M{"$set": updateuser}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(false)
	result := UserCollection.FindOneAndUpdate(context.Background(), filter,update,options)
	if result.Err() == mongo.ErrNoDocuments {
        detail.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
        return
	}
	detail.JSON(http.StatusOK, updateuser)
}
