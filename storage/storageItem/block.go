package storageItem

import (
	. "NKNDataPump/common"
	"NKNDataPump/storage/pumpDataTypes"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse"
	"time"
)

const (
	STORE_BLOCK_ITEM_UPDATE_NEXT_HASH = "STORE_BLOCK_ITEM_UPDATE_NEXT_HASH"
)

type BlockItem struct {
	pumpDataTypes.Block
	Validator string
	Timestamp string
}

func (b *BlockItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{
		STORE_BLOCK_ITEM_UPDATE_NEXT_HASH: b.updateNextHash,
	}
}

func (b *BlockItem) Table() string {
	return "t_blocks"
}

func (b *BlockItem) FieldList() []string {
	return []string{
		"height",
		"hash",
		"prev_hash",
		"next_hash",
		"validator",
		"time",
		"transaction_root",
		"size",
		"transaction_count",
	}
}

func (b *BlockItem) StatementSqlValue() []string {
	return []string{
		Fmt2Str(b.Height),
		b.Hash,
		b.PrevBlockHash,
		b.NextBlockHash,
		b.Validator,
		b.Timestamp,
		b.TransactionsRoot,
		Fmt2Str(b.Size),
		Fmt2Str(b.TxCount),
	}
}

func (b *BlockItem) MappingFrom(data interface{}, _ interface{}) {
	block := data.(*rpcApiResponse.Block)
	blockHeader := block.Result.Header
	b.Height = blockHeader.Height
	b.Hash = block.Result.Hash
	b.PrevBlockHash = blockHeader.PrevBlockHash
	b.Timestamp = Fmt2Str(time.Unix(int64(blockHeader.Timestamp), 0))
	b.TransactionsRoot = blockHeader.TransactionsRoot
}

func (b *BlockItem) updateNextHash(hash interface{}) (pSql string, execVal []interface{}, err error) {
	hashStr := hash.(string)
	if "" == hashStr {
		err = &GatewayError{Code: GW_ERR_DATA_TYPE}
		return
	}

	b.NextBlockHash = hashStr
	values := b.StatementSqlValue()

	pSql = PrepareUpdateSql(b.FieldList(), b.Table()) + " where `height`='" + Fmt2Str(b.Height) + "'"
	execVal = StringSlice2InterfaceSlice(values)

	return
}
