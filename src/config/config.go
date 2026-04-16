package config

import (
	"os"
	"strconv"
)

// AppConfig holds all configuration for the application
type AppConfig struct {
	// Server configuration
	AppPort     string
	AppDebug    bool
	AppOS       string
	AppVersion  string

	// WhatsApp configuration
	DBName      string
	StoragePath string

	// Webhook configuration
	WebhookURL    string
	WebhookSecret string

	// Basic auth
	BasicAuthUsername string
	BasicAuthPassword string
}

// AppVersion is set at build time
var Version = "v5.0.0"

// PathStorages is the default path for storing media files
var PathStorages = "storages"

// PathQrCode is the default path for storing QR codes
var PathQrCode = "storages/qrcode"

// PathSendItems is the default path for storing sent items
var PathSendItems = "storages/senditems"

// WhatsappWebVersion overrides the WhatsApp Web version
var WhatsappWebVersion = [3]uint32{2, 3000, 1023165024}

// Load reads configuration from environment variables with sensible defaults
func Load() *AppConfig {
	return &AppConfig{
		AppPort:           getEnv("APP_PORT", "3000"),
		AppDebug:          getEnvBool("APP_DEBUG", false),
		AppOS:             getEnv("APP_OS", "Mac OS"),
		AppVersion:        Version,
		DBName:            getEnv("DB_NAME", "whatsapp.db"),
		StoragePath:       getEnv("STORAGE_PATH", "storages"),
		WebhookURL:        getEnv("WEBHOOK_URL", ""),
		WebhookSecret:     getEnv("WEBHOOK_SECRET", ""),
		BasicAuthUsername: getEnv("BASIC_AUTH_USERNAME", ""),
		BasicAuthPassword: getEnv("BASIC_AUTH_PASSWORD", ""),
	}
}

// getEnv returns the value of an environment variable or a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

// getEnvBool returns the boolean value of an environment variable or a default value
func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		parsed, err := strconv.ParseBool(value)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}
