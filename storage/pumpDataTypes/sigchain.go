package pumpDataTypes

import "github.com/nknorg/nkn/pb"

type Sigchain struct {
	Id 			string
	SigIndex	uint32
	Addr        string
	NextPubkey  string
	TxHash		string
	SigData		string
	Height		uint32
	SigAlgo 	pb.SigAlgo
	Vrf 		string
	Proof 		string
	Timestamp	string
}

