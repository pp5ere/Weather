package usecase

import (
	"Weather/controller"
	"Weather/helper"
	"net/http"

	"github.com/gorilla/mux"
)


func (a *API) endPoints()  {
	a.Router.HandleFunc("/weather", a.findAll)
	a.Router.HandleFunc("/weather/{data}", a.findByData)
	a.Router.HandleFunc("/maxmintemp/{data}", a.findMaxMinTempCPerDay)
	a.Router.HandleFunc("/iotdata", a.findDataFromIoT)
	a.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(a.ConfigFile.ReactAppFolder))))
}

//Initialize the route
func Initialize(w *controller.Controllers,c *helper.Config) (*mux.Router, error) {
	thisRoute := mux.NewRouter()
	thisAPI := &API{Router: thisRoute,Weather: w.Weather, ConfigFile: c}
	thisAPI.endPoints()	
	return thisRoute, nil
}