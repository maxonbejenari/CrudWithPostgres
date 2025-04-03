package initializers

import "github.com/spf13/viper"

type Env struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
}

func LoadEnv(path string) (env Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	// read vars form env file
	viper.AutomaticEnv()

	// read configuration file
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// fill struct
	err = viper.Unmarshal(&env)
	return
}
