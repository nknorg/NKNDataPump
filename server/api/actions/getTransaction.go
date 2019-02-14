package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
	"NKNDataPump/network/chainDataTypes"
	"NKNDataPump/storage/storageItem"
)

var GetTransactionAPI IRestfulAPIAction = &getTransaction{
}

type getTransaction struct {
	restfulAPIBase
}

func (g *getTransaction) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/transaction/hash/:" + LOWERCASE_WORD_HASH
}

func (g *getTransaction) Action(ctx *gin.Context) {
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
	}

	tx, err := dbHelper.QueryTransactionByHash(hash)
	if nil != err {
		if err.(*GatewayError).Code == GW_ERR_NO_SUCH_DATA {
			response.Success(map[string] interface{} {
				"transaction": nil,
			})
		} else {
			response.InternalServerError("get transaction error")
		}
		return
	}

	var transfers []storageItem.TransferItem
	var sc []storageItem.SigchainItem
	switch tx.TxType {
	case chainDataTypes.Coinbase:
		transfers, _, err = dbHelper.QueryTransferByTxHash(tx.Hash)

	case chainDataTypes.TransferAsset:
		transfers, _, err = dbHelper.QueryTransferByTxHash(tx.Hash)

	case chainDataTypes.Commit:
		sc, err = dbHelper.QuerySigchainForTx(tx.Hash)

	default:
	}

	if nil != err {
		response.InternalServerError("get transaction detail error")
		return
	}

	response.Success(map[string] interface{} {
		"transaction": tx,
		"sigchain": sc,
		"transfer": transfers,
	})
}
