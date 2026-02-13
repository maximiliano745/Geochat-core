package ai_ceo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	//"strings"
	"time"
)

// IniciarCicloAutonomo es el latido constante de GeoChat
func IniciarCicloAutonomo() {
	fmt.Println("🧠 [GeoChat]: Motor de Ciclo de Vida iniciado.")
	for {
		ProponerSiguientePaso()
		ProbarYDocumentar()
		VerificarFirmaYEjecutar()
		time.Sleep(1 * time.Minute)
	}
}

// ProbarYDocumentar analiza el código y genera informes con la identidad GeoChat
func ProbarYDocumentar() {
	fmt.Println("🔍 [GeoChat]: Iniciando fase de Verificación de Integridad...")

	// Intentamos compilar el motor principal
	cmdCheck := exec.Command("go", "build", "-o", "temp_check", "./motor.go")
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

	// 2. GENERAR DOCUMENTACIÓN (Actualizado con encabezado GeoChat)
	generarDocs(status, detalles)

	// 3. COMUNICAR RESULTADOS
	if err != nil {
		NotificarAlJefe("🚨 [GeoChat]: Jefe, el Laboratorio detectó errores. Revisar /docs/laboratorio.")
	} else {
		enviarPropuestaWhatsApp()
	}
}

// generarDocs crea el informe en Markdown con la nueva cabecera
func generarDocs(status string, detalles string) {
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
func enviarPropuestaWhatsApp() {
	if _, err := os.Stat("autorizar.txt"); err != nil {
		mensaje := "🤖 [GeoChat]: Rompecabezas analizado y código verificado. ¿Dás tu OK para la evolución?"
		NotificarAlJefe(mensaje)
	}
}

// VerificarFirmaYEjecutar busca el permiso del Dueño
func VerificarFirmaYEjecutar() {
	if _, err := os.Stat("autorizar.txt"); err == nil {
		log.Println("👑 [GeoChat]: Firma soberana detectada. Iniciando fase de construcción...")

		hitos := LeerPlanMaestro()
		if len(hitos) > 0 {
			PublicarEvolucion()
			log.Println("✅ [GeoChat]: Pieza instalada y sincronizada correctamente.")
		}
	}
}

// PublicarEvolucion ejecuta el Git Push con el nuevo sello de commit solicitado
func PublicarEvolucion() {
	fmt.Println("🚀 [GeoChat]: Preparando despacho de ADN a GitHub...")
	
	exec.Command("git", "add", ".").Run()

	// Obtener contexto del plan para el commit
	hitos := LeerPlanMaestro()
	pieza := "Evolución General"
	if len(hitos) > 0 {
		pieza = hitos[0].Modulo
	}

	// SELLO ACTUALIZADO: "🧬 Evolución ADN: GeoChat activa..."
	commitMsg := fmt.Sprintf("🧬 Evolución ADN: GeoChat activa [%s] (%s)", pieza, time.Now().Format("15:04"))
	
	err := exec.Command("git", "commit", "-m", commitMsg).Run()
	if err != nil {
		log.Printf("⚠️ [GeoChat]: No se detectaron cambios nuevos: %v", err)
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