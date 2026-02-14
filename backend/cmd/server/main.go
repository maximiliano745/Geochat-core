package main

import (
	"geochat/internal/ai/ai_ceo"
	"geochat/internal/ai/ai_friend"
	"geochat/internal/api"
	"geochat/internal/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. CARGA DE ENTORNO SOBERANO
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ Usando variables de entorno del sistema.")
	}

	verificarVariablesCriticas()
	database.ConectarDB()

	// 2. DESPERTAR ENTIDADES (CEO con 15% y Friend IA)
	ceoInstancia := ai_ceo.NewCEO()
	ceoInstancia.ID = "GeoChat-CEO-Soberano"
	friendIA := &ai_friend.FriendIA{}

	// 3. CONFIGURACIÓN DEL MOTOR DE RUTAS
	r := gin.Default()

	// --- 🛡️ MIDDLEWARE CORS TOTAL (Túnel libre para Codespaces) ---
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// --- 🔗 VÍNCULO DE CAPAS (Sin duplicados, sin conflictos) ---
	// api.GetDashboardData registra: GET /api/dashboard/stats
	// api.SetupCEOEndpoints registra: POST /api/ceo/autorizar, /api/chat, etc.
	api.GetDashboardData(r) 
	api.SetupCEOEndpoints(r, ceoInstancia, friendIA)

	// 4. INICIO DE OPERACIONES AUTÓNOMAS
	// El CEO escanea, pero tu firma en la Interfaz de Autorización es la Ley.
	go ceoInstancia.EscanearOportunidades()
	
	log.Println("🌍 GeoChat Core iniciado en 0.0.0.0:8080. El Pueblo es el dueño.")
	
	// Bind total para Codespaces
	r.Run("0.0.0.0:8080")
}

func verificarVariablesCriticas() {
	keys := []string{"MASTER_KEY", "DATABASE_URL"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA: Variable %s no detectada.", k)
		}
	}
}