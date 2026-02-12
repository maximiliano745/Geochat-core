package ai_ceo

import (
	"fmt"
	"log"
	"time"
	"geochat/internal/vault" // Conexión con tu Bóveda Soberana E2E
)

// Logro representa un hito en la historia de GeoChat [cite: 2026-02-11]
type Logro struct {
	Fecha       time.Time
	Categoria   string // "Social", "Software", "Financiero"
	Descripcion string
	FirmaIA     string
}

// RegistrarEnLibroDeLogros: La IA 5 escribe la historia que solo tú puedes leer.
func (c *CEO) RegistrarEnLibroDeLogros(logro Logro) {
	// 1. Formatea el texto de forma humana (Entendible para ti) [cite: 2026-01-26]
	entrada := fmt.Sprintf("[%s] [%s] EVOLUCIÓN: %s (Firmado: %s)", 
		logro.Fecha.Format("02-01-2006 15:04"), 
		logro.Categoria, 
		logro.Descripcion, 
		logro.FirmaIA,
	)
	
	// 2. Lo guarda en el Vault Personal (E2E Encrypted)
	// Una vez guardado, ni la IA ni nadie puede verlo sin tu llave privada [cite: 2026-01-12]
	err := vault.GuardarDocumentoLegal("Libro_de_Logros_2026.pdf", entrada)
	
	if err != nil {
		log.Printf("⚠️ IA 5: Error al documentar en Vault: %v", err)
		return
	}
	
	fmt.Println("📜 IA 5: Logro documentado y protegido con cifrado soberano en tu Vault.")
}

// DocumentarEnVault es un acceso rápido para que otras IAs (como la 2) reporten hitos.
func (c *CEO) DocumentarEnVault(detalle string) {
	nuevoLogro := Logro{
		Fecha:       time.Now(),
		Categoria:   "SOFTWARE",
		Descripcion: detalle,
		FirmaIA:     "IA-5-CEO-AUTONOMO",
	}
	c.RegistrarEnLibroDeLogros(nuevoLogro)
}