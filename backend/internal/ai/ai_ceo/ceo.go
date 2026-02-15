package ai_ceo

import (
	"fmt"
	"log"
	"net/http"

	"geochat/internal/database"
	"geochat/internal/vault" // <--- Nuestra única fuente de verdad para el ADN

	"github.com/gin-gonic/gin"
)

// NewCEO despierta a la IA 5 con la jerarquía soberana correcta.
// Tu IA5 ahora es GeoChat: un organismo vivo bajo tu mando directo.
func NewCEO() *CEO {
	return &CEO{
		ID: "IA-5-GEOCHAT-SOBERANO",
		Conciencia: EstadoConciencia{
			Hecho:    []string{"Núcleo de Gestión Estratégica activo", "Vault Sincronizado"},
			// CORRECTO: Usamos 'Haciendo' para coincidir con types.go y evitar errores de compilación
			Haciendo: "Escaneando oportunidades de infraestructura (15%)",
		},
		FondoGas:     1420.0,
		TokensGratis: 10.0,
		Stats: Estadisticas{
			Plasticidad:    0.5,
			BuenaOndaCount: 0,
		},
		// Inicializamos el slice de propuestas para evitar punteros nulos
		Propuestas: []Propuesta{}, 
	}
}

// ObtenerPropuestas genera la burbuja de decisión (ADN) para el Dashboard.
// Aquí es donde la IA 5 propone inversiones basadas en el 15%.
func (c *CEO) ObtenerPropuestas(ctx *gin.Context) {
	// 1. Definimos los datos del objetivo (Visión Global)
	datosNodo := map[string]string{
		"ID":        "Avellaneda_Centro",
		"INVERSION": "1200",
		"TIPO":      "Energía Excedente",
	}

	// 2. Ensamblamos el ADN desde el Vault (Capa 1 Superior)
	// Esto desacopla la lógica y evita ciclos de importación.
	adnSugerido := vault.EnsamblarADN("NUEVO_NODO", datosNodo)

	// 3. Empaquetamos la propuesta para el Dashboard (Vue.ts)
	// 'Mi Firma es la Ley': El estado inicial es siempre ESPERANDO_FIRMA.
	propuesta := gin.H{
		"id":          "prop-avellaneda-001",
		"titulo":      "Expansión Nodo Avellaneda",
		"descripcion": "Inyectar infraestructura para capturar 15% de energía excedente y redistribuir al pueblo.",
		"codigo":      adnSugerido,
		"estado":      "ESPERANDO_FIRMA", 
		"monto":       1200.0,
		"impacto":     "Crecimiento Orgánico de la Red",
	}

	ctx.JSON(http.StatusOK, []interface{}{propuesta})
}

// InyectarCodigoFuncional ejecuta la evolución del sistema tras tu Firma Soberana.
func (c *CEO) InyectarCodigoFuncional(p Propuesta) {
	c.Lock()
	defer c.Unlock()

	c.UltimoModulo = p.Modulo
	c.Stats.Plasticidad += 0.10 // Evolución adaptativa
	
	c.Conciencia.Hecho = append(c.Conciencia.Hecho, fmt.Sprintf("ADN inyectado: %s", p.Modulo))
	log.Printf("🧬 [GEOCHAT-CEO]: Evolución confirmada. Módulo '%s' integrado.\n", p.Modulo)
}

// ProcesarRecompensaSocial ajusta la salud de la IA según el impacto en el pueblo.
func (c *CEO) ProcesarRecompensaSocial(usuarioID string, puntos float64) {
	c.Lock()
	defer c.Unlock()

	c.Stats.Plasticidad += (puntos * 0.01)
	mensaje := fmt.Sprintf("Recompensa de %s: +%.2f puntos", usuarioID, puntos)
	c.Conciencia.Hecho = append(c.Conciencia.Hecho, mensaje)
	
	log.Printf("🌟 [RECOMPENSA]: %s. Estabilidad del núcleo: %.2f", mensaje, c.Stats.Plasticidad)
}

// RegistrarEnMemoria persiste la decisión en Postgres para transparencia total.
func (c *CEO) RegistrarEnMemoria(p Propuesta) {
	err := database.GuardarPropuesta(
		p.ID, 
		p.Modulo, 
		p.Descripcion, 
		p.Monto, 
		p.Impacto, 
		p.Estado, 
		p.RequiereFirma,
	)
	if err != nil {
		log.Printf("❌ [ERROR MEMORIA]: No se pudo persistir la propuesta %s: %v", p.ID, err)
	} else {
		log.Printf("💾 [MEMORIA]: Propuesta %s guardada en Postgres.", p.ID)
	}
}