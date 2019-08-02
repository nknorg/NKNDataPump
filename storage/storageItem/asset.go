package storageItem

import (
	. "github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/network/chainDataTypes/rpcApiResponse/transactionPayload"
	"github.com/nknorg/NKNDataPump/storage/pumpDataTypes"
)

type AssetItem struct {
	pumpDataTypes.Asset
}

func (a *AssetItem) FieldList() []string {
	return []string{
		"hash",
		"amount",
		"name",
		"description",
		"asset_precision",
		"asset_type",
		"time",
		"height",
	}
}

func (a *AssetItem) StatementSqlValue() []string {
	return []string{
		a.Hash,
		a.Amount,
		a.Name,
		a.Description,
		Fmt2Str(a.AssetPrecision),
		Fmt2Str(a.AssetType),
		a.Timestamp,
		Fmt2Str(a.Height),
	}
}

func (a *AssetItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{}
}

func (a *AssetItem) Table() string {
	return "t_assets"
}

func (a *AssetItem) MappingFrom(data interface{}, extData interface{}) {
	regInfo := data.(transactionPayload.Register)
	assetInfo := regInfo.Asset

	a.AssetType = assetInfo.AssetType
	a.AssetPrecision = assetInfo.Precision
	a.Name = assetInfo.Name
	a.Description = assetInfo.Description
	a.Amount = Fmt2Str(regInfo.Amount)

	tx := extData.(TransactionItem)

	a.Hash = tx.Hash
	a.Height = tx.Height
	a.Timestamp = tx.Timestamp
}
