package capacontroladores

import (
	"context"
	"fmt"

	"servidorstreaming/capaFachadaServices/fachada"
	pb "servidorstreaming/proto"
)

// ControladorStreaming implementa la interfaz gRPC AudioStreamingServiceServer.
type ControladorStreaming struct {
	pb.UnimplementedAudioStreamingServiceServer
	fachada *fachada.FachadaStreaming
}

// NuevoControladorStreaming crea una instancia del controlador gRPC.
func NuevoControladorStreaming(f *fachada.FachadaStreaming) *ControladorStreaming {
	return &ControladorStreaming{fachada: f}
}

// ListarAudios implementa el RPC que retorna todos los audios disponibles.
func (c *ControladorStreaming) ListarAudios(ctx context.Context, req *pb.SolicitudListarAudios) (*pb.RespuestaListarAudios, error) {
	fmt.Printf("\n📥 [gRPC] ListarAudios - Solicitante: %s\n", req.Solicitante)

	audios, err := c.fachada.ListarAudios(req.Solicitante)
	if err != nil {
		return nil, fmt.Errorf("error listando audios: %v", err)
	}

	// Convertir modelos internos a mensajes protobuf
	var pbAudios []*pb.InfoAudio
	for _, a := range audios {
		pbAudios = append(pbAudios, &pb.InfoAudio{
			NombreArchivo: a.NombreArchivo,
			Titulo:        a.Titulo,
			Artista:       a.Artista,
			Genero:        a.Genero,
		})
	}

	fmt.Printf("   ✅ Retornando %d audios\n", len(pbAudios))
	return &pb.RespuestaListarAudios{Audios: pbAudios}, nil
}

// ObtenerMetadatosAudio implementa el RPC que retorna metadatos de un audio.
func (c *ControladorStreaming) ObtenerMetadatosAudio(ctx context.Context, req *pb.SolicitudMetadatos) (*pb.RespuestaMetadatos, error) {
	fmt.Printf("\n📥 [gRPC] ObtenerMetadatos - Archivo: %s\n", req.NombreArchivo)

	meta, err := c.fachada.ObtenerMetadatos(req.NombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo metadatos: %v", err)
	}

	fmt.Printf("   ✅ Metadatos retornados para: %s\n", meta.Titulo)
	return &pb.RespuestaMetadatos{
		NombreArchivo: meta.NombreArchivo,
		Titulo:        meta.Titulo,
		Artista:       meta.Artista,
		Genero:        meta.Genero,
		TamanoBytes:   meta.TamanoBytes,
	}, nil
}

// ReproducirAudio implementa el RPC de streaming - envía el audio en fragmentos.
func (c *ControladorStreaming) ReproducirAudio(req *pb.SolicitudReproduccion, stream pb.AudioStreamingService_ReproducirAudioServer) error {
	fmt.Printf("\n📥 [gRPC] ReproducirAudio - Cliente: %s | Archivo: %s\n",
		req.NicknameCliente, req.NombreArchivo)

	// Obtener canal de fragmentos desde la fachada
	fragmentos, err := c.fachada.ObtenerFragmentosAudio(req.NombreArchivo, req.NicknameCliente)
	if err != nil {
		return fmt.Errorf("error iniciando streaming: %v", err)
	}

	// Enviar cada fragmento al cliente via stream gRPC
	for fragmento := range fragmentos {
		err := stream.Send(&pb.FragmentoAudio{
			Datos:       fragmento.Datos,
			NumeroChunk: int32(fragmento.NumeroChunk),
			EsUltimo:    fragmento.EsUltimo,
		})
		if err != nil {
			return fmt.Errorf("error enviando fragmento %d: %v", fragmento.NumeroChunk, err)
		}

		if fragmento.NumeroChunk%10 == 0 {
			fmt.Printf("   📦 Enviando chunk #%d...\n", fragmento.NumeroChunk)
		}
	}

	fmt.Printf("   ✅ Streaming completado para: %s\n", req.NombreArchivo)
	return nil
}
