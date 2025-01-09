package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JuanMartinCoder/Module_ServerHealthMonitor/internal"
)

func RenderRoutes(mux *http.ServeMux) {
	mux.HandleFunc(RoutesInstance.MAIN, HandlerMain)
}

func main() {
	log.Println("Server Health Monitor is running...")
	go func() {
		for {
			sysInfo := internal.GetSystemInfo()
			fmt.Printf("%+v\n", sysInfo)
			time.Sleep(time.Second * 3)

		}
	}()

	time.Sleep(time.Second * 30)
}
