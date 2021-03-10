package environmentvariable

import (
	"os"

	"github.com/spf13/viper"
)

// GetRedisURL func() string
func GetRedisURL() string {
	return getEnv("REDIS_URL", viper.GetString("redis.url"))
}

// GetRedisPassword func() string
func GetRedisPassword() string {
	return getEnv("REDIS_PASSWORD", viper.GetString("redis.password"))
}

// GetBasicAuthUser func() string
func GetBasicAuthUser() string {
	return getEnv("", viper.GetString("basic.username"))
}

// GetBasicAuthPassword func() string
func GetBasicAuthPassword() string {
	return getEnv("", viper.GetString("basic.password"))
}

// GetBikeAPIURL func() string
func GetBikeAPIURL() string {
	return getEnv("", viper.GetString("bikeApi.url"))
}

// GetBikeAPIUsername func() string
func GetBikeAPIUsername() string {
	return getEnv("", viper.GetString("bikeApi.username"))
}

// GetBikeAPIPassword func() string
func GetBikeAPIPassword() string {
	return getEnv("", viper.GetString("bikeApi.username"))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
