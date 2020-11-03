package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
)

// Config for broker
type Config struct {
	APIHost        string     `toml:"api_host"`
	APIRateLimit   int64      `toml:"api_rate_limit"`
	WebsocketHost  string     `toml:"websocket_host"`
	RedisURL       string     `toml:"redis_url"`
	DataBaseURL    string     `toml:"database_url"`
	JwtSecret      string     `toml:"jwt_secret"`
	FactoryAddress string     `toml:"factory_address"`
	WatcherID      int        `toml:"watcher_id"`
	BlockChain     blockchain `toml:"blockchain"`
}

type blockchain struct {
	Interval    duration `toml:"interval"`
	Timeout     duration `toml:"timeout"`
	ProviderURL string   `toml:"provider_url"`
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
