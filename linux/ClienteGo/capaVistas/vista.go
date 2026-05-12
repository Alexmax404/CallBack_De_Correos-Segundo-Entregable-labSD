package capavistas

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Vista gestiona la interacción con el usuario por consola.
type Vista struct {
	lector *bufio.Reader
}

// NuevaVista crea una instancia de la vista.
func NuevaVista() *Vista {
	return &Vista{lector: bufio.NewReader(os.Stdin)}
}

// LeerTexto lee una línea de texto del usuario.
func (v *Vista) LeerTexto(mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := v.lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

// LeerEntero lee un número entero del usuario.
func (v *Vista) LeerEntero(mensaje string) int {
	for {
		texto := v.LeerTexto(mensaje)
		numero, err := strconv.Atoi(texto)
		if err == nil {
			return numero
		}
		fmt.Println("⚠️  Por favor ingresa un número válido.")
	}
}

// MostrarBienvenida muestra el encabezado del sistema.
func (v *Vista) MostrarBienvenida() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║     🎵 CLIENTE DE STREAMING DE AUDIO     ║")
	fmt.Println("║        Universidad del Cauca - FIET      ║")
	fmt.Println("║         Sistemas Distribuidos            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
}

// MostrarMenuLogin muestra el menú de login.
func (v *Vista) MostrarMenuLogin() int {
	fmt.Println("──────────────────────────────────────────")
	fmt.Println("  MENÚ PRINCIPAL")
	fmt.Println("──────────────────────────────────────────")
	fmt.Println("  1. Iniciar sesión")
	fmt.Println("  0. Salir")
	fmt.Println("──────────────────────────────────────────")
	return v.LeerEntero("Selecciona una opción: ")
}

// MostrarMenuPrincipal muestra el menú principal del cliente autenticado.
func (v *Vista) MostrarMenuPrincipal(nickname string) int {
	fmt.Println()
	fmt.Println("──────────────────────────────────────────")
	fmt.Printf("  👤 Usuario: %s\n", nickname)
	fmt.Println("──────────────────────────────────────────")
	fmt.Println("  1. Ver tipos de audio (géneros)")
	fmt.Println("  2. Ver todos los audios")
	fmt.Println("  3. Consultar metadatos de un audio")
	fmt.Println("  4. Reproducir un audio")
	fmt.Println("  0. Cerrar sesión")
	fmt.Println("──────────────────────────────────────────")
	return v.LeerEntero("Selecciona una opción: ")
}

// MostrarGeneros muestra la lista de géneros disponibles.
func (v *Vista) MostrarGeneros(generos []string) {
	fmt.Println()
	fmt.Println("🎸 GÉNEROS DISPONIBLES:")
	fmt.Println("──────────────────────────────────────────")
	if len(generos) == 0 {
		fmt.Println("  No hay géneros disponibles.")
	} else {
		for i, g := range generos {
			fmt.Printf("  %d. %s\n", i+1, g)
		}
	}
	fmt.Println("──────────────────────────────────────────")
}

// MostrarListaAudios muestra todos los audios con su índice.
func (v *Vista) MostrarListaAudios(audios []struct {
	NombreArchivo string
	Titulo        string
	Artista       string
	Genero        string
}) {
	fmt.Println()
	fmt.Println("🎵 AUDIOS DISPONIBLES:")
	fmt.Println("──────────────────────────────────────────")
	if len(audios) == 0 {
		fmt.Println("  No hay audios disponibles.")
	} else {
		for i, a := range audios {
			fmt.Printf("  %d. %s\n", i+1, a.Titulo)
			fmt.Printf("     🎤 Artista : %s\n", a.Artista)
			fmt.Printf("     🎸 Género  : %s\n", a.Genero)
			fmt.Printf("     📄 Archivo : %s\n", a.NombreArchivo)
			fmt.Println()
		}
	}
	fmt.Println("──────────────────────────────────────────")
}

// MostrarMetadatos muestra los metadatos completos de un audio.
func (v *Vista) MostrarMetadatos(nombreArchivo, titulo, artista, genero string, tamano int64) {
	fmt.Println()
	fmt.Println("🔍 METADATOS DEL AUDIO:")
	fmt.Println("──────────────────────────────────────────")
	fmt.Printf("  📄 Archivo : %s\n", nombreArchivo)
	fmt.Printf("  🎵 Título  : %s\n", titulo)
	fmt.Printf("  🎤 Artista : %s\n", artista)
	fmt.Printf("  🎸 Género  : %s\n", genero)
	fmt.Printf("  💾 Tamaño  : %d bytes (%.2f MB)\n", tamano, float64(tamano)/1024/1024)
	fmt.Println("──────────────────────────────────────────")
}

// MostrarError muestra un mensaje de error.
func (v *Vista) MostrarError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
}

// MostrarMensaje muestra un mensaje informativo.
func (v *Vista) MostrarMensaje(mensaje string) {
	fmt.Println("ℹ️ ", mensaje)
}

// PedirSeleccionAudio pide al usuario que seleccione un número de la lista.
func (v *Vista) PedirSeleccionAudio(max int) int {
	for {
		opcion := v.LeerEntero(fmt.Sprintf("Selecciona el número del audio (1-%d): ", max))
		if opcion >= 1 && opcion <= max {
			return opcion
		}
		fmt.Printf("⚠️  Opción inválida. Ingresa un número entre 1 y %d.\n", max)
	}
}
