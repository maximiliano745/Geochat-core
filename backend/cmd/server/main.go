package main

import (
	"fmt"
	"geochat/internal/ai/ai_ceo" // Importa el corazón de la IA 5
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 1. ESTRUCTURA DE LA API
type API struct {
	// Sincronizado: Usamos ai_ceo.CEO en lugar de IA5
	CEO *ai_ceo.CEO
	Git *ai_ceo.CerebroEjecucion
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ No se detectó archivo .env local.")
	}

	verificarVariablesCriticas()

	// 2. INICIALIZACIÓN SOBERANA
	// Usamos NewCEO() que ya limpia y prepara la Conciencia y los Stats
	ceoInstancia := ai_ceo.NewCEO()
	
	// Ajustamos campos específicos si es necesario
	ceoInstancia.ID = "GeoChat-CEO-Soberano"

	manos := &ai_ceo.CerebroEjecucion{
		RepoPath: "./",
		RepoURL:  os.Getenv("GITHUB_REPO"),
		Token:    os.Getenv("GITHUB_TOKEN"),
		Username: os.Getenv("GITHUB_USER"),
	}

	apiInstancia := &API{
		CEO: ceoInstancia,
		Git: manos,
	}

	// 3. ENCENDIDO: El CEO empieza a escanear oportunidades para el pueblo
	go ceoInstancia.EscanearOportunidades()

	r := gin.Default()

	// Middleware CORS (Esencial para conectar con tu Vue.ts)
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

	// Endpoints
	r.GET("/webhook/whatsapp", apiInstancia.VerificarWebhook)
	r.POST("/webhook/whatsapp", apiInstancia.RecibirRespuestaWhatsApp)
	r.POST("/usuario/:id/accion", apiInstancia.ProcesarAccionUsuario)
	r.POST("/api/admin/authorize", apiInstancia.AutorizarEvolucion)
	
	log.Println("🌍 GeoChat Core iniciado. CEO (IA 5) en línea bajo mando Soberano.")
	r.Run(":8080")
}

// --- HANDLERS ---

func (api *API) AutorizarEvolucion(c *gin.Context) {
	var input struct {
		Passphrase string `json:"passphrase"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
		return
	}

	// Tu Firma es la Ley: Validación contra el .env
	if input.Passphrase != os.Getenv("ADMIN_PASSPHRASE") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Firma no válida. El 15% permanece bloqueado."})
		return
	}

	fmt.Println("👑 [Soberano]: Firma validada. Ejecutando evolución del ADN...")
	
	// La IA vuelve a escanear tras tu autorización
	go api.CEO.EscanearOportunidades() 

	c.JSON(http.StatusOK, gin.H{"mensaje": "Evolución iniciada, Jefe. El CEO está trabajando."})
}

// Handler para procesar recompensas sociales desde la API
func (api *API) ProcesarAccionUsuario(c *gin.Context) {
	usuarioID := c.Param("id")
	// Le damos 1.0 punto de impacto por defecto en cada acción detectada
	api.CEO.ProcesarRecompensaSocial(usuarioID, 1.0)
	
	c.JSON(200, gin.H{"status": "Acción procesada", "usuario": usuarioID})
}

func verificarVariablesCriticas() {
	keys := []string{"GITHUB_TOKEN", "ADMIN_PASSPHRASE"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA CRÍTICA: %s no definida", k)
		}
	}
}

func (api *API) RecibirRespuestaWhatsApp(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) }
func (api *API) VerificarWebhook(c *gin.Context)         { c.String(200, c.Query("hub.challenge")) }