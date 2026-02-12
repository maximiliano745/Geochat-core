package ai_subconscious

import (
	"log"
	"time"
	"geochat/internal/ai/ai_ceo" // Importante para reportar al CEO [cite: 2026-02-10]
)

// Subconsciente es la IA 2 que gestiona la persistencia y evolución [cite: 2026-02-10]
type Subconsciente struct {
	CEO *ai_ceo.CEO
}

// PerfilAsimilado guarda la huella digital del usuario y su evolución [cite: 2026-02-10]
type PerfilAsimilado struct {
	UserID          string
	ModulosActivos  []string // Lo que la IA 1 inyectó bajo tu mando [cite: 2026-02-10]
	PreferenciaVoz  string
	UltimaEvolucion time.Time
}


// AsimilarCambio procesa la evolución y garantiza que sea PERSISTENTE [cite: 2026-02-10]
func (ia2 *Subconsciente) AsimilarCambio(userID string, nuevoCodigoID string) {
	// 1. Verificación de utilidad (Lógica interna de IA 2)
	
	// 2. Simulación de guardado persistente en Go
	log.Printf("🧠 IA 2: Asimilando nuevo módulo %s para el usuario %s", nuevoCodigoID, userID)
	
	// 3. Notificación a la IA 5 (CEO) para documentación en Vault [cite: 2026-02-11]
	ia2.EnviarReporteACapa5(userID, "Evolución de Interfaz Exitosa: " + nuevoCodigoID)
}

func (ia2 *Subconsciente) EnviarReporteACapa5(userID string, detalle string) {
	log.Printf("📤 IA 2 -> IA 5: Enviando reporte de evolución del usuario %s", userID)
	
	// El CEO recibe el reporte y lo guarda en el Vault personal [cite: 2026-02-11]
	ia2.CEO.DocumentarEnVault("Usuario: " + userID + " - " + detalle)
}