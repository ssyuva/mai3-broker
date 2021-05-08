package model

type Perpetual struct {
	LiquidityPoolAddress string `json:"liquidityPoolAddress" db:"liquidity_pool_address"`
	PerpetualIndex       int64  `json:"perpetualIndex"  db:"perpetual_index"`
	Symbol               string `json:"symbol"  db:"symbol"`
	OperatorAddress      string `json:"operatorAddress"  db:"operator_address"`
	CollateralSymbol     string `json:"collateralSymbol" db:"collateral_symbol"`
	CollateralAddress    string `json:"collateralAddress" db:"collateral_address"`
	CollateralDecimals   int32  `json:"collateralDecimals" db:"collateral_decimals"`
	IsPublished          bool   `json:"isPublished"       db:"is_published"`
	BlockNumber          int64  `json:"-" db:"block_number"`
}
