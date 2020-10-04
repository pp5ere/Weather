package repository

import (
	"Weather/entity"
	"Weather/middlewares/util"
	"Weather/repository/log"
	"context"
	"database/sql"
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

//Insert inserts a new weather
func (c *SqliteDB) Insert(w *entity.Weather) error {
	connection := c.connection
	addColumnIfNotExist("hi", c, updateHeatIndex)
	addColumnIfNotExist("dewpoint", c, updateDewPoint)
	dp := util.DewPoint(w.TempC,w.Hum)
	hiC := util.FahrenheitToCelsius(util.CalculateHeatIndex(w.TempF,w.Hum))
	res, err := connection.Exec(`insert into weather(data, tempc, tempf, hum, pres, alt, hi, dewpoint) values(?,?,?,?,?,?,?,?)`, time.Now(),w.TempC,w.TempF,w.Hum,w.Pres,w.Alt,hiC, dp);if err != nil {
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
	var weathers []*entity.Weather
	connection := c.connection
	rows, err := connection.Query("SELECT * FROM weather where date(data, 'localtime') = date(?) order by data asc", d); if err != nil {
		return nil, err
	}
	for rows.Next(){
		var w entity.Weather
		err = rows.Scan(&w.ID, &w.Data, &w.TempC, &w.TempF, &w.Hum, &w.Pres, &w.Alt, &w.Hi, &w.DewPoint)
		weathers = append(weathers, &w)
	}
	defer rows.Close()
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
	msg = "Execute insert into weather: ID: " + strconv.FormatInt(id, 10) +
			 " TempC = " + strconv.FormatFloat(w.TempC, 'f', 2, 64) + 
			 " TempF: " + strconv.FormatFloat(w.TempF, 'f', 2, 64) + 
			 " Hum: " + strconv.FormatFloat(w.Hum,'f',2,64) + 
			 " Pres: " + strconv.FormatFloat(w.Pres,'f',2,64) + 
			 " Alt: " + strconv.FormatFloat(w.Alt,'f',2,64)+
			 " Hi: " + strconv.FormatFloat(w.Hi,'f',2,64)+
			 " DewPoint: " + strconv.FormatFloat(w.DewPoint,'f',2,64)
	return log.WriteLog(msg)
}

func addColumnIfNotExist(colName string, c *SqliteDB, f func(*sql.Tx)error) (*sql.Tx, error) {
	var exist bool
	var tx *sql.Tx
	connection := c.connection
	col := "%"+colName+"%"
	err := connection.QueryRow(`SELECT EXISTS (SELECT * FROM sqlite_master WHERE tbl_name = 'weather' AND sql LIKE ?)as exist`, col).Scan(&exist); if err != nil {
		log.WriteLog(err.Error())
		return tx, err
	}
	if !exist  {
		ctx := context.Background()
		tx, err = connection.BeginTx(ctx, nil); if err != nil {
			log.WriteLog(err.Error())
			return tx, err
		}
		_, err = tx.ExecContext(ctx, `alter table weather add column ` + colName + ` float;`);if err != nil {
			tx.Rollback()
			log.WriteLog(err.Error())
			return tx, err
		}
			
		//err = updateHeatIndex(tx);if err != nil {
		err = f(tx);if err != nil {
			tx.Rollback()
			log.WriteLog(err.Error())
			return tx, err
		}
		err = tx.Commit();if err != nil {
			log.WriteLog(err.Error())
			return tx, err
		}
	}
	
	return tx, err
}

func updateHeatIndex(c *sql.Tx) error{
	rows, err := c.Query(`select id, tempf, hum from weather where hi is null order by id;`);if err != nil {
		log.WriteLog(err.Error())
		return err
	}
	for rows.Next(){
		var w entity.Weather
		err := rows.Scan(&w.ID, &w.TempF, &w.Hum);if err != nil {
			log.WriteLog(err.Error())
			return err
		}
		hiC := util.FahrenheitToCelsius(util.CalculateHeatIndex(w.TempF, w.Hum))
		_, err = c.Exec("update weather set hi = ? where id=?;", hiC, w.ID);if err != nil {
			log.WriteLog(err.Error())
			return err
		}
	}
	return err
}

func updateDewPoint(c *sql.Tx) error{
	rows, err := c.Query(`select id, tempc, hum from weather where dewpoint is null order by id;`);if err != nil {
		log.WriteLog(err.Error())
		return err
	}
	for rows.Next(){
		var w entity.Weather
		err := rows.Scan(&w.ID, &w.TempC, &w.Hum);if err != nil {
			log.WriteLog(err.Error())
			return err
		}
		dp := util.DewPoint(w.TempC,w.Hum)
		_, err = c.Exec("update weather set dewpoint = ? where id=?;", dp, w.ID);if err != nil {
			log.WriteLog(err.Error())
			return err
		}
	}
	return err
}