package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"tgBot/token"
	"time"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func KnowWeather(city string) (string, error) {
	req, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + token.WEATHER_API + "&q=" + city + "&days=1&aqi=no&alerts=no")
	if err != nil {
		return "", errors.New("try different city")
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", errors.New("weather api not available")
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return "", errors.New("error while unmarshalling")
	}
	if len(weather.Forecast.Forecastday) < 1 {
		return "", errors.New("wrong city")
	}
	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	weatherNow := fmt.Sprintf(
		"%s, %s, %.0f°C, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
	weatherFuture := ""
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if !date.Before(time.Now()) {
			weatherFuture += fmt.Sprintf(
				"%s - %.0f°C, Rain: %.0f%%, %s\n",
				date.Format("15:04"),
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)
		}
	}
	result := weatherNow + weatherFuture
	return result, nil
}
