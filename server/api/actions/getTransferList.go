package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetTransferListAPI IRestfulAPIAction = &getTransferList{
}

type getTransferList struct {
	restfulAPIBase
}

func (g *getTransferList) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/transfer/list/page/:" + LOWERCASE_WORD_ID
}

func (g *getTransferList) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)
	page := uint32(0)

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_ID: &page,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	transferList, _, err := dbHelper.QueryTransferByUnionIdx(page)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(transferList)
}
