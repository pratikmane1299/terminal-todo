package main

import (
	"log"

	"github.com/pratikmane1299/terminal-todo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal("error occured", err)
	}
}
