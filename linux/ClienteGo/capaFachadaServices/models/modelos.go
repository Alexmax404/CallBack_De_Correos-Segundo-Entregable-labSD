package models

// Usuario representa al cliente autenticado en el sistema.
type Usuario struct {
	Nickname   string
	Contrasena string
}

// InfoAudio contiene la información básica de un audio para mostrar en el menú.
type InfoAudio struct {
	NombreArchivo string
	Titulo        string
	Artista       string
	Genero        string
}

// MetadatosAudio contiene la información completa de un audio.
type MetadatosAudio struct {
	NombreArchivo string
	Titulo        string
	Artista       string
	Genero        string
	TamanoBytes   int64
}
