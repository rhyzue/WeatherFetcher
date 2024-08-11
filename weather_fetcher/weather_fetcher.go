package weather_fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
}

const API_KEY = "4dec8f8f0cc99ebd610e4b912745e4bd"

var (
	openweather_path_prefix = "https://api.openweathermap.org/"
)

// mustEmbedUnimplementedWeatherFetcherServer implements WeatherFetcherServer.
func (s *Server) mustEmbedUnimplementedWeatherFetcherServer() {
	panic("unimplemented")
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func (s *Server) GetLocation(ctx context.Context, input *StringValue) (*LocationOptions, error) {
	log.Println("Recieved request for location")

	cityName := input.GetValue()
	request_path := fmt.Sprintf("%sgeo/1.0/direct?q=%s&limit=5&appid=%s", openweather_path_prefix, cityName, API_KEY)
	resp, err := http.Get(request_path)
	if err != nil {
		fmt.Printf("Error making request to %s: %s", request_path, err)
		return nil, nil
	}
	defer resp.Body.Close()

	var response_body []LocationAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&response_body)
	if err != nil {
		fmt.Println("Error parsing response")
		return nil, nil
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
	log.Println("Recieved request for current weather")

	request_path := fmt.Sprintf("%sdata/2.5/weather?q=Toronto&appid=%s", openweather_path_prefix, API_KEY)

	resp, err := http.Get(request_path)
	if err != nil {
		fmt.Printf("Error making request to %s: %s", request_path, err)
		return nil, nil
	}
	defer resp.Body.Close()

	response_body := CurrentWeatherAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response_body)
	if err != nil {
		fmt.Println("Error parsing response")
		return nil, nil
	}

	if len(response_body.Weather) == 0 {
		fmt.Println("Error: Unexpected response")
		fmt.Println(PrettyPrint(response_body))
		return nil, nil
	}

	weather := Weather{
		Name:        response_body.Weather[0].Main,
		Description: response_body.Weather[0].Description,
	}

	fmt.Printf("Got result - Name: %s, Description:%s", weather.Name, weather.Description)

	return &weather, nil
}
