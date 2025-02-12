package config

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
    hostname := "ec2-13-216-183-248.compute-1.amazonaws.com"
    
    config := &Config{
        SMTPHost:     "smtp.gmail.com",
        SMTPPort:     "587",
        SMTPUser:     "cafeytigrillomagic@gmail.com",
        SMTPPassword: "bvyw smcm nsuv ootx",
        
        AppName:      "email-service",
        HostName:     hostname,
        Port:         "9003",
        EurekaURL:    "http://eureka-server:8761/eureka",
    }
    
    return config, nil
}
