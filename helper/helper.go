package helper

import (
	"Weather/middlewares/util"
	"encoding/json"
	"log"
	"os"
)

// Config is a helper to open json config file
type Config struct{
	URLFromIoTWebService string `json:"URLFromIoTWebService"`
    PathLog string				`json:"PathLog"`
	PathSqliteDB string			`json:"PathSqliteDB"`
	APIPort string				`json:"APIPort"`
	APIHost string				`json:"APIHost"`
	ReactAppFolder string		`json:"ReactAppFolder"`
	ReactAppPort string			`json:"ReactAppPort"`
}

//LoadFromConfigFile loads json config file
func LoadFromConfigFile() (*Config, error) {
	var c Config
	file, err := loadJSONFile(util.RootDir + "config.json");if err != nil {
		return nil, err
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)	
	err = jsonParser.Decode(&c);if err != nil {
		log.Fatal(err)
	}
	c.PathLog = util.RootDir + c.PathLog
	c.PathSqliteDB = util.RootDir + c.PathSqliteDB
	c.APIPort = ":" + c.APIPort

	return &c, nil
}

//LoadFile returns a new file or open a exists file
func LoadFile(fileName string) (*os.File, error) {
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	return os.Create(fileName)
}

func loadJSONFile(fileName string) (*os.File, error) {
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return os.Open(fileName)
	}
	return os.Create(fileName)
}