package main

import (
	"log"
	"os"

	"geochat/internal/ai/ai_ceo"
	"geochat/internal/ai/ai_friend" // <--- IMPORT DESCOMENTADO Y LISTO
	"geochat/internal/api"
	"geochat/internal/database"
	"geochat/internal/vault"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. CARGA DE ENTORNO SOBERANO
	// Cargamos el .env para que MASTER_KEY y DATABASE_URL estén disponibles
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ Usando variables de entorno del sistema (no se encontró .env).")
	}

	verificarVariablesCriticas()
	
	// 2. INFRAESTRUCTURA DE DATOS Y ADN
	database.ConectarDB() // Conexión a Postgres
	vault.InitEsquemas()  // Inicializa moldes de ADN en el Vault

	// 3. DESPERTAR ENTIDADES (Instancias Reales)
	// Creamos al CEO y a la Friend IA para pasarlas al puente de la API
	ceoInstancia := ai_ceo.NewCEO()
	ceoInstancia.ID = "GeoChat-CEO-Soberano"
	
	friendInstancia := &ai_friend.FriendIA{ID: "Friend-Soberano"}

	// 4. CONFIGURACIÓN DEL MOTOR DE RUTAS (GIN)
	r := gin.Default()

	// --- 🛡️ MIDDLEWARE CORS TOTAL (Esencial para Codespaces y Frontend Vue.ts) ---
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

	// --- 🔗 VÍNCULO DE CAPAS (Puente API) ---
	// Cargamos las rutas del Dashboard y los endpoints del CEO/Friend
	api.GetDashboardData(r) 
	api.SetupCEOEndpoints(r, ceoInstancia, friendInstancia)

	// 5. INICIO DE OPERACIONES AUTÓNOMAS
	// El CEO empieza a escanear oportunidades en segundo plano
	go ceoInstancia.EscanearOportunidades()
	
	log.Println("🌍 GeoChat Core iniciado en 0.0.0.0:8080. El Pueblo es el dueño.")
	
	// 6. LANZAMIENTO
	// Usamos 0.0.0.0 para asegurar visibilidad en entornos cloud/Codespaces
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}

// verificarVariablesCriticas asegura que el sistema no arranque "ciego"
func verificarVariablesCriticas() {
	keys := []string{"MASTER_KEY", "DATABASE_URL"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA: Variable %s no detectada en el entorno.", k)
		}
	}
}