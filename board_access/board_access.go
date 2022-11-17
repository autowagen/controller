package board_access

import (
	"github.com/autowagen/controller/config"
	"github.com/autowagen/hoverboard-api"
	"log"
)

var board hoverboard.Hoverboard

func InitBoard(mockedBoard bool) {
	serialPort := config.GetHoverboardSerialPort()

	log.Printf("serialPort: %v", serialPort)

	if mockedBoard {
		board = MockedBoard
		return
	}

	h, err := hoverboard.NewHoverboard(serialPort)
	if err != nil {
		log.Fatalf("failed to initialize hoverboard-api: %v", err)
	}
	board = h

}

func GetBoard() hoverboard.Hoverboard {
	if board == nil {
		log.Fatalf("board_access not initialized")
	}
	return board
}
