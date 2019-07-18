package chainDataTypes

type TransactionType byte

const (
	Coinbase      TransactionType = 0x00
	Pay           TransactionType = 0x01
	TransferAsset TransactionType = 0x10
	RegisterAsset TransactionType = 0x11
	IssueAsset    TransactionType = 0x12
	BookKeeper    TransactionType = 0x20
	Prepaid       TransactionType = 0x40
	Withdraw      TransactionType = 0x41
	Commit        TransactionType = 0x42
)

type TransactionAttributeUsage byte

const (
	Nonce          TransactionAttributeUsage = 0x00
	Script         TransactionAttributeUsage = 0x20
	DescriptionUrl TransactionAttributeUsage = 0x81
	Description    TransactionAttributeUsage = 0x90
)

const (
	CoinbaseN      string = "COINBASE_TYPE"
	SigChainN      string = "SIG_CHAIN_TXN_TYPE"
	GenerateId     string = "GENERATE_ID_TYPE"
	TransferAssetN string = "TRANSFER_ASSET_TYPE"
)
