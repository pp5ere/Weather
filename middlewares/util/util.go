package util

import (
	"log"
	"math"
	"os"
	"path/filepath"
	"time"
)

var (
	//RootDir defines where the binary is running
	RootDir = getRootDir()
)

const (
	//DriveSqlite3DB defines the drive of database
	DriveSqlite3DB = "sqlite3"
	//Success Code for response 
	Success = 1
	//Empty Code for response
	Empty 	= 2
	//Error Code for response
	Error   = 3
	//Unauth Code for response
	Unauth = 9
	
)

func getRootDir() string {
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
	}	
	return rootDir + "/"
}

//StrToDate parse string date to time type
func StrToDate(dateStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", dateStr);if err != nil {
		return t, err
	}
	return t, err
}

//CalculateHeatIndex calculate the Heat Index
func CalculateHeatIndex(tempF float64, humRel float64) float64{
	Hif := 0.5 * (tempF + 61.0 + ((tempF-68.0)*1.2) + (humRel*0.094))
	if Hif >= 80{
		Hif = -42.379 + 2.04901523*tempF + 10.14333127*humRel - .22475541*tempF*humRel - .00683783*tempF*tempF - .05481717*humRel*humRel + .00122874*tempF*tempF*humRel + .00085282*tempF*humRel*humRel - .00000199*tempF*tempF*humRel*humRel
		if humRel < 13 && (tempF >=80 && tempF <= 112){
			Adj := ((13-humRel)/4.0)*math.Sqrt((17-math.Abs(tempF-95.0))/17.0)
			Hif = Hif - Adj
		}
		if humRel > 85 && (tempF >=80 && tempF <= 87){
			Adj := ((humRel-85)/10.0) * ((87-tempF)/5.0)
			Hif = Hif + Adj
		}
	}

	Hif = math.Round(Hif*100)/100.0
	return Hif
}

//FahrenheitToCelsius convert Fahrenheit temperature to Celsius temperature
func FahrenheitToCelsius(tempF float64) float64{
	tempC := (tempF - 32) * 5/9.0
	return math.Round(tempC*100)/100.0
}

//CelsiusToFahrenheit convert Celsius temperature to Fahrenheit temperature
func CelsiusToFahrenheit(tempC float64) float64{
	tempF := (tempC * 9/5.0) + 32
	return math.Round(tempF*100)/100.0
}

//DewPoint calculate Dew Point in Celsius
func DewPoint(tempC, hum float64) float64{
	DP := (math.Pow((hum/100.0),(1/8.0))*(112+0.9*tempC))+(0.1*tempC)-112
	return math.Round(DP*100)/100.0
}
