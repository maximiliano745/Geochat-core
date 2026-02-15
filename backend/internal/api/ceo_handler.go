package api

import (
	"net/http"
	"os"

	"geochat/internal/ai/ai_ceo"    // El paquete del CEO
	"geochat/internal/ai/ai_friend" // El paquete del Friend
	"github.com/gin-gonic/gin"
)

// SetupCEOEndpoints conecta el Dashboard con la IA 5 y la IA Friend.
// Usamos alias (elCEO, elFriend) para evitar conflictos con los nombres de paquete.
func SetupCEOEndpoints(r *gin.Engine, elCEO *ai_ceo.CEO, elFriend *ai_friend.FriendIA) {

	// 1. CHAT SOBERANO (Conexión con IA Friend)
	r.POST("/api/chat", func(c *gin.Context) {
		var req struct {
			UserID string `json:"user_id"`
			Texto  string `json:"texto"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mensaje inválido"})
			return
		}

		// La IA Friend evalúa el impacto social/energía
		puntuacion := elFriend.EvaluarInteraccion(req.UserID, req.Texto)
		
		c.JSON(http.StatusOK, gin.H{
			"respuesta": "Mensaje procesado por la Red Soberana",
			"puntos":    puntuacion,
		})
	})

	// 2. PROPUESTAS DEL CEO (IA 5)
	// Aquí es donde el CEO inyecta las Burbujas de ADN al Dashboard
	r.GET("/api/ceo/propuestas", func(c *gin.Context) {
		// CORRECCIÓN: Llamamos al método que ya maneja el contexto de Gin
		// Este método vive en internal/ai/ai_ceo/ceo.go
		elCEO.ObtenerPropuestas(c) 
	})

	// 3. AUTORIZACIÓN SOBERANA [cite: 2026-02-10]
	r.POST("/api/ceo/autorizar", func(c *gin.Context) {
		var req struct {
			PropuestaID string `json:"propuesta_id"`
			Firma       string `json:"firma"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan datos de autorización"})
			return
		}

		// REGLA DE ORO: MI FIRMA ES LA LEY
		masterKey := os.Getenv("MASTER_KEY")
		if req.Firma != masterKey {
			c.JSON(http.StatusForbidden, gin.H{"error": "Firma soberana inválida"})
			return
		}

		// Respuesta de éxito tras validar tu firma
		c.JSON(http.StatusOK, gin.H{
			"status":  "AUTORIZADO",
			"mensaje": "Firma validada. ADN listo para inyección en el Core",
		})
	})
}