package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rhyzue/WeatherFetcher/weather_fetcher"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client could not connect: %s", err)
	}

	defer conn.Close()

	c := pb.NewWeatherFetcherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetWeather(ctx, &pb.City{Name: "Toronto"})
	if err != nil {
		log.Fatalf("could not get weather: %v", err)
	}

	log.Printf("Weather: %s, Description: %s", r.GetName(), r.GetDescription())

}
