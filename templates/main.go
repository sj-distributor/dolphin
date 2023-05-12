package templates

var Main = `package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/cors"
	"github.com/urfave/cli"

	"{{.Config.Package}}/gen"
	"{{.Config.Package}}/src"
)

// const defaultPort = "8080"

func main() {
	app := cli.NewApp()
	app.Name = "dolphin"
	app.Usage = "This tool is for generating "
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		startCmd,
		migrateCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

var startCmd = cli.Command{
	Name:  "start",
	Usage: "start api server",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "cors",
			Usage: "Enable cors",
		},
		cli.StringFlag{
			Name:   "p,port",
			Usage:  "Port to listen to",
			Value:  "80",
			EnvVar: "PORT",
		},
	},
	Action: func(ctx *cli.Context) error {
		cors := ctx.Bool("cors")
		port := ctx.String("port")
		if err := startServer(cors, port); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		return nil
	},
}

var migrateCmd = cli.Command{
	Name:  "migrate",
	Usage: "migrate schema database",
	Action: func(ctx *cli.Context) error {
		fmt.Println("starting migration")
		if err := automigrate(); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		fmt.Println("migration complete")
		return nil
	},
}

func automigrate() error {
	db := gen.NewDBFromEnvVars("")
	return db.AutoMigrate()
}

func startServer(enableCors bool, port string) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	db := gen.NewDBFromEnvVars("")
	defer db.Close()

	eventController, err := gen.NewEventController()
	if err != nil {
		return err
	}

	mux := gen.GetHTTPServeMux(src.New(db, &eventController), db)

	var handler http.Handler
	if enableCors {
		handler = cors.AllowAll().Handler(mux)
	} else {
		handler = mux
	}

	h := &http.Server{Addr: ":" + port, Handler: handler}

	go func() {
		log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
		log.Fatal(h.ListenAndServe())
	}()

	<-stop

	log.Println("\nShutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err = h.Shutdown(ctx)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	log.Println("Server gracefully stopped")

	err = db.Close()
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	log.Println("Database connection closed")

	return nil

}
`
