package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	zipCode := os.Getenv("ZIP_CODE")

	//First we get the current temperature
	temp := getTemperature(zipCode, apiKey)
	//Then we calculate the color
	lowTemp := float32(0)
	highTemp := float32(100)
	color := calculateTemperatureColor(temp, lowTemp, highTemp)
	fmt.Println(color)

	//Then we call WLED to set the color
}

type Color struct {
	Red   int
	Green int
	Blue  int
}

type OpenWeatherMapResponse struct {
	Coordinates Coordinate `json:"coord"`
	Weather     []Weather  `json:"weather"`
	Base        string     `json:"base"`
	Main        Main       `json:"main"`
	Visibility  float32    `json:"visibility"`
	Wind        Wind       `json:"wind"`
	Clouds      Clouds     `json:"clouds"`
	Name        string     `json:"name"`
}

type Coordinate struct {
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  float32 `json:"pressure"`
	Humidity  float32 `json:"humidity"`
}

type Wind struct {
	Speed  float32 `json:"speed"`
	Degree float32 `json:"deg"`
	Gust   float32 `json:"gust"`
}

type Clouds struct {
	All float32 `json:"all"`
}

func getTemperature(zipCode string, apiKey string) float32 {
	jsonBody := []byte("")

	url := "https://api.openweathermap.org/data/2.5/weather/?zip=" + zipCode + "&appid=" + apiKey + "&units=imperial"
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var weather OpenWeatherMapResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weather)

	resp.Body.Close()

	fmt.Println(resp)
	temp := weather.Main.Temp
	return temp

}

func calculateTemperatureColor(currTemp float32, lowTemp float32, highTemp float32) Color {

	//Red
	red := int(((currTemp - lowTemp) / (highTemp - lowTemp)) * (255 - 0))

	//Green
	green := 0

	//Blue
	blue := int(-((currTemp-lowTemp)/(highTemp-lowTemp))*(255-0) + 255)

	color := Color{Red: red,
		Green: green,
		Blue:  blue,
	}
	return color
}
