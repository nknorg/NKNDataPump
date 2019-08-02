package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetLastPayListAPI IRestfulAPIAction = &getLastPayList{
}

type getLastPayList struct {
	restfulAPIBase
}

func (g *getLastPayList) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/last/pays"
}

func (g *getLastPayList) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)

	payList, _, err := dbHelper.QueryLastPayList()

	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(payList)

	return
}
