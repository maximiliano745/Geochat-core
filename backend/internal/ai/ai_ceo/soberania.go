package ai_ceo

import "time"

const (
	SoberaniaKeys = true
	Transparencia = "E2E_VAULT"
)

func (c *CEO) ValidarSoberania() bool {
	return SoberaniaKeys
}

func (p *PerfilLider) RegistrarFirma(id string) {
	p.HistorialFirma = append(p.HistorialFirma, id)
	p.UltimaActividad = time.Now()
}