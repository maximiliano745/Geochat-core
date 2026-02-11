package main

import (
    "geochat/internal/ai/ai_ceo"
    "geochat/internal/ai/ai_friend"
    "geochat/internal/ai/ai_infra"
    "geochat/internal/ai/ai_security"
    "geochat/internal/api"
    //"geochat/internal/finance"

    "github.com/gin-gonic/gin"
)

func main() {
    // 1. Inicialización de Seguridad Soberana
    ai_security.InitializeVault()

    // 2. Creación de la instancia ÚNICA de la Socia
    ceo := ai_ceo.NewCEO()

    // 3. Lanzamiento de todos los hilos del Organismo Vivo
    go ceo.EjecutarCicloDesarrollo()
    go ai_infra.StartMeshOptimizing()
    go ai_friend.ListenToUsers()

    // 4. Configuración del servidor Gin (API)
    r := gin.Default()
    
    // Conecta los Endpoints de Laboratorio e Inyección
    api.SetupCEOEndpoints(r, ceo) 

    // Endpoint de Propuestas para el Dashboard
    r.GET("/api/proposals", func(c *gin.Context) {
        p := ceo.DisenarModulo("Geocerca Sector-7")
        c.JSON(200, gin.H{
            "propuestas":   []ai_ceo.Propuesta{p},
            "tokensGratis": ceo.TokensGratis,
            "fondoGas":     ceo.FondoGas,
        })
    })

    // Endpoint de Autorización: "Mi Firma es la Ley"
    r.POST("/api/autorizar", func(c *gin.Context) {
        var body struct { Firma string; ID string }
        if err := c.BindJSON(&body); err == nil && body.Firma == "MI_FIRMA_ES_LA_LEY" {
            for _, p := range ceo.Propuestas {
                if p.ID == body.ID {
                    _ = ceo.InyectarCodigoFuncional(p)
                    c.JSON(200, gin.H{"status": "Inyectado exitosamente"})
                    return
                }
            }
        }
        c.JSON(403, gin.H{"error": "Firma inválida o propuesta no encontrada"})
    })

    // 5. El Organismo empieza a respirar en el puerto 3000
    r.Run(":3000")
}