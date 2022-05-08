package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("Taking only one flag track or search, but found: %s", args)
	}

	if args[0] == "track" {
		TryTrack()
	} else if args[0] == "search" {
		TryGlobalSearch()
	} else {
		log.Fatalf("Do not recognise command %s", args[0])
	}
}