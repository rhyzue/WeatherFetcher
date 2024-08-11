package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/rhyzue/WeatherFetcher/weather_fetcher"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func RequestCurrentWeather(ctx context.Context, c pb.WeatherFetcherClient) {
	var cityName string
	fmt.Print("Enter city name: ")
	fmt.Scan(&cityName)

	locations, err := c.GetLocation(ctx, &pb.StringValue{Value: cityName})
	if err != nil || len(locations.GetLocations()) == 0 {
		log.Fatalf("could not find city: %v", err)
		return
	}

	fmt.Println("Found the following results: ")
	for i, location := range locations.GetLocations() {
		fmt.Printf("[%d]: %s, %s - (%v, %v)\n", i+1, location.City, location.Country, location.Latitude, location.Longitude)
	}

	var cityInput string
	fmt.Print("Confirm city selection: ")
	fmt.Scan(&cityInput)

	if cityInput == "quit" {
		os.Exit(0)
	}

	cityOpt, err := strconv.Atoi(cityInput)
	if err != nil || cityOpt < 1 || cityOpt > len(locations.GetLocations()) {
		log.Fatalf("Invalid city selection.")
	}

	loc := locations.GetLocations()[cityOpt-1]
	weatherInput := pb.Location{
		City:      loc.City,
		Longitude: loc.Longitude,
		Latitude:  loc.Latitude,
		Country:   loc.Country,
	}

	r, err := c.GetWeather(ctx, &weatherInput)
	if err != nil {
		log.Fatalf("could not get weather: %v", err)
	}

	fmt.Printf("Weather: %s, Description: %s\n", r.GetName(), r.GetDescription())
}

func main() {

	//get user inputs
	var api_key string
	fmt.Print("Enter your api key: ")
	fmt.Scan(&api_key)

	//setup connection
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client could not connect: %s", err)
		return
	}

	defer conn.Close()

	c := pb.NewWeatherFetcherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "api-key", api_key)

	for {
		var user_opt string
		fmt.Println("--------------------------------")
		fmt.Println("Welcome to WeatherFetcher. Enter 'quit' to exit the application.")
		fmt.Println("The following operations are supported: ")
		fmt.Println("[1] - Retrieve current weather")
		fmt.Println("--------------------------------")
		fmt.Print("Select an operation: ")
		fmt.Scan(&user_opt)

		switch user_opt {
		case "quit":
			return
		case "1":
			RequestCurrentWeather(ctx, c)
		}
	}

}
