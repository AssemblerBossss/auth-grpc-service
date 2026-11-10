package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"false"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    string        `yaml:"port" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	cleanenv.ReadConfig("config.yml", &cfg)
	if path != "" {
		panic("config file path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}

	var cfg Config
	cleanenv.ReadConfig("config.yml", &cfg)
	return &cfg

}

// fetchConfigPath fetches config path from command line or environment variable
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string
	// --config="patj/to-/config.yaml "
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
}
