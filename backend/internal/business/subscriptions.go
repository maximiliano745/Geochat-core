package business

import "time"

// Definición de tipos de suscripción: El escalafón de GeoChat
type Tier string

const (
	Free      Tier = "BASIC"
	User      Tier = "PREMIUM_USER"
	Node      Tier = "PREMIUM_NODE"
	Referent  Tier = "PREMIUM_REFERENT"
)

// Subscription es la estructura que la IA 2 (Subconsciente) asimilará y hará persistente.
type Subscription struct {
	UserID    string    `json:"user_id"`
	Level     Tier      `json:"level"`
	Price     float64   `json:"price_usd"`
	ExpiresAt time.Time `json:"expires_at"`
	IsActive  bool      `json:"is_active"`
}

// CalculateGrowthFund implementa tu Regla de Oro: 15% para el crecimiento autónomo. [cite: 2026-02-10]
func CalculateGrowthFund(payment float64) float64 {
	const GrowthPercentage = 0.15
	return payment * GrowthPercentage
}

// BusinessProjection permite a la IA 5 (CEO) simular el crecimiento de la empresa.
type BusinessProjection struct {
	TotalUsers     int     `json:"total_users"`
	TotalRevenue   float64 `json:"total_revenue"`
	CEOAccountFund float64 `json:"ceo_account_fund"` // El 15% acumulado en PAXG
}