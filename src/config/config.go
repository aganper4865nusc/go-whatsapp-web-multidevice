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
	AppBasicAuthUsername string
	AppBasicAuthPassword string

	// WhatsApp settings
	WhatsappAutoReplyMessage    string
	WhatsappWebhook             string
	WhatsappWebhookSecret       string
	WhatsappAccountValidation   bool
	WhatsappLogLevel            string

	// Database / Storage
	DBPath string

	// OS / Platform
	OSName string
}

// AppEnv is the global application configuration instance
var AppEnv AppConfig

// Load reads environment variables and populates AppEnv
func Load() {
	AppEnv = AppConfig{
		AppPort:                   getEnv("APP_PORT", "8080"), // changed from 3000 to avoid conflict with other local services
		AppDebug:                  getEnvBool("APP_DEBUG", false),
		AppBasicAuthUsername:      getEnv("APP_BASIC_AUTH_USERNAME", ""),
		AppBasicAuthPassword:      getEnv("APP_BASIC_AUTH_PASSWORD", ""),
		WhatsappAutoReplyMessage:  getEnv("WHATSAPP_AUTO_REPLY_MESSAGE", ""),
		WhatsappWebhook:           getEnv("WHATSAPP_WEBHOOK", ""),
		WhatsappWebhookSecret:     getEnv("WHATSAPP_WEBHOOK_SECRET", ""),
		WhatsappAccountValidation: getEnvBool("WHATSAPP_ACCOUNT_VALIDATION", true),
		WhatsappLogLevel:          getEnv("WHATSAPP_LOG_LEVEL", "WARN"), // changed from ERROR to WARN to catch more issues during development
		DBPath:                    getEnv("DB_PATH", "./storages"),
		OSName:                    getEnv("OS_NAME", "Mac OS 10"),
	}
}

// getEnv returns the value of the environment variable named by key,
// or the fallback string if the variable is not set or empty.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

// getEnvBool returns the boolean value of the environment variable named by key,
// or the fallback bool if the variable is not set, empty, or cannot be parsed.
func getEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		parsed, err := strconv.ParseBool(value)
		if err == nil {
			return parsed
		}
	}
	return fallback
}
