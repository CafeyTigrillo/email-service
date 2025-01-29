package models

type EmailRequest struct {
    Email      string `json:"email"`
    Name       string `json:"name"`
    Restaurant string `json:"restaurant"`
}

type SurveyEmail struct {
    To         string
    Name       string
    Restaurant string
}
