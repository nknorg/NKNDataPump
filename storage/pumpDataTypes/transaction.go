package pumpDataTypes

type Transaction struct {
	Hash            string
	Height          uint32
	HeightIdxUnion  string
	TxType          int32

	Attributes 		string
	Fee				int
	Nonce			int

	AssetId         string
	UTXOInputCount  int
	UTXOOutputCount int
	Timestamp       string
}
