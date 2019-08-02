package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetTransfersByTxHashAPI IRestfulAPIAction = &getTransfersByTxHash{
}

type getTransfersByTxHash struct {
	restfulAPIBase
}

func (g *getTransfersByTxHash) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/transfer/in/transaction/:" + LOWERCASE_WORD_HASH
}

func (g *getTransfersByTxHash) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}

		//g.mutex.Unlock()
	}()

	response := apiServerResponse.New(ctx)
	hash := ""

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_HASH: &hash,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
	}

	transfers, _, err := dbHelper.QueryTransferByTxHash(hash)

	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(transfers)
}
