package usecase

import (
	"Weather/entity"
	"Weather/middlewares/util"
	"Weather/repository/log"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (a *API) findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	weather, _ := a.Weather.FindAll()
	if weather != nil{
		responseCodeResultWeather(w, util.Success, "Success", weather)	
	}else{
		responseCodeResultWeather(w, util.Empty, "Not Found", nil)
	}
}

func (a *API) findByData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	d := mux.Vars(r)
	dataStr := strings.Trim(d["data"], " ")
	data, err := util.StrToDate(dataStr);if err != nil {
		log.WriteLog(err.Error())
		responseCodeResultWeather(w, util.Error, "Only dates in this format will be accepted: yyyy-mm-dd", nil)
		
	}else{
		weather, _ := a.Weather.FindByDate(data)
		if weather != nil{
			responseCodeResultWeather(w, util.Success, "Success", weather)
		}else{
			responseCodeResultWeather(w, util.Empty, "Not Found", nil)
		}	
	}
}

func (a *API) findMaxMinTempCPerDay(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	d := mux.Vars(r)
	dataStr := strings.Trim(d["data"], " ")
	data, err := util.StrToDate(dataStr);if err != nil {
		log.WriteLog(err.Error())
		responseCodeResultWeather(w, util.Error, "Only dates in this format will be accepted: yyyy-mm-dd", nil)
	}else{
	weather, _ := a.Weather.FindMaxMinTempCPerDay(data)
		if weather != nil{
			responseCodeResultWeatherMaxMin(w, util.Success, "Success", weather)
		}else{
			responseCodeResultWeatherMaxMin(w, util.Empty, "Not Found", nil)
		}	
	}
}

func (a *API) findDataFromIoT(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	weather, _ := GetDataFromIoT(a.ConfigFile)
	weather.Data = time.Now()
	err := json.NewEncoder(w).Encode(weather);	if err != nil {
		log.WriteLog(err.Error())
	}
} 

func responseCodeResultWeather(w http.ResponseWriter, code int, msg string, ws []*entity.Weather) {
	var response ResultWeather
	response.Code = code
	response.Result = msg
	response.Weather = ws
	err := json.NewEncoder(w).Encode(response);	if err != nil {
		log.WriteLog(err.Error())
	}
}

func responseCodeResultWeatherMaxMin(w http.ResponseWriter, code int, msg string, ws []*entity.WeatherMaxMin) {
	var response ResultWeatherMaxMin
	response.Code = code
	response.Result = msg
	response.Weather = ws
	err := json.NewEncoder(w).Encode(response);	if err != nil {
		log.WriteLog(err.Error())
	}
}
