package main

import (
	"fmt"
	"os"

	capacontroladores "clientestreaming/capaControladores"
	"clientestreaming/capaFachadaServices/fachada"
	capavistas "clientestreaming/capaVistas"
)

func main() {
	vista := capavistas.NuevaVista()
	vista.MostrarBienvenida()

	// Conectar al servidor de streaming
	fachadaCliente, err := fachada.NuevaFachadaCliente()
	if err != nil {
		fmt.Printf("❌ No se pudo conectar al servidor: %v\n", err)
		os.Exit(1)
	}
	defer fachadaCliente.Cerrar()

	controlador := capacontroladores.NuevoControladorCliente(fachadaCliente, vista)

	// Bucle del menú de login
	for {
		opcion := vista.MostrarMenuLogin()

		switch opcion {
		case 1:
			if controlador.EjecutarLogin() {
				ejecutarMenuPrincipal(controlador, vista, fachadaCliente)
			}
		case 0:
			fmt.Println("\n👋 ¡Hasta luego!")
			return
		default:
			vista.MostrarMensaje("Opción inválida.")
		}
	}
}

// ejecutarMenuPrincipal gestiona el menú principal una vez autenticado el usuario.
func ejecutarMenuPrincipal(
	controlador *capacontroladores.ControladorCliente,
	vista *capavistas.Vista,
	f *fachada.FachadaCliente,
) {
	for {
		usuario := f.ObtenerUsuario()
		opcion := vista.MostrarMenuPrincipal(usuario.Nickname)

		switch opcion {
		case 1:
			controlador.EjecutarVerGeneros()
		case 2:
			controlador.EjecutarVerAudios()
		case 3:
			controlador.EjecutarConsultarMetadatos()
		case 4:
			controlador.EjecutarReproducirAudio()
		case 0:
			fmt.Println("\n🔓 Sesión cerrada.")
			return
		default:
			vista.MostrarMensaje("Opción inválida. Intenta de nuevo.")
		}
	}
}
