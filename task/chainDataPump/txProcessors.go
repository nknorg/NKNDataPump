package chainDataPump

import (
	"NKNDataPump/storage/storageItem"
	"NKNDataPump/common"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"github.com/nknorg/nkn/pb"
	nknCommon "github.com/nknorg/nkn/common"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse"
	"strconv"
)

func coinbaseProcessor(data interface{}, extData interface{}, blockInfo interface{}) (err error) {
	tx := data.(rpcApiResponse.Transaction)
	txItem := extData.(storageItem.TransactionItem)
	block := blockInfo.(*storageItem.BlockItem)

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

	//rewardTransfer.ToAddr = hex.EncodeToString(coinbase.Recipient)
	recipientUint160 := nknCommon.BytesToUint160(coinbase.Recipient)
	address, addrErr := recipientUint160.ToAddress()
	if nil != addrErr {
		common.Log.Error(err)
		return
	}

	rewardTransfer.ToAddr = address
	block.Validator = address
	rewardTransfer.AssetId = ""
	rewardTransfer.Value = common.Fmt2Str(coinbase.Amount)
	rewardTransfer.Fee = common.Fmt2Str(txItem.Fee)
	rewardTransfer.Height = txItem.Height
	rewardTransfer.Timestamp = txItem.Timestamp

	insertItems([]storageItem.IItem{rewardTransfer})

	go recordAddr(address, txItem)

	return
}

func sigchainProcessor(data interface{}, extData interface{}, blockInfo interface{}) (err error) {
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

	var sigchainItems []storageItem.IItem
	for i, v := range sigchain.Elems {
		sigElem := &storageItem.SigchainItem{}
		sigElem.MappingFrom(hex.EncodeToString(v.Signature), txItem)

		if len(v.Id) == 0 {
			if i == 0 {
				sigElem.Id = hex.EncodeToString(sigchain.SrcId)
			} else {
				if i == len(sigchain.Elems)-1 {
					sigElem.Id = hex.EncodeToString(sigchain.DestId)
				}
			}
		} else {
			sigElem.Id = hex.EncodeToString(v.Id)
		}

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

func generateIdProcessor(data interface{}, extData interface{}, blockInfo interface{}) (err error) {
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

func transferAssetProcessor(data interface{}, extData interface{}, blockInfo interface{}) (err error) {
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

	//transferAssetItem.ToAddr = hex.EncodeToString(transferAsset.Recipient)
	addressUint := nknCommon.BytesToUint160(transferAsset.Recipient)
	toAddress, addrErr := addressUint.ToAddress()
	if nil!= addrErr{
		common.Log.Error(err)
		return
	}

	transferAssetItem.ToAddr = toAddress
	transferAssetItem.AssetId = ""
	transferAssetItem.Value = common.Fmt2Str(transferAsset.Amount)

	transferAssetItem.Fee = common.Fmt2Str(txItem.Fee)
	transferAssetItem.Height = txItem.Height
	transferAssetItem.Timestamp = txItem.Timestamp

	insertItems([]storageItem.IItem{transferAssetItem})
	return
}
