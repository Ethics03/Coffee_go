package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ethics03/basic/cmd/services"
	"github.com/Ethics03/basic/db"
	"github.com/Ethics03/basic/router"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
	//Models

}

func (app *Application) Serve() error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Println("API is listening on port : ", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe() // to return the listening and serving

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")

	dbConn, err := db.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Cannot connect to Database.")
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()

	if err != nil {
		log.Fatal(err)
	}

}
