package usecase

import (
	"github.com/gorilla/mux"
	"Weather/controller"
)


func (a *API) endPoints()  {
	a.Router.HandleFunc("/weather", a.findAll)
	a.Router.HandleFunc("/weather/{data}", a.findByData)
	a.Router.HandleFunc("/maxmintemp/{data}", a.findMaxMinTempCPerDay)
	a.Router.HandleFunc("/iotdata", a.findDataFromIoT)
}

//Initialize the route
func Initialize(w *controller.Controllers) *mux.Router {
	thisRoute := mux.NewRouter()
	thisAPI := &API{Router: thisRoute,Weather: w.Weather}
	thisAPI.endPoints()	
	return thisRoute
}