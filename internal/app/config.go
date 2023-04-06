package app

import "github.com/spf13/viper"

var (
	PORT     string
	ROOT_URL string
	DB_URL   string
	JWT_KEY  string
)

func loadVars() {
	PORT = viper.GetString("PORT")
	ROOT_URL = viper.GetString("ROOT_URL")
	DB_URL = viper.GetString("DB_URL")
	JWT_KEY = viper.GetString("JWT_KEY")
}

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	loadVars()
}

func Reload() {
	viper.ReadInConfig()
	loadVars()
}
