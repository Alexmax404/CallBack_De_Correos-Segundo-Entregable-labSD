package models

// InfoAudio contiene la información básica de un audio para listar.
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

// FragmentoAudio representa un chunk del archivo de audio para streaming.
type FragmentoAudio struct {
	Datos       []byte
	NumeroChunk int
	EsUltimo    bool
}
