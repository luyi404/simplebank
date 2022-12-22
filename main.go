package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/luyi404/simplebank/api"
	db "github.com/luyi404/simplebank/db/sqlc"
	"github.com/luyi404/simplebank/gapi"
	"github.com/luyi404/simplebank/pb"
	"github.com/luyi404/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
	runGrpcServer(config, store)
	//runGinServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("failed to create a server: ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	log.Printf("server is listening at %s", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot serve: ", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("failed to create a server: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
