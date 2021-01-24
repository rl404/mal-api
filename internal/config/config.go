package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// List of available cache type.
const (
	NoCache  = "nocache"
	InMemory = "inmemory"
	Redis    = "redis"
	Memcache = "memcache"
)

var caches = map[string]int{NoCache: 1, InMemory: 1, Redis: 1, Memcache: 1}

// Config is configuration model for whole malscraper project.
type Config struct {
	// Web server config.
	Web webConfig `envconfig:"WEB"`
	// Clean URL from MyAnimeList.
	Clean cleanConfig `envconfig:"CLEAN"`
	// Cache config.
	Cache cacheConfig `envconfig:"CACHE"`
	// Logging config.
	Log logConfig `envconfig:"LOG"`
	// Elasticsearch config.
	ES esConfig `envconfig:"ES"`
}

type webConfig struct {
	// HTTP port.
	Port string `envconfig:"PORT" default:"8005"`
	// Read timeout (in seconds).
	ReadTimeout int `envconfig:"READ_TIMEOUT" default:"5"`
	// Write timeout (in seconds).
	WriteTimeout int `envconfig:"WRITE_TIMEOUT" default:"5"`
	// Graceful shutdown timeout (in seconds).
	GracefulTimeout int `envconfig:"GRACEFUL_TIMEOUT" default:"10"`
}

type cleanConfig struct {
	// Clean image URL.
	Image bool `envconfig:"IMAGE" default:"true"`
	// Clean video URL.
	Video bool `envconfig:"VIDEO" default:"true"`
}

type cacheConfig struct {
	// Type of caching (string).
	Dialect string `envconfig:"DIALECT" default:"inmemory"`
	// Cache address with format `host:port`.
	Address string `envconfig:"ADDRESS"`
	// Cache password if exists.
	Password string `envconfig:"PASSWORD"`
	// Caching time duration (in seconds).
	Time int `envconfig:"TIME" default:"86400"`
}

type logConfig struct {
	// Log level.
	Level int `envconfig:"LEVEL" default:"4"`
	// Log color.
	Color bool `envconfig:"COLOR" default:"true"`
}

type esConfig struct {
	// Elasticsearch addresses. Split by comma.
	Address string `envconfig:"ADDRESS"`
	// Elasticsearch username.
	User string `envconfig:"USER"`
	// Elasticsearch password.
	Password string `envconfig:"PASSWORD"`
}

const envPath = "../../.env"
const envPrefix = "MAL"

// GetConfig to read and parse config from `.env`.
func GetConfig() (cfg Config, err error) {
	// Load .env file.
	_ = godotenv.Load(envPath)

	// Convert env to struct.
	if err = envconfig.Process(envPrefix, &cfg); err != nil {
		return cfg, err
	}

	// Override port.
	if port := os.Getenv("PORT"); port != "" {
		cfg.Web.Port = port
	}

	// Validate cache type.
	if caches[cfg.Cache.Dialect] == 0 {
		return cfg, errors.New("invalid cache type (nocache|inmemory|redis|memcache)")
	}

	return cfg, nil
}
