package main

import (
	"geochat/internal/ai/ai_ceo"
	"geochat/internal/api" // Importante para usar la estructura API
	"log"
	//"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. CARGA DE CONFIGURACIÓN
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ Usando variables de entorno del sistema.")
	}

	// 2. INICIAR EL MOTOR DE LA IA 5 (CEO Autónomo)
	// Creamos la instancia primero para poder pasársela a la API
	miIA5 := ai_ceo.NuevoCEO()
	
	// Lanzamos su latido constante en segundo plano
	go miIA5.IniciarCicloAutonomo()

	// 3. CONFIGURACIÓN DEL SERVIDOR DE CONTROL (GIN)
	r := gin.Default()

	// 4. INICIALIZAR LA API CON LA IA CONECTADA
	servidorAPI := &api.API{
		IA5: miIA5,
	}

	// Endpoints de comunicación y mando
	r.GET("/webhook/whatsapp", verificarWebhook)
	r.POST("/webhook/whatsapp", recibirRespuesta)

	// --- 👑 NUEVA LLAVE MAESTRA (Vía API) ---
	// Usamos el método que ya definimos en internal/api/auth.go
	r.POST("/ia/autorizar", servidorAPI.AutorizarEvolucion)

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. Motor financiero y IA 5 en línea.")
	r.Run(":8080")
}

// Handlers básicos para el inicio del sistema
func verificarWebhook(c *gin.Context) {
	c.String(200, c.Query("hub.challenge"))
}

func recibirRespuesta(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}