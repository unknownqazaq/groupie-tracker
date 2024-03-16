package main

import (
	grp_trk "groupie-tracker"
	"log"
)

func main() {
	srv := new(grp_trk.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
