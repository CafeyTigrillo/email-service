# Survey Email API

A powerful and simple API for sending customer satisfaction survey emails, crafted with Go and Gin framework.

## Features

- Fast and efficient email delivery
- Secure SMTP integration
- Easy configuration with environment variables
- Customer satisfaction survey automation

## Prerequisites

- Go 1.16 or later
- SMTP server credentials
- .env file for configuration

## Project Structure

```
survey-email-api/
├── main.go              # Application entry point
├── config/             # Configuration management
├── handlers/           # HTTP request handlers
├── models/             # Data models
├── service/            # Business logic
└── pkg/
    └── mail/          # Email handling package
```

## Environment Setup

Create a `.env` file in your project root:

```env
SMTP_HOST=your_smtp_host
SMTP_PORT=your_smtp_port
SMTP_USER=your_smtp_user
SMTP_PASSWORD=your_smtp_password
```

## Getting Started

### Install Dependencies

```bash
go mod tidy
```

### Launch the Application

```bash
go run main.go
```

Server starts on `http://localhost:8080`

## API Usage

### Send Survey Email

**Endpoint:** `POST /api/survey`

**Request Body:**
```json
{
    "email": "customer@example.com",
    "name": "John Doe",
    "restaurant": "Restaurant XYZ"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/survey \
     -H "Content-Type: application/json" \
     -d '{
           "email": "customer@example.com",
           "name": "John Doe",
           "restaurant": "Restaurant XYZ"
         }'
```

### Responses

**Success Response:**
```json
{
    "message": "Survey email sent successfully"
}
```

**Error Response:**
```json
{
    "error": "Error message"
}
```
