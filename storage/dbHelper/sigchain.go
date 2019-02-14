package dbHelper

import (
	"NKNDataPump/storage/storageItem"
	"database/sql"
)

func sigchainFromRows(rows *sql.Rows) (sigchainList []storageItem.SigchainItem, count int, err error) {
	count = 0
	for rows.Next() {
		sc := storageItem.SigchainItem{}
		err = rows.Scan(
			&sc.Height,
			&sc.SigIndex,
			&sc.Addr,
			&sc.NextPubkey,
			&sc.TxHash,
			&sc.SigData,
			&sc.Timestamp,
		)

		if nil != err {
			sigchainList = nil
			return
		}

		sigchainList = append(sigchainList, sc)
		count++
	}
	return
}

func QuerySigchainForTx(txHash string) (sc []storageItem.SigchainItem, err error) {
	rows, err := db.Query("select * from t_sigchain where tx_hash=?", txHash)
	if nil != err {
		return
	}
	defer rows.Close()
	sc, _, err = sigchainFromRows(rows)

	return
}

func QuerySigchainList(page uint32) (scList []storageItem.SigchainItem, count int, err error) {

	rows, err := db.Query("select * from t_sigchain " +
		" order by height DESC limit?,120", page * 120)

	if nil != err {
		return
	}
	defer rows.Close()

	scList, count, err = sigchainFromRows(rows)
	return
}
