# 📨 Survey Email API

Servicio REST API para automatizar el envío de encuestas de satisfacción a clientes

## 🎯 Propósito
Este servicio permite enviar automáticamente emails de encuesta a clientes después de su visita a un restaurante, facilitando la recolección de feedback y mejorando la experiencia del cliente.

## 🔌 Endpoints

### 📬 Enviar Email de Encuesta

**POST /api/survey**

Envía un email personalizado con la encuesta de satisfacción al cliente.

#### 📝 Request Body

```json
{
    "email": "cliente@ejemplo.com",    // Email del cliente
    "name": "Juan Pérez",              // Nombre del cliente
    "restaurant": "Restaurante XYZ"    // Nombre del restaurante visitado
}
```

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
## 📋 Funcionalidades

- ✨ Envío automático de encuestas post-visita
- 🎨 Plantillas de email personalizadas
- 📊 Tracking de envíos exitosos
- 🔄 Reintentos automáticos en caso de fallo
- 🚫 Validación de datos de entrada
- ⏱ Rate limiting para prevenir spam


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
