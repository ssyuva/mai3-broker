package conf

import (
	"time"

	"github.com/kelseyhightower/envconfig"
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
	ChainTimeout time.Duration     `envconfig:"chain_timeout"`
	ProviderURL  string            `envconfig:"provider_url"`
	Headers      map[string]string `envconfig:"headers"`
	Password     string            `envconfig:"password"`

	SubgraphURL string `envconfig:"subgraph_url"`
}

type L2RelayerConfig struct {
	BrokerAddress            string        `envconfig:"broker_address"`
	ProviderURL              string        `envconfig:"provider_url"`
	ChainID                  int64         `envconfig:"chain_id"`
	GasPrice                 uint64        `envconfig:"gas_price"`
	L2RelayerHost            string        `envconfig:"l2_relayer_host"`
	L2Timeout                time.Duration `envconfig:"l2_timeout"`
	L2MaxTradeExpiration     time.Duration `envconfig:"l2_max_trade_expiration"`
	L2CallFunctionFeePercent uint32        `envconfig:"l2_call_function_fee_percent"`
	L2TradeFee               int64         `envconfig:"l2_trade_fee"`
	L2RelayerKey             string        `envconfig:"l2_relayer_key"`
}

var (
	// Conf for server
	Conf          Config
	L2RelayerConf L2RelayerConfig
)

// Init is parse config file
func Init() (err error) {
	return envconfig.Process("mcdex", &Conf)
}

func InitL2RelayerConf() (err error) {
	return envconfig.Process("mcdex", &L2RelayerConf)
}
