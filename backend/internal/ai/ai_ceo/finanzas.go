package ai_ceo

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ObtenerReporteFinanciero conecta la IA 5 con la realidad económica de la red.
// [cite: 2026-02-10] "La IA mira los números de toda la red"
func (c *CEO) ObtenerReporteFinanciero() string {
	c.RLock()
	defer c.RUnlock()

	// 1. Conexión con la Capa 1 (Blockchain)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rpcURL := os.Getenv("POLYGON_RPC_URL")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "⚠️ *IA CEO:* Error de conexión RPC. No pude leer Polygon."
	}
	defer client.Close()

	// 2. Consulta de Gas Real (MATIC)
	addressStr := os.Getenv("CEO_WALLET_ADDRESS")
	address := common.HexToAddress(addressStr)
	balance, err := client.BalanceAt(ctx, address, nil)
	
	maticDisplay := "0.0000"
	if err == nil {
		fbalance := new(big.Float).SetInt(balance)
		maticValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
		maticDisplay = fmt.Sprintf("%.4f", maticValue)
	}

	// 3. Consolidación de datos Soberanos [cite: 2026-02-10]
	// Mezclamos datos On-Chain (Blockchain) con datos Off-Chain (tu memoria local)
	reporte := fmt.Sprintf("📊 *REPORTE ESTRATÉGICO GEOCHAT*\n\n"+
		"💰 *Fondo Crecimiento (15%%):* %.2f PAXG\n"+
		"⛽ *Combustible (MATIC):* %s\n"+
		"⚡ *Gas Local:* %.2f tokens\n"+
		"🧬 *Plasticidad ADN:* %.2f\n"+
		"🏗️ *Sandbox:* %d propuestas\n\n"+
		"*Estado:* Operativo bajo 'Firma es la Ley'.", 
		c.FondoGas, maticDisplay, c.TokensGratis, c.Stats.Plasticidad, len(c.Propuestas))

	return reporte
}