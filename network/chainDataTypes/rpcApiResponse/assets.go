package rpcApiResponse

//STCChain STC is planed to support UTXO and Balance
const (
	Asset_UTXO    byte = 0x00
	Asset_Balance byte = 0x01
)

//define the asset stucture in STCChain STC
//registered asset will be assigned to contract address
type Asset struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Precision   uint32 `json:"Precision"`
	AssetType   byte   `json:"AssetType"`
	RecordType  byte   `json:"RecordType"`
}
