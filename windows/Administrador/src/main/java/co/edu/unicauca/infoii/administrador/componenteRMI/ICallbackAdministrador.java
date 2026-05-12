package co.edu.unicauca.infoii.administrador.componenteRMI;

import co.edu.unicauca.infoii.administrador.capaFachadaServices.DTOs.NotificacionReproduccionDTO;
import java.rmi.Remote;
import java.rmi.RemoteException;

/**
 * Interfaz RMI que el ServidorDeStreaming invoca para notificar
 * al Administrador cuando un cliente reproduce un audio.
 */
public interface ICallbackAdministrador extends Remote {

    /**
     * Recibe la notificación de reproducción de un audio.
     * @param notificacion contiene idAudio y fechaHoraReproduccion
     */
    void notificarReproduccion(NotificacionReproduccionDTO notificacion) throws RemoteException;
}
