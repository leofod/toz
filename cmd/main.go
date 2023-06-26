package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	proto "toz/pkg/proto"
	handler "toz/pkg/url/handler"
	repository "toz/pkg/url/repository"
	service "toz/pkg/url/service"
	postgres "toz/postgres"
	server "toz/server"

	"google.golang.org/grpc"
)

func main() {
	var store bool
	var rep *repository.UrlRepository

	flag.BoolVar(&store, "db", false, "Flag storage")
	flag.Parse()

	if store {
		db, err := postgres.ConnectDB(postgres.Config{
			Host:     "postgres_container",
			Port:     "5432",
			Username: "postgres",
			Password: "qweasd",
			DBName:   "postgres",
			SSLMode:  "disable",
		})
		if err != nil {
			fmt.Printf("No accees to database: %s", err.Error())
		}
		rep = repository.NewRepositoryUrlDB(db)
	} else {
		hm_full := make(map[string]string)
		hm_short := make(map[string]string)
		rep = repository.NewRepositoryUrl(hm_full, hm_short)
	}

	serv := service.NewServiceUrl(rep)
	han := handler.NewHandlerUrl(serv)

	go func() {
		s := grpc.NewServer()
		srv := handler.NewGrpcServer(serv)
		proto.RegisterUrlServer(s, srv)

		l, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal(err)
		}

		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	server := new(server.Server)
	if err := server.Run("8000", han.InitRoutes()); err != nil {
		log.Printf("listen: %s\n", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server user forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
