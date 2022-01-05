package main

import (
	"fmt"
)

func initializeWebServer() {
	err := addUserToDatabase("user@test.com")
	if err != nil {
		fmt.Println(err)
	}
}
