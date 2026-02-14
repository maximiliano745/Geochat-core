package ai_ceo

import (
	"sync"
	"time"
)

// Propuesta representa una intención de ejecución de la IA que requiere validación.
type Propuesta struct {
	ID            string    `json:"id"`
	Modulo        string    `json:"modulo"`         // Reemplaza a 'Tipo' (Ej: "Infraestructura", "Trading", "Marketing")
	Descripcion   string    `json:"descripcion"`
	Monto         float64   `json:"monto"`          // Reemplaza a 'CostoTokens' (En PAXG o USD según contexto)
	Destino       string    `json:"destino"`        // Wallet o proveedor
	Impacto       string    `json:"impacto"`        // Justificación del crecimiento para el pueblo
	Estado        string    `json:"estado"`         // "Pendiente", "Aprobado", "Rechazado", "Ejecutado"
	Fecha         time.Time `json:"fecha"`
	FirmaDigital  string    `json:"firma_digital"`  // Tu firma (Hash de validación)
	RequiereFirma bool      `json:"requiere_firma"` // El interruptor de seguridad de tus Reglas de Oro
    Status            string    // <--- AGREGAR ESTO (ej: "Pendiente", "Aprobada")
    ImpactoFinanciero float64   // <--- AGREGAR ESTO (para el 15% de crecimiento)
}

// Estadisticas monitorea la salud del organismo vivo GeoChat.
type Estadisticas struct {
	Plasticidad   float64 `json:"plasticidad"`   // Capacidad de la IA de adaptarse
	EnergiaVendida float64 `json:"energia_vendida"`
	AnchoBanda    float64 `json:"ancho_banda"`
	UsuariosActivos int    `json:"usuarios_activos"`
    BuenaOndaCount  int     `json:"buena_onda_count"` // <--- AGREGAR ESTO
}

// PerfilLider define tu relación con la IA.
type PerfilLider struct {
    ID              string
    Nombre          string
    AlineacionEtica float64
    FirmaDigital    string
    HistorialFirma  []string // <--- AGREGAR ESTO PARA EL RASTRO SOBERANO
    UltimaActividad time.Time `json:"ultima_actividad"` // <--- AGREGA ESTA LÍNEA
}

// EstadoConciencia define el modo de operación.
type EstadoConciencia struct {
    Hecho     []string `json:"hecho"`      // <--- Asegúrate que sea HECHO (en plural)
    Haciendo  string   `json:"haciendo"`
    PorHacer  []string `json:"por_hacer"`
    StatusRed string   `json:"status_red"`
    Proposito string   `json:"proposito"`
}

// CEO es la estructura principal (Capa de Inteligencia de Negocio).
type CEO struct {
	sync.RWMutex
	ID           string           `json:"id"`
	Version      string           `json:"version"`
	FondoOcioso  float64          `json:"fondo_ocioso"` // El 15% para crecimiento
	FondoGas     float64          `json:"fondo_gas"`    // Combustible operativo
	TokensGratis float64          `json:"tokens_gratis"`
	UltimoModulo string           `json:"ultimo_modulo"`
	Propuestas   []Propuesta      `json:"propuestas"`
	Stats        Estadisticas     `json:"stats"`
	Lider        PerfilLider      `json:"lider"`
	Conciencia   EstadoConciencia `json:"conciencia"`
}

// AuthRequest es para la comunicación con la interfaz Vue.ts
type AuthRequest struct {
    PropuestaID  string `json:"propuesta_id"`
    FirmaDigital string `json:"firma_digital"`
    LiderID      string `json:"lider_id"`
}