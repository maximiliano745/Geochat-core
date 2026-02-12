package main

import (
	"sync"
	"time"
)

type CEO struct {
	sync.RWMutex
	FondoGas     float64
	TokensGratis int
	Propuestas   []Propuesta
	Stats        Estadisticas
	Lider        PerfilLider
}

type Propuesta struct {
	ID                string
	Modulo            string
	Monto             float64
	ImpactoFinanciero float64 // Añadido para reportes.go
	Status            string
	RequiereFirma     bool    // Añadido para la lógica de Vue.ts
	Tipo              string
	FirmaDigital      string
}

type EvolucionSoftware struct { // Añadido para que getEvolucionesPendientes funcione
	Modulo  string
	Origen  string
	Impacto string
}

type Estadisticas struct {
	BuenaOndaCount int
	MalaOndaCount  int
	Plasticidad    float64
}

type PerfilLider struct {
	AlineacionEtica float64
	HistorialFirma  []string
	UltimaActividad time.Time
}