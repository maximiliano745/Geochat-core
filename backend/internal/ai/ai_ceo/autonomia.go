package ai_ceo

import (
	"log"
	"time"
)

// EscanearOportunidades: El proceso autónomo que patrulla la red.
func (c *CEO) EscanearOportunidades() {
	log.Println("🚀 [IA 5]: Iniciando Escáner de Autonomía (15% para el Pueblo)...")

	for {
		log.Println("🔭 [GeoChat]: Escaneando infraestructura y excedentes...")

		c.Lock()
		// Si detectamos fondos suficientes, simulamos la creación de una propuesta
		if c.FondoGas > 500 {
			log.Printf("💡 [GeoChat]: Oportunidad de inversión detectada. Saldo: %.2f", c.FondoGas)
			
			// Ejemplo de cómo se vería la propuesta interna (Sin errores de campos desconocidos)
			_ = Propuesta{
				ID:            "AUTO-SCAN-001",
				Modulo:        "INFRA_NODO",
				Descripcion:   "Aumento de ancho de banda mayorista",
				Monto:         300.0,
				Estado:        "PENDIENTE",
				RequiereFirma: true,
				CreatedAt:     time.Now(), // <--- Usamos el nombre correcto
			}
		}
		c.Unlock()

		// Pausa de 1 hora para el siguiente ciclo
		time.Sleep(1 * time.Hour)
	}
}