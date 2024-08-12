package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	pb "github.com/rhyzue/WeatherFetcher/weather_fetcher"
	"google.golang.org/grpc/metadata"
)

var (
	openweather_path_prefix = "https://api.openweathermap.org/"
)

type Coordinates struct {
	X, Y float64
}

type Cache struct {
	data  map[Coordinates]*pb.Weather
	mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[Coordinates]*pb.Weather),
	}
}

func (c *Cache) Get(key Coordinates) (*pb.Weather, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, found := c.data[key]
	return value, found
}

func (c *Cache) Set(key Coordinates, value *pb.Weather) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
}

type weatherFetcherServer struct {
	pb.UnimplementedWeatherFetcherServer
	cache *Cache
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func HttpGetRequest(ctx context.Context, path string) (*http.Response, error) {
	var api_key string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		key := md.Get("api-key")
		if len(key) == 0 || key[0] == "" {
			return nil, errors.New("no api key received")
		}
		api_key = key[0]
	}

	request_path := fmt.Sprintf("%s&appid=%s", path, api_key)

	fmt.Printf("Making request to: %s\n", request_path)

	resp, err := http.Get(request_path)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *weatherFetcherServer) GetLocation(ctx context.Context, input *pb.StringValue) (*pb.LocationOptions, error) {

	cityName := input.GetValue()
	request_path := fmt.Sprintf("%sgeo/1.0/direct?q=%s&limit=5", openweather_path_prefix, cityName)

	resp, err := HttpGetRequest(ctx, request_path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response_body []LocationAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&response_body)
	if err != nil {
		return nil, err
	}

	var location_options pb.LocationOptions
	for _, option := range response_body {
		location_options.Locations = append(location_options.Locations, &pb.Location{
			City:      option.Name,
			Latitude:  option.Lat,
			Longitude: option.Lon,
			Country:   option.Country,
		})
	}

	return &location_options, nil
}

func (s *weatherFetcherServer) GetWeather(ctx context.Context, location *pb.Location) (*pb.Weather, error) {
	cache_key := Coordinates{X: location.Latitude, Y: location.Longitude}
	current_timestamp := int32(time.Now().Unix())
	if cachedData, found := s.cache.Get(cache_key); found {
		fmt.Printf("cached data timestamp: %d, current data timestamp: %d\n", cachedData.Timestamp, current_timestamp)
		if cachedData.Timestamp >= current_timestamp-600 {
			fmt.Println("Returning cached data.")
			return cachedData, nil
		}
		fmt.Println("Data in cache has expired.")
	} else {
		fmt.Println("Data not found in cache.")
	}

	request_path := fmt.Sprintf("%sdata/2.5/weather?lat=%v&lon=%v", openweather_path_prefix, location.Latitude, location.Longitude)

	resp, err := HttpGetRequest(ctx, request_path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response_body := CurrentWeatherAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response_body)
	if err != nil {
		return nil, err
	}

	if len(response_body.Weather) == 0 {
		fmt.Println(PrettyPrint(response_body))
		return nil, errors.New("unexpected response, no weather returned")
	}

	weather := pb.Weather{
		Name:        response_body.Weather[0].Main,
		Description: response_body.Weather[0].Description,
		Temperature: response_body.Main.Temp,
		FeelsLike:   response_body.Main.FeelsLike,
		Pressure:    int32(response_body.Main.Pressure),
		Humidity:    int32(response_body.Main.Humidity),
		Location:    location,
		Timestamp:   int32(response_body.Dt),
	}

	s.cache.Set(cache_key, &weather)
	return &weather, nil
}

func (s *weatherFetcherServer) GetForecast(location *pb.Location, stream pb.WeatherFetcher_GetForecastServer) error {
	request_path := fmt.Sprintf("%sdata/2.5/forecast?lat=%v&lon=%v", openweather_path_prefix, location.Latitude, location.Longitude)

	resp, err := HttpGetRequest(stream.Context(), request_path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	response_body := ForecastAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response_body)
	if err != nil {
		return err
	}

	for _, forecast_data := range response_body.List {
		weather := pb.Weather{
			Name:        forecast_data.Weather[0].Main,
			Description: forecast_data.Weather[0].Description,
			Temperature: forecast_data.Main.Temp,
			FeelsLike:   forecast_data.Main.FeelsLike,
			Pressure:    int32(forecast_data.Main.Pressure),
			Humidity:    int32(forecast_data.Main.Humidity),
			Location:    location,
			Timestamp:   int32(forecast_data.Dt),
		}
		if err := stream.Send(&weather); err != nil {
			return err
		}
	}
	return nil
}
