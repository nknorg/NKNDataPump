package webServer

import (
	. "NKNDataPump/common"
	"NKNDataPump/server/web/const"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func Start(baseURI string, port string, webDir string) {
	if !FileExist(webDir) {
		Log.Fatalf("web directory [%s] not exist", webDir)
		os.Exit(-1)
	}

	if strings.Contains(webDir, ".") {
		Log.Fatalf("web directory contains '.' [webDir]", webDir)
		os.Exit(-1)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	setRouters(router, baseURI, webDir)
	go router.Run(":" + port)
}

func setRouters(router *gin.Engine, baseURI string, webDir string) {
	router.StaticFS(baseURI+webServerConsts.WEB_SERVER_URI_BASE, http.Dir(webDir))
}
