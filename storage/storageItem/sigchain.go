package storageItem

import (
	. "NKNDataPump/common"
	"NKNDataPump/storage/pumpDataTypes"
)

type SigchainItem struct {
	pumpDataTypes.Sigchain
}

func (s *SigchainItem) ExecBuilder() map[string]StoreCustomActions {
	return map[string]StoreCustomActions{
	}
}

func (s *SigchainItem) Table() string {
	return "t_sigchain"
}

func (s *SigchainItem) FieldList() []string {
	return []string{
		"height",
		"sig_idx",
		"addr",
		"next_pubkey",
		"tx_hash",
		"sig_data",
		"time",
	}
}

func (s *SigchainItem) StatementSqlValue() []string {
	return []string{
		Fmt2Str(s.Height),
		Fmt2Str(s.SigIndex),
		s.Addr,
		s.NextPubkey,
		s.TxHash,
		s.SigData,
		s.Timestamp,
	}
}

func (s *SigchainItem) MappingFrom(data interface{}, extData interface{}) {
	sig := data.(string)
	txItem := extData.(TransactionItem)

	s.Height = txItem.Height
	s.TxHash = txItem.Hash
	s.SigData = sig
}
