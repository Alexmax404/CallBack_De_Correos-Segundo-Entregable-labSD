# ServidorDeStreaming — Instrucciones de instalación

## 1. Instalar dependencias en Ubuntu (Debian)

```bash
# Instalar protoc
sudo apt-get install -y protobuf-compiler

# Instalar plugins de Go para protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Agregar al PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

## 2. Generar el código Go desde el archivo .proto

Desde la raíz del proyecto (ServidorStreaming/):

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/streaming.proto
```

Esto genera dos archivos en la carpeta proto/:
- streaming.pb.go       (estructuras de datos)
- streaming_grpc.pb.go  (interfaces del servidor y cliente)

## 3. Descargar dependencias

```bash
go mod tidy
```

## 4. Copiar los audios

Copia la carpeta `audios/` del ServidorDeAudiosMetadatos dentro de esta carpeta:

```bash
cp -r ../SERVIDORQUEALMACENACANCIONES/audios ./audios
```

## 5. Ejecutar el servidor

```bash
go run main/main.go
```

Deberías ver:
```
========================================
  🎵 Servidor de Streaming de Audios
        Protocolo: gRPC
========================================
✅ Servidor gRPC escuchando en :50051...
```

## 6. Estructura del proyecto

```
ServidorStreaming/
├── proto/
│   ├── streaming.proto          ← definición del servicio
│   ├── streaming.pb.go          ← generado por protoc
│   └── streaming_grpc.pb.go     ← generado por protoc
├── capaAccesoDatos/
│   └── repositorioAudios.go
├── capaFachadaServices/
│   ├── models/
│   │   └── modelos.go
│   └── fachada/
│       └── fachadaStreaming.go
├── capaControladores/
│   └── controladorStreaming.go
├── main/
│   └── main.go
├── audios/                      ← carpeta con los .mp3
├── go.mod
└── README.md
```

## Servicios gRPC disponibles

| Servicio | Tipo | Descripción |
|----------|------|-------------|
| ListarAudios | Unario | Lista todos los audios disponibles |
| ObtenerMetadatosAudio | Unario | Metadatos de un audio específico |
| ReproducirAudio | Server Streaming | Envía el audio en fragmentos de 64KB |

## Puerto
- gRPC: **50051**
