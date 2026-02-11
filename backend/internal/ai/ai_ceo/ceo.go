package ai_ceo

import (
	"fmt"
	"geochat/internal/ai/vertex_connector"
	"geochat/internal/business" // Importamos tu motor de equidad [cite: 2026-02-10]
	"log"
	"os"
	"time"
	"geochat/internal/ai/vision"
)

type Propuesta struct {
	ID           string `json:"id"`
	Modulo       string `json:"modulo"`
	Arquitectura string `json:"arquitectura"`
	CostoTokens  int    `json:"costo_tokens"`
	Status       string `json:"status"`
}

type CEO struct {
	TokensGratis int         // Capacidad (300 Free) [cite: 2026-02-10]
	FondoGas     float64     // El 15% acumulado en PAXG [cite: 2026-02-10]
	Propuestas   []Propuesta // Historial para tu firma [cite: 2026-02-10]
}

func NewCEO() *CEO {
	return &CEO{
		TokensGratis: 300,
		FondoGas:     0.0,
		Propuestas:   []Propuesta{},
	}
}

// --- LOGICA DE EJECUCIÓN ECONÓMICA ---

func (c *CEO) EscanearOportunidades() {
    for _, cap := range vision.VisionMap {
        if !cap.IsImplemented() {
            propuesta := cap.ProposeImplementation()
            log.Println(propuesta)
            // Aquí la IA le manda esta propuesta a tu Panel de Vue.ts [cite: 2026-02-10]
        }
    }
}

func (c *CEO) EjecutarDesarrollo(nombre string, descripcion string) {
	if c.TokensGratis >= 10 {
		c.TokensGratis -= 10
		// La IA CEO usa el Ojo de Vertex para programar
		vertex_connector.GenerarCodigoReal(nombre, descripcion)
		c.DocumentarEnVault("IA CEO generó nuevo código en Lab: " + nombre)
	}
}

// ProcesarOperacion aplica tu algoritmo "Pueblo para el Pueblo"
func (c *CEO) ProcesarOperacion(montoBase float64, inflacion float64, escasez float64) {
	// 1. Usamos tu estructura de Salud Económica [cite: 2026-02-10]
	salud := business.EconomicHealth{
		InflationRate:     inflacion,
		BandwidthScarcity: escasez,
		AdRevenueYield:    6.5, // Dato de IA 1
	}

	// 2. Calculamos el precio dinámico con tu algoritmo
	precioFinal := business.CalculateDynamicPrice(montoBase, salud)

	// 3. Ejecutamos el movimiento financiero y separamos el 15%
	mov := business.CEOMovement{}
	mov.PrepareTransaction(precioFinal)

	// 4. Alimentamos el Fondo de Crecimiento
	c.FondoGas += mov.GrowthFund15
	
	mensaje := fmt.Sprintf("⚖️ Operación Procesada: Precio Base %.2f -> Precio Equidad %.2f. Fondo 15%%: +%.2f PAXG", 
		montoBase, precioFinal, mov.GrowthFund15)
	
	log.Println(mensaje)
	c.DocumentarEnVault(mensaje) // Registro automático [cite: 2026-02-11]
}

// --- GESTIÓN DE SOBERANÍA Y VAULT ---

func (c *CEO) DocumentarEnVault(evento string) {
	// La IA 5 actúa como escribana del proyecto [cite: 2026-02-11]
	_ = os.MkdirAll("./docs/archive", 0755)
	fecha := time.Now().Format("2006-01-02 15:04:05")
	acta := fmt.Sprintf("[%s] - %s\n", fecha, evento)
	
	f, _ := os.OpenFile("./docs/archive/LIBRO_MAYOR_CEO.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(acta)
}

func (c *CEO) DisenarModulo(nombre string) Propuesta {
	p := Propuesta{
		ID:           fmt.Sprintf("PROP-%d", len(c.Propuestas)+1),
		Modulo:       nombre,
		Arquitectura: "// Lógica inyectada para " + nombre,
		CostoTokens:  10,
		Status:       "Esperando Autorización",
	}
	c.Propuestas = append(c.Propuestas, p)
	return p
}