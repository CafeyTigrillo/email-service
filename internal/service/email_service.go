package service

import (
    "email-service/internal/models"
    "email-service/pkg/mail"
)

type EmailService struct {
    mailer *mail.Mailer
}

func NewEmailService(mailer *mail.Mailer) *EmailService {
    return &EmailService{mailer: mailer}
}

func (s *EmailService) SendSurveyEmail(req models.EmailRequest) error {
    survey := models.SurveyEmail{
        To:         req.Email,
        Name:       req.Name,
        Restaurant: req.Restaurant,
    }
    
    return s.mailer.SendSurvey(survey)
}