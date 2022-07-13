package main

import (
	"context"
	"log"
	"os"
	"school/app"
	"school/config"
	"school/pkg"
)

func main() {
	pkg.RandomString()
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
	type msg struct {
		str string
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "msg", msg{str: "Hello from main...."})

	container := app.InitContainer(ctx)
	container.LoadRoutes()
	container.HttpServe(ctx, "8000")

}
