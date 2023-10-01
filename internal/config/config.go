package config

import (
	"github.com/spf13/viper"
)

var configuration *Configuration

type Configuration struct {
	System SystemConfiguration
	MySQL  DBConfiguration
}

type SystemConfiguration struct {
	Port                   string `default:"9295"`
	AccessTokenExpireTime  string `default:"30m"`
	RefreshTokenExpireTime string `default:"30m"`
	TokenSecretKey         string `default:"jwt-token-secret-key"`
}

type DBConfiguration struct {
	Host     string `default:"localhost"`
	Dbname   string `default:"bank"`
	Username string `default:"root"`
	Password string `default:"1234"`
	Port     string `default:"3306"`
}

func Init() (*Configuration, error) {

	bindEnvs()

	setDefault()

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func bindEnvs() {
	viper.BindEnv("system.Port", "BT_PORT")
	viper.BindEnv("system.AccessTokenExpireTime", "BT_AC_TOKEN_TIME")
	viper.BindEnv("system.RefreshTokenExpireTime", "BT_RF_TOKEN_TIME")
	viper.BindEnv("system.TokenSecretKey", "BT_JWT_SECRET")

	viper.BindEnv("mysql.Host", "BT_MYSQL_HOST")
	viper.BindEnv("mysql.Dbname", "BT_MYSQL_DBNAME")
	viper.BindEnv("mysql.Username", "BT_MYSQL_USERNAME")
	viper.BindEnv("mysql.Password", "BT_MYSQL_PASSWORD")
	viper.BindEnv("mysql.PORT", "BT_MYSQL_PORT")
}

func setDefault() {
	viper.SetDefault("system.Port", "9295")
	viper.SetDefault("system.AccessTokenExpireTime", "5m")
	viper.SetDefault("system.RefreshTokenExpireTime", "20m")
	viper.SetDefault("system.TokenSecretKey", "jwt-token-secret-key")

	viper.SetDefault("mysql.Host", "localhost")
	viper.SetDefault("mysql.Dbname", "bank")
	viper.SetDefault("mysql.Username", "root")
	viper.SetDefault("mysql.Password", "1234")
	viper.SetDefault("mysql.PORT", "3306")
}
