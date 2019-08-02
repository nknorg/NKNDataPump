package pumpDataTypes

import (
	"github.com/nknorg/NKNDataPump/network/chainDataTypes/rpcApiResponse"
)

type Block struct {
	rpcApiResponse.BlockHeader
	NextBlockHash string

	Size    int
	TxCount int
}
