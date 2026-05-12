package main

import (
	"fmt"
	"net"

	capacontroladores "servidorstreaming/capaControladores"
	"servidorstreaming/capaFachadaServices/fachada"
	pb "servidorstreaming/proto"

	"google.golang.org/grpc"
)

// Puerto donde escucha el servidor gRPC de streaming
const puerto = ":50051"

func main() {
	fmt.Println("========================================")
	fmt.Println("  🎵 Servidor de Streaming de Audios")
	fmt.Println("        Protocolo: gRPC")
	fmt.Println("========================================")

	// Inicializar fachada
	fachadaStreaming := fachada.NuevaFachadaStreaming()

	// Inicializar controlador gRPC
	controlador := capacontroladores.NuevoControladorStreaming(fachadaStreaming)

	// Crear listener TCP
	listener, err := net.Listen("tcp", puerto)
	if err != nil {
		fmt.Printf("❌ Error al abrir puerto %s: %v\n", puerto, err)
		return
	}

	// Crear servidor gRPC y registrar el servicio
	servidorGRPC := grpc.NewServer()
	pb.RegisterAudioStreamingServiceServer(servidorGRPC, controlador)

	fmt.Printf("📡 Servicios disponibles:\n")
	fmt.Printf("   - ListarAudios\n")
	fmt.Printf("   - ObtenerMetadatosAudio\n")
	fmt.Printf("   - ReproducirAudio (streaming)\n")
	fmt.Println("========================================")
	fmt.Printf("✅ Servidor gRPC escuchando en %s...\n", puerto)

	// Iniciar servidor
	if err := servidorGRPC.Serve(listener); err != nil {
		fmt.Printf("❌ Error en servidor gRPC: %v\n", err)
	}
}
