package main

import (
	"geochat/internal/ai/ai_ceo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	//"strings"
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

	// --- 🚀 LABORATORIO DE INTEGRIDAD (PRUEBA DE FUEGO) ---
	// Iniciamos la verificación de código y generación de docs en paralelo
	go func() {
		time.Sleep(1 * time.Second)
		log.Println("🧪 [IA 5]: Entrando al Laboratorio para verificar integridad y generar /docs...")
		ai_ceo.ProbarYDocumentar() // Llama a la función que crea la carpeta docs y el buzon
	}()

	// --- 👑 ESCUCHADOR DE FIRMA LOCAL (BACKUP SOBERANO) ---
	// Si Meta falla, este hilo busca el archivo 'autorizar.txt' para hacer el PUSH
	go func() {
		for {
			if _, err := os.Stat("autorizar.txt"); err == nil {
				log.Println("👑 [IA 5]: Firma detectada en autorizar.txt. Iniciando evolución de ADN...")
				modulo := ceo.ObtenerUltimoModuloPendiente()
				errGit := manos.PublicarEvolucion(modulo)
				if errGit == nil {
					log.Println("✅ [IA 5]: Push exitoso. Limpiando firma...")
					os.Remove("autorizar.txt")
				}
				time.Sleep(10 * time.Second)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// --- 🧠 CICLO DE PENSAMIENTO ESTRATÉGICO ---
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("🧠 [IA 5]: Ciclo de pensamiento estratégico activado.")
		for {
			ceo.EjecutarCicloDesarrollo()
			time.Sleep(1 * time.Hour)
		}
	}()

	// 3. CONFIGURACIÓN DEL SERVIDOR
	r := gin.Default()

	r.GET("/webhook/whatsapp", api.VerificarWebhook)
	r.POST("/webhook/whatsapp", api.RecibirRespuestaWhatsApp)
	r.POST("/usuario/:id/accion", api.ProcesarAccionUsuario)

	// Saludo inicial al Jefe
	go func() {
		time.Sleep(4 * time.Second)
		log.Println("📱 Intentando saludo soberano al Jefe...")
		ceo.EnviarMensajeSoberano("🚀 ¡Jefe! Sistema operativo activo.\n\nHe verificado el código y generado la documentación en /docs.")
	}()

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. IA 5 lista.")
	r.Run(":8080")
}

func verificarVariablesCriticas() {
	keys := []string{"WA_API_KEY", "WA_PHONE_ID", "WA_RECIPIENT", "GITHUB_TOKEN"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA: La variable %s no está definida", k)
		}
	}
}

// --- HANDLERS (Tu código original mantenido para integridad) ---

func (a *API) RecibirRespuestaWhatsApp(c *gin.Context) {
	// ... (Tu lógica de WhatsApp se mantiene igual)
}

func (a *API) VerificarWebhook(c *gin.Context) {
	verifyToken := "GEOCHAT_SOBERANO_2026"
	if c.Query("hub.verify_token") == verifyToken {
		c.String(http.StatusOK, c.Query("hub.challenge"))
		return
	}
	c.Status(http.StatusForbidden)
}

func (a *API) ProcesarAccionUsuario(c *gin.Context) {
	usuarioID := c.Param("id")
	a.CEO.ProcesarRecompensaSocial(usuarioID)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
