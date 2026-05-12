package co.edu.unicauca.infoii.administrador.capaFachadaServices.fachada;

import co.edu.unicauca.infoii.administrador.capaFachadaServices.DTOs.AudioAlmacenarDTOInput;
import co.edu.unicauca.infoii.administrador.capaFachadaServices.models.MetadatosAudio;
import co.edu.unicauca.infoii.administrador.componenteRMI.CallbackAdministradorImpl;
import co.edu.unicauca.infoii.administrador.componenteRMI.ICallbackAdministrador;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import okhttp3.*;

import java.io.File;
import java.rmi.Naming;
import java.rmi.registry.LocateRegistry;
import java.util.Collections;
import java.util.List;

/**
 * FachadaAdministrador coordina:
 *  - Las peticiones REST al ServidorDeAudiosMetadatos
 *  - El registro del Callback RMI en el ServidorDeStreaming
 */
public class FachadaAdministrador {

    // URL del ServidorDeAudiosMetadatos
    private static final String URL_SERVIDOR_AUDIOS = "http://localhost:5000";

    // Nombre RMI con el que se registra el ServidorDeStreaming
    private static final String NOMBRE_RMI_SERVIDOR = "ServidorStreaming";

    // Puerto donde corre el registro RMI del administrador
    private static final int PUERTO_RMI = 2099;

    private final OkHttpClient clienteHttp;
    private final ObjectMapper mapper;

    public FachadaAdministrador() {
        this.clienteHttp = new OkHttpClient();
        this.mapper = new ObjectMapper();
        System.out.println("🎛️  Fachada del Administrador inicializada.");
    }

    // ================================================================
    // OPERACIONES REST — ServidorDeAudiosMetadatos
    // ================================================================

    /**
     * Almacena un audio en el ServidorDeAudiosMetadatos via REST multipart.
     */
    public boolean almacenarAudio(AudioAlmacenarDTOInput dto) {
        try {
            File archivo = new File(dto.getRutaArchivo());
            if (!archivo.exists()) {
                System.out.println("❌ Archivo no encontrado: " + dto.getRutaArchivo());
                return false;
            }

            RequestBody cuerpo = new MultipartBody.Builder()
                    .setType(MultipartBody.FORM)
                    .addFormDataPart("titulo", dto.getTitulo())
                    .addFormDataPart("artista", dto.getArtista())
                    .addFormDataPart("genero", dto.getGenero())
                    .addFormDataPart("audio", archivo.getName(),
                            RequestBody.create(archivo, MediaType.parse("audio/mpeg")))
                    .build();

            Request peticion = new Request.Builder()
                    .url(URL_SERVIDOR_AUDIOS + "/canciones/almacenamiento")
                    .post(cuerpo)
                    .build();

            try (Response respuesta = clienteHttp.newCall(peticion).execute()) {
                if (respuesta.isSuccessful()) {
                    System.out.println("✅ Audio almacenado correctamente.");
                    return true;
                } else {
                    System.out.println("❌ Error del servidor: " + respuesta.code());
                    return false;
                }
            }
        } catch (Exception e) {
            System.out.println("❌ Error al almacenar audio: " + e.getMessage());
            return false;
        }
    }

    /**
     * Lista los metadatos de todos los audios almacenados via REST GET.
     */
    public List<MetadatosAudio> listarAudios() {
        try {
            Request peticion = new Request.Builder()
                    .url(URL_SERVIDOR_AUDIOS + "/canciones/listar")
                    .get()
                    .build();

            try (Response respuesta = clienteHttp.newCall(peticion).execute()) {
                if (respuesta.isSuccessful() && respuesta.body() != null) {
                    String json = respuesta.body().string();
                    List<MetadatosAudio> audios = mapper.readValue(
                            json, new TypeReference<List<MetadatosAudio>>() {});
                    return audios != null ? audios : Collections.emptyList();
                } else {
                    System.out.println("❌ Error al listar audios: " + respuesta.code());
                    return Collections.emptyList();
                }
            }
        } catch (Exception e) {
            System.out.println("❌ Error al listar audios: " + e.getMessage());
            return Collections.emptyList();
        }
    }

    // ================================================================
    // CALLBACK RMI — ServidorDeStreaming
    // ================================================================

    /**
     * Registra el Callback RMI para recibir notificaciones de reproducción.
     * El administrador crea un registro RMI local y lo exporta
     * para que el ServidorDeStreaming pueda invocarlo.
     */
    public void registrarCallbackRMI(String nicknameAdmin) {
        try {
            System.out.println("📡 Registrando Callback RMI...");

            // Crear registro RMI local en el puerto configurado
            LocateRegistry.createRegistry(PUERTO_RMI);

            // Crear e implementar el objeto remoto callback
            ICallbackAdministrador callback = new CallbackAdministradorImpl();

            // Registrar con el nombre del administrador
            String nombreRegistro = "CallbackAdmin_" + nicknameAdmin;
            Naming.rebind("rmi://localhost:" + PUERTO_RMI + "/" + nombreRegistro, callback);

            System.out.println("✅ Callback RMI registrado como: " + nombreRegistro);
            System.out.println("   Puerto RMI: " + PUERTO_RMI);
            System.out.println("   En espera de notificaciones de reproducción...");

            // Notificar al ServidorDeStreaming que este admin quiere callbacks
            registrarEnServidorStreaming(nicknameAdmin, nombreRegistro);

        } catch (Exception e) {
            System.out.println("❌ Error registrando Callback RMI: " + e.getMessage());
        }
    }

    /**
     * Notifica al ServidorDeStreaming la dirección RMI del callback
     * de este administrador via REST.
     */
    private void registrarEnServidorStreaming(String nicknameAdmin, String nombreRegistro) {
        try {
            String urlRegistro = "http://localhost:8080/streaming/callback/registrar"
                    + "?nickname=" + nicknameAdmin
                    + "&rmiUrl=rmi://localhost:" + PUERTO_RMI + "/" + nombreRegistro;

            Request peticion = new Request.Builder()
                    .url(urlRegistro)
                    .post(RequestBody.create("", null))
                    .build();

            try (Response respuesta = clienteHttp.newCall(peticion).execute()) {
                if (respuesta.isSuccessful()) {
                    System.out.println("✅ Administrador registrado en ServidorDeStreaming.");
                } else {
                    System.out.println("⚠️  ServidorDeStreaming no disponible aún " +
                            "(se intentará cuando se reproduzca un audio).");
                }
            }
        } catch (Exception e) {
            System.out.println("⚠️  ServidorDeStreaming no disponible aún: " + e.getMessage());
        }
    }
}
