package ai_friend

import (
	"sync"
)

// FriendIA es la entidad que acompaña al usuario
type FriendIA struct {
	sync.RWMutex
	ID   string
}

// EvaluarInteraccion analiza el impacto positivo del mensaje
func (f *FriendIA) EvaluarInteraccion(userID string, texto string) float64 {
	f.RLock()
	defer f.RUnlock()
	
	// Lógica básica: si el mensaje es largo, da más puntos
	if len(texto) > 10 {
		return 1.5
	}
	return 1.0
}