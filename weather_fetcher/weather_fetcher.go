package weather_fetcher

import (
	"context"
	"log"
)

type Server struct {
}

// mustEmbedUnimplementedWeatherFetcherServer implements WeatherFetcherServer.
func (s *Server) mustEmbedUnimplementedWeatherFetcherServer() {
	panic("unimplemented")
}

func (s *Server) GetWeather(ctx context.Context, city *City) (*Weather, error) {
	log.Println("Recieved request for weather")
	return &Weather{
		Name:        "Rain",
		Description: "High rainfall.",
	}, nil
}
