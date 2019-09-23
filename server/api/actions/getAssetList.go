package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetAssetListAPI IRestfulAPIAction = &getAssetList{
}

type getAssetList struct {
	restfulAPIBase
}

func (g *getAssetList) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/asset/list"
}

func (g *getAssetList) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}

	}()

	response := apiServerResponse.New(ctx)

	assetsList, _, err := dbHelper.QueryAssetList()
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(assetsList)
}
