package main

import (
	"bytes"
	"fmt"
	"github.com/TheOtherDavid/color-temp/internal/weather"
	"net/http"
	"os"
)

func main() {
	owmApiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	zipCode := os.Getenv("ZIP_CODE")
	visApiKey := os.Getenv("VISUAL_CROSSING_API_KEY")
	programMode := os.Getenv("PROGRAM_MODE")

	//First we get the current temperature
	temp := weather.GetTemperature(zipCode, owmApiKey)
	var lowTemp float32
	var highTemp float32
	//Then we get the high and low limits
	if programMode == "ROLLING_AVERAGE" {
		lowTemp, highTemp = weather.GetFiveDayAverageLowAndHigh(zipCode, visApiKey)
	} else if programMode == "ZERO_TO_ONE_HUNDRED" {
		lowTemp = 0
		highTemp = 100
	}
	//Then we calculate the color
	color := calculateTemperatureColor(temp, lowTemp, highTemp)
	fmt.Println(color)

	//Then we call WLED to set the color
	callWledWithColor(color)
}

type Color struct {
	Red   int
	Green int
	Blue  int
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

func callWledWithColor(color Color) {
	red := color.Red
	green := color.Green
	blue := color.Blue

	wledBody := `
{
    "on": true,
    "bri": 240,
    "transition": 7,
    "ps": -1,
    "pl": -1,
    "ccnf": {
        "min": 1,
        "max": 5,
        "time": 12
    },
    "nl": {
        "on": false,
        "dur": 60,
        "fade": true,
        "mode": 1,
        "tbri": 0,
        "rem": -1
    },
    "udpn": {
        "send": true,
        "recv": true
    },
    "lor": 0,
    "mainseg": 0,
    "seg": [
        {
            "id": 0,
            "start": 0,
            "stop": 600,
            "len": 600,
            "grp": 1,
            "spc": 0,
            "on": true,
            "bri": 255,
            "col": [
                [
                    %d,
                    %d,
                    %d
                ],
                [
                    0,
                    0,
                    0
                ],
                [
                    %d,
                    %d,
                    %d
                ]
            ],
            "fx": 75,
            "sx": 63,
            "ix": 103,
            "pal": 5,
            "sel": false,
            "rev": false,
            "mi": false
        }
    ]
}
`

	formattedBody := fmt.Sprintf(wledBody, red, green, blue, red, green, blue)
	fmt.Println(formattedBody)
	callWledWithJson(formattedBody)

}

func callWledWithJson(body string) {
	jsonBody := []byte(body)
	ipAddress := os.Getenv("WLED_IP_ADDRESS")

	url := "http://" + ipAddress + "/json/state"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}
