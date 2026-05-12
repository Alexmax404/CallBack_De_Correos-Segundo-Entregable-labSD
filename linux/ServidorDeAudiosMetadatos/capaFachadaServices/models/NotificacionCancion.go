package models

type NotificacionCancion struct {
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Genero  string `json:"genero"`
	Mensaje string `json:"mensaje"`
}
