package models

type EthBlock struct {
	Model
	BaseFeePerGas    string           `json:"baseFeePerGas"`
	Difficulty       string           `json:"difficulty"`
	ExtraData        string           `json:"extraData"`
	GasLimit         string           `json:"gasLimit"`
	GasUsed          string           `json:"gasUsed"`
	Hash             string           `json:"hash"`
	LogsBloom        string           `json:"logsBloom"`
	Miner            string           `json:"miner"`
	MixHash          string           `json:"mixHash"`
	Nonce            string           `json:"nonce"`
	Number           string           `json:"number" gorm:"column:number_"`
	ParentHash       string           `json:"parentHash"`
	ReceiptsRoot     string           `json:"receiptsRoot"`
	Sha3Uncles       string           `json:"sha3Uncles"`
	Size             string           `json:"size" gorm:"column:size_"`
	StateRoot        string           `json:"stateRoot"`
	Timestamp        string           `json:"timestamp"`
	TotalDifficulty  string           `json:"totalDifficulty"`
	TransactionsRoot string           `json:"transactionsRoot"`
	WithdrawalsRoot  string           `json:"withdrawalsRoot"`
	Transactions     []EthTransaction `json:"transactions" gorm:"foreignKey:EthBlockId"`
	Uncles           []Uncles         `json:"uncles" gorm:"foreignKey:EthBlockId"`
	Withdrawals      []Withdrawals    `json:"withdrawals" gorm:"foreignKey:EthBlockId"`
	Timestamps
}

type Uncles struct {
	Model
	EthBlockId uint   `json:"EthBlockId"`
	Uncles     string `json:"uncles"`
	Timestamps
}

type Withdrawals struct {
	Model
	EthBlockId     uint   `json:"EthBlockId"`
	Address        string `json:"address" gorm:"column:address_"`
	Amount         string `json:"amount"`
	Index          string `json:"index" gorm:"column:index_"`
	ValidatorIndex string `json:"validatorIndex"`
	Timestamps
}
