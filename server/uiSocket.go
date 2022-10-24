package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type message struct {
	msgType int
	buffer  []byte
}

type UI struct {
	writeChan chan message
}

var ui *UI = nil

func uiSocket(c echo.Context) error {
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

	ui = &UI{
		writeChan: make(chan message),
	}
	defer func() {
		close(ui.writeChan)
		ui = nil

		close(readChan)
		close(errorChan)
		ws.Close()
	}()

	for {
		select {
		case message := <-readChan:
			if client != nil {
				client.writeChan <- message
			}
			if message.msgType == websocket.TextMessage {
				if string(message.buffer) == "uiDisconnect" {
					return nil
				}
			}

		case message := <-ui.writeChan:
			errorChan <- ws.WriteMessage(message.msgType, message.buffer)

		case err := <-errorChan:
			if err != nil {
				return err
			}
		}
	}

}
