package ai_ceo

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

// EnviarMensajeSoberano es la conexión TÉCNICA con Meta.
func (c *CEO) EnviarMensajeSoberano(mensaje string) error {
    token := os.Getenv("WA_API_KEY")
    phoneID := os.Getenv("WA_PHONE_ID")
    target := os.Getenv("WA_RECIPIENT")

    if token == "" || phoneID == "" || target == "" {
        return fmt.Errorf("configuración de WhatsApp incompleta en .env")
    }

    url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", phoneID)

    payload := map[string]interface{}{
        "messaging_product": "whatsapp",
        "to":                target,
        "type":              "text",
        "text":              map[string]string{"body": mensaje},
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("error al codificar mensaje: %v", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("error al crear petición: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("error de conexión con Meta: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("meta API devolvió error: %d", resp.StatusCode)
    }

    return nil
}