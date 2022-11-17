package api

import (
	"encoding/json"
	"errors"
)

type CmdRequestBody struct {
	Command string
	Data    json.RawMessage
}

func HandleApiCmd(cid int, cmdJson []byte) error {
	var body CmdRequestBody
	err := json.Unmarshal(cmdJson, &body)
	if err != nil {
		return err
	}

	if body.Command == "rename" {
		return handleCmdRename(cid, body.Data)
	} else if body.Command == "set_drive" {
		return handleCmdSetDrive(body.Data)
	} else if body.Command == "abort" {
		return handleCmdAbort()
	} else {
		return errors.New("unknown command")
	}
}
