package network

const (
	HTTP_METHOD_POST = "POST"
	HTTP_METHOD_GET  = "GET"

)

type HttpHeader struct {
	Name string
	Value string
}

const (
	HTTP_CONTENT_TYPE_JSON = "HTTP_CONTENT_TYPE_JSON"
	HTTP_ACCEPT_ENCODING = "HTTP_ACCEPT_ENCODING"
	HTTP_CONTENT_ENCODING = "HTTP_CONTENT_ENCODING"
)

func GetDefaultHeader(headerName string) (header *HttpHeader) {
	defaultHeaders := map[string] HttpHeader {
		HTTP_CONTENT_TYPE_JSON: {"Content-Type", "application/json"},
		HTTP_ACCEPT_ENCODING: {"Accept-Encoding", "gzip, deflate"},
		HTTP_CONTENT_ENCODING: {"Content-Encoding", "gzip"},
	}

	if def, ok := defaultHeaders[headerName]; ok {
		header = &HttpHeader{def.Name, def.Value}
	}

	return
}

const (
	RESTFUL_API_NODE_COUNT             = "NODE_COUNT"
	RESTFUL_API_BLOCK_HEIGHT           = "BLOCK_HEIGHT"
	RESTFUL_API_BLOCK_DETAIL_BY_HEIGHT = "BLOCK_DETAIL_BY_HEIGHT"
	RESTFUL_API_BLOCK_DETAIL_BY_HASH   = "BLOCK_DETAIL_BY_HASH"
	RESTFUL_API_TX_BY_HASH             = "TX_BY_HASH"
	RESTFUL_API_ASSET_BALANCE_BY_ADDR  = "ASSET_BALANCE_BY_ADDR"
	RESTFUL_API_ASSET_UTXO_BY_ADDR     = "ASSET_UTXO_BY_ADDR"
	RESTFUL_API_ASSET_BY_HASH          = "ASSET_BY_HASH"
)


const (
	RPC_API_BLOCK_HEIGHT 			= "BLOCK_HEIGHT"
	RPC_API_BLOCK_DETAIL_BY_HEIGHT  = "BLOCK_DETAIL"
	RPC_API_TX_DETAIL 				= "TX_DETAIL"
)
