package main

import (
    "geochat/internal/ai/ai_ceo" 
    "github.com/gin-gonic/gin"
)

func ProcesarAccionUsuario(c *gin.Context) {
    // 1. OBTENER EL CEO 
    // Nota: En producción, el CEO debería venir de un Singleton o Inyectado
    ceo := ai_ceo.NewCEO() 

    // 2. OBTENER EL USUARIO
    usuarioID := c.Param("id") 
    if usuarioID == "" {
        usuarioID = "Anonimo_Soberano"
    }

    // 3. EJECUTAR: Pasamos el ID (string) y los Puntos (float64)
    // El CEO se encarga de la lógica interna
    ceo.ProcesarRecompensaSocial(usuarioID, 1.0) 
    
    c.JSON(200, gin.H{
        "status": "Soberanía Social procesada",
        "usuario": usuarioID,
    })
}