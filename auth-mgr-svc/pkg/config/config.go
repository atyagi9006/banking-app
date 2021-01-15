package config

import "github.com/spf13/viper"

func init() {
	setViperDefaults()
}

type SVCConfig struct {
	DBConfig    *SQLConfig
	OPAConfig   *OPAConfig
	RedisConfig *RedisConfig
	JWtConfig   *JWtConfig
}

type SQLConfig struct {
	Host    string
	Port    string
	DBName  string
	User    string
	Pass    string
	SSLMode string
}

type OPAConfig struct {
	Enabled  bool
	Endpoint string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}
type JWtConfig struct {
	SecretKey        string
	RefreshSecretKey string
}

func setViperDefaults() {
	viper.AutomaticEnv()
	// service configuration
	viper.SetDefault("grpc-port", 7781)
	viper.SetDefault("rest-port", 7782)
	viper.SetDefault("acc-grpc-ip4", "")
	// postgres configuration
	viper.SetDefault("postgres_host", "localhost")
	viper.SetDefault("postgres_port", "5432")
	viper.SetDefault("postgres_db", "auth")
	viper.SetDefault("postgres_user", "root")
	viper.SetDefault("postgres_password", "P@ssw0rd")
	viper.SetDefault("postgres_ssl_mode", "disable")

	// opa defaults
	viper.SetDefault("OPA_ENABLED", true)
	viper.SetDefault("OPA_ENDPOINT", "http://localhost:8181")

	//redis defaults
	viper.SetDefault("REDIS_HOST", "127.0.0.1")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_PASSWORD", "")

	//jwt defaults
	viper.SetDefault("ACCESS_SECRET", "")
	viper.SetDefault("REFRESH_SECRET", "")
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
	opaConfig := OPAConfig{
		Enabled:  viper.GetBool("OPA_ENABLED"),
		Endpoint: viper.GetString("OPA_ENDPOINT"),
	}

	redisConfig := RedisConfig{
		Host:     viper.GetString("REDIS_HOST"),
		Port:     viper.GetString("REDIS_PORT"),
		Password: viper.GetString("REDIS_PASSWORD"),
	}
	jwtConfig := JWtConfig{
		SecretKey:        viper.GetString("ACCESS_SECRET"),
		RefreshSecretKey: viper.GetString("REFRESH_SECRET"),
	}
	conf := SVCConfig{
		DBConfig:    &dbConfig,
		OPAConfig:   &opaConfig,
		RedisConfig: &redisConfig,
		JWtConfig:   &jwtConfig,
	}

	return &conf
}
