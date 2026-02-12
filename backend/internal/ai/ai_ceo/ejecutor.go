package ai_ceo

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type CerebroEjecucion struct {
	RepoPath string
	RepoURL  string
	Token    string
	Username string
}

func (c *CerebroEjecucion) PublicarEvolucion(nombreModulo string) error {
	r, err := git.PlainOpen(c.RepoPath)
	if err != nil { return err }

	w, err := r.Worktree()
	if err != nil { return err }

	_, err = w.Add(nombreModulo)
	if err != nil { return err }

	// CORRECCIÓN AQUÍ:
	// Como 'err' ya existe, usamos '=' para el Commit. 
	// Quitamos el Status innecesario que causaba el conflicto de variables.
	
	msg := fmt.Sprintf("🧬 IA 5 Evolución: %s\n\nAutorizado por Jefe vía WhatsApp", nombreModulo)
	
	commit, err := w.Commit(msg, &git.CommitOptions{}) // 'commit' es nueva, así que := funciona
	if err != nil { return err }
	
	fmt.Printf("Commit creado: %s\n", commit)

	// Aquí 'err' ya existe, así que usamos '='
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: c.Username,
			Password: c.Token,
		},
	})

	return err
}