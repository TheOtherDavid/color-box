# color-temp

This is an app to represent the current temperature in a given zip code as an RGB color ranging from 0-0-255 (pure blue) and 255-0-0 (pure red), and then send that color to a device running WLED, so the device lights up according to the current temperature.

There are multiple program_mode options to calculate the temperature color in various ways:
1: Comparison within a range of 0 to 100
2: Seven-day rolling average of high and low temperatures.
Future improvements will explore the comparison of the current temperature to historical averages for that month

Two weather APIs are currently used. The current temperature information comes from OpenWeatherMap, and requires an API key.
https://openweathermap.org/api
The historical weather information comes from VisualCrossing, and also requires an API key.
https://www.visualcrossing.com/weather-api

The output is a command to a device running WLED.
https://github.com/Aircoookie/WLED
