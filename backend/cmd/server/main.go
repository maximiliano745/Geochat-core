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
	// 1. Cargar el .env (Las llaves del reino)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Usando variables de entorno del sistema")
	}

	// 2. Inicializar IA 5 y sus herramientas de ejecución
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

	// --- 🚀 LATIDO AUTÓNOMO (El motor que hace crecer a GeoChat) ---
	// Este hilo corre en paralelo y hace que la IA analice el backlog sola.
	go func() {
		log.Println("🧠 IA 5: Ciclo de pensamiento estratégico activado.")
		for {
			ceo.EjecutarCicloDesarrollo()
			// Espera 1 hora entre análisis para no saturar, pero puedes bajarlo a 5m para pruebas.
			time.Sleep(1 * time.Hour) 
		}
	}()

	// 3. Configurar Servidor (Gin)
	r := gin.Default()

	// Endpoints de Soberanía Digital
	r.GET("/webhook/whatsapp", api.VerificarWebhook)
	r.POST("/webhook/whatsapp", api.RecibirRespuestaWhatsApp)
	r.POST("/usuario/:id/accion", api.ProcesarAccionUsuario)

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. IA 5 lista.")
	
	// Saludo inicial automático al celular del Jefe [cite: 2026-02-10]
	_ = ceo.EnviarMensajeSoberano("🚀 ¡Jefe! Sistema operativo y ciclo autónomo iniciado.\n\nComandos:\n- *STATUS*: Ver finanzas.\n- *OK*: Aprobar código.")

	r.Run(":8080")
}

// --- HANDLERS (Los nervios del sistema) ---

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
				log.Println("📊 Solicitud de Reporte Financiero...")
				// Nota: Asegúrate de tener ObtenerReporteFinanciero en tu ceo.go
				reporte := a.CEO.ObtenerReporteFinanciero()
				a.CEO.EnviarMensajeSoberano(reporte)

			case "OK", "ACEPTAR":
				log.Println("📱 Firma recibida. Iniciando evolución de código...")
				modulo := a.CEO.ObtenerUltimoModuloPendiente()
				
				// La IA 5 ejecuta la voluntad del Líder [cite: 2026-02-10]
				err := a.Git.PublicarEvolucion(modulo)
				if err == nil {
					a.CEO.EnviarMensajeSoberano("✅ ¡Misión cumplida Jefe! Código en GitHub y registrado en el Vault.")
					a.CEO.DocumentarLogro(modulo, "Aprobado vía WhatsApp")
				} else {
					a.CEO.EnviarMensajeSoberano("❌ Error en Git: " + err.Error())
				}
			
			default:
				// Opcional: Que la IA te diga que no entendió pero está atenta
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
