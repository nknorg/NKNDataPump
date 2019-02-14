package dbHelper

import (
	"NKNDataPump/common"
	"NKNDataPump/storage/storageItem"
	"database/sql"
)

func transferFromRows(rows *sql.Rows) (transferList []storageItem.TransferItem, count int, err error) {
	count = 0
	transferList = []storageItem.TransferItem{}
	for rows.Next() {
		transfer := storageItem.TransferItem{}

		err = rows.Scan(
			&transfer.Hash,
			&transfer.Height,
			&transfer.HeightTxIdx,
			&transfer.FromAddr,
			&transfer.ToAddr,
			&transfer.AssetId,
			&transfer.Value,
			&transfer.Fee,
			&transfer.Timestamp,
		)

		if nil != err {
			transferList = nil
			return
		}

		transferList = append(transferList, transfer)
		count++
	}
	return
}

func QueryTransferByTxHash(hash string) (transferList []storageItem.TransferItem, count int, err error) {
	rows, err := db.Query("select * from t_assets_transfer where hash=?", hash)
	if nil != err {
		return
	}
	defer rows.Close()

	transferList, count, err = transferFromRows(rows)

	return
}

func QueryTransferByUnionIdx(page uint32) (transferList []storageItem.TransferItem, count int, err error) {
	condition := ""
	if 0 == uint32(page) {
		condition = "limit 0," + common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY)
	} else {
		startRow := page * uint32(DEFAULT_ROW_COUNT_PER_QUERY)
		condition = "limit " + common.Fmt2Str(startRow) +
			"," + common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY)
	}

	rows, err := db.Query("select * from t_assets_transfer " +
		"order by height_tx_idx_union desc " + condition)
	if nil != err {
		return
	}
	defer rows.Close()

	transferList, count, err = transferFromRows(rows)

	return
}

func QueryTransferOut(addr string, uionIdx uint64) (transferList []storageItem.TransferItem, count int, err error) {
	rows, err := db.Query("select * from t_assets_transfer where from_addr=? and height_tx_idx_union<? "+
		"order by height_tx_idx_union desc limit 0,"+common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY), addr, uionIdx)

	if nil != err {
		return
	}
	defer rows.Close()

	transferList, count, err = transferFromRows(rows)

	return
}

func QueryTransferIn(addr string, uionIdx uint64) (transferList []storageItem.TransferItem, count int, err error) {
	rows, err := db.Query("select * from t_assets_transfer where to_addr=? and height_tx_idx_union<? "+
		"order by height_tx_idx_union desc limit 0,"+common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY), addr, uionIdx)

	if nil != err {
		return
	}
	defer rows.Close()

	transferList, count, err = transferFromRows(rows)

	return
}

func QueryTransferCountForAddr(addr string) (count int, err error) {
	rows, err := db.Query("select count(*) from " +
		"t_assets_transfer where from_addr=? or to_addr=? ", addr, addr)

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

func QueryTransferForAddr(addr string, page uint64) (transferList []storageItem.TransferItem, count int, err error) {
	rows, err := db.Query("select * from t_assets_transfer where from_addr=? or to_addr=? "+
		"order by height_tx_idx_union desc limit ?,"+common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY + 1),
			addr, addr, page * DEFAULT_ROW_COUNT_PER_QUERY)

	if nil != err {
		return
	}
	defer rows.Close()

	transferList, count, err = transferFromRows(rows)
	return
}

//func QueryTransferForAddr(addr string, page uint64) (transferList []storageItem.TransferItem, count int, err error) {
//	toRows, err := db.Query("select * from t_assets_transfer where to_addr=? "+
//		"order by height desc limit ?,"+common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY + 1),
//		addr, page * DEFAULT_ROW_COUNT_PER_QUERY)
//
//	if nil != err {
//		return
//	}
//	defer toRows.Close()
//	toList, toCount, err := transferFromRows(toRows)
//
//	fromRows, err := db.Query("select * from t_assets_transfer where from_addr=? "+
//		"order by height desc limit ?,"+common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY + 1),
//		addr, addr, page * DEFAULT_ROW_COUNT_PER_QUERY)
//
//	if nil != err {
//		return
//	}
//	defer fromRows.Close()
//	fromList, fromCount, err := transferFromRows(fromRows)
//
//	count = toCount + fromCount
//	transferList = append(toList, fromList...)
//	return
//}

func QueryMiningRewards(addr string, height uint32) (transferList []storageItem.TransferItem, count int, err error) {
	rows, err := db.Query("select * from t_assets_transfer where from_addr='' and to_addr=? and height<? "+
		"order by height_tx_idx_union desc limit 0," + common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY), addr, height)

	if nil != err {
		return
	}
	defer rows.Close()

	transferList, count, err = transferFromRows(rows)

	return
}
