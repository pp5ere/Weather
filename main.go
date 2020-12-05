package main

import (
	"Weather/controller"
	"Weather/helper"
	"Weather/middlewares/util"
	"Weather/repository"
	"Weather/repository/log"
	"Weather/usecase"
	"fmt"
	l "log"
	"net/http"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
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
			log.WriteLog(err.Error());if err != nil {
				fmt.Println("Error to write log: "+err.Error())
			}
		}else{		
			err := repo.CreateTable("weather"); if err != nil {			
				log.WriteLog(err.Error());if err != nil {
					fmt.Println("Error to write log: "+err.Error())
				}
				return
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
}

// Execute this function to insert weather into database
func Execute()  {
	c, err := helper.LoadFromConfigFile();if err != nil {
		l.Fatal(err)
	}else{
		repo, err := repository.New(util.DriveSqlite3DB, c.PathSqliteDB);if err != nil {
			log.WriteLog(err.Error());if err != nil {
				fmt.Println("Error to write log: "+err.Error())
			}
		}else{	
			w, err := usecase.GetDataFromIoT();if err != nil {
				log.WriteLog(err.Error());if err != nil {
					fmt.Println("Error to write log: "+err.Error())
				}
			}else{
				controllers := controller.New(repo)
				if ((w.TempC < -50) || (w.TempC == 23.39 && w.TempF == 74.1 && w.Hum == 30.03 && w.Pres == 618.94 )){
					log.WriteLog(fmt.Sprintf("Invalid data values: TempC = %.2f TempF = %.2f Hum = %.2f Pres = %.2f",w.TempC, w.TempF, w.Hum, w.Pres));if err != nil {
						fmt.Println("Error to write log: "+err.Error())
					}			
				}else{
					err := controllers.Weather.Insert(w);if err != nil {
						log.WriteLog(err.Error());if err != nil {
							fmt.Println("Error to write log: "+err.Error())
						}
					}
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