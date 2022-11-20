package hoverboard_controller

import (
	"github.com/autowagen/controller/board_access"
	"github.com/autowagen/controller/config"
	"github.com/autowagen/controller/web"
)

func Start() {
	config.Init()
	board_access.InitBoard(config.GetHoverboardMock())
	web.InitWeb()
}
