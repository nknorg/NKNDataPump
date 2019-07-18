package rpcApiResponse

//type BlockHeader struct {
//	Version          uint32 `json:"Version"`
//	PrevBlockHash    string `json:"PrevBlockHash"`
//	TransactionsRoot string `json:"TransactionsRoot"`
//	Timestamp        uint32 `json:"Timestamp"`
//	Height           uint32 `json:"Height"`
//	ConsensusData    uint64 `json:"ConsensusData"`
//	Hash             string `json:"Hash"`
//
//	Program struct {
//		Code      string `json:"Code"`
//		Parameter string `json:"Parameter"`
//	} `json:"Program"`
//}
//
//type Block struct {
//	Base
//	Result struct {
//		Hash     	    string        `json:"hash"`
//		Header			BlockHeader     `json:"header"`
//		Transactions 	[]Transaction `json:"Transactions"`
//	} `json:"result"`
//}

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
