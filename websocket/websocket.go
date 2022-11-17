package websocket

import (
	"github.com/autowagen/controller/api"
	"github.com/autowagen/controller/clients"
	"github.com/autowagen/controller/util/notify"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func Handle(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		cid := clients.RegisterClient("websocket", ws.Request().RemoteAddr)

		log.Printf("ws open")

		closeNotify := notify.NewNotify()

		// reader
		go func() {
			for !closeNotify.Check() {
				msg := ""
				err := websocket.Message.Receive(ws, &msg)
				if err != nil {
					closeNotify.Notify()
					return
				}

				log.Printf("ws msg %v", msg)

				err = api.HandleApiCmd(cid, []byte(msg))
				if err != nil {
					log.Fatalf("error while handline cmd=%v:", msg, err)
				}
			}
		}()

		// writer
		go func() {
			for !closeNotify.Check() {
				err := websocket.Message.Send(ws, "Hello, Client!")
				if err != nil {
					closeNotify.Notify()
				}
				time.Sleep(1 * time.Second)
			}
		}()

		closeNotify.Wait()

		log.Printf("ws close")

		clients.DeregisterClient(cid)

		/*
			for {
				// Write
				err := websocket.Message.Send(ws, "Hello, Client!")
				if err != nil {
					c.Logger().Error(err)
					return
				}

				// Read
				msg := ""
				err = websocket.Message.Receive(ws, &msg)
				if err != nil {
					c.Logger().Error(err)
					return
				}
				fmt.Printf("%s\n", msg)
			}
		*/
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
