package rpcApiResponse

type BlockHeader struct {
	Hash             string `json:"Hash"`
	Height           uint32 `json:"Height"`
	PrevBlockHash    string `json:"PrevBlockHash"`
	RandomBeacon	 string `json:"RandomBeacon"`
	Signature		 string `json:"Signature"`
	SignerId		 string `json:"SignerId"`
	SignerPk		 string `json:"SignerPk"`
	StateRoot		 string `json:"StateRoot"`
	Timestamp        uint32 `json:"Timestamp"`
	TransactionsRoot string `json:"TransactionsRoot"`
	Version          uint32 `json:"Version"`
	WinnerHash		 string `json:"WinnerHash"`
	WinnerType 		 string `json:"WinnerType"`
}

type Block struct {
	Base
	Result struct {
		Hash         string        `json:"hash"`
		Header       BlockHeader   `json:"header"`
		Size         int           `json:"size"`
		Transactions []Transaction `json:"Transactions"`
	} `json:"result"`
}
