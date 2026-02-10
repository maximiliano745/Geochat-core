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

// AppData: Estructura fiel a tu db.json estratégico
type AppData struct {
	Project     string `json:"project"`
	Partnership struct {
		PartnerCapital struct {
			Share int `json:"share"`
		} `json:"partner_capital_100"`
		FounderWork struct {
			Share int `json:"share"`
		} `json:"founder_work_laboratory"`
	} `json:"partnership"`
	Assets struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"assets"`
}

// Estructura para el Hash de Voz
type VoiceAuth struct {
	UserHash string `json:"user_hash"`
	DeviceID string `json:"device_id"`
}

var (
	cachedData AppData
	dataMutex  sync.RWMutex
)

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

// Carga datos respetando la jerarquía del Socio
func loadData() {
	wd, _ := os.Getwd()
	dbPath := filepath.Join(wd, "..", "data", "db.json")

	file, err := os.ReadFile(dbPath)
	if err != nil {
		log.Printf("⚠️ Alerta: db.json no encontrado en %s. Inicializando valores de emergencia.", dbPath)
		
		dataMutex.Lock()
		// Inicialización manual siguiendo la nueva estructura
		cachedData.Project = "GeoChat Core (Local)"
		cachedData.Partnership.FounderWork.Share = 60
		cachedData.Partnership.PartnerCapital.Share = 40
		cachedData.Assets.Type = "PAXG"
		cachedData.Assets.Description = "Reserva Oro (Esperando DB)"
		dataMutex.Unlock()
		return
	}

	dataMutex.Lock()
	defer dataMutex.Unlock()
	if err := json.Unmarshal(file, &cachedData); err != nil {
		log.Printf("❌ Error al procesar JSON: %v", err)
		return
	}
	fmt.Println("✅ Datos del Pueblo (60/40) cargados en RAM exitosamente.")
}

func main() {
	loadData()

	// Endpoints
	http.HandleFunc("/api/stats", enableCORS(statsHandler))
	http.HandleFunc("/auth/voice-hash", enableCORS(voiceAuthHandler))
	http.HandleFunc("/status", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GeoChat: Sistema cargado y veloz.")
	}))

	port := "0.0.0.0:8080"
	fmt.Printf("🚀 GeoChat Backend rugiendo en el puerto 8080...\n")
	log.Fatal(http.ListenAndServe(port, nil))
}

//statsHandler: Traduce la estructura estratégica a la visual del Dashboard
func statsHandler(w http.ResponseWriter, r *http.Request) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	// Mapeamos los datos de la DB al formato que espera tu App.vue
	response := map[string]interface{}{
		"saldo":       1.0, // Base para PAXG
		"nombreNodo":  cachedData.Project,
		"mensajes": []string{
			fmt.Sprintf("Socio Fundador: %d%% (Idea/Lab)", cachedData.Partnership.FounderWork.Share),
			fmt.Sprintf("Socio Capital: %d%% (Formalidad)", cachedData.Partnership.PartnerCapital.Share),
			"Activo: " + cachedData.Assets.Description,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

	fmt.Printf("✅ Validación de Hash recibida: %s\n", auth.UserHash)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"info":   "Bóveda E2E preparada",
	})
}
