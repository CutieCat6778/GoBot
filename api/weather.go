package api

import (
	"cutiecat6778/discordbot/class"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Weather struct {
	HttpClient *http.Client
}

var (
	CurrentWeatherURL string = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&units=%v&appid=" + class.WeatherKey
	RoadRiskURL       string = "https://api.openweathermap.org/data/2.5/roadrisk?lat=%v&lon=%v&dt=%v&appid=" + class.WeatherKey
	ForecastURL       string = "https://api.openweathermap.org/data/2.5/forecast?lat=%v&lon=%v&units=%v&appid=" + class.WeatherKey
)

func NewWeather() Weather {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return Weather{
		HttpClient: &client,
	}
}

func UnitsConverter(units string) string {
	if units == "celsius" {
		return "metric"
	} else {
		return "imperial"
	}
}

func (handler Weather) GetCurrentWeather(lat float64, long float64, units string) CurrentWeatherStruct {

	url := fmt.Sprintf(CurrentWeatherURL, lat, long, UnitsConverter(units))

	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch: ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read: ", err)
	}

	defer resp.Body.Close()

	res := CurrentWeatherStruct{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Error while formating json: ", err)
	}

	return res
}

func (handler Weather) GetRoadRisk(lat float64, long float64) RoadRiskStruct {
	url := fmt.Sprintf(CurrentWeatherURL, lat, long, time.Now().Unix())

	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch: ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read: ", err)
	}

	defer resp.Body.Close()

	res := RoadRiskStruct{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Error while formating json: ", err)
	}

	return res
}

func (handler Weather) GetForecast(lat float64, long float64, units string) CurrentWeatherStruct {

	url := fmt.Sprintf(CurrentWeatherURL, lat, long, UnitsConverter(units))

	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch: ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read: ", err)
	}

	defer resp.Body.Close()

	res := CurrentWeatherStruct{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal("Error while formating json: ", err)
	}

	return res
}

func URLConverter(id string) string {
	url := "http://openweathermap.org/img/wn/%v@4x.png"

	url = fmt.Sprintf(url, id)
	return url
}
