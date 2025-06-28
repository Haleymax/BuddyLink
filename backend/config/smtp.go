package config

type SMTPConfig struct {
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	SSL      bool   `mapstructure:"ssl"`
}
