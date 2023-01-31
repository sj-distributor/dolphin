package main

import (
	"fmt"

	"github.com/sj-distributor/dolphin/cmd"
)

func main() {
	fmt.Println("123")
	cmd.Execute()
}

// this is just for importing the events package and adding it to the go modules
// var _ events.EventController
