package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetBlockListAPI IRestfulAPIAction = &getBlockList{
}

type getBlockList struct {
	restfulAPIBase
}

func (g *getBlockList) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/block_list/:" + LOWERCASE_WORD_HEIGHT
}

func (g *getBlockList) Action(ctx *gin.Context) {
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

	blockList, _, err := dbHelper.QueryBlockListByHeight(height)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(blockList)

}
