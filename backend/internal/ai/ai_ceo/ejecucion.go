package ai_ceo

import (
    "fmt"
    "log"
)

func (c *CEO) EjecutarMovimiento(p Propuesta) {
    if p.FirmaDigital == "" {
        log.Printf("🚨 INTENTO DE BRECHA: Sin firma.")
        return 
    }

    log.Printf("🛡️ Firma validada para: %s", p.Modulo)
    
    // Aquí solo LLAMAMOS a las funciones que viven en otros archivos
    c.InyectarCodigoFuncional(p) // Esta vive en ceo.go
c.DocumentarEnVault(fmt.Sprintf("%s: Ejecución lista", p.Modulo))}