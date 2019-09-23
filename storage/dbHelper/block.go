package dbHelper

import (
	. "github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/storage/pumpDataTypes"
	"github.com/nknorg/NKNDataPump/storage/storageItem"
	"database/sql"
	"errors"
)

func blocksFromRows(rows *sql.Rows) (blocks []storageItem.BlockItem, count int, err error) {
	count = 0
	for rows.Next() {
		block := storageItem.BlockItem{}

		err = rows.Scan(
			&block.Height,
			&block.Hash,
			&block.PrevBlockHash,
			&block.NextBlockHash,
			&block.Signature,
			&block.SignerId,
			&block.SignerPk,
			&block.StateRoot,
			&block.Validator,
			&block.Timestamp,
			&block.TransactionsRoot,
			&block.WinnerHash,
			&block.Size,
			&block.TxCount,
		)

		if nil != err {
			blocks = nil
			return
		}

		blocks = append(blocks, block)
		count++
	}

	return
}

func QueryCurrentBlockHeight() (height int, err error) {
	rows, err := db.Query("select height from t_blocks order by height DESC limit 0,1")
	if nil != err {
		height = 0
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&height)
		break
	}

	return
}

func QueryBlockCount() (count int, err error) {
	rows, err := db.Query("select count(*) from t_blocks")
	if nil != err {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		break
	}

	return
}

func queryBlocks(pSql string, args []interface{}) (blocks []storageItem.BlockItem, count int, err error) {
	rows, err := db.Query("select * from t_blocks "+pSql, args...)
	if nil != err {
		return
	}
	defer rows.Close()

	blocks, count, err = blocksFromRows(rows)
	return
}

func queryABlock(pSql string, args []interface{}) (block *storageItem.BlockItem, err error) {
	rows, err := db.Query("select * from t_blocks "+pSql, args...)
	if nil != err {
		return
	}
	defer rows.Close()

	blocks, count, err := blocksFromRows(rows)
	if nil != err {
		return
	}

	if 1 != count {
		err = &GatewayError{Code: GW_ERR_DATA_TYPE}
		return
	}

	block = &blocks[0]
	return
}

func QueryNewestBlockList() (blockList []storageItem.BlockItem, count int, err error) {
	condition := "where height>0 order by height desc limit 0," + Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY)
	blockList, count, err = queryBlocks(condition, nil)
	return
}

func QueryBlockListByHeight(height uint32) (blockList []storageItem.BlockItem, count int, err error) {
	condition := ""
	args := StringSlice2InterfaceSlice([]string{Fmt2Str(height)})

	if uint32(0) == height {
		condition = "order by height desc limit 0," + Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY)
		args = nil
	} else {
		condition = "where height<? order by height desc limit 0," + Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY)
	}

	blockList, count, err = queryBlocks(condition, args)
	return
}

func QueryBlockByHeight(height uint32) (block *storageItem.BlockItem, err error) {
	condition := "where height=? order by height desc"
	args := StringSlice2InterfaceSlice([]string{Fmt2Str(height)})
	block, err = queryABlock(condition, args)
	return
}

func QueryBlockByHash(hash string) (block *storageItem.BlockItem, err error) {
	condition := "where hash=? limit 0,1"
	block, err = queryABlock(condition, StringSlice2InterfaceSlice([]string{hash}))
	return
}

func BlockHashExist(hash string) bool {
	rows, err := db.Query("select count(*) from t_blocks where hash=?", hash)
	if nil != err {
		return false
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		rows.Scan(&count)
	}

	return 1 == count
}

func BlockHeightExist(height uint32) bool {
	rows, err := db.Query("select count(*) from t_blocks where height=?", height)
	if nil != err {
		return false
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		rows.Scan(&count)
	}

	return 1 == count
}

func QueryNoNextHashBlock() (blocks []pumpDataTypes.Block, err error) {
	rows, err := db.Query("select height,hash from t_blocks where next_hash='' order by height desc")
	defer func() {
		if r := recover(); nil != r {
			err = errors.New(Fmt2Str(r))
			Log.Trace(r)
		}
	}()
	if nil != err {
		return
	}
	defer rows.Close()

	for rows.Next() {
		newBlock := pumpDataTypes.Block{}

		_ = rows.Scan(&newBlock.Height, &newBlock.Hash)

		blocks = append(blocks, newBlock)
	}

	return
}
