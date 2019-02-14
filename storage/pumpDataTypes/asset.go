package pumpDataTypes

type Asset struct {
	Hash           string
	Amount         string
	Name           string
	Description    string
	AssetPrecision uint32
	AssetType      byte
	Timestamp      string
	Height         uint32
}
