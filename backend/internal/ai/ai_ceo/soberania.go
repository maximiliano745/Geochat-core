package ai_ceo

import (
	"time"
)

// Constantes de Filosofía GeoChat [cite: 2026-02-10]
const (
	// Nota: Si PorcentajePueblo ya está en coherencia.go, bórralo de aquí.
	SoberaniaKeys    = true      // El usuario siempre tiene sus llaves privadas [cite: 2026-01-12]
	Transparencia    = "E2E_VAULT" // Todo registro es cifrado punto a punto
)

// ValidarSoberania asegura que la IA no ejecute nada si se viola la ética de GeoChat
func (c *CEO) ValidarSoberania() bool {
	// Regla de Oro: Si la soberanía de las llaves es false, el sistema se detiene.
	if !SoberaniaKeys {
		return false
	}
	return true
}

// RegistrarFirma agrega la firma al historial del Líder después de una autorización exitosa.
// Este método ahora actúa sobre la estructura que ya vive en ceo.go
func (p *PerfilLider) RegistrarFirma(idTransaccion string) {
	p.HistorialFirma = append(p.HistorialFirma, idTransaccion)
	p.UltimaActividad = time.Now()
}