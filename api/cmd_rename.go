package api

import (
	"encoding/json"
	"github.com/autowagen/controller/clients"
	"log"
)

type RenameData struct {
	Name string
}

func handleCmdRename(cid int, jsonData json.RawMessage) error {
	var data RenameData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}

	log.Printf("rename cid=%v name=%v", cid, data.Name)

	clients.RenameClient(cid, data.Name)

	return nil
}
