package main

import (
	"fmt"
	"geochat/internal/ai/ai_ceo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// 1. ESTRUCTURA DE LA API
type API struct {
	// CAMBIO: Ahora usamos el nombre exacto que definiste: CEOEngine
	CEO *ai_ceo.CEOEngine 
	Git *ai_ceo.CerebroEjecucion
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ No se detectó archivo .env local.")
	}

	verificarVariablesCriticas()

	// 2. INICIALIZACIÓN
	// NuevoCEO() devuelve un *CEOEngine, ahora coincide con la estructura API
	ceoInstancia := ai_ceo.NuevoCEO() 
	
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

	// 3. ENCENDIDO
	go ceoInstancia.IniciarCicloAutonomo()

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

	r.GET("/webhook/whatsapp", apiInstancia.VerificarWebhook)
	r.POST("/webhook/whatsapp", apiInstancia.RecibirRespuestaWhatsApp)
	r.POST("/usuario/:id/accion", apiInstancia.ProcesarAccionUsuario)
	r.POST("/api/admin/authorize", apiInstancia.AutorizarEvolucion)
	
	log.Println("🌍 GeoChat Core iniciado en puerto 8080. CEOEngine en línea.")
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

	fmt.Println("👑 [Soberano]: Firma validada.")
	// Ahora Go encuentra este método porque el tipo CEOEngine lo tiene definido
	go api.CEO.EjecutarEvolucionADN()

	c.JSON(http.StatusOK, gin.H{"mensaje": "Evolución iniciada, Jefe."})
}

func verificarVariablesCriticas() {
	keys := []string{"GITHUB_TOKEN", "ADMIN_PASSPHRASE"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA: %s no definida", k)
		}
	}
}

func (api *API) RecibirRespuestaWhatsApp(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) }
func (api *API) VerificarWebhook(c *gin.Context)         { c.String(200, c.Query("hub.challenge")) }
func (api *API) ProcesarAccionUsuario(c *gin.Context)    { c.JSON(200, gin.H{"status": "ok"}) }