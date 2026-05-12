package capaaccedoadatos

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type RepositorioCanciones struct {
	mu sync.Mutex
}

var (
	instancia *RepositorioCanciones
	once      sync.Once
)

func GetRepositorioCanciones() *RepositorioCanciones {
	once.Do(func() {
		instancia = &RepositorioCanciones{}
	})
	return instancia
}

func (r *RepositorioCanciones) GuardarCancion(titulo string, genero string, artista string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := os.MkdirAll("audios", os.ModePerm); err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s_%s_%s.mp3", titulo, genero, artista)
	filePath := filepath.Join("audios", fileName)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}

	return nil
}
