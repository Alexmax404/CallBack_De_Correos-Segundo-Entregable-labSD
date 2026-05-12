package capaaccesodatos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"servidorstreaming/capaFachadaServices/models"
)

// RepositorioAudios gestiona el acceso a los archivos de audio almacenados en disco.
type RepositorioAudios struct {
	mu          sync.Mutex
	carpetaBase string
}

var (
	instancia *RepositorioAudios
	once      sync.Once
)

// GetRepositorioAudios aplica patrón Singleton.
func GetRepositorioAudios() *RepositorioAudios {
	once.Do(func() {
		instancia = &RepositorioAudios{
			carpetaBase: "audios",
		}
	})
	return instancia
}

// ListarAudios lee el directorio de audios y retorna la info de cada archivo.
func (r *RepositorioAudios) ListarAudios() ([]models.InfoAudio, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	entries, err := os.ReadDir(r.carpetaBase)
	if err != nil {
		return nil, fmt.Errorf("error leyendo carpeta de audios: %v", err)
	}

	var audios []models.InfoAudio
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".mp3") {
			continue
		}
		// El nombre sigue el formato: titulo_genero_artista.mp3
		info := parsearNombreArchivo(entry.Name())
		audios = append(audios, info)
	}
	return audios, nil
}

// ObtenerMetadatos retorna los metadatos y tamaño de un audio específico.
func (r *RepositorioAudios) ObtenerMetadatos(nombreArchivo string) (models.MetadatosAudio, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	filePath := filepath.Join(r.carpetaBase, nombreArchivo)
	info, err := os.Stat(filePath)
	if err != nil {
		return models.MetadatosAudio{}, fmt.Errorf("audio no encontrado: %v", err)
	}

	meta := parsearNombreArchivo(nombreArchivo)
	return models.MetadatosAudio{
		NombreArchivo: nombreArchivo,
		Titulo:        meta.Titulo,
		Artista:       meta.Artista,
		Genero:        meta.Genero,
		TamanoBytes:   info.Size(),
	}, nil
}

// LeerFragmentosAudio lee el archivo y lo divide en fragmentos de tamaño fijo.
// Retorna un canal por donde se envían los fragmentos.
func (r *RepositorioAudios) LeerFragmentosAudio(nombreArchivo string, tamanoChunk int) (<-chan models.FragmentoAudio, error) {
	filePath := filepath.Join(r.carpetaBase, nombreArchivo)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo audio '%s': %v", nombreArchivo, err)
	}

	canal := make(chan models.FragmentoAudio)

	go func() {
		defer close(canal)
		totalBytes := len(data)
		numeroChunk := 0

		for inicio := 0; inicio < totalBytes; inicio += tamanoChunk {
			fin := inicio + tamanoChunk
			if fin > totalBytes {
				fin = totalBytes
			}

			fragmento := models.FragmentoAudio{
				Datos:       data[inicio:fin],
				NumeroChunk: numeroChunk,
				EsUltimo:    fin >= totalBytes,
			}
			canal <- fragmento
			numeroChunk++
		}
	}()

	return canal, nil
}

// parsearNombreArchivo extrae titulo, genero y artista del nombre del archivo.
// Formato esperado: titulo_genero_artista.mp3
func parsearNombreArchivo(nombreArchivo string) models.InfoAudio {
	sinExtension := strings.TrimSuffix(nombreArchivo, ".mp3")
	partes := strings.Split(sinExtension, "_")

	titulo, genero, artista := sinExtension, "Desconocido", "Desconocido"
	if len(partes) >= 3 {
		titulo = partes[0]
		genero = partes[1]
		artista = strings.Join(partes[2:], "_")
	}

	return models.InfoAudio{
		NombreArchivo: nombreArchivo,
		Titulo:        titulo,
		Artista:       artista,
		Genero:        genero,
	}
}
