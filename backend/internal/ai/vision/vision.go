package vision

import "fmt"

// MacroCapability es la brújula de la IA 5 [cite: 2026-02-10]
type MacroCapability interface {
	IsImplemented() bool
	ProposeImplementation() string
	GetName() string 
}

// HardwareSalesModule: Módulo de Venta de Packs de Hardware
type HardwareSalesModule struct {
	ID          string // <-- AGREGADO: Ahora el compilador ya no dirá "unknown field"
	Status      string // "VISION" o "REAL"
	Target      string // "Motos", "Autos", "Hogares"
	MissingCode string // Descripción técnica
}

// Implementación de los métodos para cumplir con la Interfaz
func (m *HardwareSalesModule) IsImplemented() bool {
	return m.Status == "REAL"
}

func (m *HardwareSalesModule) GetName() string {
	return fmt.Sprintf("Venta Hardware para %s", m.Target)
}

func (m *HardwareSalesModule) ProposeImplementation() string {
	return fmt.Sprintf("IA 5: Jefe, sugiero implementar '%s'. Falta código para: %s.", 
		m.GetName(), m.MissingCode)
}

// --- EL MAPA DE VISIÓN ---
// Aquí es donde la IA 5 monitorea los 'huecos' en blanco
var VisionMap = map[string]MacroCapability{
	"Venta_Motos": &HardwareSalesModule{
		ID:          "Venta_Motos", // ✅ Ahora este campo ya existe arriba
		Status:      "VISION",
		Target:      "Motos",
		MissingCode: "gateway_payment.go",
	},
	"Venta_Hogares": &HardwareSalesModule{
		ID:          "Venta_Hogares",
		Status:      "VISION",
		Target:      "Hogares",
		MissingCode: "home_mesh_logic.go",
	},
}
