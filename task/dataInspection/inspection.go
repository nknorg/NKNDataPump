package dataInspection

import (
	"github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/network"
	"github.com/nknorg/NKNDataPump/storage"
	"github.com/nknorg/NKNDataPump/storage/dbHelper"
	"github.com/nknorg/NKNDataPump/storage/pumpDataTypes"
	"github.com/nknorg/NKNDataPump/storage/storageItem"
	"time"
	"github.com/nknorg/NKNDataPump/network/rpcRequest"
	"github.com/nknorg/NKNDataPump/network/chainDataTypes/rpcApiResponse"
)

func Start() {
	go dataIntegrityInspection()
}

//for now it only recover next block hash missing
func dataIntegrityInspection() {
	for {
		noNextHashBlocks, err := dbHelper.QueryNoNextHashBlock()
		if nil != err {
			time.Sleep(time.Second)
			continue
		}

		blockCount := len(noNextHashBlocks)
		if blockCount > 2 {
			common.Log.Info("data is missing, need data recover.")
			dataRecover(noNextHashBlocks, blockCount)
		}

		time.Sleep(time.Second)
	}
}

func dataRecover(blocks []pumpDataTypes.Block, count int) {
	for _, damagedBlock := range blocks[1:] {
		if "" != damagedBlock.Hash {
			nextHashRecover(damagedBlock)
		} else {
			common.Log.Errorf("this block is very damaged. can not recover for now.")
		}
	}
}

func nextHashRecover(block pumpDataTypes.Block) {
	data, err := rpcRequest.Api.Call(network.RPC_API_BLOCK_DETAIL_BY_HEIGHT, block.Height+1,
		false, common.NETWORK_RETRY_TIMES)

	if nil != err {
		common.Log.Trace(err)
		return
	}

	nextBlock := data.(*rpcApiResponse.Block)
	blockItem := &storageItem.BlockItem{}
	blockItem.MappingFrom(nextBlock, nil)

	damagedBlockItem, err := dbHelper.QueryBlockByHeight(block.Height)
	if nil != err || nil == damagedBlockItem {
		common.Log.Trace(err)
		return
	}

	storage.Exec(storageItem.STORE_BLOCK_ITEM_UPDATE_NEXT_HASH, blockItem.Hash, damagedBlockItem)
}
