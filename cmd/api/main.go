package main

import (
	"log"

	"github.com/VolodimirKorpan/go_kobi/config"
	"github.com/VolodimirKorpan/go_kobi/server"
)

func main() {
	cfg := config.Get()

	app, err := server.NewApp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err = app.Run(cfg.HTTPPort); err != nil {
		log.Fatalf("%s", err.Error())
	}
}