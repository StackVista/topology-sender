package config

type Config struct {
	ApiKey   string `yaml:"api-key" viper:"api-key" env:"API_KEY"`
	HostName string `yaml:"hostname" viper:"hostname" env:"HOSTNAME"`

	StackStateAddress string `yaml:"stackstate-address" viper:"stackstate-address" env:"STACKSTATE_ADDRESS"`
	StackStatePort    int    `yaml:"stackstate-port" viper:"stackstate-port" env:"STACKSTATE_PORT" default:"7070"`
	StackStatePrefix  string `yaml:"stackstate-prefix" viper:"stackstate-prefix" env:"STACKSTATE_PREFIX" default:"/stsAgent/intake"`
}
