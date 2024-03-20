package main

import (
	groupie_tracker "groupie-tracker/internal/server"
	"log"
)

func main() {
	srv := new(groupie_tracker.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
