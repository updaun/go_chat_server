package main

import (
	"go_chat_server/network"
)

func main() {
	n := network.NewServer()
	n.StartServer()
}
