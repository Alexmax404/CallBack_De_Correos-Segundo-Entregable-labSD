package fachada

import (
	"fmt"

	capaaccesodatos "servidorstreaming/capaAccesoDatos"
	"servidorstreaming/capaFachadaServices/models"
)

// TamanoChunk define el tamaño de cada fragmento de audio en bytes (64KB)
const TamanoChunk = 64 * 1024

// FachadaStreaming coordina el acceso a los audios y el envío de fragmentos.
type FachadaStreaming struct {
	repo *capaaccesodatos.RepositorioAudios
}

// NuevaFachadaStreaming crea e inicializa la fachada de streaming.
func NuevaFachadaStreaming() *FachadaStreaming {
	fmt.Println("🎵 Inicializando fachada de streaming...")
	return &FachadaStreaming{
		repo: capaaccesodatos.GetRepositorioAudios(),
	}
}

// ListarAudios retorna la lista de audios disponibles en el servidor.
func (f *FachadaStreaming) ListarAudios(solicitante string) ([]models.InfoAudio, error) {
	fmt.Printf("📋 [ListarAudios] Solicitado por: %s\n", solicitante)
	audios, err := f.repo.ListarAudios()
	if err != nil {
		return nil, err
	}
	fmt.Printf("   → %d audios disponibles\n", len(audios))
	return audios, nil
}

// ObtenerMetadatos retorna los metadatos completos de un audio.
func (f *FachadaStreaming) ObtenerMetadatos(nombreArchivo string) (models.MetadatosAudio, error) {
	fmt.Printf("🔍 [ObtenerMetadatos] Archivo: %s\n", nombreArchivo)
	return f.repo.ObtenerMetadatos(nombreArchivo)
}

// ObtenerFragmentosAudio retorna un canal con los fragmentos del audio solicitado.
func (f *FachadaStreaming) ObtenerFragmentosAudio(nombreArchivo string, nicknameCliente string) (<-chan models.FragmentoAudio, error) {
	fmt.Printf("▶️  [ReproducirAudio] Cliente: %s | Archivo: %s\n", nicknameCliente, nombreArchivo)
	return f.repo.LeerFragmentosAudio(nombreArchivo, TamanoChunk)
}
