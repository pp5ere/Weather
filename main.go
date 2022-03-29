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
var (
	c *helper.Config
	repo *repository.SqliteDB
	controllers *controller.Controllers
)

func main() {
	config, err := helper.LoadFromConfigFile();if err != nil {
		l.Fatal(err)
	}else{
		c = config
		repo, err = repository.New(util.DriveSqlite3DB, c.PathSqliteDB); if err != nil{
			l.Fatal(err)
		}
		controllers = controller.New(repo)
		go startGorillaMux(c)
		log.WriteLog("Back End Application started...")
		Execute(c)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		cronJob := cron.New()
		cronJob.Start()
		cronJob.AddFunc("@every 5m", CallExecute) //Wait 5 minutes and Execute
		wg.Wait()
	}
}

func CallExecute()  {
	Execute(c)
}

func startGorillaMux(c *helper.Config){			
	err := repo.CreateTable("weather"); if err != nil {			
		log.WriteLog(err.Error());if err != nil {
			fmt.Println("Error to write log: "+err.Error())
		}
	}else{
		r, err := usecase.Initialize(controllers, c); if err != nil {
			log.WriteLog(err.Error())
		}else{
			err := http.ListenAndServe(c.APIHost + c.APIPort, r);if err != nil {
				log.WriteLog(err.Error())
				l.Fatal(err)
			}
		}
	}
	
}

// Execute this function to insert weather into database
func Execute(c *helper.Config)  {	
	w, err := usecase.GetDataFromIoT(c);if err != nil {
		log.WriteLog(err.Error());if err != nil {
			fmt.Println("Error to write log: "+err.Error())
		}
	}else{		
		if ((w.TempC < -50) || (w.TempC == 23.39 && w.TempF == 74.1 && w.Hum == 30.03 && w.Pres == 618.94 )){
			log.WriteLog(fmt.Sprintf("Invalid data values: TempC = %.2f TempF = %.2f Hum = %.2f Pres = %.2f Alt = %2f Hi = %2f DewPoint = %2f",w.TempC, w.TempF, w.Hum, w.Pres, w.Alt, w.Hi, w.DewPoint));if err != nil {
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

/*CROSSCOMPILE
https://mtarnawa.org/2018/08/23/cross-compile-gorm-with-sqlite-for-raspberry-pi-arm7-and-odroid-arm64/
install 
sudo apt update && apt install gcc-arm-linux-gnueabihf
compile to rapiberry pi:
GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc go build -o weather main.go
compile to FreeBSD
env GOOS=freebsd GOARCH=amd64 go build
*/