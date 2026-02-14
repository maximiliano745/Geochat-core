package ai_ceo

import (
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
)

// API ahora vive dentro del paquete ai_ceo, no necesita importar nada de afuera.
type API struct {
    CEO *CEO // Antes decía IA5, ahora CEO como en el resto del ADN
}

func (a *API) AutorizarEvolucion(c *gin.Context) {
    var req AuthRequest 
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de solicitud inválido"})
        return
    }

    // Usamos la variable que viene del .env
    masterKeySecret := os.Getenv("ADMIN_PASSPHRASE")
    
    // Sincronizamos con el nombre del campo en types.go (FirmaDigital)
    if req.FirmaDigital != masterKeySecret {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "error",
            "message": "Firma Soberana Inválida - Acceso Denegado",
        })
        return
    }

    if a.CEO != nil {
        // Ejecuta la lógica en segundo plano
        go a.CEO.EjecutarCicloDesarrollo() 
        
        c.JSON(http.StatusOK, gin.H{
            "status":  "success",
            "message": "Firma Aceptada. Evolución de GeoChat en curso...",
        })
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Motor CEO no iniciado"})
    }
}