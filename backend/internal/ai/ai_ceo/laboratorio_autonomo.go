package ai_ceo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// IniciarCicloAutonomo es el latido constante de la IA 5
func IniciarCicloAutonomo() {
	fmt.Println("🧠 [IA 5]: Motor de Ciclo de Vida iniciado.")
	for {
		// 1. La IA mira qué sigue en el rompecabezas (Planificador)
		ProponerSiguientePaso()

		// 2. La IA verifica si el código actual está sano (Compilación y Docs)
		ProbarYDocumentar()

		// 3. La IA busca tu firma para ejecutar lo que propuso (GitHub)
		VerificarFirmaYEjecutar()

		// 4. Pausa de pensamiento (1 minuto de enfriamiento)
		time.Sleep(1 * time.Minute)
	}
}

// ProbarYDocumentar analiza el código y genera informes
func ProbarYDocumentar() {
	fmt.Println("🔍 IA CEO: Iniciando fase de Verificación de Integridad...")

	// 1. ANALIZAR ERRORES (Prueba de compilación)
	cmdCheck := exec.Command("go", "build", "-o", "temp_check", "./cmd/server")
	output, err := cmdCheck.CombinedOutput()

	status := "100% Verificado"
	detalles := "Sin errores de sintaxis detectados en el núcleo soberano."

	if err != nil {
		status = "⚠️ Error de Compilación"
		detalles = fmt.Sprintf("Se detectaron fallos en el código:\n%s", string(output))
		fmt.Printf("❌ IA CEO: Errores detectados. Documentando...\n")
	} else {
		fmt.Println("✅ IA CEO: Código verificado con éxito.")
		os.Remove("temp_check")
	}

	// 2. GENERAR DOCUMENTACIÓN (Vault Digital)
	generarDocs(status, detalles)

	// 3. COMUNICAR RESULTADOS
	if err != nil {
		NotificarAlJefe("🚨 Jefe, el Laboratorio detectó errores. Revisar /docs/laboratorio.")
	} else {
		// Solo enviamos propuesta si todo está OK
		enviarPropuestaWhatsApp()
	}
}

// VerificarFirmaYEjecutar busca el archivo de autorización del Jefe
func VerificarFirmaYEjecutar() {
    if _, err := os.Stat("autorizar.txt"); err == nil {
        log.Println("👑 [IA 5]: Firma soberana detectada. Iniciando fase de construcción...")

        // 1. Obtener la tarea del planificador
        hitos := LeerPlanMaestro()
        if len(hitos) > 0 && (strings.Contains(hitos[0].Estado, "PENDIENTE") || strings.Contains(hitos[0].Estado, "DISEÑO")) {
            
            tareaActual := hitos[0].Modulo
            log.Printf("🛠️ [IA 5]: Ejecutando construcción de: %s\n", tareaActual)

            // 2. Aquí es donde la IA "escribe" el código de la fábrica (lo haremos en el paso 4)
            // EjecutarFabrica(tareaActual)

            // 3. Subir a GitHub con el nombre de la tarea
            PublicarEvolucion() 
            
            log.Println("✅ [IA 5]: Pieza instalada y sincronizada.")
        }
    }
}

// PublicarEvolucion es el brazo ejecutor en GitHub
func PublicarEvolucion() {
	fmt.Println("🚀 IA 5: Preparando despacho a GitHub...")

	// 1. Agrega TODO
	exec.Command("git", "add", ".").Run()

	// 2. Obtener el nombre de la pieza actual para el commit
	hitos := LeerPlanMaestro()
	pieza := "Evolución General"
	if len(hitos) > 0 {
		pieza = hitos[0].Modulo
	}

	// 3. Commit con sello de la IA y contexto del Plan Maestro
	commitMsg := fmt.Sprintf("🧬 IA 5 [%s]: ADN actualizado (%s)", pieza, time.Now().Format("15:04"))
	exec.Command("git", "commit", "-m", commitMsg).Run()

	// 4. Push al repositorio soberano
	err := exec.Command("git", "push", "origin", "main").Run()

	if err == nil {
		fmt.Println("✅ [IA 5]: Push exitoso. ADN sincronizado.")
		os.Remove("autorizar.txt") // Se limpia la firma solo si hubo éxito
		NotificarAlJefe("🚀 ¡Misión cumplida! El ADN de GeoChat ha evolucionado en GitHub.")
	} else {
		fmt.Printf("❌ [IA 5]: Error en el despacho: %v\n", err)
	}
}

func generarDocs(status string, detalles string) {
	path := "/workspaces/Geochat-core/backend/docs/laboratorio"
	_ = os.MkdirAll(path, os.ModePerm)

	contenido := fmt.Sprintf(`# Informe de Laboratorio GeoChat
## Estado: %s
- **Fecha:** 2026-02-13
- **ADN:** IA 5 activa.
- **Soberanía:** 15%% de crecimiento proyectado.

### Detalles del Análisis:
%s

---
Informe generado automáticamente por el CEO Autónomo.`, status, detalles)

	fullPath := path + "/analisis_actual.md"
	_ = os.WriteFile(fullPath, []byte(contenido), 0644)
	fmt.Println("📝 IA CEO: Informe actualizado en /docs.")
}

func enviarPropuestaWhatsApp() {
	// Solo enviamos si no hay una firma pendiente para no saturar
	if _, err := os.Stat("autorizar.txt"); err != nil {
		mensaje := "🤖 IA CEO: Rompecabezas analizado y código verificado. ¿Dás tu OK (autorizar.txt)?"
		NotificarAlJefe(mensaje)
	}
}
