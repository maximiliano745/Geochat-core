package ai_ceo

import (
	"fmt"
)

// Constante sagrada del sistema: El 15% es para el Pueblo y el crecimiento.
const PorcentajePueblo = 0.15

// AnalizarCoherencia actúa como el Pepito Grillo de la IA 5 [cite: 2026-02-10]
// Valida que cada paso del CEO respete la descentralización y tu mando.
func (c *CEO) AnalizarCoherencia(p Propuesta) (bool, string) {
	// Bloqueo de seguridad para lectura: GeoChat es un organismo vivo que no admite datos corruptos.
	c.RLock()
	defer c.RUnlock()

	// 1. REGLA DE ORO: Respetar el 15% del Pueblo
	// Comparamos el monto de la propuesta contra el 15% del FondoGas disponible.
	margenPermitido := c.FondoGas * PorcentajePueblo
	if p.Monto > margenPermitido {
		return false, fmt.Sprintf("⚠️ Alerta: El monto solicitado (%.2f) excede el margen del 15%% del fondo (Máximo: %.2f).", p.Monto, margenPermitido)
	}

	// 2. FILTRO ARQUITECTÓNICO: GeoChat distribuye riesgo [cite: 2026-01-12]
	// Si el CEO propone algo que huela a control central, se bloquea.
	if p.Modulo == "Privatización" || p.Modulo == "Centralización" {
		return false, "⚠️ Bloqueo de Seguridad: Esta propuesta concentra el riesgo. GeoChat nació para ser del pueblo y distribuido."
	}

	// 3. VÍNCULO SOBERANO: Alineación Ética del Líder
	// Ahora que 'Lider' existe en types.go, accedemos a tu perfil sin errores.
	if c.Lider.AlineacionEtica < 0.5 {
		return true, "💡 Propuesta aceptada, pero tu alineación ética es baja (< 0.5). Sugiero activar el 'Modo Tesla' para equilibrar el impacto social."
	}

	return true, "✅ Coherencia confirmada. La propuesta está alineada con la Visión Global de GeoChat."
}