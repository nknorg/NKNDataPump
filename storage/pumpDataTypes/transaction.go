package pumpDataTypes

type Transaction struct {
	Hash            string
	Height          uint32
	HeightIdxUnion  string
	//TxType          chainDataTypes.TransactionType
	TxType          string

	Attributes 		string
	Fee				int
	Nonce			int

	AssetId         string
	UTXOInputCount  int
	UTXOOutputCount int
	Timestamp       string
}
