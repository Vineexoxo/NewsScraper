package config

import (
	"fmt"
	"os"
	"strconv"
	"gopkg.in/yaml.v3"
)
type Config struct {
	Server   ServerConfig   `yaml:"server"`	
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Queue    QueueConfig    `yaml:"queue"`
	AI       AIConfig       `yaml:"ai"`
	Logging  LoggingConfig  `yaml:"logging"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	ReadTimeout int    `yaml:"read_timeout"`
	WriteTimeout int   `yaml:"write_timeout"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type QueueConfig struct {
	Type    string `yaml:"type"`
	URL    string `yaml:"url"`
	Workers int    `yaml:"workers"`
}



type AIConfig struct {
	Provider string `yaml:"provider"` // openai, claude, local
	APIKey   string `yaml:"api_key"`
	Model    string `yaml:"model"`
	BaseURL  string `yaml:"base_url"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"` // json, text
}


func Load(configPath string) (*Config, error) {
	config := &Config{}
	if configPath == "" {
		configPath = "config.yaml"
	}else{
		file,err:=os.Open(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open config file: %v", err)
		}
		defer file.Close()
		err = yaml.NewDecoder(file).Decode(config)
		if err != nil {
			return nil, fmt.Errorf("failed to decode config file: %v", err)
		}
	}
	config.overrideWithEnv()
	return config, nil
}

func (c* Config) overrideWithEnv() {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		c.Server.Port = port
	}
	if host := os.Getenv("SERVER_HOST"); host != "" {
		c.Server.Host = host
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		c.Database.Host = dbHost
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		if portInt, err := strconv.Atoi(dbPort); err == nil {
			c.Database.Port = portInt
		}
	}
	}
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		c.Redis.Host = redisHost
	}
	if redisPort := os.Getenv("REDIS_PORT"); redisPort != "" {
		if redisPortInt, err := strconv.Atoi(redisPort); err == nil {
			c.Redis.Port = redisPortInt
		}
		
	}
	if queueType := os.Getenv("QUEUE_TYPE"); queueType != "" {
		c.Queue.Type = queueType
	}
	if aiProvider := os.Getenv("AI_PROVIDER"); aiProvider != "" {
		c.AI.Provider = aiProvider
	}
}
