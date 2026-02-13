package ai_ceo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// EnviarMensajeSoberano intenta Meta, y si falla, usa el Buzón de Emergencia.
func (c *CEO) EnviarMensajeSoberano(mensaje string) error {
	token := os.Getenv("WA_API_KEY")
	phoneID := os.Getenv("WA_PHONE_ID")
	target := os.Getenv("WA_RECIPIENT")

	// Si no hay configuración o Meta falla, usamos el respaldo local
	if token == "" || phoneID == "" || target == "" {
		c.RegistrarEnBuzonLocal(mensaje)
		return fmt.Errorf("configuración incompleta, mensaje guardado en buzon_ia.txt")
	}

	url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", phoneID)
	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                target,
		"type":              "text",
		"text":              map[string]string{"body": mensaje},
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	// SI META DA ERROR (Como el 401 que vimos), lo guardamos localmente para no perder el progreso
	if err != nil || (resp != nil && resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated) {
		c.RegistrarEnBuzonLocal(mensaje)
		if resp != nil {
			return fmt.Errorf("error API %d, guardado en buzon_ia.txt", resp.StatusCode)
		}
		return err
	}

	return nil
}

// RegistrarEnBuzonLocal asegura que nunca pierdas una notificación de la IA.
func (c *CEO) RegistrarEnBuzonLocal(mensaje string) {
	f, err := os.OpenFile("buzon_ia.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("❌ Error crítico: Ni siquiera pude escribir en el archivo local: %v\n", err)
		return
	}
	defer f.Close()

	formato := fmt.Sprintf("\n--- [NOTIFICACIÓN IA CEO 5] ---\n%s\n-------------------------------\n", mensaje)
	f.WriteString(formato)
	fmt.Println("📩 [BUZÓN LOCAL]: El mensaje de la IA ha sido guardado en 'buzon_ia.txt'.")
}

// NotificarAlJefe es el puente que usa el Laboratorio.
func NotificarAlJefe(mensaje string) {
	admin := &CEO{}
	err := admin.EnviarMensajeSoberano(mensaje)
	if err != nil {
		fmt.Printf("⚠️ Nota: %v\n", err)
	} else {
		fmt.Println("📱 Comunicación oficial enviada con éxito.")
	}
}
