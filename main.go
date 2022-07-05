package main

import (
	"log"
	"net/http"
	"os"
	"school/config"
	"school/pkg"

	"github.com/gorilla/mux"
)

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "database":
			config.MigrateDB(config.GetDB())
			log.Print("\n Database loaded...")
			os.Exit(1)
		default:
			log.Print("\n Starting server....")
		}

	}

	router := mux.NewRouter()
	pkg.Routes(router)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Server started at :8000")
	}

}
