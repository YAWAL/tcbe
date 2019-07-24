package app

import (
	"log"
	"net/http"

	"github.com/YAWAL/tcbe/src/handlers"

	"github.com/YAWAL/tcbe/src/config"
	"github.com/YAWAL/tcbe/src/db"
	"github.com/YAWAL/tcbe/src/routers"
)

func Load() (err error) {

	// read configs
	log.Println("Reading configs")
	conf, err := config.LoadConfig("config.json")
	if err != nil {
		return err
	}

	// connect to Mongo
	log.Println("Connecting to Mongo")
	err = db.Mongo(&conf.DB)
	if err != nil {
		return err
	}

	// fill database with initial data
	err = db.FillInitData(&conf.DB)
	if err != nil {
		return err
	}

	log.Println("Starting the application")
	err = http.ListenAndServe(conf.Port, routers.Router(handlers.Init(&conf.DB)))
	if err != nil {
		return err
	}
	return nil
}
