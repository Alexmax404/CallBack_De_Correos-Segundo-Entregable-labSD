package co.edu.unicauca.infoii.administrador.capaControladores;

import co.edu.unicauca.infoii.administrador.capaFachadaServices.DTOs.AudioAlmacenarDTOInput;
import co.edu.unicauca.infoii.administrador.capaFachadaServices.fachada.FachadaAdministrador;
import co.edu.unicauca.infoii.administrador.capaFachadaServices.models.MetadatosAudio;
import co.edu.unicauca.infoii.administrador.capaVistas.VistaAdministrador;

import java.util.List;

/**
 * ControladorAdministrador orquesta las operaciones entre la vista y la fachada.
 * Aplica el patrón MVC.
 */
public class ControladorAdministrador {

    private final FachadaAdministrador fachada;
    private final VistaAdministrador vista;
    private String nicknameAdmin;

    public ControladorAdministrador(FachadaAdministrador fachada, VistaAdministrador vista) {
        this.fachada = fachada;
        this.vista = vista;
    }

    /**
     * Gestiona el proceso de login del administrador y registra el Callback RMI.
     */
    public boolean ejecutarLogin() {
        System.out.println("\n🔐 INICIO DE SESIÓN — ADMINISTRADOR");
        System.out.println("──────────────────────────────────────────");
        String nickname = vista.leerTexto("  Nickname   : ");
        String contrasena = vista.leerTexto("  Contraseña : ");

        if (nickname.isEmpty() || contrasena.length() < 4) {
            vista.mostrarError("Nickname vacío o contraseña muy corta (mínimo 4 caracteres).");
            return false;
        }

        this.nicknameAdmin = nickname;
        System.out.println("✅ Sesión iniciada como administrador: " + nickname);

        // Registrar el Callback RMI para recibir notificaciones de reproducción
        fachada.registrarCallbackRMI(nicknameAdmin);
        return true;
    }

    /**
     * Solicita los datos del audio al usuario y los envía al servidor via REST.
     */
    public void ejecutarAlmacenarAudio() {
        String[] datos = vista.pedirDatosAudio();
        String titulo = datos[0];
        String artista = datos[1];
        String genero = datos[2];
        String rutaArchivo = datos[3];

        System.out.println("\n📤 Enviando audio al servidor...");

        AudioAlmacenarDTOInput dto = new AudioAlmacenarDTOInput(
                titulo, artista, genero, rutaArchivo);

        boolean exito = fachada.almacenarAudio(dto);
        if (!exito) {
            vista.mostrarError("No se pudo almacenar el audio. Verifica la ruta y que el servidor esté activo.");
        }
    }

    /**
     * Obtiene y muestra la lista de audios con todos sus metadatos.
     */
    public void ejecutarListarAudios() {
        System.out.println("\n📋 Consultando audios al servidor...");
        List<MetadatosAudio> audios = fachada.listarAudios();
        vista.mostrarListaAudios(audios);
    }

    public String getNicknameAdmin() {
        return nicknameAdmin;
    }
}
