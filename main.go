package main

import (
	"log"
	"os"

	"github.com/izaakdale/phoneNormalizer/api"
	db "github.com/izaakdale/phoneNormalizer/db/sqlc"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbConn *sqlx.DB
var store db.Store
var server *api.Server

func init() {

	var err error
	dbConn, err = sqlx.Connect("postgres", "user=root password=secret dbname=phoneNormalizer sslmode=disable")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	store = db.NewStore(dbConn.DB)
	server, err = api.NewServer(store)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()

}

func main() {

	// here we want to normalize the numbers
	// get all numbers in db.Numbers array
	// iterate each phone number and number
	// check if it is numeric, replace non numeric with "'"

}
