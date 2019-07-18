package chainDataPump

import (
	"NKNDataPump/storage/storageItem"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse/transactionPayload"
	"NKNDataPump/common"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"NKNDataPump/storage"
	"github.com/nknorg/nkn/pb"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse"
	"strconv"
)

func commitProcessor(data interface{}, extData interface{}) (err error) {
	txItem := extData.(storageItem.TransactionItem)
	commitPayload := transactionPayload.Commit{}
	err = common.JsonPointer2Struct(data, &commitPayload)

	if nil != err {
		common.Log.Fatal(err)
		return
	}

	sigchain := &pb.SigChain{}
	chainByte, _ := hex.DecodeString(commitPayload.SigChain)
	err = proto.Unmarshal(chainByte, sigchain)

	if nil != err {
		common.Log.Error(err)
		return
	}

	var sigchainItems []storageItem.IItem
	for idx, v := range sigchain.Elems {
		sig := &storageItem.SigchainItem{}
		sig.MappingFrom(hex.EncodeToString(v.Signature), txItem)
		//sig.Addr = hex.EncodeToString(v.Addr)
		sig.NextPubkey = hex.EncodeToString(v.NextPubkey)
		sig.SigIndex = uint32(idx)
		sig.Timestamp = txItem.Timestamp

		sigchainItems = append(sigchainItems, sig)
	}

	insertItems(sigchainItems)

	return
}

//func transferProcessor(data interface{}, extData interface{}) (err error) {
//	tx := data.(rpcApiResponse.Transaction)
//	txItem := extData.(storageItem.TransactionItem)
//
//	currentUTXOOut := tx.Outputs
//	refUTXO := tx.UTXOInputs
//
//	//check current utxo
//	//currentUTXOCount := len(currentUTXOOut)
//	//if currentUTXOCount > 2 {
//	//	common.Log.Fatalf("not normal nkn utxo model in block %d!!", txItem.Height)
//	//	os.Exit(0)
//	//}
//
//	if 0 == len(refUTXO) {
//		common.Log.Error("no ref utxo shown!!")
//		os.Exit(0)
//	}
//
//	//get ref utxos
//	var refUTXOList []rpcApiResponse.TxoutputInfo
//
//	for _, v := range refUTXO {
//		utxo, err := getRefUTXO(v.ReferTxID, v.ReferTxOutputIndex)
//		if nil != err {
//			//todo: error process
//			common.Log.Errorf("some thing wrong when try to get ref utxos."+
//				"err [%v] . block height [%d] . tx hash [%s]", err, txItem.Height, txItem.Hash)
//			return err
//		}
//
//		if "" == utxo.Address {
//			utxo.Address, err = common.ScriptHashToAddress(utxo.ProgramHash)
//
//			if nil != err {
//				//todo: error process
//				common.Log.Tracef("some thing wrong when calc address."+
//					"err [%v] . block height [%d] . tx hash [%s]", err, txItem.Height, txItem.Hash)
//				return err
//			}
//		}
//
//		refUTXOList = append(refUTXOList, utxo)
//
//		go recordAddr(utxo.Address, txItem)
//	}
//
//	//calc the transfer
//	txItem.AssetId = currentUTXOOut[0].AssetID
//	transferItems := utxoTransferCalc(refUTXOList, currentUTXOOut, txItem)
//	if nil != transferItems {
//		err = insertItems(transferItems)
//	}
//
//	return
//}

func registerAssetPayloadProcessor(data interface{}, extData interface{}) (err error) {
	txItem := extData.(storageItem.TransactionItem)
	regPayload := transactionPayload.Register{}
	err = common.JsonPointer2Struct(data, &regPayload)

	assetItem := storageItem.AssetItem{}
	assetItem.MappingFrom(regPayload, txItem)
	txItem.AssetId = assetItem.Hash

	insertItems([]storageItem.IItem{
		&assetItem,
	})

	return
}

func assetIssueProcessor(data interface{}, extData interface{}) (err error) {
	issueItem := storageItem.AssetIssueItem{}
	issueItem.MappingFrom(data, extData)

	storage.IgnoreKeyInsert([]storageItem.IItem{
		&issueItem,
	})

	insertItems([]storageItem.IItem{
		&issueItem,
	})

	go recordAddr(issueItem.IssueTo, extData.(storageItem.TransactionItem))

	return
}

//func payProcessor(data interface{}, extData interface{}) (err error) {
//	tx := data.(rpcApiResponse.Transaction)
//	txItem := extData.(storageItem.TransactionItem)
//	payPayload := transactionPayload.Pay{}
//	common.JsonPointer2Struct(tx.Payload, &payPayload)
//
//	var payList []storageItem.IItem
//
//	for _, o := range tx.Outputs {
//		newPayItem := &storageItem.PayItem{}
//		newPayItem.MappingFrom(payPayload, txItem)
//
//		newPayItem.Payee = o.Address
//		newPayItem.Value = o.Value
//
//		payList = append(payList, newPayItem)
//	}
//
//	err = insertItems(payList)
//
//	return
//}

