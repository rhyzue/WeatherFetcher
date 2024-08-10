package main

import (
	"log"
	"net"

	pb "github.com/rhyzue/WeatherFetcher/weather_fetcher"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting Server")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWeatherFetcherServer(grpcServer, &pb.Server{})
	log.Printf("server listening at %v", lis.Addr())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
