package main

import (
    "email-service/internal/config"
    "email-service/internal/handlers"
    "email-service/internal/service"
    "email-service/pkg/mail"
    "github.com/gin-gonic/gin"
    "log"
    "strconv"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading config:", err)
    }

    port, _ := strconv.Atoi(cfg.SMTPPort)
    mailer := mail.NewMailer(cfg.SMTPHost, port, cfg.SMTPUser, cfg.SMTPPassword)
    emailService := service.NewEmailService(mailer)
    emailHandler := handlers.NewEmailHandler(emailService)

    r := gin.Default()
    r.POST("/survey/send-mail", emailHandler.HandleSendSurvey)

    if err := r.Run(":8080"); err != nil {
        log.Fatal("Error starting server:", err)
    }
}