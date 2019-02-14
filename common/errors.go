package common

import "fmt"

type GatewayError struct {
	Code     int
	UserData interface{}
}

func (self *GatewayError) FmtOutput() string {
	var desc = errorDescription(self.Code)
	if nil != self.UserData {
		return fmt.Sprintf(desc+": %v", self.UserData)
	}
	return desc
}

func (self *GatewayError) Error() string {
	return self.FmtOutput()
}

const (
	GW_SUCCESS = iota
	GW_ERR_FILE_NOT_FOUND
	GW_ERR_CAN_NOT_READ_FILE
	GW_ERR_JSON_UNMARSHAL
	GW_ERR_CONFIG
	GW_ERR_DATA_TYPE
	GW_ERR_NO_SUCH_DATA
	GW_ERR_NO_SUCH_METHOD
	GW_ERR_INDEX_OUT_OF_RANGE
	GW_ERR_END
)

var errorDesc = map[int]string{
	GW_SUCCESS:                "ok",
	GW_ERR_FILE_NOT_FOUND:     "file not found",
	GW_ERR_CAN_NOT_READ_FILE:  "can not read file",
	GW_ERR_JSON_UNMARSHAL:     "json unmarshal failed",
	GW_ERR_CONFIG:             "set gateway config failed",
	GW_ERR_DATA_TYPE:          "data type error",
	GW_ERR_NO_SUCH_DATA:       "data not found",
	GW_ERR_NO_SUCH_METHOD:     "no such method",
	GW_ERR_INDEX_OUT_OF_RANGE: "index out of range",
}

func errorDescription(errorCode int) string {
	if errorCode >= GW_ERR_END {
		return ""
	}

	return errorDesc[errorCode]
}
