package api

import (
	"geochat/internal/vault"
	"strings"
)

// EnsamblarADN toma un esquema por ID y rellena sus variables
func EnsamblarADN(esquemaID string, valores map[string]string) string {
	// 1. Buscamos el molde en la biblioteca que acabamos de arreglar
	esquema, existe := vault.BibliotecaADN[esquemaID]
	if !existe {
		return "// ❌ Error: Esquema no encontrado en el Vault"
	}

	codigoFinal := esquema.Template

	// 2. Rellenamos los campos (ID, INVERSION, etc.)
	for _, campo := range esquema.Campos {
		placeholder := "{{" + campo + "}}"
		valor := valores[campo]
		codigoFinal = strings.ReplaceAll(codigoFinal, placeholder, valor)
	}

	// 3. Inyectamos la Regla de Oro (15%) por defecto
	codigoFinal = strings.ReplaceAll(codigoFinal, "{{REGLA_ORO}}", "0.15")

	return codigoFinal
}