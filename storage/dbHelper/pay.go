package dbHelper

import (
	"database/sql"
	"github.com/nknorg/NKNDataPump/storage/storageItem"
)

func payFromRows(rows *sql.Rows) (pays []storageItem.PayItem, count int, err error) {
	count = 0
	for rows.Next() {
		pay := storageItem.PayItem{}
		height := 0

		err = rows.Scan(
			&height,
			&pay.Height,
			&pay.Hash,
			&pay.Payer,
			&pay.Payee,
			&pay.Value,
		)

		if nil != err {
			pays = nil
			return
		}

		pays = append(pays, pay)
		count++
	}
	return
}

func QueryLastPayList() (assetList []storageItem.PayItem, count int, err error) {
	rows, err := db.Query("select * from t_pay where 1 order by id desc limit 0,10")
	if nil != err {
		return
	}
	defer rows.Close()

	assetList, count, err = payFromRows(rows)
	return
}
