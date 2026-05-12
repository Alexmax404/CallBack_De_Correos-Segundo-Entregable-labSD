package co.edu.unicauca.infoii.correo.componenteListener;

import co.edu.unicauca.infoii.correo.DTO.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.commons.Simulacion;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Service;

/**
 * Componente que escucha la cola RabbitMQ y simula el envío de un correo
 * cuando se registra un nuevo audio en el servidor de audios y metadatos.
 */
@Service
public class MessageConsumer {

    @RabbitListener(queues = "notificaciones_canciones")
    public void receiveMessage(CancionAlmacenarDTOInput cancion) {
        System.out.println("\n╔══════════════════════════════════════════╗");
        System.out.println("║     📧 NUEVO CORREO RECIBIDO EN COLA     ║");
        System.out.println("╚══════════════════════════════════════════╝");
        System.out.println("📨 Preparando envío de correo electrónico...");

        Simulacion.simular(10000, "Enviando correo...");

        System.out.println("\n✅ Correo enviado al cliente con los siguientes datos:");
        System.out.println("──────────────────────────────────────────");
        System.out.println("🎵 ID del Audio    : " + cancion.getIdAudio());
        System.out.println("🎵 Título          : " + cancion.getTitulo());
        System.out.println("🎤 Artista         : " + cancion.getArtista());
        System.out.println("🎸 Género          : " + cancion.getGenero());
        System.out.println("📅 Fecha/Hora      : " + cancion.getFechaHoraRegistro());
        System.out.println("📢 Mensaje         : " + cancion.getMensaje());
        System.out.println("💫 Frase motivadora: " + cancion.getFraseMotivadora());
        System.out.println("──────────────────────────────────────────\n");
    }
}