func coinbaseProcessor(data interface{}, extData interface{}) (err error) {
	tx := data.(rpcApiResponse.Transaction)
	txItem := extData.(storageItem.TransactionItem)

	coinbase := &pb.Coinbase{}
	chainByte, _ := hex.DecodeString(tx.PayloadData)
	err = proto.Unmarshal(chainByte, coinbase)

	if nil != err {
		common.Log.Error(err)
		return
	}

	rewardTransfer := &storageItem.TransferItem{}
	unionBaseIdx, _ := strconv.ParseUint(txItem.HeightIdxUnion, 10, 64)
	rewardTransfer.Hash = txItem.Hash
	rewardTransfer.HeightTxIdx = common.Fmt2Str(unionBaseIdx)
	rewardTransfer.FromAddr = hex.EncodeToString(coinbase.Sender)

	rewardTransfer.ToAddr = hex.EncodeToString(coinbase.Recipient)
	rewardTransfer.AssetId = ""
	rewardTransfer.Value = common.Fmt2Str(coinbase.Amount)

	rewardTransfer.Fee = common.Fmt2Str(txItem.Fee)
	rewardTransfer.Height = txItem.Height
	rewardTransfer.Timestamp = txItem.Timestamp

	insertItems([]storageItem.IItem{rewardTransfer})

	go recordAddr(rewardTransfer.ToAddr, txItem)

	return
}

func sigchainProcessor(data interface{}, extData interface{}) (err error) {
	tx := data.(rpcApiResponse.Transaction)
	txItem := extData.(storageItem.TransactionItem)

	sigchainTx := &pb.SigChainTxn{}
	chainByte, _ := hex.DecodeString(tx.PayloadData)
	err = proto.Unmarshal(chainByte, sigchainTx)

	if nil != err {
		common.Log.Error(err)
		return
	}

	sigchain := &pb.SigChain{}
	err = proto.Unmarshal(sigchainTx.SigChain, sigchain)

	if nil != err {
		common.Log.Error(err)
		return
	}

	//somehash := hex.EncodeToString(sigchain.BlockHash)
	//nonce := sigchain.Nonce
	//dataSize := sigchain.DataSize
	//srcId := hex.EncodeToString(sigchain.SrcId)
	//srcPubKey := hex.EncodeToString(sigchain.SrcPubkey)
	//destId := hex.EncodeToString(sigchain.DestId)
	//destPubKey := hex.EncodeToString(sigchain.DestPubkey)

	var sigchainItems []storageItem.IItem
	for i, v := range sigchain.Elems {
		sigElem := &storageItem.SigchainItem{}
		sigElem.MappingFrom(hex.EncodeToString(v.Signature), txItem)

		sigElem.Id = hex.EncodeToString(v.Id)
		sigElem.SigIndex = uint32(i)
		sigElem.NextPubkey = hex.EncodeToString(v.NextPubkey)
		sigElem.SigAlgo = v.SigAlgo
		sigElem.Vrf = hex.EncodeToString(v.Vrf)
		sigElem.Proof = hex.EncodeToString(v.Proof)
		sigElem.Timestamp = txItem.Timestamp

		sigchainItems = append(sigchainItems, sigElem)
	}

	insertItems(sigchainItems)

	return
}

func generateIdProcessor(data interface{}, extData interface{}) (err error) {
	tx := data.(rpcApiResponse.Transaction)
	txItem := extData.(storageItem.TransactionItem)

	genId := &pb.GenerateID{}
	chainByte, _ := hex.DecodeString(tx.PayloadData)
	err = proto.Unmarshal(chainByte, genId)

	if nil != err {
		common.Log.Error(err)
		return
	}

	generateIdItem := &storageItem.GenerateIdItem{}
	generateIdItem.MappingFrom(nil, txItem)
	generateIdItem.RegistrationFee = genId.RegistrationFee
	generateIdItem.PublicKey = hex.EncodeToString(genId.PublicKey)

	insertItems([]storageItem.IItem{generateIdItem,})

	return
}

func transferAssetProcessor(data interface{}, extData interface{}) (err error) {
	tx := data.(rpcApiResponse.Transaction)
	txItem := extData.(storageItem.TransactionItem)

	transferAsset := &pb.TransferAsset{}
	chainByte, _ := hex.DecodeString(tx.PayloadData)
	err = proto.Unmarshal(chainByte, transferAsset)

	if nil != err {
		common.Log.Error(err)
		return
	}

	transferAssetItem := &storageItem.TransferItem{}
	unionBaseIdx, _ := strconv.ParseUint(txItem.HeightIdxUnion, 10, 64)
	transferAssetItem.Hash = txItem.Hash
	transferAssetItem.HeightTxIdx = common.Fmt2Str(unionBaseIdx)
	transferAssetItem.FromAddr = hex.EncodeToString(transferAsset.Sender)

	transferAssetItem.ToAddr = hex.EncodeToString(transferAsset.Recipient)
	transferAssetItem.AssetId = ""
	transferAssetItem.Value = common.Fmt2Str(transferAsset.Amount)

	transferAssetItem.Fee = common.Fmt2Str(txItem.Fee)
	transferAssetItem.Height = txItem.Height
	transferAssetItem.Timestamp = txItem.Timestamp

	insertItems([]storageItem.IItem{transferAssetItem})
	return
}
