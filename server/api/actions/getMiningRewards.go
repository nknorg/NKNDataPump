package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
)

var GetMiningRewardsAPI IRestfulAPIAction = &getMiningRewards{
}

type getMiningRewards struct {
	restfulAPIBase
}

func (g *getMiningRewards) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/mining/rewards/:" + LOWERCASE_WORD_ADDRESS + "/:" + LOWERCASE_WORD_HEIGHT
}

func (g *getMiningRewards) Action(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)
	address := ""
	height := uint32(0)

	paramMap := map[string]interface{}{
		LOWERCASE_WORD_ADDRESS: &address,
		LOWERCASE_WORD_HEIGHT: &height,
	}

	err := g.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	transferList, _, err := dbHelper.QueryMiningRewards(address, height)
	if nil != err {
		response.InternalServerError(nil)
		return
	}

	response.Success(transferList)
}
