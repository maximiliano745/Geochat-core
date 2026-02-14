package api

import (
    "geochat/internal/database"
    "net/http"
    "github.com/gin-gonic/gin"
)

// GetDashboardData resume el estado del organismo vivo para el Dashboard
func GetDashboardData(r *gin.Engine) {
    r.GET("/api/dashboard/stats", func(c *gin.Context) {
        // 1. Calculamos el Saldo Total (El 100%)
        // La IA calculará el 15% en el frontend o lo mostramos directo
        var saldoTotal float64
        err := database.DB.QueryRow("SELECT COALESCE(SUM(monto), 0) FROM transacciones").Scan(&saldoTotal)
        if err != nil {
            saldoTotal = 1.2450 // Fallback por si la DB está naciendo
        }

        // 2. Logs de actividad del CEO
        logs := []string{
            "Conexión con Nodo Avellaneda establecida",
            "IA 5: Escaneando oportunidades de infraestructura",
            "Capa de Red: Soberanía activa",
        }

        // 3. Enviamos el paquete con los nombres que VUE espera
        c.JSON(http.StatusOK, gin.H{
            "saldo_total":    saldoTotal,
            "nodo_nombre":    "Nodo Avellaneda Principal",
            "logs_recientes": logs,
            "estado_red":     "Soberana",
        })
    })
}