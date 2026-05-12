package main

import (
	"fmt"
	"net/http"

	capacontroladores "servidorquealmacenacanciones/capaControladores"
	"servidorquealmacenacanciones/capaFachadaServices/fachada"
)

func main() {
	fachadaAlmacenamiento := fachada.NuevaFachadaAlmacenamiento()
	controlador := capacontroladores.NuevoControladorAlmacenamientoCanciones(fachadaAlmacenamiento)

	mux := http.NewServeMux()
	mux.HandleFunc("/canciones/almacenamiento", controlador.AlmacenarAudioCancion)
	fmt.Println("Servidor escuchando en :5000...")
	http.ListenAndServe(":5000", mux)

}
