//package server
package main

import (
	"log"
	"time"
	"geochat/internal/ai/ai_ceo" // Ajustado a tu estructura real
)

// IniciarCicloAutonomia lanza el latido constante de la IA CEO
func IniciarCicloAutonomia(ceo *ai_ceo.CEO) {
	// Definimos el pulso: cada 1 hora la IA revisa el negocio
	ticker := time.NewTicker(1 * time.Hour)
	
	go func() {
		log.Println("🧠 IA CEO: Realizando escaneo inicial del VisionMap...")
		// Primer análisis al arrancar el servidor
		ceo.SincronizarFinanzas() // Mira la billetera real de Polygon [cite: 2026-02-10]
		ceo.EscanearOportunidades()

		for t := range ticker.C {
			log.Printf("💓 Latido de Autonomía (%s): Analizando rentabilidad y red...", t.Format("15:04"))
			
			// La IA analiza cuánta energía e internet se venden [cite: 2026-02-10]
			ceo.SincronizarFinanzas()
			ceo.EscanearOportunidades()
			
			log.Println("📢 IA CEO: Ciclo completado. Fondos monitoreados.")
		}
	}()
}

func main() {
    // 1. Inicializamos al CEO
    miCEO := ai_ceo.NewCEO()

    // 2. ACTIVAMOS EL LATIDO SOBERANO
    IniciarCicloAutonomia(miCEO)

    log.Println("--- 🚀 GeoChat en línea. Ciclo de Autonomía activado (1h). ---")
    
    // El servidor se queda escuchando (aquí iría tu configuración de Gin/HTTP)
    select {}
}