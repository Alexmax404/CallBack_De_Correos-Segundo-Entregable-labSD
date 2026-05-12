package co.edu.unicauca.infoii.administrador.capaFachadaServices.models;

/**
 * Modelo que representa los metadatos completos de un audio almacenado.
 */
public class MetadatosAudio {
    private String idAudio;
    private String titulo;
    private String artista;
    private String genero;
    private String fechaHoraRegistro;
    private String nombreArchivo;

    public MetadatosAudio() {}

    public String getIdAudio() { return idAudio; }
    public String getTitulo() { return titulo; }
    public String getArtista() { return artista; }
    public String getGenero() { return genero; }
    public String getFechaHoraRegistro() { return fechaHoraRegistro; }
    public String getNombreArchivo() { return nombreArchivo; }

    public void setIdAudio(String idAudio) { this.idAudio = idAudio; }
    public void setTitulo(String titulo) { this.titulo = titulo; }
    public void setArtista(String artista) { this.artista = artista; }
    public void setGenero(String genero) { this.genero = genero; }
    public void setFechaHoraRegistro(String f) { this.fechaHoraRegistro = f; }
    public void setNombreArchivo(String nombreArchivo) { this.nombreArchivo = nombreArchivo; }
}
