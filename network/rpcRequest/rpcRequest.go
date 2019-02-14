package rpcRequest

import (
	"net/url"
	. "NKNDataPump/common"
	. "NKNDataPump/network"
	. "NKNDataPump/config"
	"errors"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
	"NKNDataPump/network/chainDataTypes/rpcApiResponse"
	"net"
	"time"
)

var Api = new(RpcApi)

func (r *RpcApi) Build() {
	rpcUrl := PumpConfig.NodeRootURI

	r.api = map[string] rpcApiHandler {
		RPC_API_BLOCK_HEIGHT : {rpcUrl, r.getBlockHeight },
		RPC_API_BLOCK_DETAIL_BY_HEIGHT : {rpcUrl, r.getBlockDetailByHeight },
		RPC_API_TX_DETAIL : {rpcUrl, r.getTxByHash },
	}
}

type rpcRequestContext struct {
	ApiUrl       string
	RequestData  string
	ResponseData interface{}
}

type rpcRequestContextBuilder func(reqUrl string, param interface{}) *rpcRequestContext

type rpcApiHandler struct {
	Url        string
	ReqContext rpcRequestContextBuilder
}

func (r *RpcApi) buildRpcParam(method string, params interface{}) string  {
	reqData := &basicRequestData {
		rpc_CONST_ID,rpc_VERSION,method, params,
	}

	ret, _ := json.Marshal(reqData)

	return string(ret)
}

func (r *RpcApi) getBlockHeight(reqUrl string, _ interface{}) *rpcRequestContext {
	var param interface{}

	switch PumpConfig.NodeRPCVersion {
	case NODE_RPC_SERVER_V1:
		param = []interface{}{0}
		break

	case NODE_RPC_SERVER_V2:
		param = map[string] interface{}{"padding":0}
		break

	default:
		param = []interface{}{0}
		break
	}

	return &rpcRequestContext{
		ApiUrl:       reqUrl,
		RequestData:  r.buildRpcParam(rpc_api_GET_BLOCK_HEIGHT, param),
		ResponseData: &rpcApiResponse.BlockHeight{},
	}
}

func (r *RpcApi) getBlockDetailByHeight(reqUrl string, height interface{}) *rpcRequestContext {
	var param interface{}

	switch PumpConfig.NodeRPCVersion {
	case NODE_RPC_SERVER_V1:
		param = []interface{}{height}
		break

	case NODE_RPC_SERVER_V2:
		param = map[string] interface{}{"height": height}
		break

	default:
		param = []interface{}{height}
		break
	}

	return &rpcRequestContext{
		ApiUrl:       reqUrl,
		RequestData:  r.buildRpcParam(rpc_api_GET_BLOCK_DETAIL_BY_HEIGHT, param),
		ResponseData: &rpcApiResponse.Block{},
	}
}

func (r *RpcApi) getTxByHash(reqUrl string, hash interface{}) *rpcRequestContext {
	var param interface{}

	switch PumpConfig.NodeRPCVersion {
	case NODE_RPC_SERVER_V1:
		param = []interface{}{hash}
		break

	case NODE_RPC_SERVER_V2:
		param = map[string] interface{}{"hash": hash}
		break

	default:
		param = []interface{}{hash}
		break
	}

	return &rpcRequestContext{
		ApiUrl:       reqUrl,
		RequestData:  r.buildRpcParam(rpc_api_GET_TX, param),
		ResponseData: &rpcApiResponse.TransactionByHash{},
	}
}

type RpcApi struct {
	api map[string] rpcApiHandler
}

func (r *RpcApi) Call(name string, data interface{}, getRawData bool, reTryCount int) (ret interface{}, err error) {
	defer func() {
		if r := recover(); nil != r {
			Log.Error(r)
		}
	}()

	var handler = r.api[name]
	if nil == handler.ReqContext {
		err = errors.New("api [" + name + " (request url: " + handler.Url + ")] has no request context")
		return
	}
	reqCtx := handler.ReqContext(handler.Url, data)

	if nil == reqCtx.ResponseData {
		err = errors.New("api [" + name + "] has no response data struct")
		return
	}

	var jsonData []byte
	for i := 0; i < reTryCount; i++ {
		jsonData, err = r.post(reqCtx.ApiUrl, reqCtx.RequestData)
		if nil == err {
			break
		}
	}
	ret, err = responseHandler(jsonData, err, reqCtx.ResponseData)

	return
}

func (r *RpcApi) post(reqUrl string, param string) (responseData []byte, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(120 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*120)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	defer func() {
		if r := recover(); r != nil {
			Log.Error(r)
			panic(r)
		}
	}()

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(HTTP_METHOD_POST, reqUrl, bytes.NewBuffer([]byte(param)))
	req.Close = true

	content := GetDefaultHeader(HTTP_CONTENT_TYPE_JSON)
	acceptEncoding := GetDefaultHeader(HTTP_ACCEPT_ENCODING)
	contentEncoding := GetDefaultHeader(HTTP_CONTENT_ENCODING)

	req.Header.Set(content.Name, content.Value)
	req.Header.Add(acceptEncoding.Name, acceptEncoding.Value)

	response, err := client.Do(req)
	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		switch response.Header.Get(contentEncoding.Name) {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					break
				}

				if n == 0 {
					break
				}
				responseData = append(responseData, buf...)
			}

		default:
			responseData, err = ioutil.ReadAll(response.Body)
		}
	}

	return
}

func (r *RpcApi) get(reqUrl string, param url.Values) (data []byte, err error) {
	return nil, nil
}

func responseHandler(data []byte, retErr error, out interface{}) (ret interface{}, err error) {
	err = retErr

	if nil != retErr {
		return
	}

	err = json.Unmarshal(data, out)
	if nil == err {
		ret = out
	}
	return
}