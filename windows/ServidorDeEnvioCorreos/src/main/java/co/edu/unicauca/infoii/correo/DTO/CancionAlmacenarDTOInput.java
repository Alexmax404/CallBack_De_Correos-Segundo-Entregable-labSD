package co.edu.unicauca.infoii.correo.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * DTO que representa los metadatos de una canción recibidos desde la cola RabbitMQ.
 */
public class CancionAlmacenarDTOInput {

    @JsonProperty("idAudio")
    private String idAudio;

    @JsonProperty("titulo")
    private String titulo;

    @JsonProperty("artista")
    private String artista;

    @JsonProperty("genero")
    private String genero;

    @JsonProperty("fechaHoraRegistro")
    private String fechaHoraRegistro;

    @JsonProperty("mensaje")
    private String mensaje;

    @JsonProperty("fraseMotivadora")
    private String fraseMotivadora;

    public String getIdAudio() { return idAudio; }
    public String getTitulo() { return titulo; }
    public String getArtista() { return artista; }
    public String getGenero() { return genero; }
    public String getFechaHoraRegistro() { return fechaHoraRegistro; }
    public String getMensaje() { return mensaje; }
    public String getFraseMotivadora() { return fraseMotivadora; }
}
