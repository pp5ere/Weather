package entity

import "time"

//Weather is a structure for received data from bmp sensor IoT
type Weather struct {
	ID		 int 		`json:"Id,omitempty"`
	Data 	 time.Time 	`json:"Data,omitempty"`
	TempC 	 float64 	`json:"TempC,omitempty"`
	TempF 	 float64 	`json:"TempF,omitempty"`
	Hi 	  	 float64 	`json:"Hi,omitempty"`
	DewPoint float64	`json:"DewPoint,omitempty"`
	Hum   	 float64 	`json:"Hum,omitempty"`
	Pres  	 float64 	`json:"Pres,omitempty"`
	Alt   	 float64 	`json:"Alt,omitempty"`
	Msg   	 string 	`json:"Msg,omitempty"`
}
//WeatherMaxMin is a structure for max and min temperature
type WeatherMaxMin struct{
	Data 	 time.Time 	`json:"Data,omitempty"`
	MinTempC float64 	`json:"MinTempC,omitempty"`
	MaxTempC float64 	`json:"MaxTempC,omitempty"`
}