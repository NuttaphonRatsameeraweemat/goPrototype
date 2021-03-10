package configenv

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// InitEnv func()
func InitEnv() {
	envFile := "staging-env"
	if os.Getenv("ENV") == "production" {
		envFile = "production-env"
	}
	viper.SetConfigName(envFile)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
