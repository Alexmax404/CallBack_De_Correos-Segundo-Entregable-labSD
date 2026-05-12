package co.edu.unicauca.infoii.administrador.capaFachadaServices.DTOs;

import java.io.Serializable;

/**
 * DTO serializable que viaja por RMI cuando se reproduce un audio.
 * Contiene la fecha/hora de reproducción y el id del audio,
 * ambos generados en el ServidorDeStreaming.
 */
public class NotificacionReproduccionDTO implements Serializable {
    private static final long serialVersionUID = 1L;

    private String idAudio;
    private String fechaHoraReproduccion;
    private String tituloAudio;

    public NotificacionReproduccionDTO() {}

    public NotificacionReproduccionDTO(String idAudio,
                                        String fechaHoraReproduccion,
                                        String tituloAudio) {
        this.idAudio = idAudio;
        this.fechaHoraReproduccion = fechaHoraReproduccion;
        this.tituloAudio = tituloAudio;
    }

    public String getIdAudio() { return idAudio; }
    public String getFechaHoraReproduccion() { return fechaHoraReproduccion; }
    public String getTituloAudio() { return tituloAudio; }

    public void setIdAudio(String idAudio) { this.idAudio = idAudio; }
    public void setFechaHoraReproduccion(String f) { this.fechaHoraReproduccion = f; }
    public void setTituloAudio(String tituloAudio) { this.tituloAudio = tituloAudio; }
}
