package storageItem

import (
	"github.com/nknorg/NKNDataPump/common"
	"strings"
)

type StoreCustomActions func(interface{}) (string, []interface{}, error)
type ExecBuilder func() map[string]StoreCustomActions
type ItemFieldName string

type IItem interface {
	FieldList() []string
	StatementSqlValue() []string
	ExecBuilder() map[string]StoreCustomActions
	Table() string
	MappingFrom(interface{}, interface{})
}

func PrepareUpdateSql(fields []string, table string) (pSql string) {
	pSql = "update " + table + " set `" + strings.Join(fields, "`=?, `") + "`=?"

	return
}

func PrepareInsertSql(valueCount uint, fields []string, table string) (pSql string) {
	if 0 == valueCount {
		return
	}
	qMark := make([]string, len(fields))
	for i := 0; i < len(fields); i++ {
		qMark[i] = "?"
	}
	pSql = table + "(" + strings.Join(fields, ",") + ")"

	var values []string
	for i := uint(0); i < valueCount; i++ {
		values = append(values, "("+strings.Join(qMark[:], ",")+")")
	}

	pSql += " values " + strings.Join(values, ",")
	return
}

func BuildExec(act string, builder ExecBuilder, data interface{}) (pSql string, values []interface{}, err error) {
	if act := builder()[act]; nil != act {
		pSql, values, err = act(data)
	} else {
		err = &common.GatewayError{Code: common.GW_ERR_NO_SUCH_METHOD}
	}

	return
}
