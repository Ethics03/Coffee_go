package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	DB *sql.DB // points to the db
}

var dbConnect = &DB{}

const maxOpenDBConn = 10 // u can change the limits and all so that code dosent break

const maxIdleDBConn = 5

const maxDbLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {

	d, err := sql.Open("pgx", dsn) // this is the *sql.DB
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetMaxIdleConns(maxIdleDBConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dbConnect.DB = d // initialized to call it as D it is the DB inside the db struct
	return dbConnect, nil

}

func testDB(d *sql.DB) error {

	err := d.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	fmt.Println("*** Pinged Database Successfully! ***") // this is pinging the database
	return nil

}
