package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
)

// Config for broker
type Config struct {
	APIHost        string                    `toml:"api_host"`
	APIRateLimit   int64                     `toml:"api_rate_limit"`
	WebsocketHost  string                    `toml:"websocket_host"`
	RedisURL       string                    `toml:"redis_url"`
	DataBaseURL    string                    `toml:"database_url"`
	JwtSecret      string                    `toml:"jwt_secret"`
	RPCHost        string                    `toml:"rpc_host"`
	FactoryAddress string                    `toml:"factory_address"`
	BrokerAddress  string                    `toml:"broker_address"`
	WatcherID      int                       `toml:"watcher_id"`
	GasStation     gasStation                `toml:"gas_station"`
	BlockChain     blockchain                `toml:"blockchain"`
	TokenMinAmount map[string]tokenMinAmount `toml:"token_min_amounts"`
}

type tokenMinAmount struct {
	Amount float64
}

type blockchain struct {
	ChainType   string   `toml:"chain_type"`
	Interval    duration `toml:"interval"`
	Timeout     duration `toml:"timeout"`
	ProviderURL string   `toml:"provider_url"`
	Password    string   `toml:"password"`
}

type gasStation struct {
	URL      string   `toml:"url"`
	GasLevel string   `toml:"gas_level"`
	Timeout  duration `toml:"timeout"`
	GasLimit uint64   `toml:"gas_limit"`
}

var (
	confFile string
	// Conf for server
	Conf *Config
)

func init() {
	flag.StringVar(&confFile, "conf", "./broker.toml", "set broker conf path")
}

// Init is parse config file
func Init() (err error) {
	_, err = toml.DecodeFile(confFile, &Conf)
	return err
}
