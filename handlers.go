package main

import (
	"net/http"
	"time"

	"github.com/JuanMartinCoder/Module_ServerHealthMonitor/internal"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WSResponse struct {
	SysInfo   *internal.SystemInfo `json:"sysInfo"`
	TimeStamp string               `json:"timestamp"`
}

func HandlerMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main", nil)
}

var upgrader = websocket.Upgrader{}

func HandlerMonitor(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	defer ws.Close()

	for {

		sysInfo := internal.GetSystemInfo()
		wsReponse := WSResponse{
			SysInfo:   sysInfo,
			TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		}

		html := `
    <div hx-swap-oob="innerHTML:#timestamp">` + wsReponse.TimeStamp + `</div>
    <div hx-swap-oob="innerHTML:#sysInfo">
      <div class="system-section">
        <div>Hostname: ` + wsReponse.SysInfo.Hostname + `</div>
        <div>Platform: ` + wsReponse.SysInfo.Platform + `</div>
        <div>Total RAM: ` + wsReponse.SysInfo.RAMTotal + ` MB</div>
        <div>Free RAM: ` + wsReponse.SysInfo.RAMFree + ` MB</div>
        <div>Percentage used RAM: ` + wsReponse.SysInfo.RAMUsedPercent + `%</div>
      </div>
      <div class="disk-section">
        <div>Disk Total: ` + wsReponse.SysInfo.DiskTotal + ` GB</div>
        <div>Disk used: ` + wsReponse.SysInfo.DiskUsed + ` GB</div>
        <div>Disk Free: ` + wsReponse.SysInfo.DiskFree + ` GB</div>
        <div>Percentage used Disk: ` + wsReponse.SysInfo.DiskPercent + `%</div>
      </div>
      <div class="cpu-section">
        <div>Model Name: ` + wsReponse.SysInfo.CPU + `</div>
        <div>CPU Cores: ` + wsReponse.SysInfo.CPUCores + `</div>
        <div>CPU Speed: ` + wsReponse.SysInfo.CPUSpeed + `MHz</div>
      </div>
    `

		err := ws.WriteMessage(websocket.TextMessage, []byte(html))
		if err != nil {
			c.Logger().Error(err)
		}

		time.Sleep(time.Second * 2)
	}
}
