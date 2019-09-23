package chainDataPump

import (
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/nknorg/NKNDataPump/common"
)

var currentBlockHeight = 0

func Start() {
	currentBlockHeight, _ = dbHelper.QueryCurrentBlockHeight()

	common.CurrentBlockHeight = currentBlockHeight
	//go restfulDataPump()
	go rpcDataPump()
}
