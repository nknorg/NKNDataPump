	package dbHelper

import (
	"github.com/nknorg/NKNDataPump/common"
	"github.com/nknorg/NKNDataPump/storage/storageItem"
	"encoding/json"
)

type ChainMiscInfo struct {
	BlockHeight      uint32                  `json:"block_height"`
	AssetCount       uint32                  `json:"asset_count"`
	TransactionCount uint32                  `json:"transaction_count"`
	TransferCount    uint32                  `json:"transfer_count"`
	AddrCount        uint32                  `json:"addr_count"`
	AssetsList       []storageItem.AssetItem `json:"assets_list"`
}

func queryCount(table string) (count uint32, err error) {
	pSql := "select count(*) from `" + table + "`"
	rows, err := db.Query(pSql)
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

func QueryMiscInfo() (info *ChainMiscInfo, err error) {
	tableToQuery := map[string]string{
		"block_height":      new(storageItem.BlockItem).Table(),
		"asset_count":       new(storageItem.AssetItem).Table(),
		"transaction_count": new(storageItem.TransactionItem).Table(),
		"transfer_count":    new(storageItem.TransferItem).Table(),
		"addr_count":        new(storageItem.AddressItem).Table(),
	}

	resultMap := map[string]uint32{}

	for k, v := range tableToQuery {
		count, queryErr := queryCount(v)
		if nil != queryErr {
			common.Log.Debug(queryErr)
			err = queryErr
			break
		}

		resultMap[k] = count
	}

	if nil != err {
		return
	}

	countJson, err := json.Marshal(resultMap)
	if nil != err {
		return
	}

	info = new(ChainMiscInfo)
	err = json.Unmarshal(countJson, info)

	assetList, _, err := QueryAssetList()
	if nil != err {
		info.AssetsList = []storageItem.AssetItem{}
	} else {
		info.AssetsList = assetList
	}

	return
}
