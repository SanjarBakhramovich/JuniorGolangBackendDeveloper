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
	apiKey := "75686baa4ffc8898ad84046eaf112d8c"
	cityID := "601294"

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?id=%s&appid=%s", cityID, apiKey)

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

	fmt.Printf("Temperature in Nukus: %.2f °C\n", temperatureCelsius)
	fmt.Printf("Humidity in Nukus: %d %%\n", humidity)
}
