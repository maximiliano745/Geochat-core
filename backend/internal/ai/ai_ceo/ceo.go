package ai_ceo

import (
	"fmt"
	"sync"
	"time"
)

// --- ESTRUCTURAS PRINCIPALES ---

// CEO representa la Inteligencia de Gestión Estratégica (IA 5).
// Actúa bajo el mando directo del usuario (El Líder). [cite: 2026-02-10]
type CEO struct {
	sync.RWMutex
	FondoGas      float64
	TokensGratis  float64
	Propuestas    []Propuesta
	Stats         Estadisticas
	Lider         PerfilLider
	UltimoModulo  string 
}

// Propuesta define una acción que requiere validación del Líder.
type Propuesta struct {
	ID                string
	Modulo            string
	Monto             float64
	CostoTokens       float64
	Destino           string
	ImpactoFinanciero float64
	Status            string // "PENDIENTE", "APROBADO", "EJECUTADO"
	RequiereFirma     bool
	Tipo              string
	FirmaDigital      string
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

// --- MÉTODOS DE GESTIÓN Y ADADN ---

// InyectarCodigoFuncional integra el nuevo código al núcleo tras tu firma.
func (c *CEO) InyectarCodigoFuncional(p Propuesta) {
	c.Lock()
	defer c.Unlock()
	c.UltimoModulo = p.Modulo
	c.Stats.Plasticidad += 0.10
	fmt.Printf("🧬 ADN GeoChat: Módulo '%s' inyectado. Plasticidad: %.2f\n", p.Modulo, c.Stats.Plasticidad)
}

// TieneFirmaDelLider valida si el mensaje de WhatsApp contenía tu autorización.
func (p Propuesta) TieneFirmaDelLider() bool {
	return p.FirmaDigital != ""
}

// DocumentarLogro registra el éxito en el historial y el Vault [cite: 2026-02-11].
func (c *CEO) DocumentarLogro(modulo, detalle string) {
	c.Lock()
	defer c.Unlock()
	logMsg := fmt.Sprintf("[%s] %s: %s", time.Now().Format(time.RFC822), modulo, detalle)
	c.Lider.HistorialFirma = append(c.Lider.HistorialFirma, logMsg)
	c.Lider.UltimaActividad = time.Now()
}

// DocumentarEnVault es un alias para compatibilidad con ejecucion.go.

// ProcesarRecompensaSocial traduce valor humano en crecimiento técnico.
func (c *CEO) ProcesarRecompensaSocial(usuarioID string) {
	c.Lock()
	defer c.Unlock()
	c.Stats.BuenaOndaCount++
	c.Stats.Plasticidad += 0.05
	fmt.Printf("🌟 IA CEO: Resonancia de %s. Plasticidad: %.2f\n", usuarioID, c.Stats.Plasticidad)
}

// ObtenerUltimoModuloPendiente informa al ejecutor qué archivo subir a Git.
func (c *CEO) ObtenerUltimoModuloPendiente() string {
	c.RLock()
	defer c.RUnlock()
	return c.UltimoModulo
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
// getEvolucionesPendientes conecta con la lógica de crecimiento [cite: 2026-02-10].
// Permite que la IA 5 sepa qué mejoras técnicas ha detectado la IA 2.
func (c *CEO) getEvolucionesPendientes() []EvolucionSoftware {
	// Por ahora devolvemos una lista interna basada en el estado actual.
	// Esto se conectará luego con el motor de evolución profunda.
	return []EvolucionSoftware{
		{Modulo: "Core", Origen: "IA_Subconscious", Impacto: "Estable"},
	}
}