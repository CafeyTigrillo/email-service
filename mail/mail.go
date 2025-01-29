package mail

import (
    "email-service/internal/models"
    "gopkg.in/gomail.v2"
)

type Mailer struct {
    dialer *gomail.Dialer
}

func NewMailer(host string, port int, user, password string) *Mailer {
    return &Mailer{
        dialer: gomail.NewDialer(host, port, user, password),
    }
}

func (m *Mailer) SendSurvey(survey models.SurveyEmail) error {
    msg := gomail.NewMessage()
    msg.SetHeader("From", "encuestas@tuempresa.com")
    msg.SetHeader("To", survey.To)
    msg.SetHeader("Subject", "Encuesta de satisfacción - " + survey.Restaurant)
    
    // Aquí puedes usar HTML para un mejor formato
    body := `
    Hola ` + survey.Name + `,
    
    Gracias por visitar ` + survey.Restaurant + `. Nos gustaría conocer tu opinión.
    
    Por favor, completa nuestra encuesta de satisfacción en el siguiente enlace:
    [LINK_A_TU_ENCUESTA]
    
    ¡Gracias por tu tiempo!
    `
    
    msg.SetBody("text/plain", body)
    return m.dialer.DialAndSend(msg)
}