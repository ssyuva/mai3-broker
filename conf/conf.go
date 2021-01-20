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

	ChainType   string            `envconfig:"chain_type"`
	Interval    time.Duration     `envconfig:"interval"`
	Timeout     time.Duration     `envconfig:"timeout"`
	ProviderURL string            `envconfig:"provider_url"`
	Headers     map[string]string `envconfig:"headers"`
	Password    string            `envconfig:"password"`

	SubgraphURL string `envconfig:"subgraph_url"`
}

var (
	// Conf for server
	Conf Config
)

// Init is parse config file
func Init() (err error) {
	return envconfig.Process("mcdex", &Conf)
}
