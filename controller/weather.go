package controller

import (
	"Weather/entity"
	"Weather/repository"
	"time"
)

type (
	// Controllers contains the Controllers for each Entity.
	Controllers struct {
		Weather repository.Weather
	}
	// Weather contains the injected Weather interface from Repository layer.
	Weather struct {
		Repository repository.Weather
	}
)
// WeatherController contains methods that must be implemented by the injected layer.
type WeatherController interface{
	FindAll() ([]*entity.Weather, error)
	Insert(w *entity.Weather) error
	FindByDate(d time.Time) ([]*entity.Weather, error)
	FindMaxMinTempCPerDay(d time.Time) ([]*entity.WeatherMaxMin, error)
}

// New creates new Controllers for each Entity.
func New(repo *repository.SqliteDB) *Controllers{
	return &Controllers{
		Weather: newWeatherController(repo),
	}
}

// newWeatherController creates a new Weather Controller.
func newWeatherController(e *repository.SqliteDB) *Weather{
	return &Weather{
		Repository: e,
	}
}

// FindAll requests the Repository layer to return all Weather from database.
func (w *Weather) FindAll() ([]*entity.Weather, error) {
	return w.Repository.FindAll()
}

//Insert requests the Repository layer for the insertion of a new Weather in the database.
func (w *Weather) Insert(weather *entity.Weather) error {
	return w.Repository.Insert(weather)
}

//FindByDate requests the Repository layer to return all Weather specific date from database.
func (w *Weather) FindByDate(d time.Time) ([]*entity.Weather, error){
	return w.Repository.FindByDate(d)
}

//FindMaxMinTempCPerDay requests the Repository layer to return all max temperature Weather per date from database.
func (w *Weather) FindMaxMinTempCPerDay(d time.Time) ([]*entity.WeatherMaxMin, error){
	return w.Repository.FindMaxMinTempCPerDay(d)
}

