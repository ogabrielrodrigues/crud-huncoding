package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	ConnString string
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = &config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			ConnString: viper.GetString("database.dsn"),
		},
	}

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetAPIConfig() APIConfig {
	return cfg.API
}
