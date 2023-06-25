package configs

import "os"

type PostgresDBConfig struct {
	User   string `mapstructure:"POSTGRES_USER"`
	Pass   string `mapstructure:"POSTGRES_PASSWORD"`
	Host   string `mapstructure:"POSTGRES_URL"`
	Port   string `mapstructure:"POSTGRES_PORT"`
	DbName string `mapstructure:"POSTGRES_DB"`
}

type HTTPServer struct {
	Host string `mapstructure:"ADMIN_APP_HTTP_HOST"`
	Port string `mapstructure:"ADMIN_APP_HTTP_PORT"`
}

type Config struct {
	Postgre    PostgresDBConfig `mapstructure:"postgre"`
	HttpServer HTTPServer       `mapstructure:"httpServer"`
}

func LoadConfig() Config {

	//TODO: use viper

	return Config{
		PostgresDBConfig{
			User:   os.Getenv("POSTGRES_USER"),
			Pass:   os.Getenv("POSTGRES_PASSWORD"),
			Host:   os.Getenv("POSTGRES_URL"),
			Port:   os.Getenv("POSTGRES_PORT"),
			DbName: os.Getenv("POSTGRES_DB"),
		},

		HTTPServer{
			Port: os.Getenv("ADMIN_APP_HTTP_PORT"),
			Host: os.Getenv("ADMIN_APP_HTTP_HOST"),
		},
	}

}
