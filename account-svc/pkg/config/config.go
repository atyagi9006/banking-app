package config

import "github.com/spf13/viper"

func init() {
	setViperDefaults()
}

type SVCConfig struct {
	DBConfig *SQLConfig
}

type SQLConfig struct {
	Host    string
	Port    string
	DBName  string
	User    string
	Pass    string
	SSLMode string
}

type OpaConfig struct {
	Enabled  bool
	Endpoint string
}

func setViperDefaults() {
	viper.AutomaticEnv()
	// service configuration
	viper.SetDefault("grpc-port", 7777)
	viper.SetDefault("rest-port", 7778)
	viper.SetDefault("acc-grpc-ip4", "")
	// postgres configuration
	viper.SetDefault("postgres_host", "localhost")
	viper.SetDefault("postgres_port", "5432")
	viper.SetDefault("postgres_db", "my_bank")
	viper.SetDefault("postgres_user", "root")
	viper.SetDefault("postgres_password", "P@ssw0rd")
	viper.SetDefault("postgres_ssl_mode", "disable")
}

// GetConfig returns an instance of config, is used internally
func GetConfig() *SVCConfig {
	dbConfig := SQLConfig{
		Host:    viper.GetString("postgres_host"),
		Port:    viper.GetString("postgres_port"),
		DBName:  viper.GetString("postgres_db"),
		User:    viper.GetString("postgres_user"),
		Pass:    viper.GetString("postgres_password"),
		SSLMode: viper.GetString("postgres_ssl_mode"),
	}
	conf := SVCConfig{
		DBConfig: &dbConfig,
	}
	return &conf
}
