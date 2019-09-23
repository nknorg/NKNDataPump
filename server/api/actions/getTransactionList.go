package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetTransactionListAPI IRestfulAPIAction = &getTransactionList{
}

type getTransactionList struct {
	restfulAPIBase
}

func (g *getTransactionList) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/transaction/list/page/:" + LOWERCASE_WORD_ID
}

func (g *getTransactionList) Action(ctx *gin.Context) {
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

	tx, _, err := dbHelper.QueryTransactionList(page)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(tx)
}
