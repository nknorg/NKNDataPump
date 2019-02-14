package rpcApiResponse

type TransactionByHash struct {
	Base
	Result Transaction `json:"Result"`
}
