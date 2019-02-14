package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var SearchAPI IRestfulAPIAction = &search{
}

type search struct {
	restfulAPIBase
}

type searchType int

const (
	search_type_height searchType = iota
	search_type_hash
	search_type_addr
	search_type_unknown
)

type redirectToType int

const (
	redirect_to_block_by_hash redirectToType = iota
	redirect_to_block_by_height
	redirect_to_tx_by_hash
	redirect_to_address

	redirect_to_nowhere = -1
)

func (s *search) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/search/:" + LOWERCASE_WORD_VALUE
}

func (s *search) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)

	v := ""
	paramMap := map[string]interface{}{
		LOWERCASE_WORD_VALUE: &v,
	}

	err := s.getUrlParam(paramMap, ctx)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	searchActionMap := map[searchType]func(string) (redirectToType, error){
		search_type_hash:   searchHash,
		search_type_addr:   searchAddr,
		search_type_height: searchBlockByHeight,
	}

	searchAction, actExist := searchActionMap[getSearchType(v)]
	if !actExist {
		response.BadRequest(nil)
		return
	}

	retData, err := searchAction(v)
	if nil != err {
		response.BadRequest(nil)
		return
	}

	response.Success(retData)
}

func searchBlockByHeight(v string) (redirectToType, error) {
	height, err := strconv.Atoi(v)
	if nil != err || height < 0 {
		return redirect_to_nowhere, &GatewayError{Code: GW_ERR_NO_SUCH_DATA}
	}

	if dbHelper.BlockHeightExist(uint32(height)) {
		return redirect_to_block_by_height, nil
	}

	return redirect_to_nowhere, &GatewayError{Code: GW_ERR_NO_SUCH_DATA}
}

func searchAddr(addr string) (redirectToType, error) {
	if dbHelper.AddressExist(addr) {
		return redirect_to_address, nil
	}

	return redirect_to_nowhere, &GatewayError{Code: GW_ERR_NO_SUCH_DATA}
}

func searchHash(hash string) (redirectToType, error) {
	if dbHelper.BlockHashExist(hash) {
		return redirect_to_block_by_hash, nil
	}

	if dbHelper.TransactionHashExist(hash) {
		return redirect_to_tx_by_hash, nil
	}

	return redirect_to_nowhere, &GatewayError{Code: GW_ERR_NO_SUCH_DATA}
}

func getSearchType(v string) searchType {
	strLen := strings.Count(v, "") - 1

	switch {
	case 64 == strLen:
		return search_type_hash

	case 15 >= strLen:
		return search_type_height

	case 15 < strLen && strLen < 64:
		return search_type_addr

	default:
		return search_type_unknown
	}
}
