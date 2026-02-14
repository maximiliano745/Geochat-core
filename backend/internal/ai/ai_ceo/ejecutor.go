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

	// El mensaje de commit refleja tu autorización soberana
	msg := fmt.Sprintf("🧬 IA 5 Evolución: %s\n\nAutorizado por Jefe", nombreModulo)
	
	hash, err := w.Commit(msg, &git.CommitOptions{}) 
	if err != nil { return err }
	
	fmt.Printf("✅ Commit creado exitosamente: %s\n", hash)

	// Push al repositorio remoto (GitHub)
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: c.Username,
			Password: c.Token,
		},
	})

	return err
}