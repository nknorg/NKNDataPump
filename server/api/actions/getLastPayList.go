package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
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
