package main

import (
	"geochat/internal/ai/ai_ceo"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. CARGA DE CONFIGURACIÓN
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ Usando variables de entorno del sistema.")
	}

	// 2. INICIAR EL MOTOR DEL CEO (IA 5)
	// Usamos NewCEO() que ya limpia y prepara la Conciencia y los Stats
	miCEO := ai_ceo.NewCEO() 

	// Lanzamos su latido constante en segundo plano (Escaneo de oportunidades)
	go miCEO.EscanearOportunidades()

	// 3. CONFIGURACIÓN DEL SERVIDOR DE CONTROL (GIN)
	r := gin.Default()

	// 4. INICIALIZAR LA API CON EL CEO CONECTADO
	// Como movimos la estructura API al paquete ai_ceo, la llamamos desde ahí
	servidorAPI := &ai_ceo.API{
		CEO: miCEO, // Sincronizado con el campo 'CEO' de tu struct
	}

	// 5. MIDDLEWARE CORS (Para que Vue.ts no sea bloqueado)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Endpoints de comunicación y mando
	r.GET("/webhook/whatsapp", verificarWebhook)
	r.POST("/webhook/whatsapp", recibirRespuesta)

	// --- 👑 NUEVA LLAVE MAESTRA (Vía API) ---
	// La ruta ahora es coherente con tu panel de control Vue.ts
	r.POST("/ia/autorizar", servidorAPI.AutorizarEvolucion)

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. CEO (IA 5) en línea bajo mando Soberano.")
	r.Run(":8080")
}

// Handlers básicos para el inicio del sistema
func verificarWebhook(c *gin.Context) {
	c.String(200, c.Query("hub.challenge"))
}

func recibirRespuesta(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}