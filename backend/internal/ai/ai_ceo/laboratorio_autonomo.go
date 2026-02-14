package ai_ceo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// IniciarCicloAutonomo es el latido constante de GeoChat.
// Ahora pertenece a *CEO para tener acceso a toda la red.
func (c *CEO) IniciarCicloAutonomo() {
	fmt.Println("🧠 [GeoChat]: Motor de Ciclo de Vida iniciado.")
	for {
		// 1. Proponer nuevos pasos basados en el backlog
		c.EjecutarCicloDesarrollo() 

		// 2. Verificar que el código actual compile
		c.ProbarYDocumentar()

		// 3. Si detecta la firma, ejecuta el Git Push
		c.VerificarFirmaYEjecutar()

		time.Sleep(1 * time.Minute)
	}
}

// ProbarYDocumentar analiza el código y genera informes para el Vault.
func (c *CEO) ProbarYDocumentar() {
	fmt.Println("🔍 [GeoChat]: Iniciando fase de Verificación de Integridad...")

	// Intentamos compilar el servidor
	cmdCheck := exec.Command("go", "build", "-o", "temp_check", "./cmd/server/main.go")
	output, err := cmdCheck.CombinedOutput()

	status := "100% Verificado"
	detalles := "Sin errores de sintaxis detectados en el núcleo soberano."

	if err != nil {
		status = "⚠️ Error de Compilación"
		detalles = fmt.Sprintf("Se detectaron fallos en el código:\n%s", string(output))
		fmt.Printf("❌ [GeoChat]: Errores detectados en el ADN.\n")
	} else {
		fmt.Println("✅ [GeoChat]: Código verificado con éxito.")
		os.Remove("temp_check")
	}

	// Guardamos el informe en el Vault Digital
	c.generarDocs(status, detalles)

	if err != nil {
		c.EnviarMensajeSoberano("🚨 [GeoChat]: Jefe, el Laboratorio detectó errores. Revisar Vault.")
	}
}

// generarDocs crea el informe en Markdown dentro de tu Vault.
func (c *CEO) generarDocs(status string, detalles string) {
	path := "/workspaces/Geochat-core/backend/docs/laboratorio"
	_ = os.MkdirAll(path, os.ModePerm)

	contenido := fmt.Sprintf(`# 🧬 Informe de Laboratorio GeoChat
## Estado del Sistema: %s
- **Fecha:** %s
- **Identidad:** GeoChat activa y evolucionando.
- **Soberanía:** Fondo de Gas actual: %.2f PAXG.

### Detalles del Análisis Técnico:
%s

---
*Informe generado automáticamente por GeoChat (CEO Autónomo).*`, 
	status, time.Now().Format("2006-01-02"), c.FondoGas, detalles)

	fullPath := path + "/analisis_actual.md"
	_ = os.WriteFile(fullPath, []byte(contenido), 0644)
	
	// Actualizamos la conciencia del CEO para el Dashboard
	c.Lock()
	c.Conciencia.Hecho = append(c.Conciencia.Hecho, fmt.Sprintf("Análisis de laboratorio: %s", status))
	c.Unlock()
}

// VerificarFirmaYEjecutar busca el permiso manual o digital.
func (c *CEO) VerificarFirmaYEjecutar() {
	// Buscamos si hay alguna propuesta marcada como "LISTO_PARA_PUSH" o el archivo manual
	if _, err := os.Stat("autorizar.txt"); err == nil {
		log.Println("👑 [GeoChat]: Firma soberana detectada. Iniciando publicación...")
		c.PublicarEvolucion()
	}
}

// PublicarEvolucion sincroniza el ADN con GitHub.
func (c *CEO) PublicarEvolucion() {
	fmt.Println("🚀 [GeoChat]: Preparando despacho de ADN a GitHub...")
	
	exec.Command("git", "add", ".").Run()

	commitMsg := fmt.Sprintf("🧬 Evolución ADN: GeoChat activa [%s] (%s)", c.UltimoModulo, time.Now().Format("15:04"))
	
	err := exec.Command("git", "commit", "-m", commitMsg).Run()
	if err != nil {
		log.Printf("⚠️ [GeoChat]: No hay cambios para commit: %v", err)
		return
	}

	errPush := exec.Command("git", "push", "origin", "main").Run()

	if errPush == nil {
		fmt.Println("✅ [GeoChat]: ADN sincronizado con éxito.")
		os.Remove("autorizar.txt")
		c.EnviarMensajeSoberano("🚀 ¡ADN Evolucionado! El código ha sido actualizado en GitHub exitosamente.")
	} else {
		log.Printf("❌ [GeoChat]: Error en Push: %v\n", errPush)
	}
}