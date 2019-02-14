package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetSigchainAPI IRestfulAPIAction = &getSigchainList{
}

type getSigchainList struct {
	restfulAPIBase
}

func (g *getSigchainList) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/signature_chain/list/:" + LOWERCASE_WORD_PAGE
}

func (g *getSigchainList) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)
	page := uint32(0)

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_PAGE: &page,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	sigchainList, _, err := dbHelper.QuerySigchainList(page)
	if nil != err {
		Log.Error(err)
		response.InternalServerError(nil)
		return
	}

	response.Success(sigchainList)
}

