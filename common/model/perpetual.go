package model

type Perpetual struct {
	ID                    int64  `json:"id"               db:"id"               primaryKey:"true" gorm:"primary_key"`
	PerpetualAddress      string `json:"perpetualAddress"  db:"perpetual_address"`
	Symbol                string `json:"symbol"            db:"symbol"`
	CollateralTokenSymbol string `json:"collateralTokenSymbol"   db:"collateral_token_symbol"`
	ContractSizeSymbol    string `json:"contractSizeSymbol" db:"contract_size_symbol"`
	PriceDecimals         int    `json:"priceDecimals"     db:"price_decimals"`
	PriceSymbol           string `json:"priceSymbol"       db:"price_symbol"`
	AmountDecimals        int    `json:"amountDecimals"    db:"amount_decimals"`
	IsPublished           bool   `json:"isPublished"       db:"is_published"`
	BrokerAddress         string `json:"brokerAddress"     db:"broker_address"`
}
