package server

import (
	. "NKNDataPump/config"
	"NKNDataPump/server/api"
	"NKNDataPump/server/web"
	"strconv"
	"github.com/gin-gonic/gin"
	"NKNDataPump/common"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)

	//param from config file
	apiServerPort := strconv.FormatUint(uint64(PumpConfig.APIServerPort), 10)
	webServerPort := strconv.FormatUint(uint64(PumpConfig.WebServerPort), 10)

	webDir := PumpConfig.WebDir
	serviceBaseURI := PumpConfig.ServiceBaseURI

	//start api server
	common.Log.Info("start api server @ ", apiServerPort)
	apiServer.Start(serviceBaseURI, apiServerPort)

	//start web server
	webServer.Start(serviceBaseURI, webServerPort, webDir)

}
