package main

import (
	"context"
	"fmt"
	"io"
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

func requestLocation(ctx context.Context, c pb.WeatherFetcherClient) *pb.Location {
	cityName := getUserInput("Enter city name: ")

	locations, err := c.GetLocation(ctx, &pb.StringValue{Value: cityName})
	if err != nil || len(locations.GetLocations()) == 0 {
		log.Fatalf("could not find city: %v", err)
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

	return locations.GetLocations()[cityOpt-1]
}

func printWeather(weather *pb.Weather) {
	fmt.Printf("Weather for %s, %s, timestamp %d: \n", weather.Location.City, weather.Location.Country, weather.Timestamp)
	fmt.Printf("\tname: %s\n", weather.Name)
	fmt.Printf("\tdescription: %s\n", weather.Description)
	fmt.Printf("\ttemperature: %v\n", weather.Temperature)
	fmt.Printf("\tfeels_like: %v\n", weather.FeelsLike)
	fmt.Printf("\tpressure: %d\n", weather.Pressure)
	fmt.Printf("\thumidity: %d\n", weather.Humidity)
}

func requestCurrentWeather(ctx context.Context, c pb.WeatherFetcherClient) {
	loc := requestLocation(ctx, c)

	weather, err := c.GetWeather(ctx, loc)
	if err != nil {
		log.Fatalf("could not get weather: %v", err)
	}

	printWeather(weather)
}

func requestForecast(ctx context.Context, c pb.WeatherFetcherClient) {
	loc := requestLocation(ctx, c)

	stream, err := c.GetForecast(ctx, loc)
	if err != nil {
		log.Fatalf("c.GetForecast failed: %v", err)
	}
	fmt.Println("Receiving forecast as stream...")
	for {
		weather, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("c.GetForecast failed: %v", err)
		}
		printWeather(weather)
	}
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
		fmt.Println("[2] - Stream 5-day forecast (3-hour step)")

		fmt.Println("--------------------------------")
		user_opt := getUserInput("Select an operation: ")

		switch user_opt {
		case "quit":
			return
		case "1":
			requestCurrentWeather(ctx, c)
		case "2":
			requestForecast(ctx, c)
		}
	}

}
