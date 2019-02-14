package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetAddressListAPI IRestfulAPIAction = &getAddressList{
}

type getAddressList struct {
	restfulAPIBase
}

func (g *getAddressList) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/address/list/page/:" + LOWERCASE_WORD_ID
}

func (g *getAddressList) Action(ctx *gin.Context) {
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

	addressList, _, err := dbHelper.QueryAddressesList(page)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(addressList)
}
