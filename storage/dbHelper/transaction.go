package dbHelper

import (
	"github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/storage/storageItem"
	"database/sql"
)

func transactionsFromRows(rows *sql.Rows) (txList []storageItem.TransactionItem, count int, err error) {
	count = 0
	for rows.Next() {
		tx := storageItem.TransactionItem{}

		err = rows.Scan(
			&tx.Hash,
			&tx.Height,
			&tx.HeightIdxUnion,
			&tx.TxType,

			&tx.Attributes,
			&tx.Fee,
			&tx.Nonce,

			&tx.AssetId,
			&tx.UTXOInputCount,
			&tx.UTXOOutputCount,
			&tx.Timestamp,
			&tx.ParseStatus,
		)

		if nil != err {
			txList = nil
			return
		}

		txList = append(txList, tx)
		count++
	}
	return
}

func TransactionHashExist(hash string) bool {
	rows, err := db.Query("select count(*) from t_transactions where hash=?", hash)
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

func QueryTransactionByHash(hash string) (tx *storageItem.TransactionItem, err error) {
	rows, err := db.Query("select * from t_transactions where hash=?", hash)
	if nil != err {
		return
	}
	defer rows.Close()

	txList, count, err := transactionsFromRows(rows)
	if nil != err {
		return
	}

	if 0 == count {
		err = &common.GatewayError{Code: common.GW_ERR_NO_SUCH_DATA}
		return
	}
	tx = &txList[0]
	return
}

func QueryTransactionList(page uint32) (txList []storageItem.TransactionItem, count int, err error) {
	condition := ""
	if 0 == uint32(page) {
		condition = "limit 0," + common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY)
	} else {
		startRow := page * uint32(DEFAULT_ROW_COUNT_PER_QUERY)
		condition = "limit " + common.Fmt2Str(startRow) +
			"," + common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY)
	}

	rows, err := db.Query("select * from t_transactions " +
		"order by height_idx_union DESC " + condition)
	if nil != err {
		return
	}
	defer rows.Close()

	txList, count, err = transactionsFromRows(rows)
	return
}

func QueryTransactionsForBlock(height uint32) (txList []storageItem.TransactionItem, count int, err error) {
	rows, err := db.Query("select * from t_transactions where height=? " +
		"order by height_idx_union ASC", height)
	if nil != err {
		return
	}
	defer rows.Close()

	txList, count, err = transactionsFromRows(rows)
	return
}
