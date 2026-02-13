package api

import (
	"geochat/internal/ai/ai_ceo"
	"geochat/internal/ai/ai_friend"
	"geochat/internal/database"
	"geochat/internal/models"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// SetupCEOEndpoints conecta el Dashboard, la IA Friend y Postgres con el Core
func SetupCEOEndpoints(r *gin.Engine, ceo *ai_ceo.CEO, friend *ai_friend.FriendIA) {

	// 1. CHAT SOBERANO
	r.POST("/api/chat", func(c *gin.Context) {
		var req struct {
			UserID string `json:"user_id"`
			Texto  string `json:"texto"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mensaje inválido"})
			return
		}
		puntuacion := friend.EvaluarInteraccion(req.UserID, req.Texto)
		c.JSON(http.StatusOK, gin.H{
			"respuesta": "Mensaje procesado por la Red Soberana",
			"energia":   puntuacion,
		})
	})

	// 2. PROPUESTAS DEL CEO (IA 5) - CORREGIDO PARA POSTGRES
	r.GET("/api/ceo/propuestas", func(c *gin.Context) {
		// Validamos que la conexión a Postgres esté viva
		if database.DB == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Memoria (Postgres) no inicializada"})
			return
		}

		var propuestas []models.Task

		// Consulta SQL para traer las propuestas del 15%
		rows, err := database.DB.Query("SELECT id, titulo, descripcion, estado FROM tasks WHERE estado = 'propuesta_ceo'")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la memoria del CEO"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var t models.Task
			if err := rows.Scan(&t.ID, &t.Titulo, &t.Descripcion, &t.Estado); err != nil {
				continue
			}
			propuestas = append(propuestas, t)
		}

		c.JSON(http.StatusOK, propuestas)
	})

	// 3. LABORATORIO (Archivos físicos)
	r.GET("/api/ceo/laboratorio", func(c *gin.Context) {
		labPath := "./internal/lab"
		files, err := ioutil.ReadDir(labPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo acceder al Laboratorio"})
			return
		}

		var archivosPropuestos []map[string]string
		for _, f := range files {
			if !f.IsDir() {
				contenido, _ := os.ReadFile(filepath.Join(labPath, f.Name()))
				archivosPropuestos = append(archivosPropuestos, map[string]string{
					"nombre": f.Name(),
					"codigo": string(contenido),
				})
			}
		}
		c.JSON(http.StatusOK, archivosPropuestos)
	})

	// 4. AUTORIZACIÓN SOBERANA [cite: 2026-02-10]
	r.POST("/api/ceo/autorizar", func(c *gin.Context) {
		var req struct {
			NombreArchivo string `json:"nombre_archivo"`
			Firma         string `json:"firma"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
			return
		}

		// MI FIRMA ES LA LEY
		masterKey := os.Getenv("MASTER_KEY")
		if req.Firma != masterKey {
			c.JSON(http.StatusForbidden, gin.H{"error": "Firma soberana rechazada"})
			return
		}

		pathOrigen := filepath.Join("./internal/lab", req.NombreArchivo)
		pathDestino := filepath.Join("./internal/core", req.NombreArchivo)

		if err := os.Rename(pathOrigen, pathDestino); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo en el despliegue físico del ADN"})
			return
		}

		ceo.DocumentarEnVault("EVOLUCIÓN_CORE: " + req.NombreArchivo)
		c.JSON(http.StatusOK, gin.H{"status": "ADN actualizado con éxito."})
	})
}