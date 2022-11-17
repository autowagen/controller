package web

import "github.com/autowagen/controller/board_access"

type InfoResponse struct {
	Steer          int16
	Speed          int16
	SpeedL         int16
	SpeedR         int16
	BatteryVoltage int16
}

func getInfo() InfoResponse {
	status := board_access.GetBoard().GetStatus()
	return InfoResponse{
		Steer:          status.Cmd1,
		Speed:          status.Cmd2,
		SpeedL:         status.SpeedLMaes,
		SpeedR:         status.SpeedRMaes,
		BatteryVoltage: status.BatVoltage,
	}
}
