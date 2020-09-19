package main

import (
	"Weather/controller"
	"Weather/helper"
	"Weather/middlewares/util"
	"Weather/repository"
	"Weather/repository/log"
	"Weather/usecase"
	l "log"
	"sync"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
	"net/http"
)

func main() {
	go startGorillaMux()
	log.WriteLog("Application started...")
	Execute()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	cronJob := cron.New()
	cronJob.Start()
	cronJob.AddFunc("@every 5m", Execute) //Wait 5 minutes and Execute
	wg.Wait()
}

func startGorillaMux(){	
	c, err := helper.LoadFromConfigFile();if err != nil {
		l.Fatal(err)
	}else{
		repo, err := repository.New(util.DriveSqlite3DB, c.PathSqliteDB);if err != nil {
			log.WriteLog(err.Error())
		}else{				
			controllers := controller.New(repo)
			r := usecase.Initialize(controllers)
			err := http.ListenAndServe(c.APIHost + c.APIPort, r);if err != nil {
				log.WriteLog(err.Error())
				l.Fatal(err)
			}    					
		}
	}
}

// Execute this function to insert weather into database
func Execute()  {
	c, err := helper.LoadFromConfigFile();if err != nil {
		l.Fatal(err)
	}else{
		repo, err := repository.New(util.DriveSqlite3DB, c.PathSqliteDB);if err != nil {
			log.WriteLog(err.Error())
		}else{	
			w, err := usecase.GetDataFromIoT();if err != nil {
				log.WriteLog(err.Error())
			}else{
				controllers := controller.New(repo)
				err := controllers.Weather.Insert(w);if err != nil {
					log.WriteLog(err.Error())
				}
			}
		}
	}
}

/*CROSSCOMPILE
compile to rapiberry pi:
env GOOS=linux GOARCH=arm GOARM=5 go build
compile to FreeBSD
env GOOS=freebsd GOARCH=amd64 go build
*/