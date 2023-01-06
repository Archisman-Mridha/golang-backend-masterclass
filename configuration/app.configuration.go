package configuration

import "github.com/spf13/viper"

type AppConfiguration struct {

	POSTGRES_DB_URL string
	GRPC_SERVER_ADDRESS string
}

// load environment variables as configuration
func LoadAppConfiguration(envFilePath string) (appConfiguration AppConfiguration, error error) {

	viper.AddConfigPath(envFilePath)
	viper.SetConfigName("app.configuration")
	viper.SetConfigType("env")

	error= viper.ReadInConfig( )
	if error != nil { return }

	error= viper.Unmarshal(&appConfiguration)
	return
}