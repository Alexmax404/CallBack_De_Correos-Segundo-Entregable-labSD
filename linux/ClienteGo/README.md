# Cliente de Streaming — Instrucciones

## 1. Generar código desde el proto

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/streaming.proto
```

## 2. Descargar dependencias

```bash
go mod tidy
```

## 3. Ejecutar

```bash
go run main/main.go
```

> ⚠️ El ServidorDeStreaming debe estar corriendo en :50051 antes de ejecutar el cliente.

## 4. Menú del cliente

```
╔══════════════════════════════════════════╗
║     🎵 CLIENTE DE STREAMING DE AUDIO     ║
╚══════════════════════════════════════════╝

  MENÚ PRINCIPAL
  1. Iniciar sesión
  0. Salir

  [Tras el login]
  1. Ver tipos de audio (géneros)
  2. Ver todos los audios
  3. Consultar metadatos de un audio
  4. Reproducir un audio
  0. Cerrar sesión
```

## 5. Estructura del proyecto

```
ClienteGo/
├── proto/
│   ├── streaming.proto          ← igual al del servidor
│   ├── streaming.pb.go          ← generado por protoc
│   └── streaming_grpc.pb.go     ← generado por protoc
├── capaFachadaServices/
│   ├── DTOs/
│   │   └── loginDTOInput.go
│   ├── models/
│   │   └── modelos.go
│   └── fachada/
│       └── fachadaCliente.go    ← conexión gRPC
├── capaControladores/
│   └── controladorCliente.go
├── capaVistas/
│   └── vista.go                 ← menú en consola
├── main/
│   └── main.go
└── go.mod
```

## 6. Configuración

Si el servidor de streaming está en otra máquina, cambia en `fachadaCliente.go`:

```go
const DireccionServidor = "192.168.0.X:50051"
```
