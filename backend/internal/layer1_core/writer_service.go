package layer1_core

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriterService es el brazo ejecutor de la IA.
// Gestiona la escritura física de parches y nuevos módulos.
type WriterService struct {
	// Directorio base donde se inyectará el código (normalmente ./internal/core)
	BaseDir string 
}

// InyectarCodigo realiza la escritura física tras validar la autoridad del Jefe.
// [cite: 2026-02-10] "Sin mi Sí en tiempo real, no ejecuta ningún movimiento".
func (ws *WriterService) InyectarCodigo(nombreArchivo string, contenido string, firma string) error {
	
	// REGLA DE ORO: Validación de Seguridad [cite: 2026-02-10]
	// Comparamos la firma recibida por WhatsApp con la constante de mando.
	if firma != "MI_FIRMA_ES_LA_LEY" {
		return fmt.Errorf("❌ ERROR CRÍTICO: Intento de inyección no autorizada. Firma inválida")
	}

	// Construimos la ruta completa de forma segura
	path := filepath.Join(ws.BaseDir, nombreArchivo)

	// Verificamos si el directorio existe, si no, lo creamos
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	// Escritura del archivo con permisos de sistema (0644)
	// Esto reemplaza el código anterior o crea uno nuevo.
	err := os.WriteFile(path, []byte(contenido), 0644)
	if err != nil {
		return fmt.Errorf("❌ Error físico al escribir en disco: %v", err)
	}

	fmt.Printf("✅ [SISTEMA]: Módulo '%s' inyectado correctamente en el Core.\n", nombreArchivo)
	return nil
}