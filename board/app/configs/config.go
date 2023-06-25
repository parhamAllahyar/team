package configs

import "os"

type PostgresDBConfig struct {
	User   string `mapstructure:"POSTGRES_USER"`
	Pass   string `mapstructure:"POSTGRES_PASSWORD"`
	Host   string `mapstructure:"POSTGRES_URL"`
	Port   string `mapstructure:"POSTGRES_PORT"`
	DbName string `mapstructure:"POSTGRES_DB"`
}

type GrpcServer struct {
	Host string `mapstructure:"BOARD_APP_GRPC_HOST"`
	Port string `mapstructure:"BOARD_APP_GRPC_PORT"`
}

type Config struct {
	Postgre    PostgresDBConfig `mapstructure:"postgre"`
	GrpcServer GrpcServer       `mapstructure:"grpcServer"`
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

		GrpcServer{
			Port: os.Getenv("BOARD_APP_GRPC_HOST"),
			Host: os.Getenv("BOARD_APP_GRPC_PORT"),
		},
	}

}
