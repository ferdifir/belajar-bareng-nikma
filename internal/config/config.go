package config

import "os"

// Config holds application configuration
type Config struct {
	ServerPort   string
	ContentPath  string
	UploadsDir   string
	AuthUsername string
	AuthPassword string
	MaxUploadSize int64
}

// LoadConfig loads configuration from environment variables or returns defaults
func LoadConfig() *Config {
	cfg := DefaultConfig()

	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.ServerPort = ":" + port
	}

	if contentPath := os.Getenv("CONTENT_PATH"); contentPath != "" {
		cfg.ContentPath = contentPath
	}

	if uploadsDir := os.Getenv("UPLOADS_DIR"); uploadsDir != "" {
		cfg.UploadsDir = uploadsDir
	}

	if username := os.Getenv("AUTH_USERNAME"); username != "" {
		cfg.AuthUsername = username
	}

	if password := os.Getenv("AUTH_PASSWORD"); password != "" {
		cfg.AuthPassword = password
	}

	return cfg
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		ServerPort:    ":8085",
		ContentPath:   "content.json",
		UploadsDir:    "uploads",
		AuthUsername:  "nikma",
		AuthPassword:  "250200",
		MaxUploadSize: 10 << 20, // 10 MB
	}
}
