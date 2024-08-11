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

func getUserInput(prompt string) string {
	var res string
	fmt.Print(prompt)
	fmt.Scan(&res)

	if res == "quit" {
		os.Exit(0)
	}
	return res
}

func requestCurrentWeather(ctx context.Context, c pb.WeatherFetcherClient) {
	cityName := getUserInput("Enter city name: ")

	locations, err := c.GetLocation(ctx, &pb.StringValue{Value: cityName})
	if err != nil || len(locations.GetLocations()) == 0 {
		log.Fatalf("could not find city: %v", err)
		return
	}

	fmt.Println("Found the following results: ")
	for i, location := range locations.GetLocations() {
		fmt.Printf("[%d]: %s, %s - (%v, %v)\n", i+1, location.City, location.Country, location.Latitude, location.Longitude)
	}

	cityInput := getUserInput("Confirm city selection: ")

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

	weather, err := c.GetWeather(ctx, &weatherInput)
	if err != nil {
		log.Fatalf("could not get weather: %v", err)
	}

	fmt.Printf("Got weather for %s: \n", cityName)
	fmt.Printf("%+v\n", weather)
}

func main() {

	//get user inputs
	api_key := getUserInput("Enter your api key: ")

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
		fmt.Println("--------------------------------")
		fmt.Println("Welcome to WeatherFetcher. Enter 'quit' to exit the application.")
		fmt.Println("The following operations are supported: ")
		fmt.Println("[1] - Retrieve current weather")
		fmt.Println("--------------------------------")
		user_opt := getUserInput("Select an operation: ")

		switch user_opt {
		case "quit":
			return
		case "1":
			requestCurrentWeather(ctx, c)
		}
	}

}
