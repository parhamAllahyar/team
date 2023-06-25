package configs

import "os"

type NEOConfig struct {
	User string `mapstructure:"NEO_USER"`
	Pass string `mapstructure:"NEO_PASSWORD"`
	Port string `mapstructure:"NEO_PORT"`
	Host string `mapstructure:"NEO_HOST"`
}

type GrpcServer struct {
	Host string `mapstructure:"WORKSPACE_APP_GRPC_HOST"`
	Port string `mapstructure:"WORKSPACE_APP_GRPC_PORT"`
}

type Config struct {
	NeoConfig  NEOConfig  `mapstructure:"postgre"`
	GrpcServer GrpcServer `mapstructure:"grpcServer"`
}

func LoadConfig() Config {

	//TODO: use viper

	return Config{
		NEOConfig{
			User: os.Getenv("NEO_USER"),
			Pass: os.Getenv("NEO_PASSWORD"),
			Port: os.Getenv("NEO_PORT"),
		},

		GrpcServer{
			// Port: os.Getenv("WORKSPACE_APP_GRPC_PORT"),
			// Host: os.Getenv("WORKSPACE_APP_GRPC_HOST"),
			Port: "4044",
		},
	}

}
