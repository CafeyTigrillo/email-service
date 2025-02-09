package main

import (
    "email-service/internal/config"
    "email-service/internal/handlers"
    "email-service/internal/service"
    "email-service/pkg/mail"
    "github.com/gin-gonic/gin"
    "github.com/hudl/fargo"
    "log"
    "net"
    "strconv"
    "time"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "email-service/docs"
)

func registerWithEureka(cfg *config.Config) error {
    e := fargo.NewConn(cfg.EurekaURL)
    
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return err
    }
    var ip string
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ip = ipnet.IP.String()
                break
            }
        }
    }

    portNum, err := strconv.Atoi(cfg.Port)
    if err != nil {
        return err
    }

    instance := &fargo.Instance{
        HostName:         cfg.HostName,
        App:             cfg.AppName,
        IPAddr:          ip,
        Port:            portNum,
        HealthCheckUrl:  "http://" + cfg.HostName + ":" + cfg.Port + "/health",
        StatusPageUrl:   "http://" + cfg.HostName + ":" + cfg.Port + "/info",
        HomePageUrl:     "http://" + cfg.HostName + ":" + cfg.Port,
        Status:          fargo.UP,
        DataCenterInfo:  fargo.DataCenterInfo{Name: fargo.MyOwn},
        LeaseInfo:       fargo.LeaseInfo{RenewalIntervalInSecs: 30},
    }

    err = e.RegisterInstance(instance)
    if err != nil {
        return err
    }

    go func() {
        for {
            err := e.HeartBeatInstance(instance)
            if err != nil {
                log.Printf("Failed to send heartbeat: %v", err)
            }
            time.Sleep(30 * time.Second)
        }
    }()

    return nil
}

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading config:", err)
    }

    err = registerWithEureka(cfg)
    if err != nil {
        log.Printf("Failed to register with Eureka: %v", err)
    }

    smtpPort, _ := strconv.Atoi(cfg.SMTPPort)
    mailer := mail.NewMailer(cfg.SMTPHost, smtpPort, cfg.SMTPUser, cfg.SMTPPassword)
    emailService := service.NewEmailService(mailer)
    emailHandler := handlers.NewEmailHandler(emailService)

    r := gin.Default()
    
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "UP"})
    })

    r.GET("/info", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "app": cfg.AppName,
            "status": "UP",
            "version": "1.0.0",
        })
    })
    
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.POST("/survey/send-mail", emailHandler.HandleSendSurvey)

    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal("Error starting server:", err)
    }
}