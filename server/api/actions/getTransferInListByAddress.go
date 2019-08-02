package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetTransferInListForAddressAPI IRestfulAPIAction = &getTransferInListForAddress{
}

type getTransferInListForAddress struct {
	restfulAPIBase
}

func (g *getTransferInListForAddress) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/address/transfer/in/list/:" + LOWERCASE_WORD_ADDRESS + "/:" + LOWERCASE_WORD_ID
}

func (g *getTransferInListForAddress) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)
	addr := ""
	unionIdx := uint64(0)

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_ADDRESS: &addr,
		LOWERCASE_WORD_ID:      &unionIdx,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	transferList, _, err := dbHelper.QueryTransferIn(addr, unionIdx)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(transferList)
}
