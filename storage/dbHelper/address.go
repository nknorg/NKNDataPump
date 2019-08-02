package dbHelper

import (
	"github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/storage/storageItem"
	"database/sql"
)

func addressesFromRows(rows *sql.Rows) (addresses []storageItem.AddressItem, count int, err error) {
	count = 0
	for rows.Next() {
		addr := storageItem.AddressItem{}

		err = rows.Scan(
			&addr.Hash,
			&addr.FirstUsedTime,
			&addr.FirstUsedHeight,
		)

		if nil != err {
			addresses = nil
			return
		}

		addresses = append(addresses, addr)
		count++
	}
	return
}

func QueryAddressesList(page uint32) (addrList []storageItem.AddressItem, count int, err error) {
	rows, err := db.Query("select * from t_addr order by first_used_block_height desc limit ?,"+
		common.Fmt2Str(DEFAULT_ROW_COUNT_PER_QUERY), page*uint32(DEFAULT_ROW_COUNT_PER_QUERY))
	if nil != err {
		return
	}
	defer rows.Close()

	addrList, count, err = addressesFromRows(rows)
	return
}

func AddressExist(addr string) bool {
	rows, err := db.Query("select count(*) from t_addr where hash=?", addr)
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
