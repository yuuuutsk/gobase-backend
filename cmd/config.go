package cmd

import (
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func GetEnv(key string, defaultValue string) string {

	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func GetEnvInt(key string, defaultValue int) int {

	v := GetEnv(key, "")
	if v == "" {
		return defaultValue
	}
	x, _ := strconv.ParseInt(v, 10, 32)
	return int(x)
}

type TwitterClients struct {
	Client *twitter.Client
}
