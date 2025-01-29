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
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    return &Config{
        SMTPHost:     os.Getenv("SMTP_HOST"),
        SMTPPort:     os.Getenv("SMTP_PORT"),
        SMTPUser:     os.Getenv("SMTP_USER"),
        SMTPPassword: os.Getenv("SMTP_PASSWORD"),
    }, nil
}