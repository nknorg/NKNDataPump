package storageItem

import (
	. "github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/storage/pumpDataTypes"
)

type GenerateIdItem struct {
	pumpDataTypes.GenerateId
}

func (g *GenerateIdItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{
	}
}

func (g *GenerateIdItem) Table() string {
	return "t_generate_id"
}

func (g *GenerateIdItem) FieldList() []string {
	return []string{
		"height",
		"tx_hash",
		"public_key",
		"registration_fee",
		"time",
	}
}

func (g *GenerateIdItem) StatementSqlValue() []string {
	return []string{
		Fmt2Str(g.Height),
		g.TxHash,
		g.PublicKey,
		Fmt2Str(g.RegistrationFee),
		g.Timestamp,
	}
}

func (g *GenerateIdItem) MappingFrom(data interface{}, extData interface{}) {
	txItem := extData.(TransactionItem)

	g.Height = txItem.Height
	g.TxHash = txItem.Hash
	g.Timestamp = txItem.Timestamp
}
