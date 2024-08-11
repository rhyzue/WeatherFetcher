package weather_fetcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc/metadata"
)

type Server struct {
}

var (
	openweather_path_prefix = "https://api.openweathermap.org/"
)

const API_KEY = "d867e6413489af4d75238bab6cf3cf25"

// mustEmbedUnimplementedWeatherFetcherServer implements WeatherFetcherServer.
func (s *Server) mustEmbedUnimplementedWeatherFetcherServer() {
	panic("unimplemented")
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

	log.Printf("Making request to: %s", request_path)

	resp, err := http.Get(request_path)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) GetLocation(ctx context.Context, input *StringValue) (*LocationOptions, error) {

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

	var location_options LocationOptions
	for _, option := range response_body {
		location_options.Locations = append(location_options.Locations, &Location{
			City:      option.Name,
			Latitude:  option.Lat,
			Longitude: option.Lon,
			Country:   option.Country,
		})
	}

	return &location_options, nil
}

func (s *Server) GetWeather(ctx context.Context, location *Location) (*Weather, error) {

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

	weather := Weather{
		Name:        response_body.Weather[0].Main,
		Description: response_body.Weather[0].Description,
		Temperature: response_body.Main.Temp,
		FeelsLike:   response_body.Main.FeelsLike,
		Pressure:    int32(response_body.Main.Pressure),
		Humidity:    int32(response_body.Main.Humidity),
	}
	return &weather, nil
}
