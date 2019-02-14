package storageItem

import (
	. "NKNDataPump/common"
)

type AddressItem struct {
	Hash            string
	FirstUsedTime   string
	FirstUsedHeight uint32
}

func (a *AddressItem) FieldList() []string {
	return []string{
		"hash",
		"first_used_time",
		"first_used_block_height",
	}
}

func (a *AddressItem) StatementSqlValue() []string {
	return []string{
		a.Hash,
		a.FirstUsedTime,
		Fmt2Str(a.FirstUsedHeight),
	}
}

func (a *AddressItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{}
}

func (a *AddressItem) Table() string {
	return "t_addr"
}

func (a *AddressItem) MappingFrom(data interface{}, extData interface{}) {
	addr := data.(string)
	tx := extData.(TransactionItem)

	a.Hash = addr
	a.FirstUsedTime = tx.Timestamp
	a.FirstUsedHeight = tx.Height
}
