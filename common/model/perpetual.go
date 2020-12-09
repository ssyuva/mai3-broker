package model

type Perpetual struct {
	ID                int64  `json:"-"               db:"id"   primaryKey:"true" gorm:"primary_key"`
	PerpetualAddress  string `json:"perpetualAddress"  db:"perpetual_address"`
	GovernorAddress   string `json:"governorAddress"  db:"governor_address"`
	ShareToken        string `json:"shareToken"  db:"share_token"`
	OperatorAddress   string `json:"operatorAddress"  db:"operator_address"`
	CollateralSymbol  string `json:"collateralSymbol" db:"collateral_symbol"`
	OracleAddress     string `json:"oracleAddress" db:"oracle_address"`
	CollateralAddress string `json:"collateralAddress" db:"collateral_address"`
	IsPublished       bool   `json:"isPublished"       db:"is_published"`
	BlockNumber       int64  `json:"-" db:"block_number"`
}
