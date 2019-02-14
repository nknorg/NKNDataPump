package dbHelper

import (
	"NKNDataPump/storage/storageItem"
	"database/sql"
)

func assetsFromRows(rows *sql.Rows) (assets []storageItem.AssetItem, count int, err error) {
	count = 0
	for rows.Next() {
		asset := storageItem.AssetItem{}

		err = rows.Scan(
			&asset.Hash,
			&asset.Amount,
			&asset.Name,
			&asset.Description,
			&asset.AssetPrecision,
			&asset.AssetType,
			&asset.Timestamp,
			&asset.Height,
		)

		if nil != err {
			assets = nil
			return
		}

		assets = append(assets, asset)
		count++
	}
	return
}

func QueryAsset(hash string) (asset *storageItem.AssetItem, err error) {
	rows, err := db.Query("select * from t_assets	 where hash=?", hash)
	if nil != err {
		return
	}
	defer rows.Close()

	assets, count, err := assetsFromRows(rows)
	if nil != err || 1 != count {
		return
	}

	asset = &assets[0]
	return
}

func QueryAssetList() (assetList []storageItem.AssetItem, count int, err error) {
	rows, err := db.Query("select * from t_assets where 1 order by height asc")
	if nil != err {
		return
	}
	defer rows.Close()

	assetList, count, err = assetsFromRows(rows)
	return
}
