package vertex_connector

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"cloud.google.com/go/vertexai/genai" 
)

// GenerarCodigoReal invoca al "Ojo de Vertex" y guarda el resultado en el Laboratorio
func GenerarCodigoReal(nombreArchivo string, prompt string) (string, error) {
	ctx := context.Background()
	project := os.Getenv("GCP_PROJECT_ID") 
	location := "us-central1"
	modelName := "gemini-1.5-pro" 

	client, err := genai.NewClient(ctx, project, location)
	if err != nil {
		return "", fmt.Errorf("error conectando a Vertex: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	
	// Instrucción estricta para asegurar código limpio
	systemPrompt := "Actúa como el CEO de GeoChat. Escribe solo código Go funcional para: " + prompt + 
	                ". No incluyas explicaciones ni markdown, solo el código puro."

	resp, err := model.GenerateContent(ctx, genai.Text(systemPrompt))
	if err != nil {
		return "", err
	}

	codigo := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	// Guardar en el Laboratorio para que el Dueño lo vea en el Dashboard [cite: 2026-02-10]
	pathLab := filepath.Join("./internal/lab", nombreArchivo+".go")
	_ = os.MkdirAll("./internal/lab", 0755)
	err = os.WriteFile(pathLab, []byte(codigo), 0644)
	if err != nil {
		return "", fmt.Errorf("error guardando en laboratorio: %v", err)
	}

	return codigo, nil
}