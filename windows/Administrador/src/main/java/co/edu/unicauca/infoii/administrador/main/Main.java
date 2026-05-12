package co.edu.unicauca.infoii.administrador.main;

import co.edu.unicauca.infoii.administrador.capaControladores.ControladorAdministrador;
import co.edu.unicauca.infoii.administrador.capaFachadaServices.fachada.FachadaAdministrador;
import co.edu.unicauca.infoii.administrador.capaVistas.VistaAdministrador;

/**
 * Punto de entrada del Administrador.
 * Inicializa las capas y ejecuta el bucle del menú.
 */
public class Main {

    public static void main(String[] args) {
        VistaAdministrador vista = new VistaAdministrador();
        vista.mostrarBienvenida();

        FachadaAdministrador fachada = new FachadaAdministrador();
        ControladorAdministrador controlador = new ControladorAdministrador(fachada, vista);

        // Bucle del menú de login
        while (true) {
            int opcionLogin = vista.mostrarMenuLogin();

            switch (opcionLogin) {
                case 1:
                    if (controlador.ejecutarLogin()) {
                        ejecutarMenuPrincipal(controlador, vista);
                    }
                    break;
                case 0:
                    System.out.println("\n👋 ¡Hasta luego!");
                    System.exit(0);
                    break;
                default:
                    vista.mostrarMensaje("Opción inválida.");
            }
        }
    }

    /**
     * Ejecuta el menú principal una vez autenticado el administrador.
     */
    private static void ejecutarMenuPrincipal(
            ControladorAdministrador controlador,
            VistaAdministrador vista) {

        while (true) {
            int opcion = vista.mostrarMenuPrincipal(controlador.getNicknameAdmin());

            switch (opcion) {
                case 1:
                    controlador.ejecutarAlmacenarAudio();
                    break;
                case 2:
                    controlador.ejecutarListarAudios();
                    break;
                case 0:
                    System.out.println("\n🔓 Sesión cerrada.");
                    return;
                default:
                    vista.mostrarMensaje("Opción inválida. Intenta de nuevo.");
            }
        }
    }
}
