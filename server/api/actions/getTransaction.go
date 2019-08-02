package apiServerAction

import (
	. "github.com/nknorg/NKNDataPump/common"
	. "github.com/nknorg/NKNDataPump/server/api/const"
	"github.com/nknorg/NKNDataPump/server/api/response"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
	"github.com/nknorg/NKNDataPump/storage/storageItem"
	"github.com/nknorg/nkn/pb"
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
		if r := recover(); nil != r {
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
			response.Success(map[string]interface{}{
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
	case int32(pb.COINBASE_TYPE):
		transfers, _, err = dbHelper.QueryTransferByTxHash(tx.Hash)

	case int32(pb.TRANSFER_ASSET_TYPE):
		transfers, _, err = dbHelper.QueryTransferByTxHash(tx.Hash)

	case int32(pb.SIG_CHAIN_TXN_TYPE):
		sc, err = dbHelper.QuerySigchainForTx(tx.Hash)

	default:
	}

	if nil != err {
		response.InternalServerError("get transaction detail error")
		return
	}

	response.Success(map[string]interface{}{
		"transaction": tx,
		"sigchain":    sc,
		"transfer":    transfers,
	})
}
