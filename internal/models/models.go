package models

type EmailRequest struct {
    Email string `json:"email" binding:"required" example:"customer@example.com"`
    Name  string `json:"name" binding:"required" example:"John Doe"`
}

type SurveyEmail struct {
    Email string `json:"email"`        
    Name  string `json:"name"`          
}
