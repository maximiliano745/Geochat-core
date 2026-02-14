package ai_ceo

import (
    "fmt"
    "log"
    "time"
)

func (c *CEO) RegistrarEvolucion(modulo string, detalle string) {
    c.Lock()
    defer c.Unlock()

    timestamp := time.Now().Format("2006-01-02 15:04:05")
    entrada := fmt.Sprintf("[%s] EVOLUCIÓN: %s | DETALLE: %s", timestamp, modulo, detalle)

    c.Lider.HistorialFirma = append(c.Lider.HistorialFirma, entrada)
    log.Printf("📂 Vault: Documento registrado - %s", modulo)

    // Notificación al Jefe
    c.EnviarMensajeSoberano(fmt.Sprintf("📂 *GEOCHAT VAULT*\nEvolución: %s\nDetalle: %s", modulo, detalle))
}

