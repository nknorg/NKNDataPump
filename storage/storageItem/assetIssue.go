package storageItem

import (
	. "NKNDataPump/common"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse"
)

type AssetIssueItem struct {
	AssetId   string
	Timestamp string
	IssueTo   string
	Value     string
	Height    uint32
}

func (a *AssetIssueItem) FieldList() []string {
	return []string{
		"asset_id",
		"issue_time",
		"issue_to",
		"value",
		"height",
	}
}

func (a *AssetIssueItem) StatementSqlValue() []string {
	return []string{
		a.AssetId,
		a.Timestamp,
		a.IssueTo,
		a.Value,
		Fmt2Str(a.Height),
	}
}

func (a *AssetIssueItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{}
}

func (a *AssetIssueItem) Table() string {
	return "t_assets_issue_record"
}

func (a *AssetIssueItem) MappingFrom(data interface{}, extData interface{}) {
	orgTx := data.(rpcApiResponse.Transaction)
	txItem := extData.(TransactionItem)

	txOut := orgTx.Outputs[0]

	a.Height = txItem.Height
	a.AssetId = txOut.AssetID
	a.Value = txOut.Value
	a.Timestamp = txItem.Timestamp
	a.IssueTo = txOut.Address
}
