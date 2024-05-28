package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/Baiguoshuai1/shadiaosocketio/websocket"
)

func main() {
	server := shadiaosocketio.NewServer(*websocket.GetDefaultWebsocketTransport())

	server.On(shadiaosocketio.OnConnection, func(c *shadiaosocketio.Channel) {
		fmt.Println("connected! id:", c.Id(), c.LocalAddr().Network()+" "+c.LocalAddr().String()+
			" --> "+c.RemoteAddr().Network()+" "+c.RemoteAddr().String())

		_ = c.Emit("message", "hi client !!!")

	})

	server.On(shadiaosocketio.OnDisconnection, func(c *shadiaosocketio.Channel, reason websocket.CloseError) {
		fmt.Println("disconnect", c.Id(), "code:", reason.Code, "text:", reason.Text)
	})

	server.On("/message", func(c *shadiaosocketio.Channel, msg string) {
		fmt.Println("on message:", "mes:", msg)

		_ = c.Emit("message", "some message")

	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)
	host := "0.0.0.0:8080"
	fmt.Println("starting ...")
	log.Panic(http.ListenAndServe(host, serveMux))
}
