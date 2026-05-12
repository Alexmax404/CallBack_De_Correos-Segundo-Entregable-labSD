package fachada

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"clientestreaming/capaFachadaServices/models"
	pb "clientestreaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// DireccionServidor es la IP y puerto del ServidorDeStreaming
const DireccionServidor = "localhost:50051"

// TamanoBuffer tamaño del buffer para recibir fragmentos de audio
const TamanoBuffer = 64 * 1024

// FachadaCliente gestiona la conexión gRPC y las operaciones del cliente.
type FachadaCliente struct {
	conexion *grpc.ClientConn
	stub     pb.AudioStreamingServiceClient
	usuario  *models.Usuario
}

// NuevaFachadaCliente crea la conexión gRPC con el servidor de streaming.
func NuevaFachadaCliente() (*FachadaCliente, error) {
	fmt.Printf("🔌 Conectando al servidor de streaming en %s...\n", DireccionServidor)

	conexion, err := grpc.Dial(
		DireccionServidor,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con el servidor: %v", err)
	}

	fmt.Println("✅ Conexión gRPC establecida")
	return &FachadaCliente{
		conexion: conexion,
		stub:     pb.NewAudioStreamingServiceClient(conexion),
	}, nil
}

// Cerrar cierra la conexión gRPC.
func (f *FachadaCliente) Cerrar() {
	if f.conexion != nil {
		f.conexion.Close()
	}
}

// Login valida las credenciales del usuario.
// Por ahora valida que el nickname no esté vacío y la contraseña tenga al menos 4 caracteres.
func (f *FachadaCliente) Login(nickname string, contrasena string) error {
	if nickname == "" {
		return fmt.Errorf("el nickname no puede estar vacío")
	}
	if len(contrasena) < 4 {
		return fmt.Errorf("la contraseña debe tener al menos 4 caracteres")
	}

	f.usuario = &models.Usuario{
		Nickname:   nickname,
		Contrasena: contrasena,
	}

	fmt.Printf("✅ Sesión iniciada como: %s\n", nickname)
	return nil
}

// ObtenerUsuario retorna el usuario autenticado.
func (f *FachadaCliente) ObtenerUsuario() *models.Usuario {
	return f.usuario
}

// ListarAudios solicita al servidor la lista de audios disponibles.
func (f *FachadaCliente) ListarAudios() ([]models.InfoAudio, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	respuesta, err := f.stub.ListarAudios(ctx, &pb.SolicitudListarAudios{
		Solicitante: f.usuario.Nickname,
	})
	if err != nil {
		return nil, fmt.Errorf("error al listar audios: %v", err)
	}

	var audios []models.InfoAudio
	for _, a := range respuesta.Audios {
		audios = append(audios, models.InfoAudio{
			NombreArchivo: a.NombreArchivo,
			Titulo:        a.Titulo,
			Artista:       a.Artista,
			Genero:        a.Genero,
		})
	}
	return audios, nil
}

// ObtenerMetadatos solicita los metadatos de un audio específico.
func (f *FachadaCliente) ObtenerMetadatos(nombreArchivo string) (models.MetadatosAudio, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := f.stub.ObtenerMetadatosAudio(ctx, &pb.SolicitudMetadatos{
		NombreArchivo: nombreArchivo,
	})
	if err != nil {
		return models.MetadatosAudio{}, fmt.Errorf("error obteniendo metadatos: %v", err)
	}

	return models.MetadatosAudio{
		NombreArchivo: resp.NombreArchivo,
		Titulo:        resp.Titulo,
		Artista:       resp.Artista,
		Genero:        resp.Genero,
		TamanoBytes:   resp.TamanoBytes,
	}, nil
}

// ReproducirAudio recibe el audio del servidor via streaming y lo guarda localmente.
func (f *FachadaCliente) ReproducirAudio(nombreArchivo string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Iniciar stream gRPC
	stream, err := f.stub.ReproducirAudio(ctx, &pb.SolicitudReproduccion{
		NombreArchivo:   nombreArchivo,
		NicknameCliente: f.usuario.Nickname,
	})
	if err != nil {
		return fmt.Errorf("error iniciando reproducción: %v", err)
	}

	// Crear archivo temporal donde guardar el audio recibido
	archivoLocal := "temp_" + nombreArchivo
	archivo, err := os.Create(archivoLocal)
	if err != nil {
		return fmt.Errorf("error creando archivo temporal: %v", err)
	}
	defer archivo.Close()

	fmt.Printf("📥 Recibiendo audio '%s'...\n", nombreArchivo)
	totalChunks := 0
	totalBytes := 0

	// Recibir cada fragmento del stream
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error recibiendo fragmento: %v", err)
		}

		// Escribir fragmento al archivo local
		_, err = archivo.Write(fragmento.Datos)
		if err != nil {
			return fmt.Errorf("error escribiendo fragmento: %v", err)
		}

		totalChunks++
		totalBytes += len(fragmento.Datos)

		if totalChunks%10 == 0 {
			fmt.Printf("   📦 Fragmentos recibidos: %d | Bytes: %d\n", totalChunks, totalBytes)
		}

		if fragmento.EsUltimo {
			break
		}
	}

	fmt.Printf("✅ Audio recibido completamente!\n")
	fmt.Printf("   Total fragmentos : %d\n", totalChunks)
	fmt.Printf("   Total bytes       : %d\n", totalBytes)
	fmt.Printf("   Archivo guardado  : %s\n", archivoLocal)
	fmt.Println("🎵 [Simulando reproducción del audio...]")

	return nil
}

// ObtenerGeneros retorna los géneros únicos de los audios disponibles.
func (f *FachadaCliente) ObtenerGeneros() ([]string, error) {
	audios, err := f.ListarAudios()
	if err != nil {
		return nil, err
	}

	// Usar mapa para eliminar duplicados
	generosMap := make(map[string]bool)
	for _, a := range audios {
		generosMap[a.Genero] = true
	}

	var generos []string
	for g := range generosMap {
		generos = append(generos, g)
	}
	return generos, nil
}
