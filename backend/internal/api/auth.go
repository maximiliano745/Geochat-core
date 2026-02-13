package api

import (
	"net/http"
	"os"
	"geochat/internal/ai/ai_ceo" // Ahora SÍ se usará abajo
	"github.com/gin-gonic/gin"
)

// 1. DEFINIMOS LA ESTRUCTURA AQUÍ (Para que (a *API) funcione)
type API struct {
	IA5 *ai_ceo.CEOEngine 
}

// 2. EL MÉTODO DE AUTORIZACIÓN
func (a *API) AutorizarEvolucion(c *gin.Context) {
	var req AuthRequest // Se asume que AuthRequest ya está definido en otro archivo del paquete
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de solicitud inválido"})
		return
	}

	masterKeySecret := os.Getenv("ADMIN_PASSPHRASE")
	
	if req.MasterKey != masterKeySecret {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Firma Soberana Inválida - Acceso Denegado",
		})
		return
	}

	// AQUÍ USAMOS ai_ceo (Esto quita el error de 'imported and not used')
	if a.IA5 != nil {
		go a.IA5.EjecutarEvolucionADN() 
		
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Firma Aceptada. Evolución en curso...",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Motor IA no iniciado"})
	}
}