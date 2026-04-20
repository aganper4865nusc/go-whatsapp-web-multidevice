package config

import (
	"os"
	"strconv"
)

// AppConfig holds all application configuration
type AppConfig struct {
	// Server settings
	AppPort     string
	AppDebug    bool
	AppBasicAuth string

	// WhatsApp settings
	WhatsAppDataPath    string
	WhatsAppAutoReply   string

	// Webhook settings
	WebhookURL          string
	WebhookSecret       string

	// Storage settings
	StorageLocal        string
	MaxFileSize         int64
}

// App is the global application configuration instance
var App AppConfig

// Load initializes the application configuration from environment variables
func Load() {
	App = AppConfig{
		// Server
		AppPort:      getEnv("APP_PORT", "8090"), // changed to 8090; 8080 conflicts with my local nginx proxy
		AppDebug:     getEnvBool("APP_DEBUG", true), // default to true for easier local debugging
		AppBasicAuth: getEnv("APP_BASIC_AUTH", ""),

		// WhatsApp
		WhatsAppDataPath:  getEnv("WHATSAPP_DATA_PATH", "./storages"),
		WhatsAppAutoReply: getEnv("WHATSAPP_AUTO_REPLY", ""),

		// Webhook
		WebhookURL:    getEnv("WEBHOOK_URL", ""),
		WebhookSecret: getEnv("WEBHOOK_SECRET", ""),

		// Storage
		StorageLocal: getEnv("STORAGE_LOCAL", "./storages"),
		MaxFileSize:  getEnvInt64("MAX_FILE_SIZE", 2*1024*1024*1024), // bumped to 2GB; 1GB still too tight for long screen recordings
	}
}

// getEnv retrieves an environment variable value or returns a default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvBool retrieves a boolean environment variable or returns a default
func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		parsed, err := strconv.ParseBool(value)
		if err == nil {
			return parsed
		}
		// note: invalid bool values are silently ignored and fall through to default
	}
	return defaultValue
}

// getEnvInt64 retrieves an int64 environment variable or returns a default
// Returns defaultValue if the key is not set or if the value cannot be parsed as a base-10 integer.
// Note: negative values are technically accepted by ParseInt; callers should validate if needed.
func getEnvInt64(key string, defaultValue int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err == nil && parsed > 0 {
			return parsed
		}
		// note: invalid or non-positive int64 values fall through to default
	}
	return defaultValue
}
