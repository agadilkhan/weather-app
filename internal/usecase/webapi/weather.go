package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/evrone/go-clean-template/internal/entity"
	"net/http"
)

type WeatherWebApi struct {
	apiKey string
}

func New(apiKey string) *WeatherWebApi {
	return &WeatherWebApi{
		apiKey: apiKey,
	}
}

const NotFound = "404 Not Found"

type WeatherDataResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

func (w *WeatherWebApi) Get(ctx context.Context, city string) (entity.WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, w.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return entity.WeatherData{}, fmt.Errorf("failed to Get err: %v", err)
	}

	if resp.Status == NotFound {
		return entity.WeatherData{}, fmt.Errorf("not found")
	}

	defer resp.Body.Close()

	var data WeatherDataResponse
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return entity.WeatherData{}, fmt.Errorf("failed to Decode err: %v", err)
	}

	res := entity.WeatherData{
		City:      data.Name,
		Temp:      data.Main.Temp,
		TempMax:   data.Main.TempMax,
		TempMin:   data.Main.TempMin,
		FeelsLike: data.Main.FeelsLike,
		Humidity:  data.Main.Humidity,
		Pressure:  data.Main.Pressure,
		WindSpeed: data.Wind.Speed,
	}

	return res, nil
}
