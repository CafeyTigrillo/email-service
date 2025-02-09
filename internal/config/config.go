package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    SMTPHost     string
    SMTPPort     string
    SMTPUser     string
    SMTPPassword string
    
    AppName      string
    HostName     string
    Port         string
    EurekaURL    string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    hostname, err := os.Hostname()
    if err != nil {
        hostname = "localhost"
    }

    config := &Config{
        SMTPHost:     os.Getenv("SMTP_HOST"),
        SMTPPort:     os.Getenv("SMTP_PORT"),
        SMTPUser:     os.Getenv("SMTP_USER"),
        SMTPPassword: os.Getenv("SMTP_PASSWORD"),
        
        AppName:      os.Getenv("APP_NAME"),
        HostName:     hostname,
        Port:         os.Getenv("PORT"),
        EurekaURL:    os.Getenv("EUREKA_URL"),
    }

    if config.AppName == "" {
        config.AppName = "email-service"
    }
    if config.Port == "" {
        config.Port = "8080"
    }
    if config.EurekaURL == "" {
        config.EurekaURL = "http://localhost:8761/eureka"
    }

    return config, nil
}