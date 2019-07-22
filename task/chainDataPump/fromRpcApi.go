package chainDataPump

import (
	"NKNDataPump/network"
	"NKNDataPump/common"
	"sync"
	"time"
	"NKNDataPump/network/rpcRequest"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse"
	"NKNDataPump/storage/storageItem"
	"NKNDataPump/storage"
	"NKNDataPump/network/chainDataTypes"
)

func rpcDataPump() {
	pumpBlockHeight()
}

func pumpBlockHeight() {
	step := 200
	var lastItem *storageItem.BlockItem = nil

	for {
		ret, err := rpcRequest.Api.Call(network.RPC_API_BLOCK_HEIGHT, nil,
			false, common.NETWORK_RETRY_TIMES)

		if nil != err {
			common.Log.Error(err)
			time.Sleep(time.Second)
			continue
		}

		blockHeight := ret.(*rpcApiResponse.BlockHeight)

		targetHeight := blockHeight.Result
		thisTargetHeight := 0
		blockCountToGet := step

		wg := &sync.WaitGroup{}

		blockHeightDiff := 0

		if targetHeight > currentBlockHeight {
			blockHeightDiff = targetHeight - currentBlockHeight
			blockItems := make([]storageItem.BlockItem, step, step)
			thisTargetHeight = currentBlockHeight + step

			if thisTargetHeight > targetHeight {
				thisTargetHeight = targetHeight
			}

			blockCountToGet = thisTargetHeight - currentBlockHeight
			wg.Add(blockCountToGet)

			for i := currentBlockHeight; i < thisTargetHeight && i <= targetHeight; i++ {
				go getBlockDetailByHeight(i, &blockItems[i-currentBlockHeight], wg)
			}

			wg.Wait()

			saveBlocks(blockItems, blockCountToGet, lastItem)
			lastItem = &blockItems[blockCountToGet-1]

			currentBlockHeight += blockCountToGet
			common.CurrentBlockHeight = currentBlockHeight
		}

		if blockHeightDiff > 50 {
			time.Sleep(time.Millisecond * 10)
		} else {
			time.Sleep(time.Millisecond * 5)
		}
	}
}

func saveBlocks(blockItems []storageItem.BlockItem, itemCnt int, prevLastBlock *storageItem.BlockItem) {
	if nil != prevLastBlock {
		storage.Exec(storageItem.STORE_BLOCK_ITEM_UPDATE_NEXT_HASH, blockItems[0].Hash, prevLastBlock)
	}

	var second storageItem.BlockItem
	for i := 0; i < itemCnt-1; i++ {
		second = blockItems[i+1]
		blockItems[i].NextBlockHash = second.Hash
	}

	var data []storageItem.IItem
	for i := 0; i < itemCnt; i++ {
		data = append(data, &blockItems[i])
	}

	insertItems(data)
}

func getBlockDetailByHeight(height int, blockItem *storageItem.BlockItem, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	ret, err := rpcRequest.Api.Call(network.RPC_API_BLOCK_DETAIL_BY_HEIGHT, interface{}(height),
		false, common.NETWORK_RETRY_TIMES)

	if nil != err {
		blockItem.Height = uint32(height)
		common.Log.Error(err)
		return
	}

	//to block & raw block struct
	block := ret.(*rpcApiResponse.Block)
	blockItem.MappingFrom(block, nil)

	parseTransactions(block.Result.Transactions, blockItem)
}

func parseTransactions(tx []rpcApiResponse.Transaction, blockItem *storageItem.BlockItem) {
	processorMap := map[string]func(interface{}, interface{}) error{
		chainDataTypes.SigChainN:      sigchainProcessor,
		chainDataTypes.CoinbaseN:      coinbaseProcessor,
		chainDataTypes.GenerateId:     generateIdProcessor,
		chainDataTypes.TransferAssetN: transferAssetProcessor,
	}

	var txItems []storageItem.IItem

	for i, v := range tx {

		txItem := storageItem.TransactionItem{ParseStatus: storageItem.TRANSACTION_PARSE_STATUS_INIT}
		txItem.MappingFrom(v, blockItem)
		txItem.HeightIdxUnion = common.Fmt2Str((uint64(blockItem.Height) << 32) + uint64(i)<<16)

		processor := processorMap[v.TxType]
		var err error
		if nil != processor {
			err = processor(v, txItem)
		}

		if nil != err {
			common.Log.Error(err)
			txItem.ParseStatus = storageItem.TRANSACTION_PARSE_STATUS_FAILED
		} else {
			txItem.ParseStatus = storageItem.TRANSACTION_PARSE_STATUS_SUCCESS
		}
		txItems = append(txItems, &txItem)
	}

	if 0 != len(txItems) {
		insertItems(txItems)
	}
}

func insertItems(items []storageItem.IItem) (err error) {
	_, err = storage.IgnoreKeyInsert(items)
	if nil != err {
		common.Log.Error(err)
	}
	return
}

func recordAddr(addr string, tx storageItem.TransactionItem) {
	if false != common.GAddrList[addr] {
		return
	}

	addrItem := storageItem.AddressItem{}
	addrItem.MappingFrom(addr, tx)

	err := insertItems([]storageItem.IItem{
		&addrItem,
	})

	if nil == err {
		common.GAddrList[addr] = true
	} else {
		common.Log.Error(err)
	}
}
