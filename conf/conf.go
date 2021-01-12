package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
)

// Config for broker
type Config struct {
	APIHost       string     `toml:"api_host"`
	APIRateLimit  int64      `toml:"api_rate_limit"`
	WebsocketHost string     `toml:"websocket_host"`
	RedisURL      string     `toml:"redis_url"`
	DataBaseURL   string     `toml:"database_url"`
	JwtSecret     string     `toml:"jwt_secret"`
	RPCHost       string     `toml:"rpc_host"`
	ReaderAddress string     `toml:"reader_address"`
	BrokerAddress string     `toml:"broker_address"`
	WatcherID     int        `toml:"watcher_id"`
	WhiteList     []string   `toml:"white_list"`
	GasStation    gasStation `toml:"gas_station"`
	BlockChain    blockchain `toml:"blockchain"`
	Subgraph      subgraph   `toml:"subgraph"`
}

type blockchain struct {
	ChainType   string   `toml:"chain_type"`
	Interval    duration `toml:"interval"`
	Timeout     duration `toml:"timeout"`
	ProviderURL string   `toml:"provider_url"`
	Password    string   `toml:"password"`
}

type gasStation struct {
	GasPrice uint64 `toml:"gas_price"`
	GasLimit uint64 `toml:"gas_limit"`
}

type subgraph struct {
	URL string `toml:"url"`
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
