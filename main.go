package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/luyi404/simplebank/api"
	db "github.com/luyi404/simplebank/db/sqlc"
	"github.com/luyi404/simplebank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("failed to create a server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
