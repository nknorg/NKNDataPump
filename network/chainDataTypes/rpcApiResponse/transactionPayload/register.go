package transactionPayload

import (
	"github.com/nknorg/NKNDataPump/network/chainDataTypes/rpcApiResponse"
)

type Register struct {
	Asset  rpcApiResponse.Asset `json:"asset"`
	Amount string               `json:"amount"`
}
