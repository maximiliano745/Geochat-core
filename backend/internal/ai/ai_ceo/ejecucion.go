package ai_ceo

import (
	"log"
	"geochat/internal/finance" // Capa 4: Ejecución de billetera
)

// EjecutarMovimiento es el último paso del flujo soberano.
// Ningún PAXG se mueve y ningún código muta sin pasar por aquí.
func (c *CEO) EjecutarMovimiento(p Propuesta) {
	// 1. Verificación Crítica: ¿El líder puso su firma digital?
	// La firma debe coincidir con la 'KEY_SOBERANA' almacenada en el Vault.
	if !p.TieneFirmaDelLider() {
		log.Printf("🚨 INTENTO DE BRECHA: La IA intentó ejecutar '%s' sin firma.", p.ID)
		log.Fatal("❌ ACCESO DENEGADO: Intento de ejecución sin la Palabra del Líder. El sistema se detiene por seguridad.")
		return 
	}

	// 2. Si la firma es válida, se procede con la ejecución económica o técnica.
	log.Printf("🛡️ Firma validada. Ejecutando mandato del Líder: %s", p.Modulo)
	
	// Si la propuesta es financiera, libera los fondos desde la Capa 4
	if p.CostoTokens > 0 {
		exito := finance.EjecutarTransferencia(p.Destino, p.CostoTokens)
		if !exito {
			log.Println("⚠️ Error en la capa financiera. Transacción abortada.")
			return
		}
	}

	// 3. Inyectar el cambio en el ADN del sistema (Capa 1/2)
	c.InyectarCodigoFuncional(p)

	// 4. Documentar el éxito en el Libro de Logros (IA 5)
	c.DocumentarEnVault(p.Modulo + " ejecutado y validado.")
}