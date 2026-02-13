package ai_ceo

import (
	"fmt"
	"geochat/internal/database"
	"log"
	"sync"
	"time"
)

// --- ESTRUCTURAS PRINCIPALES ---

// CEO representa la Inteligencia de Gestión Estratégica (IA 5).
// Actúa bajo el mando directo del usuario (El Líder). [cite: 2026-02-10]
type CEO struct {
	sync.RWMutex
	FondoGas     float64
	TokensGratis float64
	Propuestas   []Propuesta
	Stats        Estadisticas
	Lider        PerfilLider
	UltimoModulo string
}

// Propuesta define una acción que requiere validación del Líder.
type Propuesta struct {
	ID                string  `json:"id"`
	Modulo            string  `json:"modulo"`
	Monto             float64 `json:"monto"`
	CostoTokens       float64 `json:"costo_tokens"`
	Destino           string  `json:"destino"`
	ImpactoFinanciero float64 `json:"impacto_financiero"`
	Status            string  `json:"status"` // "PENDIENTE", "APROBADO", "EJECUTADO"
	RequiereFirma     bool    `json:"requiere_firma"`
	Tipo              string  `json:"tipo"`
	FirmaDigital      string  `json:"firma_digital"`
}

// --- CONSTRUCTOR ---

func NewCEO() *CEO {
	return &CEO{
		FondoGas:     100.0,
		TokensGratis: 10.0,
		Propuestas:   []Propuesta{},
		Stats: Estadisticas{
			BuenaOndaCount: 0,
			Plasticidad:    1.0,
		},
		Lider: PerfilLider{
			AlineacionEtica: 1.0,
			UltimaActividad: time.Now(),
		},
		UltimoModulo: "core_logic.go",
	}
}

// --- MÉTODOS DE GESTIÓN Y ADN (CAPA 4 & 5) ---

// DetectarOportunidadTrading persiste oportunidades de mercado en Postgres.
// Es la base del crecimiento del 15%. [cite: 2026-02-10]
func (c *CEO) DetectarOportunidadTrading(precioActual float64) {
	if precioActual < 2000.0 { // Umbral de compra estratégica (Ejemplo PAXG)
		titulo := "Oportunidad de Trading: PAXG"
		descripcion := fmt.Sprintf("Precio actual: %.2f. Por debajo de la media. ¿Reforzamos reserva con el 15%%?", precioActual)

		// Insertamos en la tabla 'tasks' de Postgres para que aparezca en el Panel Admin
		query := `INSERT INTO tasks (titulo, descripcion, estado) VALUES ($1, $2, 'propuesta_ceo')`
		
		if database.DB != nil {
			_, err := database.DB.Exec(query, titulo, descripcion)
			if err != nil {
				log.Printf("⚠️ IA 5 Error: No se pudo persistir oportunidad en Postgres: %v", err)
			} else {
				log.Println("📊 IA 5: Oportunidad de mercado detectada y guardada en Memoria Relacional.")
			}
		}
	}
}

// InyectarCodigoFuncional integra el nuevo código al núcleo tras tu firma soberana.
func (c *CEO) InyectarCodigoFuncional(p Propuesta) {
	c.Lock()
	defer c.Unlock()
	c.UltimoModulo = p.Modulo
	c.Stats.Plasticidad += 0.10
	fmt.Printf("🧬 ADN GeoChat: Módulo '%s' inyectado. Plasticidad: %.2f\n", p.Modulo, c.Stats.Plasticidad)
}

// ProcesarRecompensaSocial traduce valor humano en crecimiento técnico.
func (c *CEO) ProcesarRecompensaSocial(usuarioID string) {
	c.Lock()
	defer c.Unlock()
	c.Stats.BuenaOndaCount++
	c.Stats.Plasticidad += 0.05
	fmt.Printf("🌟 IA CEO: Resonancia de %s. Plasticidad: %.2f\n", usuarioID, c.Stats.Plasticidad)
}

// --- ESTRUCTURAS SECUNDARIAS ---

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

type EvolucionSoftware struct {
	Modulo, Origen, Impacto string
}

func (c *CEO) GetEvolucionesPendientes() []EvolucionSoftware {
	return []EvolucionSoftware{
		{Modulo: "Core", Origen: "IA_Subconscious", Impacto: "Estable"},
	}
}