package main

import (
	"database/sql"
	"log"

	"github.com/jasonliuhe/simplebank/api"
	db "github.com/jasonliuhe/simplebank/db/sqlc"
	"github.com/jasonliuhe/simplebank/util"
	_ "github.com/lib/pq"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("cannot create server.", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}