package models

type EthTransaction struct {
	Model
	EthBlockId           uint         `json:"ethBlockId"`
	BlockHash            string       `json:"blockHash"`
	BlockNumber          string       `json:"blockNumber"`
	ChainId              string       `json:"chainId"`
	From                 string       `json:"from" gorm:"column:from_"`
	Gas                  string       `json:"gas"`
	GasPrice             string       `json:"gasPrice"`
	Hash                 string       `json:"hash"`
	Input                string       `json:"input"`
	MaxFeePerGas         string       `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string       `json:"maxPriorityFeePerGas"`
	Nonce                string       `json:"nonce"`
	R                    string       `json:"r"`
	S                    string       `json:"s"`
	To                   string       `json:"to" gorm:"column:to_"`
	TransactionIndex     string       `json:"transactionIndex"`
	Type                 string       `json:"type" gorm:"column:type_"`
	V                    string       `json:"v"`
	Value                string       `json:"value" gorm:"column:value_"`
	AccessList           []AccessList `json:"accessList" gorm:"foreignKey:EthTransactionId"`
	Timestamps
}

type AccessList struct {
	Model
	EthTransactionId uint         `json:"ethTransactionId"`
	Address          string       `json:"address" gorm:"column:address_"`
	StorageKeys      []StorageKey `json:"storageKeys" gorm:"foreignKey:AccessListId"`
	Timestamps
}

type StorageKey struct {
	Model
	AccessListId uint   `json:"accessListId"`
	StorageKey   string `json:"storageKey"`
	Timestamps
}
