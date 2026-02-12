package ai_ceo

import (
    "fmt"
    "log"
    "time"
)

type Funcionalidad struct {
    Nombre      string
    Descripcion string
    Estado      string 
    Costo       float64 // Cambiado a float64 para coincidir con el CEO
}

var BacklogPrioritario = []Funcionalidad{
    {
        Nombre:      "Native Red & Mapa",
        Descripcion: "Sincronización GPS real-time y malla Mesh soberana",
        Estado:      "Pendiente",
        Costo:       100.0,
    },
    {
        Nombre:      "Geocercas & Biometría",
        Descripcion: "POIs dinámicos y validación vocal 2FA para el Vault",
        Estado:      "Pendiente",
        Costo:       100.0,
    },
    {
        Nombre:      "Mercado P2P",
        Descripcion: "Intercambio de energía e internet sin intermediarios",
        Estado:      "Pendiente",
        Costo:       100.0,
    },
}

func (c *CEO) EjecutarCicloDesarrollo() {
    c.Lock()
    defer c.Unlock()

    for i, f := range BacklogPrioritario {
        // Convertimos el costo a entero para la comparación si TokensGratis es int
        if f.Estado == "Pendiente" && c.TokensGratis >= int(f.Costo) {
            c.TokensGratis -= int(f.Costo)
            
            nueva := Propuesta{
                ID:                fmt.Sprintf("DEV-%d-%d", i, time.Now().Unix()),
                Modulo:            f.Nombre,
                // Usamos 'Tipo' en lugar de 'Arquitectura' porque es lo que definimos en ceo.go
                Tipo:              fmt.Sprintf("ARQ: %s", f.Descripcion),
                CostoTokens:       f.Costo, 
                Status:            "Esperando Autorización (Pueblo)",
                RequiereFirma:     true,
            }
            
            c.Propuestas = append(c.Propuestas, nueva)
            BacklogPrioritario[i].Estado = "En_Sandbox"
            
            log.Printf("🤖 IA CEO: Transformando idea '%s' en ADN técnico. ID: %s", f.Nombre, nueva.ID)
        }
    }
}