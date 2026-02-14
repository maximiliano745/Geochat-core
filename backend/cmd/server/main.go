package main

import (
	"fmt"
	"geochat/internal/ai/ai_ceo"
	"geochat/internal/database" // Importamos tu paquete de Postgres
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 1. ESTRUCTURA DE LA API
type API struct {
	CEO *ai_ceo.CEO
	Git *ai_ceo.CerebroEjecucion
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ No se detectó archivo .env local.")
	}

	verificarVariablesCriticas()

	// --- 🟢 CAPA 4: CONEXIÓN A LA MEMORIA RELACIONAL ---
	database.ConectarDB()

	// 2. INICIALIZACIÓN SOBERANA DEL CEO
	ceoInstancia := ai_ceo.NewCEO()
	ceoInstancia.ID = "GeoChat-CEO-Soberano"

	// --- 💎 TEST DE MEMORIA ETERNA (Grabación de Prueba) ---
	testPropuesta := ai_ceo.Propuesta{
		ID:            "TEST-001",
		Modulo:        "INFRAESTRUCTURA",
		Descripcion:   "Prueba de persistencia soberana",
		Monto:         100.50,
		Impacto:       "Verificar conexión Capa 4",
		Estado:        "ESPERANDO_FIRMA",
		RequiereFirma: true,
	}

	// El CEO usa el método desglosado para evitar ciclos de importación
	err := database.GuardarPropuesta(
		testPropuesta.ID,
		testPropuesta.Modulo,
		testPropuesta.Descripcion,
		testPropuesta.Monto,
		testPropuesta.Impacto,
		testPropuesta.Estado,
		testPropuesta.RequiereFirma,
	)

	if err != nil {
		log.Printf("❌ Fallo en la memoria: %v", err)
	} else {
		log.Println("💎 ÉXITO: Registro grabado en la Memoria Eterna.")
	}
	// -------------------------------------------------------

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

	// 3. ENCENDIDO: El CEO empieza a escanear oportunidades
	go ceoInstancia.EscanearOportunidades()

	r := gin.Default()

	// Middleware CORS
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

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. CEO en línea bajo mando Soberano.")
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

	if input.Passphrase != os.Getenv("ADMIN_PASSPHRASE") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Firma no válida."})
		return
	}

	fmt.Println("👑 [Soberano]: Firma validada. Ejecutando evolución...")
	go api.CEO.EscanearOportunidades()
	c.JSON(http.StatusOK, gin.H{"mensaje": "Evolución iniciada."})
}

func (api *API) ProcesarAccionUsuario(c *gin.Context) {
	usuarioID := c.Param("id")
	api.CEO.ProcesarRecompensaSocial(usuarioID, 1.0)
	c.JSON(200, gin.H{"status": "Acción procesada", "usuario": usuarioID})
}

func verificarVariablesCriticas() {
	keys := []string{"GITHUB_TOKEN", "ADMIN_PASSPHRASE", "DATABASE_URL"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA CRÍTICA: %s no definida", k)
		}
	}
}

func (api *API) RecibirRespuestaWhatsApp(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) }
func (api *API) VerificarWebhook(c *gin.Context)         { c.String(200, c.Query("hub.challenge")) }