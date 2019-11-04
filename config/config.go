package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	serverAddress  string
	logger         LoggerConfig
	postgresConfig PostgresConfig
}

var config *Config

func Load() {
	viper.SetDefault(envServerAddress, ":8080")
	viper.SetDefault(envLogLevel, "debug")
	viper.SetDefault(envLoggerFormat, "text")

	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	config = &Config{
		serverAddress: getString(envServerAddress),
		logger: LoggerConfig{
			format: getString(envLoggerFormat),
			level:  getString(envLogLevel),
		},
		postgresConfig: PostgresConfig{
			host:        getString(envDatabaseHost),
			port:        getInt(envDatabasePort),
			name:        getString(envDatabaseName),
			username:    getString(envDatabaseUsername),
			password:    getString(envDatabasePassword),
			maxOpenConn: getInt(envDatabaseMaxOpenConn),
			maxIdleConn: getInt(envDatabaseMaxIdleConn),
			logEnabled:  getBool(envDatabaseEnableLog),
		},
	}
}

func Logger() LoggerConfig {
	return config.logger
}

func ServerAddress() string {
	return config.serverAddress
}

func Postgres() PostgresConfig {
	return config.postgresConfig
}

func checkEnvKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		log.Fatalf("%v env key is not set", key)
	}
}

func getString(key string) string {
	checkEnvKey(key)

	return viper.GetString(key)
}

func getInt(key string) int {
	str := getString(key)

	v, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("%v key is not valid int", key)
	}

	return v
}

func getBool(key string) bool {
	str := getString(key)

	v, err := strconv.ParseBool(str)
	if err != nil {
		log.Fatalf("%v key with value %s is not valid bool", key, str)
	}

	return v
}
