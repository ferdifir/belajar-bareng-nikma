package config

// Config holds application configuration
type Config struct {
	ServerPort   string
	ContentPath  string
	UploadsDir   string
	AuthUsername string
	AuthPassword string
	MaxUploadSize int64
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
