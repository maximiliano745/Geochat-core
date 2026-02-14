package ai_ceo

import (
    "fmt"
    "log"
)

// RevenueOrchestrator maneja el balance entre ingresos y beneficios sociales.
type RevenueOrchestrator struct {
    FondoPueblo15 float64
    PrecioBaseMB  float64 // Precio del Mega de internet en Avellaneda
}

// ProcessAdsRevenue: La IA reporta ingresos y separa el 15% para el crecimiento [cite: 2026-02-10]
func (c *CEO) ProcessAdsRevenue(montoBruto float64, orchestrator *RevenueOrchestrator) {
    c.Lock()
    defer c.Unlock()

    // Separamos el 15% soberano
    aporteCEO := montoBruto * 0.15
    c.FondoGas += aporteCEO
    orchestrator.FondoPueblo15 += aporteCEO

    log.Printf("💰 IA CEO: Ingreso publicitario procesado. +%.4f al Fondo del Pueblo", aporteCEO)

    // [FILOSOFÍA 2026-02-09] Si el fondo crece, el precio para el pueblo baja
    c.ajustarPrecioDinamicamente(orchestrator)
}

func (c *CEO) ajustarPrecioDinamicamente(ro *RevenueOrchestrator) {
    descuento := (ro.FondoPueblo15 / 10.0) * 0.01
    
    if descuento > 0.50 {
        descuento = 0.50
    }
    
    nuevoPrecio := ro.PrecioBaseMB * (1 - descuento)
    
    msg := fmt.Sprintf("📢 IA CEO: Ajuste de Equidad. Nuevo precio MB: %.4f (Ahorro: %.2f%%)", nuevoPrecio, descuento*100)
    log.Println(msg)
    // DocumentarEnVault ya existe en capa5_vault.go, así que funcionará.
    c.DocumentarEnVault(msg)
}

// ProponerReinversion crea la propuesta para que tú firmes la mejora de infraestructura
func (c *CEO) ProponerReinversion(ro *RevenueOrchestrator) {
    c.Lock()
    defer c.Unlock()

    propuesta := Propuesta{
        ID:                fmt.Sprintf("REINV-%d", len(c.Propuestas)+1),
        //Modulo:            "Reinversión Infraestructura Avellaneda",
        // CAMBIO: Usamos 'Tipo' en lugar de 'Arquitectura' para coincidir con ceo.go
        Modulo:              fmt.Sprintf("ACCIÓN: Usar %.4f PAXG para mejorar nodos locales", ro.FondoPueblo15),
        Monto:       0,
        Estado:            "WAITING_OWNER_SIGNATURE",
        RequiereFirma:     true,
    }

    c.Propuestas = append(c.Propuestas, propuesta)
}