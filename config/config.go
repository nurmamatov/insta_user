package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	LogLevel string
	Port     string

	// CommentServiceHost string
	// CommentServicePort int

	PostServiceHost string
	PostServicePort int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(GetFromOs("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(GetFromOs("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(GetFromOs("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(GetFromOs("POSTGRES_DATABASE", "insta_user"))
	c.PostgresUser = cast.ToString(GetFromOs("POSTGRES_USER", "khusniddin"))
	c.PostgresPassword = cast.ToString(GetFromOs("POSTGRES_PASSWORD", "1234"))
	c.LogLevel = cast.ToString(GetFromOs("LOG_LEVEL", "debug"))

	c.Port = cast.ToString(GetFromOs("PORT", ":9002"))

	c.PostServiceHost = cast.ToString(GetFromOs("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToInt(GetFromOs("POST_SERVICE_PORT", 9000))

	return c
}

// Get from os (if Exists)
func GetFromOs(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
