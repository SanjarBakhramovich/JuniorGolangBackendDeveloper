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
}

func main() {
	apiKey := "GitGuardian"
	cityName := "Nukus" // Replace with the name of your hometown

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityName, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var weather WeatherResponse
	err = json.Unmarshal(data, &weather)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	temperatureCelsius := weather.Main.Temperature - 273.15 // Convert Kelvin to Celsius
	humidity := weather.Main.Humidity

	fmt.Printf("Temperature in %s: %.2f °C\n", cityName, temperatureCelsius)
	fmt.Printf("Humidity in %s: %d %%\n", cityName, humidity)
}
