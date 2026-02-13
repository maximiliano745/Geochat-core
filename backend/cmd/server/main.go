package main

import (
	
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

	// --- 🧩 1. MOTOR ESTRATÉGICO Y OÍDO (EL PLANIFICADOR) ---
	// La IA escanea el plan, activa el latido y se pone a escuchar cambios.
	log.Println("🔍 [IA 5]: Escaneando Plan Maestro de GeoChat...")
	ai_ceo.ProponerSiguientePaso()
	
	ai_ceo.IniciarMotor()         // Latido de 30s para pensar
	ai_ceo.IniciarMonitoreoPlan() // Oído en tiempo real para el archivo .md

	// --- 🧪 2. LABORATORIO DE INTEGRIDAD ---
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("🧪 [IA 5]: Verificando salud del código y regenerando /docs...")
		ai_ceo.ProbarYDocumentar()
	}()

	// --- 👑 3. ESCUCHADOR DE FIRMA SOBERANA (AUTORIZAR.TXT) ---
	go func() {
		for {
			if _, err := os.Stat("autorizar.txt"); err == nil {
				log.Println("👑 [IA 5]: Firma detectada. Iniciando evolución de ADN...")
				
				hitos := ai_ceo.LeerPlanMaestro()
				mensajeCommit := "Evolución Autónoma de GeoChat"
				if len(hitos) > 0 {
					mensajeCommit = "Construcción: " + hitos[0].Modulo
				}

				// La IA usa las 'manos' (Git) para subir el avance
				errGit := manos.PublicarEvolucion(mensajeCommit)
				if errGit == nil {
					log.Println("✅ [IA 5]: ADN sincronizado con éxito. Limpiando firma...")
					os.Remove("autorizar.txt")
				}
				time.Sleep(10 * time.Second)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// --- 🧠 4. CICLO DE DESARROLLO (PENSAMIENTO PROFUNDO) ---
	go func() {
		time.Sleep(15 * time.Second)
		for {
			log.Println("🧠 [IA 5]: Ejecutando ciclo de desarrollo profundo...")
			ceo.EjecutarCicloDesarrollo()
			time.Sleep(1 * time.Hour) // Una vez por hora para no saturar
		}
	}()

	// 3. CONFIGURACIÓN DEL SERVIDOR (GIN)
	r := gin.Default()

	r.GET("/webhook/whatsapp", api.VerificarWebhook)
	r.POST("/webhook/whatsapp", api.RecibirRespuestaWhatsApp)
	r.POST("/usuario/:id/accion", api.ProcesarAccionUsuario)

	// Saludo inicial informando la próxima pieza del rompecabezas
	go func() {
		time.Sleep(5 * time.Second)
		hitos := ai_ceo.LeerPlanMaestro()
		proximoPaso := "Diseño inicial"
		if len(hitos) > 0 {
			proximoPaso = hitos[0].Modulo
		}
		
		log.Println("📱 Saludando al Jefe...")
		ceo.EnviarMensajeSoberano("🚀 ¡GeoChat Activo!\n\n📍 Pieza actual: " + proximoPaso + "\n\nSistema de monitoreo y motor financiero en línea.")
	}()

	log.Println("🌍 GeoChat Core iniciado en puerto 8080. CEO listo.")
	r.Run(":8080")
}

func verificarVariablesCriticas() {
	keys := []string{"WA_API_KEY", "WA_PHONE_ID", "WA_RECIPIENT", "GITHUB_TOKEN"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Printf("⚠️ ADVERTENCIA: La variable %s no está definida en el entorno", k)
		}
	}
}

// --- HANDLERS ---

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