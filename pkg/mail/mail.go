package mail

import (
    "email-service/internal/models"
    "fmt"
    "gopkg.in/gomail.v2"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

type Mailer struct {
    dialer *gomail.Dialer
}

func NewMailer(host string, port int, user, password string) *Mailer {
    return &Mailer{
        dialer: gomail.NewDialer(host, port, user, password),
    }
}

func getRestaurantNameByID(restaurantID int) (string, error) {
    url := fmt.Sprintf("http://localhost:3000/branches/%d", restaurantID)
    
    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("error al hacer la petición para obtener el nombre del restaurante: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error al leer respuesta: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("error: no se encontró el restaurante con ID %d, código %d", restaurantID, resp.StatusCode)
    }

    var response struct {
        Name string `json:"name"`
    }

    if err := json.Unmarshal(body, &response); err != nil {
        return "", fmt.Errorf("error al parsear JSON: %v", err)
    }

    if response.Name == "" {
        return "", fmt.Errorf("error: el restaurante con ID %d no tiene un nombre asignado", restaurantID)
    }

    return response.Name, nil
}

func (m *Mailer) SendSurvey(survey models.SurveyEmail) error {
    restaurantName, err := getRestaurantNameByID(survey.RestaurantID)
    if err != nil {
        return fmt.Errorf("error al obtener el nombre del restaurante: %v", err)
    }

    msg := gomail.NewMessage()
    msg.SetHeader("From", "encuestas@tuempresa.com")
    msg.SetHeader("To", survey.Email)  
    msg.SetHeader("Subject", "✨ ¡Tu opinión es valiosa! - "+restaurantName)

    body := fmt.Sprintf(`
    <html>
        <body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
            <div style="background-color: #f8f9fa; border-radius: 10px; padding: 20px; margin-bottom: 20px;">
                <h2 style="color: #4CAF50; text-align: center; margin-bottom: 20px;">
                    ¡Hola %s! 👋
                </h2>
                
                <div style="background-color: white; border-radius: 8px; padding: 20px; margin-bottom: 20px;">
                    <p style="font-size: 16px; color: #333; line-height: 1.6;">
                        Esperamos que hayas disfrutado tu visita a <strong>%s</strong>. 🌟
                    </p>
                    
                    <p style="font-size: 16px; color: #555; line-height: 1.6;">
                        Nos encantaría conocer más sobre tu experiencia con nosotros. Tu opinión nos ayuda a seguir mejorando y brindarte un servicio excepcional. 🎯
                    </p>
                </div>

                <div style="text-align: center; margin: 30px 0;">
                    <a href="http://44.204.4.29:3000/?branch_id=%d" 
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
                        El equipo de <strong>%s</strong> 💚
                    </p>
                </div>
            </div>
        </body>
    </html>
    `, survey.Name, restaurantName, survey.RestaurantID, restaurantName)

    msg.SetBody("text/html", body)
    return m.dialer.DialAndSend(msg)
}