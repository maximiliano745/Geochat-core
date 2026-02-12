package ai_ceo

import (
	"fmt"
	"sync"
	"time"
)

// Estructura Maestra del CEO [cite: 2026-02-10]
type CEO struct {
	sync.RWMutex
	FondoGas     float64
	TokensGratis int
	Propuestas   []Propuesta
	Stats        Estadisticas
	Lider        PerfilLider
}

// Propuesta con TODOS los campos necesarios para la ejecución
type Propuesta struct {
	ID                string
	Modulo            string
	Monto             float64
	CostoTokens       float64 // <--- ESTO ARREGLA EL ERROR 'undefined: CostoTokens'
	Destino           string  // <--- ESTO ARREGLA EL ERROR 'undefined: Destino'
	ImpactoFinanciero float64
	Status            string
	RequiereFirma     bool
	Tipo              string
	FirmaDigital      string
}

// ProcesarRecompensaSocial: La IA detecta valor social y lo traduce en crecimiento.
// [cite: 2026-02-10] "GeoChat es de pueblo para el pueblo".
func (c *CEO) ProcesarRecompensaSocial(usuarioID string) {
	c.Lock()
	defer c.Unlock()

	// 1. Aumentamos las estadísticas globales de "Buena Onda"
	c.Stats.BuenaOndaCount++
	
	// 2. La plasticidad del sistema mejora (el organismo aprende)
	c.Stats.Plasticidad += 0.05
	
	// 3. Documentamos que el usuario aportó valor
	fmt.Printf("🌟 IA CEO: Usuario %s generó resonancia positiva. Plasticidad: %.2f\n", 
		usuarioID, c.Stats.Plasticidad)
}


// NewCEO inicializa la entidad estratégica bajo tu mando directo.
func NewCEO() *CEO {
	return &CEO{
		FondoGas:     100.0, // El 15% inicial para empezar a crecer
		TokensGratis: 10,    // Incentivo inicial de Buena Onda
		Propuestas:   []Propuesta{},
		Stats: Estadisticas{
			BuenaOndaCount: 0,
			Plasticidad:    1.0,
		},
		Lider: PerfilLider{
			AlineacionEtica: 1.0, // El sistema nace en total armonía contigo
		},
	}
}

// Métodos de apoyo que usa ejecucion.go
func (p *Propuesta) TieneFirmaDelLider() bool {
	return p.FirmaDigital != ""
}

func (c *CEO) InyectarCodigoFuncional(p Propuesta) {
	// Simulación de plasticidad del sistema
}


// Estructuras secundarias
type Estadisticas struct {
	BuenaOndaCount int
	MalaOndaCount  int
	Plasticidad    float64
}

type PerfilLider struct {
	AlineacionEtica float64
	HistorialFirma  []string
	UltimaActividad time.Time
}

// EvolucionSoftware representa una mejora técnica detectada por el Subconsciente (IA 2)
type EvolucionSoftware struct {
    Modulo  string
    Origen  string
    Impacto string
}

// getEvolucionesPendientes conecta con la lógica de crecimiento [cite: 2026-02-10]
func (c *CEO) getEvolucionesPendientes() []EvolucionSoftware {
    // Por ahora devolvemos una lista estática para que el sistema compile
    return []EvolucionSoftware{
        {Modulo: "Core", Origen: "IA_Subconscious", Impacto: "Estable"},
    }
}