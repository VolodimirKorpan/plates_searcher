package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MysqlAddr     string `envconfig:"MYSQL_ADDR"`
	MysqlUser     string `envconfig:"MYSQL_USER"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD"`
	MysqlDB       string `envconfig:"MYSQL_DB"`
	HTTPPort      string `envconfig:"PORT"`
	HashSalt      string `envconfig:"HASH_SALT"`
	SigningKey    string `envconfig:"SIGNING_KEY"`
	TokenTTL      int64  `envconfig:"TOKEN_TTL"`
	PgAddr        string `envconfig:"PG_ADDR"`
	PgPort        string `envconfig:"PG_PORT"`
	PgUser        string `envconfig:"PG_USER"`
	PgPassword    string `envconfig:"PG_PASSWORD"`
	PgDB          string `envconfig:"PG_DB"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
