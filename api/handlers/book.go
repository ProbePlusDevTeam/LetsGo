package handlers

import (
	"sync"
	"context"
    "net/http"
	"strconv"
	"log"
    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/ProbePlusDevTeam/LetsGo/config"
    "github.com/ProbePlusDevTeam/LetsGo/api/models"
)

var BookCollection *mongo.Collection = database.GetCollection("Books")
var Book1C *mongo.Collection = database.GetCollection("BookCn1")
var Book2C *mongo.Collection = database.GetCollection("BookCn2")
var CombinedBookC *mongo.Collection = database.GetCollection("CombinedBookC")

/*Function to get all the books details */
func GetBooks(c *gin.Context){  
	var wg sync.WaitGroup   
	filter :=bson.M{}
	
	var books []models.Book        
    bookCh := make(chan models.Book)           // Channel to receive books from Goroutines
	
	wg.Add(1)             //wg.Add(1) is used to add a Goroutine to the WaitGroup
    go func() {                               // Start a Goroutine for fetching books
        cursor, _ := BookCollection.Find(context.Background(), filter)

        for cursor.Next(context.Background()) {
            var book models.Book
            if err := cursor.Decode(&book); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
                return
            }
            bookCh <- book
        }
        wg.Done()      
    }()
   
    go func() {            // Start a Goroutine to close the channel when all books are processed
        wg.Wait()       // Wait for all Goroutines to finish & the main function will be unblocked, allowing the program to proceed.
        close(bookCh)   // Close the channel when all books are processed
    }()
	for book := range bookCh {
        books = append(books, book)
    }

	// for cursor.Next(context.Background()) {
    //     var book models.Book
    //     if err := cursor.Decode(&book); err != nil {
    //         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
    //         return
    //     }
    //     books = append(books, book)
    // }
    c.JSON(http.StatusOK, books) //send an HTTP response in JSON format the HTTP status code & data to be serialized and sent as JSON
}



/*Function to get a book detail*/
func GetBook(c *gin.Context){
	var book models.Book
	id:= c.Param("id")
	filter := bson.M{"id": id}
	err := BookCollection.FindOne(context.Background(), filter).Decode(&book)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
        return
    }
	c.JSON(http.StatusOK, book)
}

/*function to create a Book detail */
func PostBook(detail *gin.Context){
	var newbook models.Book
	if error := detail.BindJSON(&newbook); error!=nil{     //If the JSON data in the request body is successfully parsed & matches the structure of newbook variable, this function call will return nil (indicating no error).
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	_, err := BookCollection.InsertOne(context.Background(), newbook)
    if err != nil {
        log.Fatalf("Failed to insert book: %v", err)
        return 
    }
	detail.JSON(http.StatusCreated,newbook)
}


/*function to update a Book detail */
func UpdateBook(detail *gin.Context) {
	id := detail.Param("id")
	var updatebook models.Book
	if error := detail.BindJSON(&updatebook); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatebook}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(false)
	result := BookCollection.FindOneAndUpdate(context.Background(), filter,update,options)
	if result.Err() == mongo.ErrNoDocuments {
        detail.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
        return
	}
	detail.JSON(http.StatusOK, updatebook)
}


/*Function to delete a book detail */
func DelBook(detail *gin.Context) {
    id := detail.Param("id")
	filter := bson.M{"id": id}
	result,_ := BookCollection.DeleteOne(context.Background(), filter)
    
    if result.DeletedCount == 1 {
        detail.JSON(http.StatusOK, gin.H{"message": "Details successfully deleted"})
    } else {
        detail.JSON(http.StatusNotFound, gin.H{"msg": "Book Not Found"})
    }
}


//Applying sorting and filtering 
func PostBook1(detail *gin.Context){
	var newbook models.Book1
	if error := detail.BindJSON(&newbook); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	_, err := Book1C.InsertOne(context.Background(), newbook)
    if err != nil {
        log.Fatalf("Failed to insert book: %v", err)
        return 
    }
	detail.JSON(http.StatusCreated,newbook)
}

func PostBook2(detail *gin.Context){
	var newbook models.Book2
	if error := detail.BindJSON(&newbook); error!=nil{     
		detail.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	}
	_, err := Book2C.InsertOne(context.Background(), newbook)
    if err != nil {
        log.Fatalf("Failed to insert book: %v", err)
        return 
    }
	detail.JSON(http.StatusCreated,newbook)
}

func GetDetails(detail *gin.Context){

	minprice:= detail.DefaultQuery("min","0")
	maxprice:= detail.DefaultQuery("max","1000")
	minPrice,_ := strconv.ParseFloat(minprice, 64)
	maxPrice,_ := strconv.ParseFloat(maxprice, 64)
	filter:= bson.M{"price": bson.M{"$gte": minPrice, "$lte": maxPrice}}

	sortField := detail.DefaultQuery("sort","price")
	sortOptions := options.Find().SetSort(bson.M{sortField: 1}).SetCollation(&options.Collation{Locale: "en_US", Strength: 2})
	cursor,_ := Book1C.Find(context.Background(), filter,sortOptions)

	var books []models.Book1
	for cursor.Next(context.Background()) {
        var book models.Book1
        if err := cursor.Decode(&book); err != nil {
            detail.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
            return
        }
        books = append(books, book)
    }
    detail.JSON(http.StatusOK, books)
	// sort :=bson.M{}

}

func CreateNewCollection(detail *gin.Context){

    pipeline := []bson.M{
        bson.M{
            "$lookup": bson.M{
                "from":         "BookCn1",    // The name of the first collection
                "localField":   "id",
                "foreignField": "id",
                "as":           "collection1Data",
            },
        },
        bson.M{
            "$lookup": bson.M{
                "from":         "BookCn2",    // The name of the second collection
                "localField":   "id",
                "foreignField": "id",
                "as":           "collection2Data",
            },
        },
        bson.M{
			"$project": bson.M{
                "ID":             bson.M{"$first": "$collection1Data.id"},
                "Title":           bson.M{"$first": "$collection1Data.title"},
                "Author":         bson.M{"$first": "$collection2Data.author"},
				"Date":       bson.M{"$first": "$collection2Data.published_date"},
            },
        },
    }
	
	cursor, err := Book2C.Aggregate(context.Background(), pipeline)
    if err != nil {
        detail.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform aggregation"})
        return
    }


	// Create an array to hold the merged documents
    var mergedData []bson.M

    // Decode the results into the mergedData slice
    if err := cursor.All(context.Background(), &mergedData); err != nil {
        detail.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode aggregation results"})
        return
    }
	var mergedDataAsInterface []interface{}
	for _, doc := range mergedData {
		mergedDataAsInterface = append(mergedDataAsInterface, doc)
	}
    // Insert the merged data into the new collection
    if _, err := CombinedBookC.InsertMany(context.Background(), mergedDataAsInterface); err != nil {
        detail.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert merged data into the new collection"})
        return
    }

    detail.JSON(http.StatusOK, mergedData)
}
    
