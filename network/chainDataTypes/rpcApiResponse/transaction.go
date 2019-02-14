package rpcApiResponse

import . "NKNDataPump/network/chainDataTypes"

type TxAttributeInfo struct {
	Usage TransactionAttributeUsage
	Data  string
}

type UTXOTxInputInfo struct {
	ReferTxID          string
	ReferTxOutputIndex int
}

type BalanceTxInputInfo struct {
	AssetID     string
	Value       string
	ProgramHash string
}

type TxoutputInfo struct {
	AssetID     string `json:"AssetID"`
	Value       string `json:"Value"`
	Address     string `json:"Address"`
	ProgramHash string `json:"ProgramHash"`
}

type TxoutputMap struct {
	Key   string
	Txout []TxoutputInfo
}

type AmountMap struct {
	Key   string
	Value string
}

type ProgramInfo struct {
	Code      string
	Parameter string
}

type Transaction struct {
	TxType         TransactionType      `json:"TxType"`
	PayloadVersion byte                 `json:"PayloadVersion"`
	Payload        interface{}          `json:"Payload"`
	Attributes     []TxAttributeInfo    `json:"Attributes"`
	UTXOInputs     []UTXOTxInputInfo    `json:"Inputs"`
	Outputs        []TxoutputInfo       `json:"Outputs"`
	Programs       []ProgramInfo        `json:"Programs"`

	//AssetOutputs      []TxoutputMap `json:"AssetOutputs"`
	//AssetInputAmount  []AmountMap   `json:"AssetInputAmount"`
	//AssetOutputAmount []AmountMap   `json:"AssetOutputAmount"`

	Hash string `json:"hash"`
}

