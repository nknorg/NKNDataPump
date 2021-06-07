package apiServerAction

import (
	"github.com/gin-gonic/gin"
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
)

var GetTransferForAddress IRestfulAPIAction = &getTransferForAddress{}

type getTransferForAddress struct {
	restfulAPIBase
}

func (g *getTransferForAddress) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/transfer/address/:" + LOWERCASE_WORD_ADDRESS + "/:" + LOWERCASE_WORD_PAGE
}

func (g *getTransferForAddress) Action(ctx *gin.Context) {
	defer func() {
		if r := recover(); nil != r {
			Log.Error(r)
		}
	}()
	Log.Info("api: /transfer/address")
	response := apiServerResponse.New(ctx)
	addr := ""
	page := uint64(0)

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_ADDRESS: &addr,
		LOWERCASE_WORD_PAGE:    &page,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}
	Log.Infof("params, addr: %s, page: %d", addr, page)

	Log.Info("query start.")
	transferCount, _ := dbHelper.QueryTransferCountForAddr(addr)
	transferList, _, err := dbHelper.QueryTransferForAddr(addr, page)

	Log.Info("query end.")

	if nil != err {
		Log.Error(err)
		response.InternalServerError(nil)
		return
	}

	response.Success(map[string]interface{}{
		"list":  transferList,
		"count": transferCount,
	})
}
