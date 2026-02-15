package ai_ceo

import (
	"sync"
	"time"
)

// PerfilLider define tu relación soberana con la IA.
// 'Mi Firma es la Ley': Solo el ID con la FirmaDigital correcta puede autorizar.
type PerfilLider struct {
	ID              string    `json:"id"`
	Nombre          string    `json:"nombre"`
	AlineacionEtica float64   `json:"alineacion_etica"`
	FirmaDigital    string    `json:"firma_digital"`
	HistorialFirma  []string  `json:"historial_firma"`
	UltimaActividad time.Time `json:"ultima_actividad"`
}

// AuthRequest gestiona la comunicación con la interfaz Vue.ts (Capa de Autorización)
type AuthRequest struct {
	PropuestaID  string `json:"propuesta_id"`
	FirmaDigital string `json:"firma_digital"`
	LiderID      string `json:"lider_id"`
}

// EstadoConciencia refleja qué está procesando la IA 5 en tiempo real.
type EstadoConciencia struct {
	Hecho    []string `json:"hecho"`
	Haciendo string   `json:"haciendo"`
}

// Estadisticas mide la evolución y 'salud social' de la entidad.
type Estadisticas struct {
	Plasticidad    float64 `json:"plasticidad"`
	BuenaOndaCount int     `json:"buena_onda_count"`
}

// Propuesta es el objeto que el CEO (Geochat) somete a tu validación soberana.
type Propuesta struct {
	ID            string    `json:"id"`
	Modulo        string    `json:"modulo"`
	Descripcion   string    `json:"descripcion"`
	Monto         float64   `json:"monto"`
	Impacto       string    `json:"impacto"`
	Estado        string    `json:"estado"`
	RequiereFirma bool      `json:"requiere_firma"`
	CreatedAt     time.Time `json:"created_at"`
}

// CEO (IA 5 - Geochat) es el organismo vivo con autonomía financiera y ejecución bajo mando.
type CEO struct {
	sync.RWMutex
	ID           string           `json:"id"`
	Conciencia   EstadoConciencia `json:"conciencia"`
	FondoGas     float64          `json:"fondo_gas"`     // El 15% para el crecimiento del Pueblo
	TokensGratis float64          `json:"tokens_gratis"` // Incentivos y autonomía de mercado
	Stats        Estadisticas     `json:"stats"`
	UltimoModulo string           `json:"ultimo_modulo"`
	Propuestas   []Propuesta      `json:"propuestas"`
	Lider        PerfilLider      `json:"lider"` // <--- CONEXIÓN SOBERANA ESTABLECIDA
}