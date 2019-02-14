package apiServerAction

import (
	. "NKNDataPump/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IRestfulAPIAction interface {
	URI(serverBaseURI string) string
	Action(ctx *gin.Context)
	getUrlParam(paramMap map[string]interface{}, ctx *gin.Context) error
}

type restfulAPIBase struct {
}

func (r *restfulAPIBase) getUrlParam(paramMap map[string]interface{}, ctx *gin.Context) error {
	paramGetters := map[VariableKind]func(paramName string, target interface{}, ctx *gin.Context) error{
		UintKind:   r.getUrlIntParam,
		StringKind: r.getUrlStringParam,
		BoolKind:   r.getUrlBoolParam,
	}

	var err error
	for k, v := range paramMap {
		vKind := InterfaceKindPtrCompatible(v)
		getter, ok := paramGetters[vKind]
		if ok {
			err = getter(k, v, ctx)
		}

		if nil != err {
			break
		}

	}

	return err
}

func (r *restfulAPIBase) getUrlIntParam(paramName string, target interface{}, ctx *gin.Context) error {
	data := ctx.Param(paramName)
	if "" == data {
		return &GatewayError{Code: GW_ERR_NO_SUCH_DATA}
	}

	switch t := target.(type) {
	case *uint32:
		v, err := strconv.ParseUint(data, 10, 32)
		if nil != err {
			return err
		}
		*t = uint32(v)

	case *uint64:
		v, err := strconv.ParseUint(data, 10, 64)
		if nil != err {
			return err
		}
		*t = uint64(v)

	default:
		return &GatewayError{Code: GW_ERR_DATA_TYPE}
	}

	return nil
}

func (r *restfulAPIBase) getUrlStringParam(paramName string, target interface{}, ctx *gin.Context) error {
	data := ctx.Param(paramName)
	if "" == data {
		return &GatewayError{Code: GW_ERR_NO_SUCH_DATA}
	}

	switch t := target.(type) {
	case *string:
		*t = data
	default:
		return &GatewayError{Code: GW_ERR_DATA_TYPE}
	}
	return nil
}

func (r *restfulAPIBase) getUrlBoolParam(paramName string, target interface{}, ctx *gin.Context) error {
	data := ctx.Param(paramName)
	if "" == data {
		return &GatewayError{Code: GW_ERR_NO_SUCH_DATA}
	}

	switch t := target.(type) {
	case *bool:
		*t = "true" == data
	default:
		return &GatewayError{Code: GW_ERR_DATA_TYPE}
	}
	return nil
}
