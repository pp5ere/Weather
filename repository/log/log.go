package log

import (
	"Weather/middlewares/util"
	"fmt"
	"log"
	"os"
	"time"
)

//WriteLog saves messagens into the log.txt
func WriteLog(msg string) error {	
	dirLog := util.RootDir + "log/"
	if _, err := os.Stat(dirLog); os.IsNotExist(err) {
		err := os.MkdirAll(dirLog, 0755);if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(dirLog+getNameLog(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666);if err != nil {		
		return err
	}	
	log.SetOutput(file)
	log.Println(msg)
	defer file.Close()
	
	return err
}

func getNameLog() string {
	y, m, d := time.Now().Date()	
	return fmt.Sprint(y)+"-"+fmt.Sprint(int(m))+"-"+fmt.Sprint(d)+".log"
}