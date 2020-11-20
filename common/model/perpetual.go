package model

type Perpetual struct {
	ID                    int64  `json:"-"               db:"id"   primaryKey:"true" gorm:"primary_key"`
	PerpetualAddress      string `json:"perpetualAddress"  db:"perpetual_address"`
	OracleAddress         string `json:"oracleAddress" db:"oracle_address"`
	Symbol                string `json:"symbol"            db:"symbol"`
	CollateralTokenSymbol string `json:"collateralTokenSymbol"   db:"collateral_token_symbol"`
	CollateralAddress     string `json:"collateralAddress" db:"collateral_address"`
	BrokerAddress         string `json:"brokerAddress" db:"broker_address"`
	Version               int    `json:"version" db:"version"`
	IsPublished           bool   `json:"isPublished"       db:"is_published"`
	BlockNumber           int64  `json:"-" db:"block_number"`
}
