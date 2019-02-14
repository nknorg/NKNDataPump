package storage

import (
	"NKNDataPump/common"
	"NKNDataPump/storage/dbHelper"
	"NKNDataPump/storage/storageItem"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"NKNDataPump/config"
)

var db *sql.DB

func Init() (err error) {
	db, err = sql.Open("mysql",
		config.PumpConfig.ServiceDBUser + ":" + config.PumpConfig.ServiceDBPwd +
			"@/" + config.PumpConfig.ServiceDBName + "?charset=utf8")

	if nil != err {
		return err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	dbHelper.Init(db)
	return
}

func IgnoreKeyInsert(items []storageItem.IItem) (lastId int64, err error) {
	pSql := getPrepareSql(items)
	if "" == pSql {
		return
	}

	pSql = "insert ignore into " + pSql

	//fmt.Println("%v", items)
	return insert(items, pSql)

}

func getPrepareSql(items []storageItem.IItem) string {
	itemCount := uint(len(items))
	if 0 == itemCount {
		return ""
	}
	return storageItem.PrepareInsertSql(itemCount, items[0].FieldList(), items[0].Table())
}

func insert(items []storageItem.IItem, pSql string) (lastId int64, err error) {
	defer func() {
		if r:=recover(); nil != r {
			common.Log.Debug(r, err, pSql, len(items), *items[0].(*storageItem.TransferItem))
			panic(r)
		}
	}()

	stmt, err := db.Prepare(pSql)
	defer stmt.Close()

	if nil != err {
		return
	}

	var values []string

	for _, v := range items {
		values = append(values, v.StatementSqlValue()...)
	}

	result, err := stmt.Exec(common.StringSlice2InterfaceSlice(values)...)
	if nil == err {
		lastId, err = result.LastInsertId()
	} else {
		common.Log.Debug(err, "!!!", pSql)
	}

	return
}

func Exec(action string, data interface{}, item storageItem.IItem) (ret interface{}, err error) {
	pSql, values, err := storageItem.BuildExec(action, item.ExecBuilder, data)
	if nil != err {
		return
	}

	stmt, err := db.Prepare(pSql)
	defer stmt.Close()

	if nil != err {
		return
	}
	ret, err = stmt.Exec(values...)

	return
}
