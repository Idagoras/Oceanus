package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	grpcserver "oceanus/gsrc/server"
	"oceanus/pb"
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
	runGrpcServer(appConfig, store)
}

func runGinServer(config config.Config, store db.Store) {
	server, err := Server.NewServer(config, store)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}

func runGrpcServer(config config.Config, store db.Store) {
	server, err := grpcserver.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOceanusServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("connot create listener")
	}
	log.Printf("start gRpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRpc server")
	}
}
