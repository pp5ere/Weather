package usecase

import (
	"Weather/entity"
	"Weather/helper"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

//GetDataFromIoT get weather data from IoT webservice
func GetDataFromIoT() (*entity.Weather, error) {
	var weather entity.Weather
	c, err := helper.LoadFromConfigFile();if err != nil {
		return nil, err
	}
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
	defer resp.Body.Close()
	return &weather,nil
}

