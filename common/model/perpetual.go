package model

type Perpetual struct {
	LiquidityPoolAddress string `json:"liquidityPoolAddress" db:"liquidity_pool_address" primaryKey:"true" gorm:"primary_key"`
	PerpetualIndex       int64  `json:"perpetualIndex"  db:"perpetual_index" primaryKey:"true" gorm:"primary_key"`
	Symbol               string `json:"symbol"  db:"symbol"`
	OperatorAddress      string `json:"operatorAddress"  db:"operator_address"`
	CollateralSymbol     string `json:"collateralSymbol" db:"collateral_symbol"`
	IsPublished          bool   `json:"isPublished"       db:"is_published"`
	BlockNumber          int64  `json:"-" db:"block_number"`
}
