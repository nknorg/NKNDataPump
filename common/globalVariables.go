package common

import "sync"

var (
	GAddrList          sync.Map
	CurrentBlockHeight = 0
)
