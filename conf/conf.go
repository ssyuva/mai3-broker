package conf

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

// Config for broker
type Config struct {
	APIHost       string   `envconfig:"api_host"`
	WebsocketHost string   `envconfig:"websocket_host"`
	DataBaseURL   string   `envconfig:"database_url"`
	JwtSecret     string   `envconfig:"jwt_secret"`
	RPCHost       string   `envconfig:"rpc_host"`
	ReaderAddress string   `envconfig:"reader_address"`
	BrokerAddress string   `envconfig:"broker_address"`
	WhiteList     []string `envconfig:"white_list"`

	GasPrice uint64 `envconfig:"gas_price"`
	GasLimit uint64 `envconfig:"gas_limit"`

	ChainType    string            `envconfig:"chain_type"`
	Interval     time.Duration     `envconfig:"interval"`
	ChainTimeout time.Duration     `envconfig:"timeout"`
	ProviderURL  string            `envconfig:"provider_url"`
	Headers      map[string]string `envconfig:"headers"`
	Password     string            `envconfig:"password"`

	SubgraphURL string `envconfig:"subgraph_url"`

	L2Timeout              duration `toml:"l2_timeout"`
	MaxTradeExpiration     duration `toml:"max_trade_expiration"`
	GasPrice               uint64   `toml:"gasPrice"`
	CallFunctionFeePercent uint32   `toml:"call_function_fee_percent"`
	TradeFee               int64    `toml:"trade_fee"`
	Key                    string   `toml:"key"`
}

var (
	// Conf for server
	Conf Config
)

// Init is parse config file
func Init() (err error) {
	return envconfig.Process("mcdex", &Conf)
}
