package main

import (
    "geochat/internal/ai/ai_ceo" // Importa el paquete del CEO
    "github.com/gin-gonic/gin"    // Si usas Gin para la API
)

func ProcesarAccionUsuario(c *gin.Context) {
    // 1. DEFINIR EL OBJETO CEO (Si no viene de otro lado)
    // Normalmente el CEO se inicializa en el main y se pasa aquí
    ceo := ai_ceo.NewCEO() 

    // 2. OBTENER EL USUARIO (Aquí defines usuarioID)
    usuarioID := c.Param("id") // O lo sacas del body/token
    
    if usuarioID == "" {
        usuarioID = "Anonimo_Soberano"
    }

    // 3. AHORA SÍ: La IA procesa la recompensa
    // Esto conectará con tu lógica de IA 1 (Resonancia)
    ceo.ProcesarRecompensaSocial(usuarioID) 
    
    c.JSON(200, gin.H{"status": "Recompensa procesada por la IA"})
}