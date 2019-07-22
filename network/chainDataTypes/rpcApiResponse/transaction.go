package rpcApiResponse


type ProgramInfo struct {
	Code      string
	Parameter string
}


type Transaction struct {
	TxType      string        `json:"TxType"`
	PayloadData string        `json:"PayloadData"`
	Attributes  string        `json:"Attributes"`
	Programs    []ProgramInfo `json:"Programs"`
	Nonce       int           `json:"nonce"`
	Fee         int           `json:"fee"`
	Hash        string        `json:"hash"`
}
