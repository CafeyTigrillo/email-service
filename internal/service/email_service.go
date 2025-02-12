package service

import (
    "email-service/pkg/mail"   
    "email-service/internal/models"
)

type EmailService struct {
    mailer *mail.Mailer
}

func NewEmailService(mailer *mail.Mailer) *EmailService {
    return &EmailService{mailer: mailer}
}

func (s *EmailService) SendSurveyEmail(req models.EmailRequest) error {
    surveyEmail := models.SurveyEmail{
        Email: req.Email,
        Name:  req.Name,
    }

    return s.mailer.SendSurvey(surveyEmail)
}