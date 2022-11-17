package api

import (
	"github.com/autowagen/controller/board_access"
)

func handleCmdAbort() error {
	board_access.GetBoard().SetSteer(0)
	board_access.GetBoard().SetSpeed(0)

	return nil
}
