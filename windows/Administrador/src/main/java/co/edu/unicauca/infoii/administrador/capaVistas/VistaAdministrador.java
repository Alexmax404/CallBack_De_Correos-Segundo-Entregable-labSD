package co.edu.unicauca.infoii.administrador.capaVistas;

import co.edu.unicauca.infoii.administrador.capaFachadaServices.models.MetadatosAudio;
import java.util.List;
import java.util.Scanner;

/**
 * Vista del Administrador — gestiona la interacción por consola.
 */
public class VistaAdministrador {

    private final Scanner scanner;

    public VistaAdministrador() {
        this.scanner = new Scanner(System.in);
    }

    public void mostrarBienvenida() {
        System.out.println();
        System.out.println("╔══════════════════════════════════════════╗");
        System.out.println("║    🎛️  PANEL DEL ADMINISTRADOR           ║");
        System.out.println("║       Universidad del Cauca - FIET       ║");
        System.out.println("║        Sistemas Distribuidos             ║");
        System.out.println("╚══════════════════════════════════════════╝");
        System.out.println();
    }

    public int mostrarMenuLogin() {
        System.out.println("──────────────────────────────────────────");
        System.out.println("  ACCESO AL SISTEMA");
        System.out.println("──────────────────────────────────────────");
        System.out.println("  1. Ingresar como administrador");
        System.out.println("  0. Salir");
        System.out.println("──────────────────────────────────────────");
        return leerEntero("Selecciona una opción: ");
    }

    public int mostrarMenuPrincipal(String nickname) {
        System.out.println();
        System.out.println("──────────────────────────────────────────");
        System.out.println("  👤 Administrador: " + nickname);
        System.out.println("──────────────────────────────────────────");
        System.out.println("  1. Almacenar nuevo audio");
        System.out.println("  2. Listar audios y metadatos");
        System.out.println("  0. Cerrar sesión");
        System.out.println("──────────────────────────────────────────");
        return leerEntero("Selecciona una opción: ");
    }

    public String[] pedirDatosAudio() {
        System.out.println("\n📁 REGISTRAR NUEVO AUDIO");
        System.out.println("──────────────────────────────────────────");
        String titulo = leerTexto("  Título del audio : ");
        String artista = leerTexto("  Artista          : ");
        String genero = leerTexto("  Género           : ");
        String ruta = leerTexto("  Ruta del archivo (.mp3): ");
        return new String[]{titulo, artista, genero, ruta};
    }

    public void mostrarListaAudios(List<MetadatosAudio> audios) {
        System.out.println();
        System.out.println("🎵 AUDIOS ALMACENADOS:");
        System.out.println("──────────────────────────────────────────");
        if (audios == null || audios.isEmpty()) {
            System.out.println("  No hay audios registrados.");
        } else {
            for (int i = 0; i < audios.size(); i++) {
                MetadatosAudio a = audios.get(i);
                System.out.println("  " + (i + 1) + ". " + a.getTitulo());
                System.out.println("     🎤 Artista    : " + a.getArtista());
                System.out.println("     🎸 Género     : " + a.getGenero());
                System.out.println("     🆔 ID         : " + a.getIdAudio());
                System.out.println("     📅 Registrado : " + a.getFechaHoraRegistro());
                System.out.println("     📄 Archivo    : " + a.getNombreArchivo());
                System.out.println();
            }
        }
        System.out.println("──────────────────────────────────────────");
    }

    public void mostrarMensaje(String mensaje) {
        System.out.println("ℹ️  " + mensaje);
    }

    public void mostrarError(String error) {
        System.out.println("❌ " + error);
    }

    public String leerTexto(String mensaje) {
        System.out.print(mensaje);
        return scanner.nextLine().trim();
    }

    public int leerEntero(String mensaje) {
        while (true) {
            try {
                System.out.print(mensaje);
                int valor = Integer.parseInt(scanner.nextLine().trim());
                return valor;
            } catch (NumberFormatException e) {
                System.out.println("⚠️  Ingresa un número válido.");
            }
        }
    }
}
