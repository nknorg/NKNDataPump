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

//func utxoTransferCalc(refUTXOList []rpcApiResponse.TxoutputInfo,
//	currentUTXOOut []rpcApiResponse.TxoutputInfo,
//	txItem storageItem.TransactionItem) (transferItems []storageItem.IItem) {
//	var outAddr []string
//	var outAddrValue []float64
//
//	addrMap := map[string]string{}
//	inAddrValueMap := map[string]float64{}
//
//	maxRefUTXO := 0.0
//	maxRefAddr := ""
//	for _, v := range refUTXOList {
//		utxoV, _ := strconv.ParseFloat(v.Value, 64)
//		addrMap[v.Address] = v.Address
//		inAddrValueMap[v.Address] += utxoV
//
//		if maxRefUTXO < utxoV {
//			maxRefAddr = v.Address
//			maxRefUTXO = utxoV
//		}
//	}
//
//	selfAddr := ""
//	outSum := 0.0
//	for _, v := range currentUTXOOut {
//		utxoV, _ := strconv.ParseFloat(v.Value, 64)
//		outSum += utxoV
//		outAddrValue = append(outAddrValue, utxoV)
//
//		if "" == v.Address {
//			addrFromProgramHash, err := common.ScriptHashToAddress(v.ProgramHash)
//
//			if nil != err {
//				//todo: error process
//				common.Log.Tracef("some thing wrong when calc address."+
//					"err [%v] . block height [%d] . tx hash [%s]", err, txItem.Height, txItem.Hash)
//				continue
//			}
//
//			v.Address = addrFromProgramHash
//		}
//
//		go recordAddr(v.Address, txItem)
//
//		if "" != addrMap[v.Address] {
//			selfAddr = v.Address
//			inAddrValueMap[v.Address] -= utxoV
//		} else {
//			outAddr = append(outAddr, v.Address)
//		}
//	}
//
//	outAddrIdx := 0
//	outCount := len(outAddr)
//
//	unionBaseIdx, _ := strconv.ParseUint(txItem.HeightIdxUnion, 10, 64)
//	if 0 == outCount {
//		common.Log.Tracef("wallet self transfer @block %d", txItem.Height)
//		//delete(inAddrValueMap, selfAddr)
//		inAddrValueMap[selfAddr] = outSum
//		outAddr = append(outAddr, selfAddr)
//	} else if 2 == outCount {
//		changeTransferItem, maxRefChangeVal, outIdx :=
//			buildChangeTransferItem(txItem, outAddrValue, outAddr, maxRefAddr, maxRefUTXO)
//
//		inAddrValueMap[maxRefAddr] = maxRefChangeVal
//
//		changeTransferItem.HeightTxIdx = common.Fmt2Str(unionBaseIdx)
//		unionBaseIdx += 1
//		transferItems = append(transferItems, changeTransferItem)
//
//		outAddrIdx = outIdx
//	}
//
//	for k, v := range inAddrValueMap {
//		transferItems = append(transferItems, &storageItem.TransferItem{
//			Hash:        txItem.Hash,
//			HeightTxIdx: common.Fmt2Str(unionBaseIdx),
//			FromAddr:    k,
//			ToAddr:      outAddr[outAddrIdx],
//			AssetId:     txItem.AssetId,
//			Value:       strconv.FormatFloat(v, 'f', 9, 64),
//			Fee:         "",
//			Timestamp:   txItem.Timestamp,
//			Height:      txItem.Height,
//		})
//		unionBaseIdx += 1
//	}
//
//	return
//}

//func buildChangeTransferItem(
//	txItem storageItem.TransactionItem,
//	outAddrValue []float64,
//	outAddr []string,
//	maxRefAddr string, maxRefInputVal float64) (
//	changeItem *storageItem.TransferItem,
//	maxRefChangeVal float64,
//	outAddrIdx int) {
//	var changeIdx int
//
//	if outAddrValue[0] > outAddrValue[1] {
//		changeIdx = 1
//		outAddrIdx = 0
//	} else {
//		changeIdx = 0
//		outAddrIdx = 1
//	}
//
//	maxRefChangeVal = maxRefInputVal - outAddrValue[changeIdx]
//
//	changeItem = &storageItem.TransferItem{
//		Hash:      txItem.Hash,
//		FromAddr:  maxRefAddr,
//		ToAddr:    outAddr[changeIdx],
//		AssetId:   txItem.AssetId,
//		Value:     strconv.FormatFloat(outAddrValue[changeIdx], 'f', 9, 64),
//		Fee:       "",
//		Timestamp: txItem.Timestamp,
//		Height:    txItem.Height,
//	}
//
//	return
//}
//
//func getRefUTXO(txHash string, utxoIdx int) (utxo rpcApiResponse.TxoutputInfo, err error) {
//	response, err := rpcRequest.Api.Call(network.RPC_API_TX_DETAIL, txHash,
//		false, common.NETWORK_RETRY_TIMES)
//
//	if nil != err {
//		return
//	}
//
//	txInfo := response.(*rpcApiResponse.TransactionByHash)
//
//	txUTXOOut := txInfo.Result.Outputs
//
//	if len(txUTXOOut) < utxoIdx+1 {
//		err = &common.GatewayError{Code: common.GW_ERR_INDEX_OUT_OF_RANGE}
//		return
//	}
//
//	utxo = txUTXOOut[utxoIdx]
//	return
//}

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
