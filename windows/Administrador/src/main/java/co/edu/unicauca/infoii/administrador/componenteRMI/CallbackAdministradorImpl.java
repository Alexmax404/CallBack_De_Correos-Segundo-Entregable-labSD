package co.edu.unicauca.infoii.administrador.componenteRMI;

import co.edu.unicauca.infoii.administrador.capaFachadaServices.DTOs.NotificacionReproduccionDTO;
import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

/**
 * Implementación del Callback RMI.
 * El ServidorDeStreaming llama a notificarReproduccion() cada vez
 * que un cliente reproduce un audio.
 */
public class CallbackAdministradorImpl extends UnicastRemoteObject
        implements ICallbackAdministrador {

    private static final long serialVersionUID = 1L;

    public CallbackAdministradorImpl() throws RemoteException {
        super();
    }

    @Override
    public void notificarReproduccion(NotificacionReproduccionDTO notificacion)
            throws RemoteException {

        System.out.println("\n╔══════════════════════════════════════════╗");
        System.out.println("║   🔔 CALLBACK - NUEVO AUDIO REPRODUCIDO  ║");
        System.out.println("╚══════════════════════════════════════════╝");
        System.out.println("  🎵 Audio ID    : " + notificacion.getIdAudio());
        System.out.println("  🎵 Título      : " + notificacion.getTituloAudio());
        System.out.println("  📅 Fecha/Hora  : " + notificacion.getFechaHoraReproduccion());
        System.out.println("══════════════════════════════════════════════\n");
    }
}
