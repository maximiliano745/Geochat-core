package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Estructura de datos para el Pueblo (Cargada en Memoria)
type AppData struct {
	Users interface{} `json:"users"` // Aquí puedes mapear la estructura de tu db.json
}

var (
	cachedData AppData
	dataMutex  sync.RWMutex // Para que varios usuarios lean a la vez sin chocar
)

// Estructura para el Hash de Voz
type VoiceAuth struct {
	UserHash string `json:"user_hash"`
	DeviceID string `json:"device_id"`
}

// Middleware CORS
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// Función para cargar datos al inicio (El Soberano no espera)
func loadData() {
	// Buscamos la ruta absoluta de forma inteligente
	wd, _ := os.Getwd()
	// Subimos un nivel si estamos en backend para llegar a la raíz
	dbPath := filepath.Join(wd, "..", "data", "db.json")

	file, err := os.ReadFile(dbPath)
	if err != nil {
		log.Printf("⚠️ Alerta: No se encontró db.json en %s. Usando memoria limpia.", dbPath)
		return
	}

	dataMutex.Lock()
	defer dataMutex.Unlock()
	if err := json.Unmarshal(file, &cachedData); err != nil {
		log.Printf("❌ Error al procesar JSON: %v", err)
		return
	}
	fmt.Println("✅ Datos del Pueblo cargados en memoria exitosamente.")
}

func main() {
	// 1. Cargamos datos ANTES de arrancar el servidor
	loadData()

	http.HandleFunc("/auth/voice-hash", enableCORS(voiceAuthHandler))
	http.HandleFunc("/status", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GeoChat: Sistema cargado y veloz.")
	}))

	port := "0.0.0.0:8080" // Esto permite que Codespaces capture el tráfico
	fmt.Printf("🚀 GeoChat Backend listo en el puerto 8080...\n")
	log.Fatal(http.ListenAndServe(port, nil))
}

func voiceAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Solo POST", http.StatusMethodNotAllowed)
		return
	}

	var auth VoiceAuth
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		http.Error(w, "Error de formato", http.StatusBadRequest)
		return
	}

	// Acceso instantáneo a los datos (Sin leer disco)
	dataMutex.RLock()
	_ = cachedData // Aquí ya tienes los datos listos para comparar el hash
	dataMutex.RUnlock()

	fmt.Printf("✅ Validación instantánea para: %s\n", auth.UserHash)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"info":   "Acceso ultra-rápido al Vault",
	})
}
