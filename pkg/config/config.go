package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	
	DBName        string `mapstructure:"DB_NAME"`
	DBSOURCE      string `mapstructure:"DB_SOURCE"`
	
}

var envs = []string{
	
	"DB_NAME",
	"DB_SOURCE",
	
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	fmt.Printf("\n\nconfig : %v\n\n", config)
	return config, nil
}
