package pumpDataTypes

import "NKNDataPump/network/chainDataTypes"

type Transaction struct {
	Hash            string
	Height          uint32
	HeightIdxUnion  string
	TxType          chainDataTypes.TransactionType
	AssetId         string
	UTXOInputCount  int
	UTXOOutputCount int
	Timestamp       string
}
