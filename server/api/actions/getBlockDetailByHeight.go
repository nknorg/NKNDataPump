package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetBlockDetailByHeightAPI IRestfulAPIAction = &getBlockDetailByHeightAPI{
}

type getBlockDetailByHeightAPI struct {
	restfulAPIBase
}

func (g *getBlockDetailByHeightAPI) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/block/detail/:" + LOWERCASE_WORD_HEIGHT
}

func (g *getBlockDetailByHeightAPI) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}

	}()

	response := apiServerResponse.New(ctx)

	height := uint32(0)
	paramMap := map[string]interface{}{
		LOWERCASE_WORD_HEIGHT: &height,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	block, err := dbHelper.QueryBlockByHeight(height)
	if nil != err {
		response.InternalServerError("query block error")
		return
	}

	transactions, _, err := dbHelper.QueryTransactionsForBlock(height)
	if nil != err {
		response.InternalServerError("query transactions error")
		return
	}

	response.Success(map[string] interface{} {
		"block": block,
		"txList": transactions,
	})
}
