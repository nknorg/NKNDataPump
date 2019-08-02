package storageItem

import (
	. "github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/network/chainDataTypes/rpcApiResponse/transactionPayload"
	"github.com/nknorg/NKNDataPump/storage/pumpDataTypes"
)

type PayItem struct {
	pumpDataTypes.Pay
}

func (p *PayItem) FieldList() []string {
	return []string{
		"id",
		"height",
		"tx_hash",
		"payer",
		"payee",
		"value",
	}
}

func (p *PayItem) StatementSqlValue() []string {
	return []string{
		"",
		Fmt2Str(p.Height),
		p.Hash,
		p.Payer,
		p.Payee,
		p.Value,
	}
}

func (p *PayItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{}
}

func (p *PayItem) Table() string {
	return "t_pay"
}

func (p *PayItem) MappingFrom(data interface{}, extData interface{}) {
	payInfo := data.(transactionPayload.Pay)
	p.Payer = payInfo.Payer

	txItem := extData.(TransactionItem)

	p.Hash = txItem.Hash
	p.Height = txItem.Height
}
