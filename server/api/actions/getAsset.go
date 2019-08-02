package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetAssetAPI IRestfulAPIAction = &getAsset{
}

type getAsset struct {
	restfulAPIBase
}

func (g *getAsset) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/asset/hash/:" + LOWERCASE_WORD_HASH
}

func (g *getAsset) Action(ctx *gin.Context) {
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

	asset, err := dbHelper.QueryAsset(hash)

	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(asset)

	return
}
