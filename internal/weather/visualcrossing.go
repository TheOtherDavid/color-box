package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type VisualCrossingHistoricalResponse struct {
	Location Location `json:"location"`
}

type Location struct {
	Values []Value `json:"values"`
	Id     string  `json:"id"`
}
type Value struct {
	Temp           float32 `json:"temp"`
	MaxT           float32 `json:"maxt"`
	MinT           float32 `json:"mint"`
	HeatIndex      float32 `json:"heatindex"`
	DateTimeString string  `json:"datetimeStr"`
}

func GetFiveDayAverageLowAndHigh(zipCode string, apiKey string) (avgLow float32, avgHigh float32) {
	currDate := time.Now().Truncate(24 * time.Hour).Format("2006-01-02")
	startDate := time.Now().Truncate(24 * time.Hour).Add(-7 * 24 * time.Hour).Format("2006-01-02")
	//Call VisualCrossing Weather
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/weatherdata/history?aggregateHours=24&combinationMethod=aggregate&startDateTime=" + startDate + "&endDateTime=" + currDate + "&maxStations=-1&maxDistance=-1&contentType=json&unitGroup=us&locationMode=single&key=" + apiKey + "&dataElements=default&locations=" + zipCode
	println(url)

	jsonBody := []byte("")
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var weather VisualCrossingHistoricalResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weather)

	resp.Body.Close()

	fmt.Println(resp)

	var totLowTemp float32
	var totHighTemp float32

	for _, value := range weather.Location.Values {
		totLowTemp = totLowTemp + value.MinT
		totHighTemp = totHighTemp + value.MaxT
	}
	numDays := float32(len(weather.Location.Values))

	avgLowTemp := totLowTemp / numDays
	avgHighTemp := totHighTemp / numDays

	return avgLowTemp, avgHighTemp
}
