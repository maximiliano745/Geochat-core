package vision

import "fmt"

// NOTA: La interfaz MacroCapability YA NO ESTÁ AQUÍ. 
// Go la toma automáticamente de vision.go porque están en el mismo paquete.

// NewRoadmapModule es una función de ayuda para expandir el VisionMap desde aquí
func NewRoadmapModule(id string, target string, missing string) *HardwareSalesModule {
	return &HardwareSalesModule{
		ID:          id,
		Status:      "VISION",
		Target:      target,
		MissingCode: missing,
	}
}

// Ejemplo de una capacidad futura que puedes inicializar desde el main
func (m *HardwareSalesModule) GetRoadmapStatus() string {
	if m.IsImplemented() {
		return fmt.Sprintf("✅ Módulo %s está activo en el Core.", m.ID)
	}
	return fmt.Sprintf("🗓️ Módulo %s programado para desarrollo con el 15%% del fondo.", m.ID)
}

// Aquí podrías agregar más tipos de módulos que cumplan con la interfaz
// Por ejemplo: módulos de publicidad o de energía solar.
