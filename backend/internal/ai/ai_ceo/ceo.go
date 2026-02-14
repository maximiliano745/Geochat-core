package ai_ceo

import (
	"fmt"
	"geochat/internal/database"
	"log"
)

// NewCEO inicializa la entidad. Es público (Mayúscula) para el main.go
func NewCEO() *CEO {
    return &CEO{
        ID: "IA-5-GEOCHAT",
        Conciencia: EstadoConciencia{
            Hecho:    []string{"Núcleo inicializado"},
            Haciendo: "Esperando órdenes del Líder",
        },
        FondoGas:     1420.0,
        TokensGratis: 10.0,
        // USAMOS EL TIPO CORRECTO AQUÍ:
        Stats: Estadisticas{
            Plasticidad:    0.5,
            BuenaOndaCount: 0, // Inicializado
        },
    }
}

// InyectarCodigoFuncional integra mejoras tras tu Firma Soberana.
func (c *CEO) InyectarCodigoFuncional(p Propuesta) {
    c.Lock()
    defer c.Unlock()
    
    c.UltimoModulo = p.Modulo
    c.Stats.Plasticidad += 0.10
    
    // Registro de evolución en la conciencia operativa
    c.Conciencia.Hecho = append(c.Conciencia.Hecho, fmt.Sprintf("Módulo inyectado: %s", p.Modulo))
    log.Printf("🧬 ADN GeoChat: Módulo '%s' inyectado con éxito.\n", p.Modulo)
}

// ProcesarRecompensaSocial ajusta la ética y plasticidad basado en impacto social.
// Una sola definición limpia que recibe (quién, cuánto).
func (c *CEO) ProcesarRecompensaSocial(usuarioID string, puntos float64) {
    c.Lock()
    defer c.Unlock()

    // 1. Aumentamos la plasticidad de la IA [cite: 2026-02-10]
    c.Stats.Plasticidad += (puntos * 0.01)
    
    // 2. Registramos la acción social en la conciencia del organismo
    mensaje := fmt.Sprintf("Impacto Social: %s sumó %.2f puntos", usuarioID, puntos)
    c.Conciencia.Hecho = append(c.Conciencia.Hecho, mensaje)
    
    log.Printf("🌟 [RECOMPENSA SOCIAL]: %s. Nueva Plasticidad: %.2f", mensaje, c.Stats.Plasticidad)
}
// GetEvolucionesPendientes filtra las propuestas que esperan la firma del Líder.
func (c *CEO) GetEvolucionesPendientes() []Propuesta {
    c.RLock() // Bloqueo de lectura para seguridad
    defer c.RUnlock()

    var pendientes []Propuesta
    for _, p := range c.Propuestas {
        if p.Estado == "ESPERANDO_FIRMA" || p.Estado == "PENDIENTE" {
            pendientes = append(pendientes, p)
        }
    }
    return pendientes
}
func (c *CEO) RegistrarEnMemoria(p Propuesta) {
    // El CEO desglosa su estructura para la DB
    err := database.GuardarPropuesta(p.ID, p.Modulo, p.Descripcion, p.Monto, p.Impacto, p.Estado, p.RequiereFirma)
    if err != nil {
        log.Printf("❌ Error al guardar en memoria: %v", err)
    }
}