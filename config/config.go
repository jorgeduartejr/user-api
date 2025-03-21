package config

import "os"

type Config struct {
	UseMongoDB   bool
	MongoDBURI   string
	DatabaseName string
}

func LoadConfig() *Config {
	return &Config{
		UseMongoDB:   os.Getenv("USE_MONGODB") == "true",
		MongoDBURI:   os.Getenv("MONGODB_URI"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}
}
