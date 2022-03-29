package repository

import (
	"Weather/entity"
	"Weather/repository/log"
	"fmt"
	"strconv"
	"time"
)

//Weather defines the methods must be implemented by injected layer.
type Weather interface{
	FindAll() ([]*entity.Weather, error)
	Insert(w *entity.Weather)  error
	FindByDate(d time.Time) ([]*entity.Weather, error)
	FindMaxMinTempCPerDay(d time.Time) ([]*entity.WeatherMaxMin, error)
}

//Insert add a new weather into db
func (c *SqliteDB) Insert(w *entity.Weather) error {
	
	connection := c.connection
	res, err := connection.Exec(`insert into weather(data, tempc, tempf, hum, pres, alt, hi, dewpoint) values(?,?,?,?,?,?,?,?)`, time.Now(),w.TempC,w.TempF,w.Hum,w.Pres,w.Alt,w.Hi, w.DewPoint);if err != nil {
		return err
	}
	id, err := res.LastInsertId();if err != nil {
		return err
	}
	
	return insertLog(id, w)
}

//FindAll returns all Weather from database order by ID
func (c *SqliteDB) FindAll() ([]*entity.Weather, error) {
	var weathers []*entity.Weather
	
	connection := c.connection
	rows, err := connection.Query("SELECT * FROM weather order by id"); if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var w entity.Weather
		err = rows.Scan(&w.ID, &w.Data, &w.TempC, &w.TempF, &w.Hum, &w.Pres, &w.Alt, &w.Hi, &w.DewPoint)
		weathers = append(weathers, &w) 
	}
	defer rows.Close()
	
	return weathers, err

}

//FindByDate returns all weather from a specific date
func (c *SqliteDB) FindByDate(d time.Time) ([]*entity.Weather, error){
	init := time.Now()
	var weathers []*entity.Weather

	connection := c.connection
	rows, err := connection.Query("SELECT * FROM weather where date(data, 'localtime') = date(?) order by data asc", d); if err != nil {
		return nil, err
	}
	//ch := make(chan entity.Weather)
	end := make(chan bool)	
	go func ()  {
		for rows.Next(){
			var w entity.Weather
			err = rows.Scan(&w.ID, &w.Data, &w.TempC, &w.TempF, &w.Hum, &w.Pres, &w.Alt, &w.Hi, &w.DewPoint)
			//ch <- w
			weathers = append(weathers, &w)
		}
		end <-true
	}()		
	
	/*go func () {			
		for w := range ch {
			nw := w
			weathers = append(weathers, &nw)
		}
		
	}()*/
	
	defer rows.Close()
	<-end
	fin := time.Now().Sub(init)
	fmt.Println(fin)
	
	return weathers, err
}

//FindMaxMinTempCPerDay returns all max temperature per day
func (c *SqliteDB) FindMaxMinTempCPerDay(d time.Time) ([]*entity.WeatherMaxMin, error){
	var weathers []*entity.WeatherMaxMin
	
	connection := c.connection
	rows, err := connection.Query(`select w.data data, min(w.tempC) minTempC, max(w.tempC) maxTempC
									from weather w
									where date(data, 'localtime') >= date(?) and
									date(data, 'localtime') <= date(?)
									group by date(w.data, 'localtime')`, d.AddDate(-1,0,0), d); 
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var w entity.WeatherMaxMin
		err = rows.Scan(&w.Data, &w.MinTempC, &w.MaxTempC)
		weathers = append(weathers, &w)
	}
	defer rows.Close()
	
	return weathers, err
}

func insertLog(id int64, w *entity.Weather) error {
	var msg string
	msg = "Execute insert into weather: ID = " + strconv.FormatInt(id, 10) +
			 " TempC = " + strconv.FormatFloat(w.TempC, 'f', 2, 64) + 
			 " TempF = " + strconv.FormatFloat(w.TempF, 'f', 2, 64) + 
			 " Hum = " + strconv.FormatFloat(w.Hum,'f',2,64) + 
			 " Pres = " + strconv.FormatFloat(w.Pres,'f',2,64) + 
			 " Alt = " + strconv.FormatFloat(w.Alt,'f',2,64)+
			 " Hi = " + strconv.FormatFloat(w.Hi,'f',2,64)+
			 " DewPoint = " + strconv.FormatFloat(w.DewPoint,'f',2,64)
	return log.WriteLog(msg)
}

//CreateTable create table weather if it not exist
func (c *SqliteDB) CreateTable(dbName string) error {
	var exist bool
	
	connection := c.connection
	err := connection.QueryRow(`SELECT EXISTS (SELECT * FROM sqlite_master WHERE tbl_name = ?)as exist`, dbName).Scan(&exist); if err != nil {
		log.WriteLog(err.Error())
		return err
	}
	if !exist{
		_, err = connection.Exec(`CREATE TABLE weather (
									id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
									data DATE NOT NULL,
									tempc float NOT NULL,
									tempf float NOT NULL,
									Hum float NOT NULL,
									Pres float NOT NULL,
									Alt float NOT NULL, 
									hi float, 
									dewpoint float)`)
		if err != nil {
			return err
		}
	}
	
	return nil
}