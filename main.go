package main

import (
	"fmt"
)

func main() {
	initializeMongodb()
	initializeSendGrid()
	initializeWebServer()
	fmt.Println("DONE!")
}
