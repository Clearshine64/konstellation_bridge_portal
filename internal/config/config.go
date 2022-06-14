package config

import (
	"os"
	"strconv"
	"strings"
)

type AppType string

func (at AppType) String() string {
	return string(at)
}

const (
	defaultPort = "1489"
)

// Config struct contains main client config
type Config struct {
	Debug              bool
	Name               string
	URL                string
	Env                string
	Port               string
	Network            string
	TLSEnable          bool
	TLSCertLocation    string
	TLSPrivKeyLocation string
	ReleaseInfo        *ReleaseInfo
	SwapInfo           *SwapInfo
	DB                 *DBConfig
}

// New set config and returns it
func New() *Config {
	var debug bool
	if os.Getenv("APP_DEBUG") == "1" {
		debug = true
	}

	tlsEnable, _ := strconv.ParseBool(os.Getenv("TLS_ENABLE"))

	cfg := &Config{
		Debug:              debug,
		Name:               os.Getenv("APP_NAME"),
		URL:                os.Getenv("APP_URL"),
		Env:                os.Getenv("APP_ENV"),
		Network:            os.Getenv("NETWORK"),
		TLSEnable:          tlsEnable,
		TLSCertLocation:    os.Getenv("TLS_CERT_LOCATION"),
		TLSPrivKeyLocation: os.Getenv("TLS_PRIV_KEY_LOCATION"),
		Port:               GetPort(),
		ReleaseInfo:        NewReleaseInfo(),
		SwapInfo:           NewSwapInfo(),
		DB:                 NewDBConfig(),
	}

	return cfg
}

func GetPort() string {
	port := defaultPort
	// GAE server will give us a port
	if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
		port = p
	}
	return port
}
