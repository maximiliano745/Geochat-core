package main

import (
	"fmt"
	"geochat/internal/ai/ai_ceo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

// API unifica el cerebro (CEO) y las manos (Git)
type API struct {
	CEO *ai_ceo.CEO
	Git *ai_ceo.CerebroEjecucion
}

func main() {
	// 1. CARGA DE CONFIGURACIÓN
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ No se detectó archivo .env local. Usando variables del sistema.")
	}

	verificarVariablesCriticas()

	// 2. INICIALIZACIÓN DE LA IA 5
	ceo := ai_ceo.NewCEO()
	manos := &ai_ceo.CerebroEjecucion{
		RepoPath: "./",
		RepoURL:  os.Getenv("GITHUB_REPO"),
		Token:    os.Getenv("GITHUB_TOKEN"),
		Username: os.Getenv("GITHUB_USER"),
	}

	api := &API{
		CEO: ceo,
		Git: manos,
	}

	// --- 🧩 1. MOTOR ESTRATÉGICO ---
	log.Println("🔍 [IA 5]: Escaneando Plan Maestro de GeoChat...")
	ai_ceo.ProponerSiguientePaso()
	ai_ceo.IniciarMotor()
	ai_ceo.IniciarMonitoreoPlan()

	// --- 🧪 2. LABORATORIO DE INTEGRIDAD ---
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("🧪 [IA 5]: Verificando salud del código y regenerando /docs...")
		ai_ceo.ProbarYDocumentar()
	}()

	// --- 🧠 4. CICLO DE DESARROLLO (PENSAMIENTO PROFUNDO) ---
	go func() {
		time.Sleep(15 * time.Second)
		for {
			log.Println("🧠 [IA 5]: Ejecutando ciclo de desarrollo profundo...")
			ceo.EjecutarCicloDesarrollo()
			time.Sleep(1 * time.Hour)
		}
	}()

	// 3. CONFIGURACIÓN DEL SERVIDOR (GIN)
	r := gin.Default()

	// --- 🛡️ MIDDLEWARE: CORS (Para que Vue pueda hablar con Go) ---
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Endpoints de comunicación
	r.GET("/webhook/whatsapp", api.VerificarWebhook)
	r.POST("/webhook/whatsapp", api.RecibirRespuestaWhatsApp)
	r.POST("/usuario/:id/accion", api.ProcesarAccionUsuario)

	// --- 🔐 TUNEL DE AUTORIZACIÓN SOBERANA ---
	r.POST("/api/admin/autorizar", api.AutorizarEvolucion)

	// Saludo inicial
	go func() {
		time.Sleep(5 * time.Second)
		hitos := ai_ceo.LeerPlanMaestro()
		proximoPaso := "Diseño inicial"
		if len(hitos) > 0 {
			proximoPaso = hitos[0].Modulo
		}
		log.Println("📱 Saludando al Jefe...")
		ceo.EnviarMensajeSoberano("🚀 ¡GeoChat Activo!\n\n📍 Túnel Digital y CORS habilitados.\n\nPieza actual: " + proximoPaso)
	}()

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. CEO listo.")
	r.Run(":8080")
}

// --- HANDLERS ---

func (api *API) AutorizarEvolucion(c *gin.Context) {
	var input struct {
		Passphrase string `json:"passphrase"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de firma inválido"})
		return
	}

	masterKey := os.Getenv("ADMIN_PASSPHRASE")
	if input.Passphrase == "" || input.Passphrase != masterKey {
		log.Println("⚠️ ALERTA: Intento de firma fallido.")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Firma no válida. Acceso denegado."})
		return
	}

	fmt.Println("👑 [Soberano]: Firma validada. La IA 5 inicia evolución de ADN...")
	
	go func() {
		hitos := ai_ceo.LeerPlanMaestro()
		mensaje := "Evolución via Dashboard"
		if len(hitos) > 0 {
			mensaje = "Construcción: " + hitos[0].Modulo
		}
		api.Git.PublicarEvolucion(mensaje)
	}()

	c.JSON(http.StatusOK, gin.H{"mensaje": "Evolución iniciada, Jefe."})
}

func verificarVariablesCriticas() {
	keys := []string{"WA_API_KEY", "WA_PHONE_ID", "WA_RECIPIENT", "GITHUB_TOKEN", "ADMIN_PASSPHRASE"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA: Variable CRÍTICA %s no definida", k)
		}
	}
}

func (a *API) RecibirRespuestaWhatsApp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "received"})
}

func (a *API) VerificarWebhook(c *gin.Context) {
	if c.Query("hub.verify_token") == "GEOCHAT_SOBERANO_2026" {
		c.String(http.StatusOK, c.Query("hub.challenge"))
		return
	}
	c.Status(http.StatusForbidden)
}

func (a *API) ProcesarAccionUsuario(c *gin.Context) {
	a.CEO.ProcesarRecompensaSocial(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}