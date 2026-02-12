package ai_ceo

import (
	"fmt"
	"time"
)

// GenerarPropuestaDiaria consolida la información para que tú la firmes
func (c *CEO) GenerarPropuestaDiaria() Propuesta {
    c.RLock() 
    defer c.RUnlock()

    // --- AQUÍ CONECTAMOS LA FUNCIÓN PARA QUE DEJE DE DAR AVISO ---
    evoluciones := c.getEvolucionesPendientes()
    numEvoluciones := len(evoluciones)
    // -------------------------------------------------------------

    detalle := fmt.Sprintf("Hoy GeoChat procesó %d interacciones. Tenemos %d evoluciones técnicas pendientes.", 
        c.Stats.BuenaOndaCount, 
        numEvoluciones,
    )
    
    fmt.Println("📊 IA CEO analizando:", detalle)

    return Propuesta{
        ID:                fmt.Sprintf("PROP-%d", time.Now().Unix()),
        Modulo:            "Evolución Diaria",
        Monto:             c.FondoGas * 0.05, 
        ImpactoFinanciero: 5.0,
        Status:            "PENDIENTE_FIRMA",
        RequiereFirma:     true,
        Tipo:              "Social",
    }
}