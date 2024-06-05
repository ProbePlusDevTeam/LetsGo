package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/Baiguoshuai1/shadiaosocketio/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI       = "mongodb://deepak:deepak123@0.0.0.0:27017/"
	dbName         = "ECG-go"
	collectionName = "live_data"
)

var (
	pingInterval = 25 * time.Second
	pingTimeout  = 20 * time.Second
	client       *mongo.Client
)

func main() {
	server := shadiaosocketio.NewServer(*websocket.GetDefaultWebsocketTransport())

	server.On(shadiaosocketio.OnConnection, func(c *shadiaosocketio.Channel, msg string) {
		fmt.Println("connected! id:", c.Id(), c.LocalAddr().Network()+" "+c.LocalAddr().String()+
			" --> "+c.RemoteAddr().Network()+" "+c.RemoteAddr().String())
		fmt.Println("on message: on connect", msg)
		messageStr := string(msg)
		fmt.Println("Recieved message: on connect", messageStr)
		// message := fmt.Sprintf(`0{"upgrades":["websocket"],"pingInterval":%s,"pingTimeout":%s,"maxPayload":1000000}`, pingInterval, pingTimeout)
		// _ = c.Emit("message", message)

	})

	server.On(shadiaosocketio.OnDisconnection, func(c *shadiaosocketio.Channel, reason websocket.CloseError) {
		fmt.Println("disconnect", c.Id(), "code:", reason.Code, "text:", reason.Text)
	})

	server.On("patch_id", func(c *shadiaosocketio.Channel, msg map[string]string) {
		fmt.Println("on message:", "mes:", msg)
		patchID := msg["patch_id"]
		log.Printf("Received patch ID: %s", patchID)

		fetchAndSendLastThreeRecordsEveryFiveSeconds(c, patchID)
		_ = c.Emit("message", "some message")

	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)
	host := "0.0.0.0:8080"
	fmt.Println("starting ...")
	log.Panic(http.ListenAndServe(host, serveMux))
}

// fetchAndSendLastThreeRecordsEveryFiveSeconds fetches and sends the last 3 records every 5 seconds
func fetchAndSendLastThreeRecordsEveryFiveSeconds(c *shadiaosocketio.Channel, patchID string) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fetchLastThreeRecords(c, patchID)
	}
}

// fetchLastThreeRecords fetches the last 3 records from the collection
func fetchLastThreeRecords(c *shadiaosocketio.Channel, patchID string) {

	fmt.Println("Fetching data for patch ID:", patchID)
	client = NewDbConnection()
	collection := client.Database("ECG-go").Collection("live_data")

	var result map[string]interface{}

	filter := bson.M{
		"PatchId": patchID,
	}
	opts := options.Find()

	// Set the sort order in FindOptions
	opts.SetSort(bson.D{{Key: "PatchId", Value: -1}})

	// Create FindOneOptions and copy relevant settings from opts
	findOneOpts := options.FindOne()
	findOneOpts.SetProjection(opts.Projection)
	findOneOpts.SetSort(opts.Sort)
	err := collection.FindOne(context.TODO(), filter, findOneOpts).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("No new data found.")
			// Sleep for some time before retrying
			time.Sleep(5 * time.Second)

		}
		log.Println("Find error:", err)

	}
	// resultJSON, err := json.Marshal(result)
	// if err != nil {
	// 	log.Println("Error marshaling JSON:", err)

	// }

	// Append the JSON result to the websocket message
	// message := append([]byte(`42["update",`), resultJSON...)
	// message = append(message, ']')
	// fmt.Println(message)
	err = c.Emit("update", result)
	if err != nil {
		log.Println("WebSocket write error:", err)

	}

	fmt.Println("Data sent successfully.")
}

// watchForChanges listens for changes in the given MongoDB collection
func watchForChanges(collection *mongo.Collection) {
	pipeline := mongo.Pipeline{}
	options := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	changeStream, err := collection.Watch(context.TODO(), pipeline, options)
	if err != nil {
		log.Fatal(err)
	}
	defer changeStream.Close(context.Background())

	fmt.Println("Listening for changes...")

	// Listen for changes
	for changeStream.Next(context.Background()) {
		var event bson.M
		if err := changeStream.Decode(&event); err != nil {
			log.Fatal(err)
		}

		// Handle the change event (insert, update, delete, etc.)
		// For example, print the change event
		fmt.Printf("Change detected: %v\n", event)

		// Uncomment the following line to process the change event data
		// processChange(event)
	}

	if err := changeStream.Err(); err != nil {
		log.Fatal(err)
	}
}

func NewDbConnection() *mongo.Client {

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("newdbconnection", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the database
	// TODO: FUTURE : handle error if not able to connect to mongodb below
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}
