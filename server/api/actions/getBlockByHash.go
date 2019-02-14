package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetBlockByHashAPI IRestfulAPIAction = &getBlockByHash{
}

type getBlockByHash struct {
	restfulAPIBase
}

func (g *getBlockByHash) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/block/hash/:" + LOWERCASE_WORD_HASH
}

func (g *getBlockByHash) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}

	}()

	response := apiServerResponse.New(ctx)
	hash := ""

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_HASH: &hash,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	block, err := dbHelper.QueryBlockByHash(hash)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(block)
}
