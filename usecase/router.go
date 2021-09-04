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
	a.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(a.ReactAppFolder))))
}

//Initialize the route
func Initialize(w *controller.Controllers) (*mux.Router, error) {
	c, err := helper.LoadFromConfigFile();if err != nil {
		return nil, err
	}
	thisRoute := mux.NewRouter()
	thisAPI := &API{Router: thisRoute,Weather: w.Weather, ReactAppFolder: c.ReactAppFolder}
	thisAPI.endPoints()	
	return thisRoute, nil
}