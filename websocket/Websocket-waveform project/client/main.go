package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	loginURL        = "https://uat2a-api.ubiqvue.com/api/v1/resources/auth/login"
	refreshTokenURL = "https://uat2a-api.ubiqvue.com/api/v1/resources/auth/refresh-token"
	webSocketURL    = "wss://uat2a-zoom.ubiqvue.com/socket.io/?EIO=4&transport=websocket"
	username        = "deepak@lifesignals.com"
	password        = "ecg123"
	// loginURL               = "https://alpha.api.ubiqvue.com/api/v1/resources/auth/login"
	// refreshTokenURL        = "https://alpha.api.ubiqvue.com/api/v1/resources/auth/refresh-token"
	// webSocketURL           = "wss://alpha-zoom.ubiqvue.com/socket.io/?EIO=4&transport=websocket"
	// username               = "alpha.cfa01+newgc10@gmail.com"
	// password               = "N@than$05"
	mongoURI               = "mongodb://deepak:deepak123@0.0.0.0:27017/"
	dbName                 = "ECG-go"
	collectionName         = "Patches_details"
	liveDataCollectionName = "live_data"
)

var (
	accessToken     string
	refreshToken    string
	patchIDs        = []string{"S3132"} // List of patch IDs
	wg              sync.WaitGroup
	MongoConnection *mongo.Client
)

func main() {
	accessToken, refreshToken = authenticateUser(username, password)
	MongoConnection = NewDbConnection()
	var EcgCollection *mongo.Collection = GetCollection(collectionName)
	var LiveDataCollection *mongo.Collection = GetCollection(liveDataCollectionName)
	// Ensure the TTL index is set up for live_data collection
	EnsureTTLIndex(LiveDataCollection)
	for _, patchID := range patchIDs {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			websocketURLWithToken := fmt.Sprintf("%s&token=%s", webSocketURL, accessToken)
			conn, err := connectWebSocket(websocketURLWithToken)
			if err != nil {
				fmt.Printf("Error connecting to WebSocket for patch ID %s: %v\n", id, err)
				return
			}
			defer conn.Close()

			handleWebSocketMessages(conn, id, EcgCollection, LiveDataCollection)
		}(patchID)
	}

	wg.Wait()
	fmt.Println("All WebSocket connections closed.")
}
func NewDbConnection() *mongo.Client {

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
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
func GetCollection(collectionName string) *mongo.Collection {
	return MongoConnection.Database(dbName).Collection(collectionName)
}
func EnsureTTLIndex(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"timestamp": 1},
		Options: options.Index().SetExpireAfterSeconds(600), // 10 minutes
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal("Error creating TTL index:", err)
	}
}
func authenticateUser(username, password string) (string, string) {
	client := &http.Client{}
	data := map[string]string{"username": username, "password": password}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error authenticating user:", err)
		return "", ""
	}
	defer resp.Body.Close()

	var authResponse map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&authResponse)

	return authResponse["data"].(map[string]interface{})["access_token"].(string),
		authResponse["data"].(map[string]interface{})["refresh_token"].(string)
}

func connectWebSocket(url string) (*websocket.Conn, error) {
	c := websocket.DefaultDialer
	conn, _, err := c.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func handleWebSocketMessages(conn *websocket.Conn, patchID string, EcgCollection, LiveDataCollection *mongo.Collection) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message from WebSocket for patch ID %s: %v\n", patchID, err)
			return
		}
		onMessage(conn, message, patchID, EcgCollection, LiveDataCollection)
	}
}

func onMessage(conn *websocket.Conn, message []byte, patchID string, EcgCollection, LiveDataCollection *mongo.Collection) {
	fmt.Println("Received message for patch ID", patchID)
	fmt.Println("Received message for patch ID", patchID, ": ", string(message))

	// Check if message contains "40{\"sid\":\""
	if strings.Contains(string(message), "40{\"sid\":\"") {
		fmt.Println("Message contains '40{\"sid\":\"'")
		sendPatchID(conn, patchID)
	}
	messageType := string(message[0])
	switch messageType {
	case "0":
		sendAccessToken(conn)
	case "2":
		sendAck(conn)
	case "4":
		if string(message[1]) == "0" {
			fmt.Println("Message contains", string(message[:10]))
			sendPatchID(conn, patchID)
		}
		if string(message[1]) == "2" {
			processData(message, patchID, EcgCollection, LiveDataCollection)
		}

	default:
		fmt.Println("Unknown message type:", messageType)
	}
}

func sendPatchID(conn *websocket.Conn, patchID string) {
	fmt.Println("Sending patch ID")
	patchIDMessage := []byte(fmt.Sprintf(`42["patch_id",{"patch_id":"%s"}]`, patchID))
	conn.WriteMessage(websocket.TextMessage, patchIDMessage)
}
func sendAccessToken(conn *websocket.Conn) {
	fmt.Println("Sending access token")
	message := []byte(fmt.Sprintf("40{\"token\":\"%s\"}", accessToken))
	conn.WriteMessage(websocket.TextMessage, message)
}

func sendAck(conn *websocket.Conn) {
	fmt.Println("Sending ack")
	message := []byte("3")
	conn.WriteMessage(websocket.TextMessage, message)
}

func processData(message []byte, patchID string, EcgCollection, LiveDataCollection *mongo.Collection) {
	var data []interface{}
	err := json.Unmarshal(message[2:], &data)
	if err != nil {
		fmt.Println("Error decoding JSON data:", err)
		return
	}
	if len(data) > 1 {
		dataMap, ok := data[1].(map[string]interface{})
		if !ok {
			fmt.Println("Error decoding data map")
			return
		}
		//just to print
		sensorData, ok := dataMap["SensorData"].([]interface{})
		if !ok {
			fmt.Println("Error decoding sensor data")
			return
		}
		//Just to print
		// fmt.Println("Full data:", dataMap)
		for _, data := range sensorData {
			// fmt.Println("_")
			dataMap, ok := data.(map[string]interface{})
			if !ok {
				fmt.Println("Error decoding sensor data map")
				continue
			}

			// Access specific sensor data fields
			ecgCHA, ok := dataMap["ECG_CH_A"].(float64)
			if ok {
				fmt.Println(patchID, "ECG_CH_A:", ecgCHA)
			}

			ecgCHB, ok := dataMap["ECG_CH_B"].(float64)
			if ok {
				fmt.Println(patchID, "ECG_CH_B:", ecgCHB)
			}
		}
		// fmt.Println("Full data of patch ID", patchID, ":", dataMap)

		// Insert dataMap into MongoDB

		_, err := EcgCollection.InsertOne(context.Background(), dataMap)
		if err != nil {
			fmt.Println("Error inserting data into MongoDB:", err)
		} else {
			fmt.Println("Data inserted into MongoDB successfully")
		}
		// Insert into live_data collection with timestamp
		dataMap["timestamp"] = time.Now()
		_, err = LiveDataCollection.InsertOne(context.Background(), dataMap)
		if err != nil {
			fmt.Println("Error inserting data into live_data MongoDB:", err)
		} else {
			fmt.Println("Data inserted into live_data MongoDB successfully")
		}
	}
}
