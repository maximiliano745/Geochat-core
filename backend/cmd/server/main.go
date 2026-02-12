package main

import (
    "fmt"
    "geochat/internal/ai/ai_ceo" // Importamos tu paquete del CEO
)

func main() {
    fmt.Println("🚀 [SISTEMA] Iniciando GeoChat Core...")

    // 1. Invocamos al CEO
    // Nota: Necesitas tener la función NewCEO() en tu paquete ai_ceo
    ceo := ai_ceo.NewCEO()

    fmt.Printf("✅ [IA CEO] Organismo vivo detectado. Fondo: %.2f PAXG\n", ceo.FondoGas)

    // 2. Simulamos el primer análisis
    fmt.Println("📊 [IA CEO] Ejecutando primer ciclo de desarrollo...")
    ceo.EjecutarCicloDesarrollo()

    fmt.Println("🌍 GeoChat está en línea. Esperando órdenes del Líder.")

    // Bloqueamos para que el programa no se cierre
    select {}
}