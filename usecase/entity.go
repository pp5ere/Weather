package usecase

import (
	"Weather/controller"
	"Weather/entity"
	"github.com/gorilla/mux"
)

// API defines api router
type API struct{
	Router *mux.Router
	Weather controller.WeatherController
	ReactAppFolder string
}

//ResultWeather show error json when its happening
type ResultWeather struct{
	Result    string   			`json:"Result"`
	Code	  int	   			`json:"Code,omitempty"`
	Weather   []*entity.Weather	`json:"Weather,omitempty"`
}

//ResultWeatherMaxMin show error json when its happening
type ResultWeatherMaxMin struct{
	Result    string   			`json:"Result"`
	Code	  int	   			`json:"Code,omitempty"`
	Weather   []*entity.WeatherMaxMin	`json:"Weather,omitempty"`
}