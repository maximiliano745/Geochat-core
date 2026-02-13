package main // IMPORTANTE: Debe ser package main

import (
	"geochat/internal/ai/ai_ceo" // Importamos la lógica que ya tienes
	"log"
	"os"
	//"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. CARGA DE CONFIGURACIÓN
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ Usando variables de entorno del sistema.")
	}

	// 2. INICIAR EL LATIDO DE LA IA 5 (Ciclo Autónomo)
	// Esto lanza el bucle de IniciarCicloAutonomo que tienes en laboratorio_autonomo.go
	go ai_ceo.IniciarCicloAutonomo()

	// 3. CONFIGURACIÓN DEL SERVIDOR DE CONTROL (GIN)
	r := gin.Default()

	// Endpoints de comunicación y mando
	r.GET("/webhook/whatsapp", verificarWebhook)
	r.POST("/webhook/whatsapp", recibirRespuesta)

	// --- 👑 NUEVA LLAVE MAESTRA (LO QUE HABLAMOS) ---
	r.POST("/ia/autorizar", func(c *gin.Context) {
		pass := c.PostForm("password")
		if pass == os.Getenv("ADMIN_PASSPHRASE") {
			os.WriteFile("autorizar.txt", []byte("OK"), 0644)
			c.JSON(200, gin.H{"status": "Firma soberana aplicada. IA 5 ejecutando..."})
		} else {
			c.JSON(403, gin.H{"error": "Contraseña incorrecta"})
		}
	})

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. Motor financiero y IA 5 en línea.")
	r.Run(":8080")
}

// Handlers básicos para que no falle la compilación
func verificarWebhook(c *gin.Context) {
	c.String(200, c.Query("hub.challenge"))
}

func recibirRespuesta(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}