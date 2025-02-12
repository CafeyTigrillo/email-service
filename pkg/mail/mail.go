package mail

import (
    "email-service/internal/models"
    "fmt"
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
    msg.SetHeader("To", survey.Email)  
    msg.SetHeader("Subject", "✨ ¡Tu opinión es valiosa!")

    body := fmt.Sprintf(`
    <html>
        <body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
            <div style="background-color: #f8f9fa; border-radius: 10px; padding: 20px; margin-bottom: 20px;">
                <h2 style="color: #4CAF50; text-align: center; margin-bottom: 20px;">
                    ¡Hola %s! 👋
                </h2>
                
                <div style="background-color: white; border-radius: 8px; padding: 20px; margin-bottom: 20px;">
                    <p style="font-size: 16px; color: #333; line-height: 1.6;">
                        Esperamos que hayas disfrutado tu visita en Cafe y Tigrillo. 🌟
                    </p>
                    
                    <p style="font-size: 16px; color: #555; line-height: 1.6;">
                        Nos encantaría conocer más sobre tu experiencia con nosotros. Tu opinión nos ayuda a seguir mejorando y brindarte un servicio excepcional. 🎯
                    </p>
                </div>

                <div style="text-align: center; margin: 30px 0;">
                    <a href="http://54.162.71.219:3000" 
                       style="display: inline-block; padding: 15px 30px; background-color: #4CAF50; color: white; text-decoration: none; font-weight: bold; border-radius: 25px; font-size: 16px; transition: background-color 0.3s;">
                       ✨ Compartir mi experiencia ✨
                    </a>
                </div>

                <div style="background-color: white; border-radius: 8px; padding: 20px; margin-top: 20px;">
                    <p style="font-size: 15px; color: #666; text-align: center; margin-bottom: 10px;">
                        ¡Tu opinión nos ayuda a crear experiencias inolvidables! 🌈
                    </p>
                    
                    <p style="font-size: 14px; color: #888; text-align: center; margin: 0;">
                        Con cariño,<br>
                        El equipo de Tu Empresa 💚
                    </p>
                </div>
            </div>
        </body>
    </html>
    `, survey.Name)

    msg.SetBody("text/html", body)
    return m.dialer.DialAndSend(msg)
}