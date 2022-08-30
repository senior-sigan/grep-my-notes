package main

import (
	"flag"
	"grepmynotes/server"
	"log"
)

func main() {
	portFlag := flag.Int("port", 3000, "Server port to listen to")
	srcFlag := flag.String("src", ".", "Path to the fodler with file")
	flag.Parse()

	err := server.Run(*portFlag, *srcFlag)
	if err != nil {
		log.Fatalf("Cannot run server: %v", err)
	}
}
