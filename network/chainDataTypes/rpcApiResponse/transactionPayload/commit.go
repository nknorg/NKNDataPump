package transactionPayload

//type SigAlgo int32
//
//type SigChainElem struct {
//	SigAlgo              SigAlgo  `protobuf:"varint,1,opt,name=SigAlgo,enum=por.SigAlgo" json:"SigAlgo,omitempty"`
//	NextPubkey           []byte   `protobuf:"bytes,2,opt,name=NextPubkey,proto3" json:"NextPubkey,omitempty"`
//	Signature            []byte   `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
//	XXX_NoUnkeyedLiteral struct{} `json:"-"`
//	XXX_unrecognized     []byte   `json:"-"`
//	XXX_sizecache        int32    `json:"-"`
//}
//
//type SigChain struct {
//	Nonce                uint32          `protobuf:"varint,1,opt,name=Nonce" json:"Nonce,omitempty"`
//	DataSize             uint32          `protobuf:"varint,2,opt,name=DataSize" json:"DataSize,omitempty"`
//	DataHash             []byte          `protobuf:"bytes,3,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
//	BlockHash            []byte          `protobuf:"bytes,4,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
//	SrcPubkey            []byte          `protobuf:"bytes,5,opt,name=SrcPubkey,proto3" json:"SrcPubkey,omitempty"`
//	DestPubkey           []byte          `protobuf:"bytes,6,opt,name=DestPubkey,proto3" json:"DestPubkey,omitempty"`
//	Elems                []*SigChainElem `protobuf:"bytes,7,rep,name=Elems" json:"Elems,omitempty"`
//	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
//	XXX_unrecognized     []byte          `json:"-"`
//	XXX_sizecache        int32           `json:"-"`
//}

type Commit struct {
	SigChain  string 	`json:"sigChain"`
	Submitter string	`json:"submitter"`
}