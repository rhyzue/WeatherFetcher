# WeatherFetcher
Application to fetch data from OpenWeather API using golang and gRPC.

## Features:
- CLI with interactive prompts.
- User can select city. Uses OpenWeatherAPI's geocoding API to retrieve list of matching cities, allowing user to confirm selection.
- Fetch the current weather of a specified city.
- Server-side streaming of 5-day forecast data (3-hour step) for a specified data.
- Supports caching for current weather data. As per OpenWeatherAPI documentation, cached data is only valid for 10 minutes.

## How to Run
### Start up Server
1. Open new terminal tab
2. `cd server`
3. `go run .`
4. Check back on this tab after running client commands for logged info.
### Start up Client
1. Open new terminal tab
2. `cd client`
3. `go run client`
4. Follow prompts.
