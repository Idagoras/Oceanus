package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"oceanus/src/Server"
	"oceanus/src/config"
	db "oceanus/src/database"
)

func main() {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(appConfig.DBDriver, appConfig.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := Server.NewServer(appConfig, store)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = server.Start(appConfig.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
