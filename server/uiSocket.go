package main

import (
	"github.com/labstack/echo/v4"
)

func uiSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// TODO
		break
	}

	return nil

}
