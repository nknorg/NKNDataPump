package apiServerAction

import (
	. "NKNDataPump/common"
	. "NKNDataPump/server/api/const"
	"NKNDataPump/server/api/response"
	"NKNDataPump/storage/dbHelper"
	"github.com/gin-gonic/gin"
	"time"
	"math/rand"
	"crypto/sha256"
	"encoding/hex"
	"NKNDataPump/storage/storageItem"
)

var GetMiscAPI IRestfulAPIAction = &getMisc{
}

type getMisc struct {
	restfulAPIBase
}

/* begin: for mock data */
type LastTenBlockInfo struct {
	Timestamp int64 `json:"timestamp"`
	TxCount int `json:"tx_count"`
	Size int `json:"size"`
}

var LastTenBlock = map[uint32] LastTenBlockInfo {}

type LatestBlockInfo struct {
	TxCount uint32 `json:"tx_count"`
	Size uint32 `json:"size"`
	Winner int	`json:"winner"`
	Tx LatestTransactionInfo `json:"tx"`

}

type LatestTransactionInfo struct {
	Hash string `json:"hash"`
	MainSigchain Sigchain `json:"main_sigchain"`
	Sigchains []Sigchain `json:"sigchains"`
}

type Sigchain struct {
	Els []SigEl `json:"els"`
}

type SigEl struct {
	Signature string `json:"signature"`
	Idx int `json:"idx"`
}

func getRandomHash(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seedBytes := []byte(str)
	randomByte := []byte{}
	bytesLen := len(seedBytes)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		randomByte = append(randomByte, seedBytes[r.Intn(bytesLen)])
	}

	result := sha256.Sum256(randomByte)

	return hex.EncodeToString(result[:])
}

func (g *getMisc)BuildMockSigchain(sigChainLen int) (sigEls []SigEl){
	nodeIdxStep := 15 / (sigChainLen - 2)
	nodeIdxStart := 0

	lastIdx := sigChainLen - 2

	for i:=0; i<sigChainLen - 1; i++ {
		idx := nodeIdxStart + rand.Intn(nodeIdxStep)

		if lastIdx == i && 6 == sigChainLen {
			idx = nodeIdxStart + rand.Intn(5)
		}

		sig := SigEl{
			Signature: getRandomHash(16),
			Idx: idx,
		}

		nodeIdxStart += nodeIdxStep

		if 0 == i {
			sig.Signature = ""
			sig.Idx = 0
			nodeIdxStart = 0
		}

		sigEls = append(sigEls, sig)
	}

	sigEls = append(sigEls, SigEl{
		Signature:"", Idx:0,
	})

	return
}

func (g *getMisc) randomSigChainLen() (sigChainLen int) {
	sigChainLen = rand.Intn(6)
	if sigChainLen <= 2 {
		sigChainLen += 3
	}

	return
}

func (g *getMisc)BuildMockData(lastTxCount uint32) *LatestBlockInfo {
	data := &LatestBlockInfo{}

	data.TxCount = uint32(rand.Intn(2) + 1)
	data.Size = 127 + (data.TxCount * 296)
	//if data.TxCount > 20 {
	//	data.TxCount = uint32(rand.Intn(20)) + 1
	//}
	//
	//if 0 == lastTxCount {
	//	lastTxCount = 1
	//}

	data.Tx.Hash = getRandomHash(32)

	sigChainLen := g.randomSigChainLen()

	data.Winner = rand.Intn(sigChainLen - 2)
	data.Tx.MainSigchain = Sigchain{g.BuildMockSigchain(sigChainLen)}

	sigChainCount := rand.Intn(2) + 1

	for i:=0; i<sigChainCount; i++ {
		sigChain := g.BuildMockSigchain(g.randomSigChainLen())

		data.Tx.Sigchains = append(data.Tx.Sigchains, Sigchain{sigChain})
	}


	return data
}


func (g *getMisc)genLastTenBlock(block uint32) (ret []LastTenBlockInfo) {
	_, ok := LastTenBlock[block]
	now := time.Now()


	if !ok {
		lastBlock := LastTenBlockInfo{}
		lastBlock.Timestamp = now.Unix()
		lastBlock.TxCount = rand.Intn(2) + 1
		lastBlock.Size = 127 + (lastBlock.TxCount * 296)

		ret = append(ret, lastBlock)
		LastTenBlock[block] = lastBlock
	}

	for i:=uint32(1); i<10; i++ {
		lastBlock := LastTenBlockInfo{}

		before := time.Duration(rand.Intn(2))
		now = now.Add(-(time.Second * (10 + before)))
		t := now.Unix()

		lastBlock.Timestamp = t
		lastBlock.TxCount = rand.Intn(2) + 1
		lastBlock.Size = 127 + (lastBlock.TxCount * 296)

		ret = append(ret, lastBlock)
		LastTenBlock[block - i] = lastBlock
	}

	return
}
/* end: for mock data */

type miscResponseInfo struct {
	dbHelper.ChainMiscInfo
	DataPerRequest int `json:"row_count_per_request"`
	Sigchain []storageItem.SigchainItem `json:"sigchain"`
	Blocks []storageItem.BlockItem `json:"blocks"`
}

var cachedMiscResponse *miscResponseInfo = nil
var lastResponseTime = time.Now()

func (g *getMisc) URI(serverURI string) string {
	return serverURI + API_SERVER_URI_BASE + "/misc"
}

func (g *getMisc) responseFromCache(response * apiServerResponse.Response) (ok bool) {
	ok = false
	if nil == cachedMiscResponse {
		return
	}
	now := time.Now()

	if now.Add(-time.Second * 5).Before(lastResponseTime) {
		response.Success(cachedMiscResponse)
		ok = true
	} else {
		lastResponseTime = time.Now()
	}

	return
}

func (g *getMisc) Action(ctx *gin.Context) {
	defer func() {
		if r:=recover(); nil != r {
			Log.Error(r)
		}
	}()

	response := apiServerResponse.New(ctx)
	if g.responseFromCache(response) {
		return
	}

	miscInfo, err := dbHelper.QueryMiscInfo()
	if nil != err {
		response.InternalServerError(nil)
		return
	}


	sigchainList, _, err := dbHelper.QuerySigchainList(0)
	blockList, _, err := dbHelper.QueryBlockListByHeight(0)

	responseInfo := &miscResponseInfo{
		*miscInfo,
		dbHelper.DEFAULT_ROW_COUNT_PER_QUERY,
		sigchainList,
		blockList,
	}

	cachedMiscResponse = responseInfo
	response.Success(responseInfo)
}
