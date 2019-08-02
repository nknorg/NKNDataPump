package apiServer

import (
	. "github.com/nknorg/NKNDataPump/server/api/actions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	apiList = []IRestfulAPIAction{
		GetAddressListAPI,

		GetAssetAPI,
		GetAssetListAPI,
		GetMiningRewardsAPI,

		GetBlockByHashAPI,
		GetBlockByHeightAPI,
		GetBlockListAPI,
		GetBlockDetailByHeightAPI,

		GetTransactionAPI,
		GetTransactionListAPI,
		GetTransactionsForBlockAPI,

		GetTransfersByTxHashAPI,
		GetTransferListAPI,
		GetTransferOutListForAddressAPI,
		GetTransferInListForAddressAPI,
		GetTransferForAddress,
		GetSigchainAPI,

		GetLastPayListAPI,

		GetMiscAPI,

		SearchAPI,
	}
)

func Start(baseURI string, port string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
				c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
			}
			c.AbortWithStatus(http.StatusOK)
		}
	})
	setApiRouters(router, baseURI)
	go router.Run(":" + port)
}

func setApiRouters(router *gin.Engine, baseURI string) {
	for _, v := range apiList {
		router.GET(v.URI(baseURI), v.Action)
	}
}
