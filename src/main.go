package main

import (
	"fmt"

	"livechat.com/lc-roler/routing"
)

func main() {
	fmt.Println("Server running")
	routing.StartHttpServer()
}
