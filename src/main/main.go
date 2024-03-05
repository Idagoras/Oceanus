package main

import (
	"bluesell/src/Server"
	"bluesell/src/config"
	db "bluesell/src/database"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
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
	err = server.Start(appConfig.ServerAddress + appConfig.Port)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
