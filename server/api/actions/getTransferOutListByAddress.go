package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetTransferOutListForAddressAPI IRestfulAPIAction = &getTransferOutListForAddress{
}

type getTransferOutListForAddress struct {
	restfulAPIBase
}

func (g *getTransferOutListForAddress) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/address/transfer/out/list/:" + LOWERCASE_WORD_ADDRESS + "/:" + LOWERCASE_WORD_ID
}

func (g *getTransferOutListForAddress) Action(ctx *gin.Context) {
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

	transferList, _, err := dbHelper.QueryTransferOut(addr, unionIdx)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(transferList)
}
