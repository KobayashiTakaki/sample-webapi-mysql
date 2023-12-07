package main

import (
	"fmt"

	"github.com/KobayashiTakaki/sample-webapi-mysql/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()
	fmt.Println(s.Serve(":8080"))
}
