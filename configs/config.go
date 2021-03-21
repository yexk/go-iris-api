package configs

import (
	"os"
	"server/configs/autoload"

	"github.com/joho/godotenv"
)

type Config struct {
	Mysql   *autoload.Mysql
	Redis   *autoload.Redis
	ApiPort string `json:"api_port"`
	AppEnv  string `json:"app_env"`
}

func Init() (config *Config, err error) {
	err = godotenv.Load()
	// mysql
	config.Mysql.Driver = env("DB_DRIVER", "mysql")
	config.Mysql.Host = env("DB_HOST", "localhost")
	config.Mysql.Port = env("DB_PORT", "3306")
	config.Mysql.Database = env("DB_DATABASE", "yexk")
	config.Mysql.Username = env("DB_USERNAME", "root")
	config.Mysql.Password = env("DB_PASSWORD", "")
	config.Mysql.Charset = env("DB_CHARSET", "utf8mb4")
	config.Mysql.Collation = env("DB_COLLATION", "utf8mb4_unicode_ci")
	config.Mysql.Prefix = env("DB_PREFIX", "")
	// redis

	// other
	config.ApiPort = env("API_PORT", ":8080")
	config.AppEnv = env("APP_ENV", "error")

	return
}

func env(key string, defaults string) (value string) {
	value = defaults
	if os.Getenv(key) != "" {
		value = os.Getenv(key)
	}
	return
}
