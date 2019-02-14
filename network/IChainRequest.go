package network

import "net/url"

type IChainRequest interface {
	Call(api string, param interface{}, getRawData bool, reTryCount int) (ret interface{}, err error)
	get(reqUrl string, param url.Values) (data []byte, err error)
	post(reqUrl string, param string) (data []byte, err error)
}
