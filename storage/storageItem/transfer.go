package storageItem

import (
	. "github.com/nknorg/NKNDataPump/common"
)

type TransferItem struct {
	Hash        string
	Height      uint32
	HeightTxIdx string
	FromAddr    string
	ToAddr      string
	AssetId     string
	Value       string
	Fee         string
	Timestamp   string
}

func (a *TransferItem) FieldList() []string {
	return []string{
		"hash",
		"height",
		"height_tx_idx_union",
		"from_addr",
		"to_addr",
		"asset_id",
		"value",
		"fee",
		"time",
	}
}

func (a *TransferItem) StatementSqlValue() []string {
	return []string{
		a.Hash,
		Fmt2Str(a.Height),
		a.HeightTxIdx,
		a.FromAddr,
		a.ToAddr,
		a.AssetId,
		a.Value,
		a.Fee,
		a.Timestamp,
	}
}

func (a *TransferItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{}
}

func (a *TransferItem) Table() string {
	return "t_assets_transfer"
}

func (a *TransferItem) MappingFrom(data interface{}, extData interface{}) {
	//method placeholder
}
