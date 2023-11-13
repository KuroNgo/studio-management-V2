package conf

import (
	"fmt"
	"time"

	env "github.com/caarlos0/env"
	viper "github.com/spf13/viper"
)

type Config struct {
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"DB_PORT"`
	DBUser     string `env:"DB_USER" envDefault:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" envDefault:"DB_NAME"`
	Port       string `env:"PORT" envDefault:"5432"`

	MaxOpenConns    int `env:"MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns    int `env:"MAX_IDLE_CONNS" envDefault:"25"`
	ConnMaxLifetime int `env:"CONN_MAX_LIFETIME" envDefault:"600"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY" `
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAX_AGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAX_AGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	Url_HTTPS    string `mapstructure:"URL_CLOUDINARY_WITH_HTTPS"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
}

var Cfg Config

func SetEnv() {
	if err := env.Parse(&Cfg); err != nil {
		fmt.Printf("Failed to parse env %v", err)
		return
	}
}

func GetEnv() Config {
	return Cfg
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
