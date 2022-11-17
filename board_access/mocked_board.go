package board_access

import (
	"github.com/autowagen/hoverboard-api"
	"log"
)

type HoverboardMockImpl struct{}

var MockedBoard hoverboard.Hoverboard = HoverboardMockImpl{}

func (m HoverboardMockImpl) Close() {
}

func (m HoverboardMockImpl) SetSteer(steer int16) {
	log.Printf("[hoverboard-mock] setSeer %v", steer)
}

func (m HoverboardMockImpl) SetSpeed(speed int16) {
	log.Printf("[hoverboard-mock] setSpeed %v", speed)
}

func (m HoverboardMockImpl) GetStatus() hoverboard.HoverboardStatus {
	return hoverboard.HoverboardStatus{
		Cmd1:       0,
		Cmd2:       0,
		SpeedRMaes: 0,
		SpeedLMaes: 0,
		BatVoltage: 0,
		BoardTemp:  0,
		CmdLed:     0,
	}
}
