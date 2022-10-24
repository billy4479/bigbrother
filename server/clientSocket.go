package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

type Client struct {
	writeChan chan message
}

var client *Client

func clientSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	readChan := make(chan message)
	errorChan := make(chan error)

	go func() {
		for {
			msgType, buffer, err := ws.ReadMessage()
			if err != nil {
				errorChan <- err
			} else {
				readChan <- message{msgType: msgType, buffer: buffer}
			}
		}
	}()

	client = &Client{
		writeChan: make(chan message),
	}
	defer func() {
		close(client.writeChan)
		client = nil
	}()

	for {
		select {
		case message := <-readChan:
			if ui != nil {
				ui.writeChan <- message
			}

			if message.msgType == websocket.TextMessage {
				if string(message.buffer) == "clientDisconnect" {
					return nil
				}
			}

		case message := <-client.writeChan:
			errorChan <- ws.WriteMessage(message.msgType, message.buffer)

		case err := <-errorChan:
			if err != nil {
				ws.Close()
				close(readChan)
				close(errorChan)
				return err
			}
		}
	}

}
