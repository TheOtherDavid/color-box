package main

func main() {
	//First we get the current temperature
	//Then we calculate the color
	//Then we call WLED to set the color
}

type Color struct {
	Red   int
	Green int
	Blue  int
}

func calculateTemperatureColor(currTemp int, lowTemp int, highTemp int) Color {

	//Red
	red := ((currTemp - lowTemp) / (highTemp - lowTemp)) * (255 - 0)

	//Green
	green := 0

	//Blue
	blue := -((currTemp-lowTemp)/(highTemp-lowTemp))*(255-0) + 255
	color := Color{Red: red,
		Green: green,
		Blue:  blue,
	}
	return color
}
