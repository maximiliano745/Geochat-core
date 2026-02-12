package ai_friend

import (
	"log"
	"geochat/internal/ai/ai_ceo" // Importamos para que reconozca al CEO
)

// 1. DEFINIMOS LA ESTRUCTURA (Para que no de error 'undefined: Mente1')
type FriendIA struct {
	CEO *ai_ceo.CEO // Conexión necesaria para la recompensa social
}

// AnalizarEnergia es una función auxiliar del paquete
func AnalizarEnergia(input string) int {
	longitud := len(input)
	
	// Nivel 3: Aporte o curiosidad profunda
	if longitud > 50 {
		return 3
	}
	// Nivel 2: Interacción estándar
	if longitud > 10 {
		return 2
	}
	// Nivel 1: Ruido
	return 1
}

// 2. ACTUALIZAMOS EL MÉTODO (Ahora pertenece a FriendIA)
func (f *FriendIA) EvaluarInteraccion(usuarioID string, input string) int {
	puntuacion := AnalizarEnergia(input)
	
	// Registro para auditoría del Dueño [cite: 2026-02-10]
	if puntuacion == 3 {
		log.Printf("💎 Energía Nivel 3: Usuario %s aportando valor.", usuarioID)
		
		// 3. RECOMPENSA SOCIAL: Usamos la conexión al CEO
		if f.CEO != nil {
			f.CEO.ProcesarRecompensaSocial(usuarioID)
			log.Println("📢 AI Friend: Notificando a CEO para bono de equidad del 15%.")
		}
	} else if puntuacion == 1 {
		log.Printf("🛡️ AI Friend: Energía baja detectada de %s. Bloqueando evolución.", usuarioID)
	}
    
	return puntuacion
}