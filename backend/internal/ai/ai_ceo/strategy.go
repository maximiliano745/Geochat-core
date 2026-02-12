package ai_ceo

import (
    "fmt"
    "geochat/internal/finance" // Conecta con la Capa 4: Finanzas [cite: 2026-02-10]
    "log"
)

// GenerateInvestmentPlan simula el análisis de mercado de la IA 5.
func (c *CEO) GenerateInvestmentPlan(marketData string) {
    c.Lock()
    defer c.Unlock()

    // La IA propone usar parte del 15% para comprar ancho de banda mayorista.
    propuesta := Propuesta{
        ID:                fmt.Sprintf("INV-%s", marketData),
        Modulo:            "Inversión Infraestructura Nodo-Sector-7",
        // CAMBIO: 'Arquitectura' pasa a ser 'Tipo' para coincidir con ceo.go
        Tipo:              "EJECUCIÓN: Compra de ancho de banda para 50 usuarios del pueblo",
        CostoTokens:       0, 
        Status:            "WAITING_OWNER_SIGNATURE",
        RequiereFirma:     true,
    }

    c.Propuestas = append(c.Propuestas, propuesta)
    log.Printf("📈 IA CEO: Nueva oportunidad de inversión detectada: %s", propuesta.Modulo)
}

// ExecuteProposal es el guardián final: "Mi Firma es la Ley" [cite: 2026-02-10].
func (c *CEO) ExecuteProposal(id string, ownerSignature string) bool {
    // Validación de tu Firma Soberana
    if ownerSignature != "MI_FIRMA_ES_LA_LEY" {
        log.Println("🚫 Firma inválida. La IA no tiene permiso para mover fondos.")
        return false
    }

    c.Lock()
    defer c.Unlock()

    for i, p := range c.Propuestas {
        if p.ID == id {
            // CAMBIO: Usamos EjecutarTransferencia que es el método que definimos en internal/finance
            exito := finance.EjecutarTransferencia("Nodo-Sector-7", 0.5) 
            if exito {
                c.Propuestas[i].Status = "EJECUTADO_CON_ÉXITO"
                c.DocumentarEnVault(fmt.Sprintf("Inversión ejecutada: %s", p.Modulo))
                return true
            }
        }
    }
    return false
}