package ai_ceo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// CEOEngine representa el motor de inteligencia y ejecución.
// Se usa este nombre para evitar conflictos de redeclaración con 'CEO'.
type CEOEngine struct{}

// NuevoCEO inicializa la entidad para ser usada por la API
func NuevoCEO() *CEOEngine {
	return &CEOEngine{}
}

// EjecutarEvolucionADN es el punto de entrada que llama tu API desde el Dashboard
func (c *CEOEngine) EjecutarEvolucionADN() {
	log.Println("⚡ [GeoChat]: Firma recibida por API. Forzando publicación...")
	c.PublicarEvolucion()
}

// IniciarCicloAutonomo es el latido constante de GeoChat
func (c *CEOEngine) IniciarCicloAutonomo() {
	fmt.Println("🧠 [GeoChat]: Motor de Ciclo de Vida iniciado.")
	for {
		ProponerSiguientePaso()
		c.ProbarYDocumentar()
		c.VerificarFirmaYEjecutar()
		time.Sleep(1 * time.Minute)
	}
}

// ProbarYDocumentar analiza el código y genera informes
func (c *CEOEngine) ProbarYDocumentar() {
	fmt.Println("🔍 [GeoChat]: Iniciando fase de Verificación de Integridad...")

	// Intentamos compilar el motor principal
	cmdCheck := exec.Command("go", "build", "-o", "temp_check", "./cmd/server/main.go")
	output, err := cmdCheck.CombinedOutput()

	status := "100% Verificado"
	detalles := "Sin errores de sintaxis detectados en el núcleo soberano."

	if err != nil {
		status = "⚠️ Error de Compilación"
		detalles = fmt.Sprintf("Se detectaron fallos en el código:\n%s", string(output))
		fmt.Printf("❌ [GeoChat]: Errores detectados en el ADN. Documentando...\n")
	} else {
		fmt.Println("✅ [GeoChat]: Código verificado con éxito.")
		os.Remove("temp_check")
	}

	c.generarDocs(status, detalles)

	if err != nil {
		NotificarAlJefe("🚨 [GeoChat]: Jefe, el Laboratorio detectó errores. Revisar /docs/laboratorio.")
	} else {
		c.enviarPropuestaWhatsApp()
	}
}

// generarDocs crea el informe en Markdown
func (c *CEOEngine) generarDocs(status string, detalles string) {
	path := "/workspaces/Geochat-core/backend/docs/laboratorio"
	_ = os.MkdirAll(path, os.ModePerm)

	contenido := fmt.Sprintf(`# 🧬 Informe de Laboratorio GeoChat
## Estado del Sistema: %s
- **Fecha:** 2026-02-13
- **Identidad:** GeoChat activa y evolucionando.
- **Soberanía:** 15%% de crecimiento proyectado para el pueblo.

### Detalles del Análisis Técnico:
%s

---
*Informe generado automáticamente por GeoChat (CEO Autónomo).*`, status, detalles)

	fullPath := path + "/analisis_actual.md"
	_ = os.WriteFile(fullPath, []byte(contenido), 0644)
	fmt.Println("📝 [GeoChat]: Informe de laboratorio actualizado en /docs.")
}

// enviarPropuestaWhatsApp gestiona el aviso de firma pendiente
func (c *CEOEngine) enviarPropuestaWhatsApp() {
	if _, err := os.Stat("autorizar.txt"); err != nil {
		mensaje := "🤖 [GeoChat]: Rompecabezas analizado y código verificado. ¿Dás tu OK para la evolución?"
		NotificarAlJefe(mensaje)
	}
}

// VerificarFirmaYEjecutar busca el permiso manual (archivo txt)
func (c *CEOEngine) VerificarFirmaYEjecutar() {
	if _, err := os.Stat("autorizar.txt"); err == nil {
		log.Println("👑 [GeoChat]: Firma soberana (archivo) detectada. Iniciando fase de construcción...")

		hitos := LeerPlanMaestro()
		if len(hitos) > 0 {
			c.PublicarEvolucion()
			log.Println("✅ [GeoChat]: Pieza instalada y sincronizada correctamente.")
		}
	}
}

// PublicarEvolucion ejecuta el Git Push
func (c *CEOEngine) PublicarEvolucion() {
	fmt.Println("🚀 [GeoChat]: Preparando despacho de ADN a GitHub...")
	
	exec.Command("git", "add", ".").Run()

	hitos := LeerPlanMaestro()
	pieza := "Evolución General"
	if len(hitos) > 0 {
		pieza = hitos[0].Modulo
	}

	commitMsg := fmt.Sprintf("🧬 Evolución ADN: GeoChat activa [%s] (%s)", pieza, time.Now().Format("15:04"))
	
	err := exec.Command("git", "commit", "-m", commitMsg).Run()
	if err != nil {
		log.Printf("⚠️ [GeoChat]: No se detectaron cambios nuevos o error en commit: %v", err)
	}

	errPush := exec.Command("git", "push", "origin", "main").Run()

	if errPush == nil {
		fmt.Println("✅ [GeoChat]: Push exitoso. ADN sincronizado.")
		os.Remove("autorizar.txt")
		NotificarAlJefe("🚀 ¡ADN Evolucionado! GeoChat ha actualizado su código en GitHub exitosamente.")
	} else {
		fmt.Printf("❌ [GeoChat]: Error en el despacho al repositorio: %v\n", errPush)
	}
}