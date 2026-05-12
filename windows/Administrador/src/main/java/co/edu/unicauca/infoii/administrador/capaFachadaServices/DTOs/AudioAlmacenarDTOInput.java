package co.edu.unicauca.infoii.administrador.capaFachadaServices.DTOs;

/**
 * DTO que encapsula los datos necesarios para almacenar un nuevo audio
 * en el ServidorDeAudiosMetadatos via REST.
 */
public class AudioAlmacenarDTOInput {
    private String titulo;
    private String artista;
    private String genero;
    private String rutaArchivo;

    public AudioAlmacenarDTOInput(String titulo, String artista,
                                   String genero, String rutaArchivo) {
        this.titulo = titulo;
        this.artista = artista;
        this.genero = genero;
        this.rutaArchivo = rutaArchivo;
    }

    public String getTitulo() { return titulo; }
    public String getArtista() { return artista; }
    public String getGenero() { return genero; }
    public String getRutaArchivo() { return rutaArchivo; }
}
