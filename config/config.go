package config

import (
	"github.com/tidwall/gjson"
	"log"
	"os"
)

type Config struct {
	Test             *string
	HoverboardConfig HoverboardConfig `json:"hoverboardConfig"`
}

type HoverboardConfig struct {
	SerialPort *string `json:"serialPort"`
}

var configJson = "{}"

func loadConfig() {
	bs, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("unable to read config.json", err)
	}

	configJson = string(bs)
}

func Init() {
	loadConfig()
}

func GetHoverboardSerialPort() string {
	v := gjson.Get(configJson, "hoverboard.serialPort").String()
	if v == "" {
		return "/dev/ttyUSB0"
	}
	return v
}
