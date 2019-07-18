package storageItem

import (
	. "NKNDataPump/common"
	"NKNDataPump/storage/pumpDataTypes"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse"
)

const (
	TRANSACTION_PARSE_STATUS_INIT    = 0
	TRANSACTION_PARSE_STATUS_SUCCESS = 1
	TRANSACTION_PARSE_STATUS_FAILED  = 2
)

type TransactionItem struct {
	pumpDataTypes.Transaction
	ParseStatus uint32
}

func (t *TransactionItem) FieldList() []string {
	return []string{
		"hash",
		"height",
		"height_idx_union",
		"tx_type",
		"attributes",
		"fee",
		"nonce",
		"asset_id",
		"utxo_input_count",
		"utxo_output_count",
		"time",
		"parse_status",
	}
}

func (t *TransactionItem) StatementSqlValue() []string {
	return []string{
		t.Hash,
		Fmt2Str(t.Height),
		Fmt2Str(t.HeightIdxUnion),
		Fmt2Str(t.TxType),
		t.Attributes,
		Fmt2Str(t.Fee),
		Fmt2Str(t.Nonce),
		t.AssetId,
		Fmt2Str(t.UTXOInputCount),
		Fmt2Str(t.UTXOOutputCount),
		t.Timestamp,
		Fmt2Str(t.ParseStatus),
	}
}

func (t *TransactionItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{}
}

func (t *TransactionItem) Table() string {
	return "t_transactions"
}

func (t *TransactionItem) MappingFrom(data interface{}, extData interface{}) {
	tx := data.(rpcApiResponse.Transaction)
	t.Hash = tx.Hash
	t.TxType = tx.TxType
	t.Attributes = tx.Attributes
	t.Fee = tx.Fee
	t.Nonce = tx.Nonce

	//t.UTXOInputCount = 0
	//t.UTXOOutputCount = 0

	block := extData.(*BlockItem)
	t.Height = block.Height
	t.Timestamp = block.Timestamp
}
