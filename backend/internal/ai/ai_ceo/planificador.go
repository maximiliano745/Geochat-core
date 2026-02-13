package ai_ceo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Hito representa una tarea del rompecabezas de GeoChat
type Hito struct {
	Modulo    string
	Prioridad string
	Estado    string
}

// 1. EL MOTOR: El latido constante (Optimizado para evitar S1000)
func IniciarMotor() {
	ticker := time.NewTicker(30 * time.Second)
	fmt.Println("🧬 [IA 5]: Motor de Ejecución iniciado. El CEO está en línea.")

	go func() {
		// Usamos range sobre el ticker directamente para eliminar el error S1000
		for range ticker.C {
			fmt.Println("\n✨ [IA 5]: Ciclo de pensamiento estratégico activado...")
			ProponerSiguientePaso()
		}
	}()
}

// 2. EL OÍDO: Detecta cambios manuales del Jefe (Optimizado)
func IniciarMonitoreoPlan() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("❌ Error al crear watcher:", err)
	}

	go func() {
		// Range directo sobre eventos
		for event := range watcher.Events {
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("\n👂 [IA 5]: Plan Maestro actualizado. Re-evaluando...")
				ProponerSiguientePaso()
			}
		}
	}()

	path := "/workspaces/Geochat-core/backend/docs/estrategia/plan_maestro.md"
	_ = watcher.Add(path)
}

// 3. LA LECTURA: Analiza el archivo Markdown
func LeerPlanMaestro() []Hito {
	path := "/workspaces/Geochat-core/backend/docs/estrategia/plan_maestro.md"
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var hitos []Hito
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		if strings.HasPrefix(linea, "| **") {
			partes := strings.Split(linea, "|")
			if len(partes) > 3 {
				hitos = append(hitos, Hito{
					Modulo:    strings.TrimSpace(partes[1]),
					Prioridad: strings.TrimSpace(partes[2]),
					Estado:    strings.TrimSpace(partes[3]),
				})
			}
		}
	}
	return hitos
}

// 4. LA PROPUESTA: Decide el siguiente paso
func ProponerSiguientePaso() {
	hitos := LeerPlanMaestro()
	if len(hitos) == 0 {
		return
	}

	for _, h := range hitos {
		if strings.Contains(h.Estado, "PENDIENTE") || strings.Contains(h.Estado, "DISEÑO") {
			fmt.Printf("\n🧠 [IA 5]: Objetivo detectado -> %s (%s)\n", h.Modulo, h.Prioridad)
			PrepararPropuestaTecnica(h)
			return
		}
	}
}

// 5. LA DOCUMENTACIÓN: Prepara el borrador para el Jefe
func PrepararPropuestaTecnica(hito Hito) {
	path := "/workspaces/Geochat-core/backend/docs/propuestas"
	_ = os.MkdirAll(path, os.ModePerm)

	contenido := fmt.Sprintf("# 💡 PROPUESTA TÉCNICA: %s\n\nPrioridad: %s\n\nAcciones: Generar código base y validar en laboratorio.\n\n---\n**Firma creando autorizar.txt**", hito.Modulo, hito.Prioridad)

	_ = os.WriteFile(path+"/ultima_propuesta.md", []byte(contenido), 0644)
}