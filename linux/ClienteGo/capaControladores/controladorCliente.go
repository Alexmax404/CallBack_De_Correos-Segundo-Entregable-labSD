package capacontroladores

import (
	"fmt"

	"clientestreaming/capaFachadaServices/fachada"
	capavistas "clientestreaming/capaVistas"
)

// ControladorCliente orquesta las operaciones entre la vista y la fachada.
type ControladorCliente struct {
	fachada *fachada.FachadaCliente
	vista   *capavistas.Vista
}

// NuevoControladorCliente crea una instancia del controlador.
func NuevoControladorCliente(f *fachada.FachadaCliente, v *capavistas.Vista) *ControladorCliente {
	return &ControladorCliente{fachada: f, vista: v}
}

// EjecutarLogin gestiona el proceso de autenticación del usuario.
func (c *ControladorCliente) EjecutarLogin() bool {
	fmt.Println("\n🔐 INICIO DE SESIÓN")
	fmt.Println("──────────────────────────────────────────")

	nickname := c.vista.LeerTexto("  Nickname   : ")
	contrasena := c.vista.LeerTexto("  Contraseña : ")

	err := c.fachada.Login(nickname, contrasena)
	if err != nil {
		c.vista.MostrarError(err)
		return false
	}
	return true
}

// EjecutarVerGeneros muestra los géneros de audio disponibles.
func (c *ControladorCliente) EjecutarVerGeneros() {
	generos, err := c.fachada.ObtenerGeneros()
	if err != nil {
		c.vista.MostrarError(err)
		return
	}
	c.vista.MostrarGeneros(generos)
}

// EjecutarVerAudios muestra todos los audios disponibles.
func (c *ControladorCliente) EjecutarVerAudios() {
	audios, err := c.fachada.ListarAudios()
	if err != nil {
		c.vista.MostrarError(err)
		return
	}

	// Convertir al tipo anónimo que espera la vista
	var audiosVista []struct {
		NombreArchivo string
		Titulo        string
		Artista       string
		Genero        string
	}
	for _, a := range audios {
		audiosVista = append(audiosVista, struct {
			NombreArchivo string
			Titulo        string
			Artista       string
			Genero        string
		}{
			NombreArchivo: a.NombreArchivo,
			Titulo:        a.Titulo,
			Artista:       a.Artista,
			Genero:        a.Genero,
		})
	}
	c.vista.MostrarListaAudios(audiosVista)
}

// EjecutarConsultarMetadatos solicita metadatos de un audio seleccionado por el usuario.
func (c *ControladorCliente) EjecutarConsultarMetadatos() {
	// Primero listar audios para que el usuario seleccione
	audios, err := c.fachada.ListarAudios()
	if err != nil {
		c.vista.MostrarError(err)
		return
	}
	if len(audios) == 0 {
		c.vista.MostrarMensaje("No hay audios disponibles.")
		return
	}

	// Mostrar lista y pedir selección
	var audiosVista []struct {
		NombreArchivo string
		Titulo        string
		Artista       string
		Genero        string
	}
	for _, a := range audios {
		audiosVista = append(audiosVista, struct {
			NombreArchivo string
			Titulo        string
			Artista       string
			Genero        string
		}{a.NombreArchivo, a.Titulo, a.Artista, a.Genero})
	}
	c.vista.MostrarListaAudios(audiosVista)

	opcion := c.vista.PedirSeleccionAudio(len(audios))
	audioSeleccionado := audios[opcion-1]

	// Obtener metadatos del audio seleccionado
	meta, err := c.fachada.ObtenerMetadatos(audioSeleccionado.NombreArchivo)
	if err != nil {
		c.vista.MostrarError(err)
		return
	}

	c.vista.MostrarMetadatos(
		meta.NombreArchivo,
		meta.Titulo,
		meta.Artista,
		meta.Genero,
		meta.TamanoBytes,
	)
}

// EjecutarReproducirAudio gestiona la selección y reproducción de un audio.
func (c *ControladorCliente) EjecutarReproducirAudio() {
	// Listar audios disponibles
	audios, err := c.fachada.ListarAudios()
	if err != nil {
		c.vista.MostrarError(err)
		return
	}
	if len(audios) == 0 {
		c.vista.MostrarMensaje("No hay audios disponibles para reproducir.")
		return
	}

	// Mostrar lista y pedir selección
	var audiosVista []struct {
		NombreArchivo string
		Titulo        string
		Artista       string
		Genero        string
	}
	for _, a := range audios {
		audiosVista = append(audiosVista, struct {
			NombreArchivo string
			Titulo        string
			Artista       string
			Genero        string
		}{a.NombreArchivo, a.Titulo, a.Artista, a.Genero})
	}
	c.vista.MostrarListaAudios(audiosVista)

	opcion := c.vista.PedirSeleccionAudio(len(audios))
	audioSeleccionado := audios[opcion-1]

	fmt.Printf("\n▶️  Iniciando reproducción de: %s\n", audioSeleccionado.Titulo)

	// Reproducir via streaming gRPC
	err = c.fachada.ReproducirAudio(audioSeleccionado.NombreArchivo)
	if err != nil {
		c.vista.MostrarError(err)
		return
	}
}
