package api // O 'package api' según dónde esté el archivo, pero el main.go usa 'main'

import (
	"geochat/internal/ai/ai_ceo"
	"geochat/internal/layer1_core"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 1. Estructura API limpia: Solo el CEO y el Git (o Writer)
type API struct {
	CEO *ai_ceo.CEO
	// Ya no existe WhatsAppService aquí [cite: 2026-02-10]
}

// 2. Estructura de Meta Cloud API (Corregida para el formato real)
type WhatsAppPayload struct {
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

func (a *API) RecibirRespuestaWhatsApp(c *gin.Context) {
	var incoming WhatsAppPayload

	if err := c.ShouldBindJSON(&incoming); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"}) // Meta quiere 200 siempre
		return
	}

	// Navegamos el JSON de Meta para llegar al texto
	if len(incoming.Entry) > 0 && len(incoming.Entry[0].Changes) > 0 {
		messages := incoming.Entry[0].Changes[0].Value.Messages
		if len(messages) > 0 {
			
			// Limpiamos la respuesta del Jefe
			respuesta := strings.ToUpper(strings.TrimSpace(messages[0].Text.Body))

			if respuesta == "OK" || respuesta == "ACEPTAR" {
				// 1. Identificamos qué módulo autorizaste
				nombre := a.CEO.ObtenerUltimoModuloPendiente()
				contenido := "// Código validado por el Líder\n// Firma: MI_FIRMA_ES_LA_LEY"
				
				// 2. LLAMAMOS AL WRITER SERVICE
				writer := &layer1_core.WriterService{BaseDir: "./internal/core"}
				
				// Tu firma se vuelve LEY
				err := writer.InyectarCodigo(nombre, contenido, "MI_FIRMA_ES_LA_LEY")
				
				if err != nil {
					// ✅ USAMOS EL CEO PARA HABLAR
					a.CEO.EnviarMensajeSoberano("❌ Jefe, falló la inyección: " + err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				// 3. Confirmación final y registro en el Vault [cite: 2026-02-11]
				a.CEO.EnviarMensajeSoberano("✅ ¡Inyectado! El nuevo módulo ya es parte del ADN de GeoChat.")
				a.CEO.DocumentarLogro(nombre, "Inyección de código aprobada vía WhatsApp")
				
				c.JSON(http.StatusOK, gin.H{"status": "inyectado"})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "esperando comando"})
}