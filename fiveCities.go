package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherResponse struct {
	Main struct {
		Temperature float64 `json:"temp"`     // Temperature in Kelvin
		Humidity    int     `json:"humidity"` // Humidity in %
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"` // Wind speed in meters per second
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"` // Cloudiness in %
	} `json:"clouds"`
}

func getWeather(cityID string, apiKey string) (WeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?id=%s&appid=%s", cityID, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	var weather WeatherResponse
	err = json.Unmarshal(data, &weather)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weather, nil
}

func main() {
	apiKey := "GitGuardian"

	cities := map[string]string{
		"Nukus":     "601294",
		"Tashkent":  "1512569",
		"Samarkand": "1216265",
		"Bukhara":   "1217662",
		"Andijan":   "1217662",
	}

	for city, cityID := range cities {
		weather, err := getWeather(cityID, apiKey)
		if err != nil {
			fmt.Printf("Error getting weather for %s: %v\n", city, err)
			continue
		}

		temperatureCelsius := weather.Main.Temperature - 273.15 // Convert Kelvin to Celsius
		humidity := weather.Main.Humidity
		windSpeed := weather.Wind.Speed
		cloudiness := weather.Clouds.All

		fmt.Printf("Weather in %s:\n", city)
		fmt.Printf("Temperature: %.2f °C\n", temperatureCelsius)
		fmt.Printf("Humidity: %d %%\n", humidity)
		fmt.Printf("Wind Speed: %.2f m/s\n", windSpeed)
		fmt.Printf("Cloudiness: %d %%\n", cloudiness)
		fmt.Println("--------------")
	}
}
