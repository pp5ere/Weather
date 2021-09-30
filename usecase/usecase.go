package usecase

import (
	"Weather/entity"
	"Weather/helper"
	"Weather/middlewares/util"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

//GetDataFromIoT get weather data from IoT webservice
func GetDataFromIoT(c * helper.Config) (*entity.Weather, error) {
	var weather entity.Weather
	client := &http.Client{
		Timeout : 20 * time.Second,
	}
	req, err := http.NewRequest("GET", c.URLFromIoTWebService, nil); if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")	
	resp, err := client.Do(req); if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf.Bytes(), &weather);if err != nil {
		return nil, err
	}
	weather.Hi = util.FahrenheitToCelsius(util.CalculateHeatIndex(weather.TempF, weather.Hum))
	weather.DewPoint = util.DewPoint(weather.TempC,weather.Hum)
	defer resp.Body.Close()
	return &weather,nil
}

