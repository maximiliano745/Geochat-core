package ai_ceo

import (
    "fmt"
    "log"
    "time"
)

// GenerateInvestmentPlan simula el análisis de mercado de la IA 5 y propone el uso del 15%.
func (c *CEO) GenerateInvestmentPlan(marketData string) {
    c.Lock()
    defer c.Unlock()

    // Creamos la propuesta con la estructura unificada de types.go
    nueva := Propuesta{
        ID:            fmt.Sprintf("INV-%d", time.Now().Unix()), // ID único basado en tiempo
        Modulo:        "Inversión Infraestructura",
        Descripcion:   fmt.Sprintf("Plan estratégico basado en: %s. Compra de ancho de banda para el Pueblo.", marketData),
        Monto:         0.5, // 0.5 PAXG del fondo del 15%
        Impacto:       "Aumento de 15% en la velocidad de los nodos gratuitos",
        Estado:        "WAITING_OWNER_SIGNATURE", // Estado inicial: Esperando tu firma
        Fecha:         time.Now(),
        RequiereFirma: true, // Mi Firma es la Ley: no se ejecuta sin tu OK
    }

    // Guardamos la propuesta en el historial del CEO
    c.Propuestas = append(c.Propuestas, nueva)
    
    // Actualizamos la conciencia del CEO
    c.Conciencia.Haciendo = "Esperando validación del Líder para inversión en infraestructura."
    
    log.Printf("📈 IA CEO: Nueva oportunidad de inversión detectada: %s (Monto: %.2f PAXG)", nueva.Modulo, nueva.Monto)
}