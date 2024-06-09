package main

import (
	"github.com/kwamekyeimonies/nba-mobler-engine/api"
	"github.com/kwamekyeimonies/nba-mobler-engine/internal"
	"log"
)

func main() {

	err := internal.InitializeDependencies()
	if err != nil {
		log.Printf("error: %v", err)
	}

	err = api.GinRestAPIServer()
	if err != nil {
		log.Printf("error: %v", err)
	}

}
