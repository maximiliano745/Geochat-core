package main

import (
	"geochat/internal/ai/ai_ceo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// API unifica el cerebro (CEO) y las manos (Git)
type API struct {
	CEO *ai_ceo.CEO
	Git *ai_ceo.CerebroEjecucion
}

func main() {
	// 1. CARGA DE CONFIGURACIÓN (Blindaje del ADN)
	// Intentamos cargar .env. Si no está, avisamos pero no matamos el proceso.
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ No se detectó archivo .env local. Usando variables de entorno del sistema.")
	}

	// Validación de Seguridad: Si faltan llaves críticas, avisamos de inmediato.
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

	// --- 🚀 LATIDO AUTÓNOMO ---
	// Lo movemos después de inicializar la API para evitar punteros nulos.
	go func() {
		// Pequeña pausa para que el servidor Gin levante primero
		time.Sleep(2 * time.Second)
		log.Println("🧠 [IA 5]: Ciclo de pensamiento estratégico activado.")
		for {
			ceo.EjecutarCicloDesarrollo()
			// 1 hora de espera para análisis profundo
			time.Sleep(1 * time.Hour)
		}
	}()

	// 3. CONFIGURACIÓN DEL SERVIDOR (Gin)
	// Quitamos el modo Debug molesto si prefieres ver todo limpio
	// gin.SetMode(gin.ReleaseMode) 
	r := gin.Default()

	// Endpoints de Soberanía Digital
	r.GET("/webhook/whatsapp", api.VerificarWebhook)
	r.POST("/webhook/whatsapp", api.RecibirRespuestaWhatsApp)
	r.POST("/usuario/:id/accion", api.ProcesarAccionUsuario)

	// Saludo inicial al Jefe (Solo si la configuración es válida)
	go func() {
		time.Sleep(3 * time.Second) // Esperamos que el puerto 8080 esté abierto
		log.Println("📱 Intentando saludo soberano al Jefe...")
		err := ceo.EnviarMensajeSoberano("🚀 ¡Jefe! Sistema operativo y ciclo autónomo iniciado.\n\nComandos:\n- *STATUS*: Ver finanzas.\n- *OK*: Aprobar código.")
		if err != nil {
			log.Printf("⚠️ No pude saludarte por WhatsApp: %v", err)
		}
	}()

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. IA 5 lista.")
	r.Run(":8080")
}

// verificarVariablesCriticas asegura que la IA tenga sus "sentidos" completos
func verificarVariablesCriticas() {
	keys := []string{"WA_API_KEY", "WA_PHONE_ID", "WA_RECIPIENT", "POLYGON_RPC_URL"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA: La variable %s no está definida en el .env", k)
		}
	}
}

// --- HANDLERS ---

func (a *API) RecibirRespuestaWhatsApp(c *gin.Context) {
	var incoming struct {
		Entry []struct {
			Changes []struct {
				Value struct {
					Messages []struct {
						Text struct {
							Body string `json:"body"`
						} `json:"text"`
					} `json:"messages"`
				} `json:"value"`
			} `json:"changes"`
		} `json:"entry"`
	}

	if err := c.ShouldBindJSON(&incoming); err != nil {
		c.Status(http.StatusOK)
		return
	}

	if len(incoming.Entry) > 0 && len(incoming.Entry[0].Changes) > 0 {
		value := incoming.Entry[0].Changes[0].Value
		if len(value.Messages) > 0 {
			respuesta := strings.ToUpper(strings.TrimSpace(value.Messages[0].Text.Body))

			switch respuesta {
			case "STATUS":
				log.Println("📊 [IA 5]: Generando Reporte Financiero...")
				reporte := a.CEO.ObtenerReporteFinanciero()
				a.CEO.EnviarMensajeSoberano(reporte)

			case "OK", "ACEPTAR":
				log.Println("📱 [IA 5]: Firma del Jefe recibida. Evolucionando ADN...")
				modulo := a.CEO.ObtenerUltimoModuloPendiente()
				
				err := a.Git.PublicarEvolucion(modulo)
				if err == nil {
					a.CEO.EnviarMensajeSoberano("✅ ¡Misión cumplida Jefe! Código en GitHub y registrado en el Vault.")
					a.CEO.DocumentarLogro(modulo, "Aprobado vía WhatsApp")
				} else {
					a.CEO.EnviarMensajeSoberano("❌ Error en Git: " + err.Error())
				}
			
			default:
				log.Printf("📩 Mensaje recibido no reconocido: %s", respuesta)
			}
		}
	}
	c.Status(http.StatusOK)
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
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Resonancia procesada"})
}