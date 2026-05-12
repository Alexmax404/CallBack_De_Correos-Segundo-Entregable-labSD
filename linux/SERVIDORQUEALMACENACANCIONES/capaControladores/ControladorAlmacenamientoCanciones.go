package capacontroladores

import (
	"fmt"
	"io"
	"net/http"
	dtos "servidorquealmacenacanciones/capaFachadaServices/DTOs"
	"servidorquealmacenacanciones/capaFachadaServices/fachada"
)

type ControladorAlmacenamientoCanciones struct {
	fachada *fachada.FachadaAlmacenamiento
}

func NuevoControladorAlmacenamientoCanciones(f *fachada.FachadaAlmacenamiento) *ControladorAlmacenamientoCanciones {
	return &ControladorAlmacenamientoCanciones{fachada: f}
}

func (thisC *ControladorAlmacenamientoCanciones) AlmacenarAudioCancion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Almacenando cancion...")

	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	// IMPORTANTE: parsear multipart
	err := r.ParseMultipartForm(50 << 20)
	fmt.Println("Form keys:", r.MultipartForm.File)
	if err != nil {
		fmt.Println("Error ParseMultipartForm:", err)
		http.Error(w, "Error al parsear formulario", http.StatusBadRequest)
		return
	}

	// 🔹 DATOS TEXTO
	titulo := r.FormValue("titulo")
	artista := r.FormValue("artista")
	genero := r.FormValue("genero")

	fmt.Println("TITULO:", titulo)
	fmt.Println("ARTISTA:", artista)
	fmt.Println("GENERO:", genero)

	// 🔹 ARCHIVO
	file, header, err := r.FormFile("audio")
	if err != nil {
		fmt.Println("Error al obtener archivo:", err)
		http.Error(w, "Error al obtener el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println("ARCHIVO RECIBIDO:", header.Filename)

	// 🔹 LEER BYTES
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		http.Error(w, "Error al leer archivo", http.StatusInternalServerError)
		return
	}

	fmt.Println("TAMAÑO AUDIO:", len(data))

	// 🔹 DTO
	dto := dtos.CancionAlmacenarDTOinput{
		Titulo:  titulo,
		Artista: artista,
		Genero:  genero,
	}

	// 🔹 LLAMADA A FACHADA
	thisC.fachada.GuardarCancion(dto.Titulo, dto.Artista, dto.Genero, data)

	fmt.Println("Canción enviada a fachada correctamente")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Canción almacenada correctamente"))
}
