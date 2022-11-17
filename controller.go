package hoverboard_controller

import (
	"github.com/autowagen/controller/board_access"
	"github.com/autowagen/controller/config"
	"github.com/autowagen/controller/web"
)

func Start(mockedBoard bool) {
	config.Init()
	board_access.InitBoard(mockedBoard)
	web.InitWeb()
}
