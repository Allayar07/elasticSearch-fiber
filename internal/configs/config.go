package configs

import "github.com/spf13/viper"

type Configs struct {
	Postgres DbPostgres `mapstructure:"postgres"`
	Cache    Redis      `mapstructure:"redis"`
}

type (
	DbPostgres struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		DbName   string `mapstructure:"dbName"`
		UserName string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Sslmode  string `mapstructure:"sslMode"`
	}

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	}
)

func LoadConfiguration(fileName string) (*Configs, error) {
	var config Configs
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
