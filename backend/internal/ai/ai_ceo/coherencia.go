package ai_ceo

import (
	"fmt"
)

// Constante sagrada del sistema (Si no está en otro archivo, se queda aquí)
const PorcentajePueblo = 0.15

// AnalizarCoherencia actúa como el Pepito Grillo de la IA 5 [cite: 2026-02-10]
// Ahora usa el tipo 'Propuesta' que definimos en ceo.go
func (c *CEO) AnalizarCoherencia(p Propuesta) (bool, string) {
	
	// 1. Verificar si respeta el 15% 
	// Usamos p.Monto que ya está en nuestra estructura maestra
	if p.Monto > (c.FondoGas * PorcentajePueblo) {
		return false, fmt.Sprintf("⚠️ Alerta: El monto %.2f excede el margen del 15%% del fondo disponible.", p.Monto)
	}

	// 2. Análisis de "Modo Tesla" vs "Concentración de Riesgo"
	// GeoChat distribuye riesgo [cite: 2026-01-12]
	if p.Modulo == "Privatización" {
		return false, "⚠️ Jefe, esto concentra el riesgo. GeoChat nació para distribuirlo. ¿Deseas proceder?"
	}

	// 3. Validación de Alineación Ética basada en el PerfilLider de ceo.go
	if c.Lider.AlineacionEtica < 0.5 {
		return true, "💡 Decisión aceptada, pero sugiero un movimiento social para equilibrar tu alineación ética."
	}

	return true, "✅ Decisión perfectamente alineada con los principios de GeoChat."
}