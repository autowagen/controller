package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/autowagen/controller/api"
	"github.com/autowagen/controller/clients"
	"github.com/autowagen/controller/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
)

type CmdRequestBody struct {
	Command string
	Data    json.RawMessage
}

func InitWeb() {
	fmt.Println("hello world")

	// TODO: register clients separately
	webCid := clients.RegisterClient("REST-API", "REST-API")

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/ws", websocket.Handle)

	e.GET("/debug/clients", func(c echo.Context) error {
		list := clients.GetClientList()
		return c.JSONPretty(200, list, "  ")
	})

	e.GET("/info", func(c echo.Context) error {
		return c.JSONPretty(200, getInfo(), "  ")
	})

	e.POST("/cmd", func(c echo.Context) error {
		if c.Request().Body == nil {
			//c.Error(errors.New("missing body"))
			return errors.New("missing body")
		}

		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return errors.New("unable to read request-body")
		}

		return api.HandleApiCmd(webCid, b)
	})

	e.Static("/", "static")
	//e.Pre(middleware.Rewrite(map[string]string{
	//	"/": "/index.html",
	//}))

	e.Logger.Fatal(e.Start(":8000"))
}
