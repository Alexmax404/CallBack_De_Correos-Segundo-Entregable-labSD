package fachada

import (
	"fmt"
	capaaccedoadatos "servidorquealmacenacanciones/capaAccedoADatos"
	"servidorquealmacenacanciones/capaFachadaServices/models"
	componenteconexioncola "servidorquealmacenacanciones/componenteConexionCola"
)

type FachadaAlmacenamiento struct {
	repo         *capaaccedoadatos.RepositorioCanciones
	conexionCola *componenteconexioncola.RabbittPublisher
}

func NuevaFachadaAlmacenamiento() *FachadaAlmacenamiento {
	fmt.Println("Inicializando fachada de almacenamiento...")
	repo := capaaccedoadatos.GetRepositorioCanciones()

	conexionCola, err := componenteconexioncola.NewRabbitPublisher()
	if err != nil {
		fmt.Println("Error al conectar con RabbitMQ:", err)
		conexionCola = nil
	}

	return &FachadaAlmacenamiento{
		repo:         repo,
		conexionCola: conexionCola,
	}
}

func (thisF *FachadaAlmacenamiento) GuardarCancion(titulo string, artista string, genero string, data []byte) error {
	fmt.Println("Guardando cancion...")
	thisF.conexionCola.PublicarNotificacion(models.NotificacionCancion{
		Titulo:  titulo,
		Artista: artista,
		Genero:  genero,
		Mensaje: "Cancion guardada exitosamente" + titulo + " de " + artista,
	})
	return thisF.repo.GuardarCancion(titulo, genero, artista, data)
}
