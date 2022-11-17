package api

import (
	"encoding/json"
	"github.com/autowagen/controller/board_access"
	"log"
)

type SetDriveData struct {
	Steer float64
	Speed float64
}

func handleCmdSetDrive(jsonData json.RawMessage) error {
	var data SetDriveData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}

	log.Printf("steer %v speed %v", data.Steer, data.Speed)

	board_access.GetBoard().SetSteer(int16(data.Steer))
	board_access.GetBoard().SetSpeed(int16(data.Speed))

	return nil
}
