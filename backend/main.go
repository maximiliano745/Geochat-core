package main

import (
	"geochat/internal/ai/ai_ceo"
	"geochat/internal/ai/ai_friend"
	"geochat/internal/ai/ai_infra"
	"geochat/internal/ai/ai_security"
	"geochat/internal/finance"
	

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicialización obligatoria
	ai_security.InitializeVault()
	ceo := ai_ceo.NewCEO(finance.GetMultisigWallet())

	go ceo.StartNegotiationLoop()
	go ai_infra.StartMeshOptimizing()
	go ai_friend.ListenToUsers()


	wallet := finance.GetMultisigWallet()
	ceo = ai_ceo.NewCEO(wallet)
	go ceo.EjecutarCicloDesarrollo()
	
	r := gin.Default()

	r.GET("/api/proposals", func(c *gin.Context) {
		p := ceo.DisenarModulo("Geocerca Sector-7")
		c.JSON(200, gin.H{
			"propuestas":   []ai_ceo.Propuesta{p},
			"tokensGratis": ceo.TokensGratis,
			"fondoGas":     ceo.FondoGas,
		})
	})

	r.POST("/api/autorizar", func(c *gin.Context) {
		var body struct { Firma string; ID string }
		if err := c.BindJSON(&body); err == nil && body.Firma == "MI_FIRMA_ES_LA_LEY" {
			for _, p := range ceo.Propuestas {
				if p.ID == body.ID {
					_ = ceo.InyectarCodigoFuncional(p)
					c.JSON(200, gin.H{"status": "Inyectado exitosamente"})
					return
				}
			}
		}
		c.JSON(403, gin.H{"error": "Firma inválida o propuesta no encontrada"})
	})

	r.Run(":3000")
}