package api

import (
	"geochat/internal/ai/ai_ceo"
	"geochat/internal/ai/ai_friend" // Importamos al Friend
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SetupCEOEndpoints conecta el Dashboard y la IA Friend con el Core [cite: 2026-02-10]
func SetupCEOEndpoints(r *gin.Engine, ceo *ai_ceo.CEO, friend *ai_friend.FriendIA) {
	
	// 1. ENDPOINT DE CHAT: Donde el AI Friend evalúa al usuario
	r.POST("/api/chat", func(c *gin.Context) {
		var req struct {
			UserID string `json:"user_id"`
			Texto  string `json:"texto"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Mensaje inválido"})
			return
		}

		// La IA Friend analiza la energía y notifica al CEO si hay mérito [cite: 2026-02-10]
		puntuacion := friend.EvaluarInteraccion(req.UserID, req.Texto)

		c.JSON(200, gin.H{
			"respuesta": "Mensaje procesado por la Red Soberana",
			"energia":   puntuacion,
			"status":    "Analizado por AI Friend",
		})
	})

	// 2. Ver propuestas en el Laboratorio (Sandbox)
	r.GET("/ceo/laboratorio", func(c *gin.Context) {
		files, _ := ioutil.ReadDir("./internal/lab")
		var propuestas []map[string]string
		
		for _, f := range files {
			contenido, _ := os.ReadFile(filepath.Join("./internal/lab", f.Name()))
			propuestas = append(propuestas, map[string]string{
				"nombre": f.Name(),
				"codigo": string(contenido),
			})
		}
		c.JSON(200, propuestas)
	})

	// 3. Autorizar: Mueve el código del Lab al Core (Producción) [cite: 2026-02-10]
	r.POST("/ceo/autorizar", func(c *gin.Context) {
		var req struct {
			NombreArchivo string `json:"nombre_archivo"`
			Firma         string `json:"firma"`
		}
		
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Solicitud inválida"})
			return
		}

		// REGLA DE ORO: Mi Firma es la Ley. [cite: 2026-02-10]
		if req.Firma != "MI_FIRMA_ES_LA_LEY" {
			c.JSON(403, gin.H{"error": "Firma soberana no reconocida"})
			return
		}

		pathOrigen := filepath.Join("./internal/lab", req.NombreArchivo)
		pathDestino := filepath.Join("./internal/core", req.NombreArchivo)

		// El acto de "Vivir": El ADN pasa al organismo real
		err := os.Rename(pathOrigen, pathDestino)
		if err != nil {
			c.JSON(500, gin.H{"error": "Fallo en el despliegue físico del ADN"})
			return
		}

		ceo.DocumentarEnVault("INYECCIÓN_CORE: " + req.NombreArchivo)

		c.JSON(200, gin.H{
			"status": "ADN de GeoChat actualizado. El módulo está en producción.",
			"file":   req.NombreArchivo,
		})
	})
}