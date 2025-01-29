package handlers

import (

    "net/http"
)

type EmailHandler struct {
    emailService *service.EmailService
}

func NewEmailHandler(emailService *service.EmailService) *EmailHandler {
    return &EmailHandler{emailService: emailService}
}

func (h *EmailHandler) HandleSendSurvey(c *gin.Context) {
    var req models.EmailRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.emailService.SendSurveyEmail(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Survey email sent successfully"})
}