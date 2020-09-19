package log

import (
	"Weather/helper"
	"log"
	"time"
)

//WriteLog saves messagens into the log.txt
func WriteLog(msg string) error {
	c, err := helper.LoadFromConfigFile();if err != nil {
		log.Fatal(err)	
	}
	fileName := c.PathLog
	file, err := helper.LoadFile(fileName); if err != nil {
		return err
	}
	defer file.Close()
	t := time.Now()
	//tf := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	tf := t.Format("2006-01-02 15:04:05")
	msg = tf + " | " + msg + "\n"
	_, err = file.WriteString(msg);if err != nil {
		return err
	}
	return err
}