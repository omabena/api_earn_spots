package main

import (
	"github.com/omabena/api_earn_spots/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
