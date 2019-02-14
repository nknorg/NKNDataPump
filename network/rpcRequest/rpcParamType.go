package rpcRequest

type basicRequestData struct {
	Id string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params interface{} `json:"params"`
}

type GetBlockHeight struct {
	Height uint
}

type GetTx struct {
	Hash string
}
