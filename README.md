# color-box

This is a simple app to represent the current temperature in a given zip code as an RGB color ranging from 0-0-255 (pure blue) and 255-0-0 (pure red), and then send that color to a device running WLED, so the device lights up according to the current temperature.

The current formula for temperature is based on a scale from 0 to 100. Future improvements will explore the comparison of the current temperature to historical averages for that month, or possibly a comparison to the temperatures over the past few days.

The weather information for this app comes from OpenWeatherMap, and requires an API key.
https://openweathermap.org/api

The output is a command to a device running WLED.
https://github.com/Aircoookie/WLED
