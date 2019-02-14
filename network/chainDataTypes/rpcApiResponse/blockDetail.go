package rpcApiResponse

type BlockHeader struct {
	Version          uint32 `json:"Version"`
	PrevBlockHash    string `json:"PrevBlockHash"`
	TransactionsRoot string `json:"TransactionsRoot"`
	Timestamp        uint32 `json:"Timestamp"`
	Height           uint32 `json:"Height"`
	ConsensusData    uint64 `json:"ConsensusData"`
	Hash             string `json:"Hash"`

	Program struct {
		Code      string `json:"Code"`
		Parameter string `json:"Parameter"`
	} `json:"Program"`
}

type Block struct {
	Base
	Result struct {
		Hash     	    string        `json:"hash"`
		Header			BlockHeader     `json:"header"`
		Transactions 	[]Transaction `json:"Transactions"`
	} `json:"result"`
}
