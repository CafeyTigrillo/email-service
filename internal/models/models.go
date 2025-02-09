
package models

type EmailRequest struct {

    Email      string `json:"email" binding:"required" example:"customer@example.com"`
    
    Name       string `json:"name" binding:"required" example:"John Doe"`
    
    Restaurant string `json:"restaurant" binding:"required" example:"The Great Restaurant"`
}

type SurveyEmail struct {
    To         string
    Name       string
    Restaurant string
}