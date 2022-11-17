package clients

import (
	"sync"
)

type Client struct {
	Cid   int
	Label string
	Info  string
}

var clients = make([]*Client, 0)
var clientsMu sync.Mutex

var nextIdentifier = 1

func RegisterClient(label string, info string) int {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	id := nextIdentifier
	nextIdentifier = nextIdentifier + 1

	client := Client{id, label, info}
	clients = append(clients, &client)

	return id
}

func DeregisterClient(cid int) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for i, client := range clients {
		if client.Cid == cid {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func RenameClient(cid int, name string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for _, client := range clients {
		if client.Cid == cid {
			client.Label = name
			break
		}
	}
}

func GetClientList() []Client {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	list := make([]Client, len(clients))
	for i, client := range clients {
		list[i] = *client
	}
	return list
}
