This is a simple implementation of a Socket.IO server using the shadiaosocketio library in Go. The server listens for incoming WebSocket connections, handles connection and disconnection events, and processes messages received from clients.

Prerequisites
Go 1.16 or later

Installation
Clone the repository:
git clone https://github.com/your-repo/socketio-server.git
cd socketio-server
Install the required Go modules:
go mod tidy

Usage
To run the server, use the following command:
go run main.go
The server will start and listen for WebSocket connections on 0.0.0.0:8080.

Code Overview
Dependencies
shadiaosocketio: A Socket.IO server implementation in Go.
websocket: WebSocket transport for the Socket.IO server.

Event Handlers
OnConnection: Triggered when a client connects to the server. Logs connection details and sends a welcome message.
OnDisconnection: Triggered when a client disconnects from the server. Logs disconnection details.
/message: Triggered when a message is received from a client on the /message endpoint. Logs the message and responds with a confirmation message.

Running the Server
To start the server, execute go run main.go. The server will begin listening for WebSocket connections
