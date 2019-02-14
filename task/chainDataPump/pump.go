package chainDataPump

import (
	"NKNDataPump/storage/dbHelper"
	"NKNDataPump/common"
)

var currentBlockHeight = 0

func Start() {
	currentBlockHeight, _ = dbHelper.QueryCurrentBlockHeight()

	common.CurrentBlockHeight = currentBlockHeight
	//go restfulDataPump()
	go rpcDataPump()
}
