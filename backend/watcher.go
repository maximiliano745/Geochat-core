package main

import (
	"log"
	"os"
	"os/exec"
	"time"
	"syscall"
	"github.com/fsnotify/fsnotify" // go get github.com/fsnotify/fsnotify
)

func watcherMain() {
	// 1. Iniciamos el servidor de GeoChat por primera vez
	cmd := startGeoChat()

	// 2. Configuramos el Vigilante
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Vigilamos la carpeta donde la IA CEO guarda sus parches autorizados [cite: 2026-02-10]
	watcher.Add("./internal/core")

	log.Println("⚡ GeoChat Vivo: Esperando parches autorizados por el Dueño...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok { return }
			
			// Detectar creación o modificación
			if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
				log.Printf("🔄 Evolución detectada en: %s", event.Name)
				
				// Breve pausa para asegurar escritura completa
				time.Sleep(500 * time.Millisecond)
				
				log.Println("🛠️ Re-compilando ADN de GeoChat...")
				
				// Matamos el proceso viejo suavemente
				if cmd.Process != nil {
					cmd.Process.Signal(syscall.SIGTERM)
					cmd.Wait() // Esperamos a que termine de cerrar
				}
				
				// Lanzamos la nueva versión evolucionada
				cmd = startGeoChat()
			}
		case err, ok := <-watcher.Errors:
			if !ok { return }
			log.Println("⚠️ Error de vigilancia:", err)
		}
	}
}

func startGeoChat() *exec.Cmd {
	// Ejecuta el comando de arranque del servidor
	cmd := exec.Command("go", "run", "./cmd/geochat/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Start()
	if err != nil {
		log.Printf("❌ Error al arrancar el Core: %v", err)
	} else {
		log.Println("✅ Servidor GeoChat Operativo")
	}
	return cmd
}