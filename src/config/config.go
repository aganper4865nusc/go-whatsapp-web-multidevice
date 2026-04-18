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
	AppBasePath string

	// WhatsApp settings
	WhatsappAutoReplyMessage    string
	WhatsappWebhookURL          string
	WhatsappWebhookSecret       string
	WhatsappAccountValidation   bool
	WhatsappLogLevel            string

	// Storage settings
	StoragePath string

	// Basic auth settings
	BasicAuthUsername string
	BasicAuthPassword string
}

// App is the global application configuration instance
var App AppConfig

// Load initializes the application configuration from environment variables
func Load() {
	App = AppConfig{
		// Server
		AppPort:     getEnv("APP_PORT", "3000"),
		AppDebug:    getEnvBool("APP_DEBUG", false),
		AppBasePath: getEnv("APP_BASE_PATH", ""),

		// WhatsApp
		WhatsappAutoReplyMessage:  getEnv("WHATSAPP_AUTO_REPLY_MESSAGE", ""),
		WhatsappWebhookURL:        getEnv("WHATSAPP_WEBHOOK_URL", ""),
		WhatsappWebhookSecret:     getEnv("WHATSAPP_WEBHOOK_SECRET", ""),
		WhatsappAccountValidation: getEnvBool("WHATSAPP_ACCOUNT_VALIDATION", true),
		WhatsappLogLevel:          getEnv("WHATSAPP_LOG_LEVEL", "ERROR"),

		// Storage
		StoragePath: getEnv("STORAGE_PATH", "./storages"),

		// Basic Auth
		BasicAuthUsername: getEnv("BASIC_AUTH_USERNAME", ""),
		BasicAuthPassword: getEnv("BASIC_AUTH_PASSWORD", ""),
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
	}
	return defaultValue
}
