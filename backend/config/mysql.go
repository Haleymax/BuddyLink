package config

type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Psw    string `mapstructure:"password"`
	DbName string `mapstructure:"dbname"`
}
