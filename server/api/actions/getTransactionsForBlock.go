package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetTransactionsForBlockAPI IRestfulAPIAction = &getTransactionsForBlock{
}

type getTransactionsForBlock struct {
	restfulAPIBase
}

func (g *getTransactionsForBlock) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/transaction/in/block/height/:" +
		LOWERCASE_WORD_HEIGHT + "/start/:" + LOWERCASE_WORD_ID
}

func (g *getTransactionsForBlock) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)
	height := uint32(0)
	startId := uint32(0)

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_HEIGHT: &height,
		LOWERCASE_WORD_ID:     &startId,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
	}

	tx, _, err := dbHelper.QueryTransactionsForBlock(height)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(tx)
}
