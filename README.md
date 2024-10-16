# Docker multi-stage build

## Que es multi-stage build
Docker Multi-stage Build (construcción en múltiples etapas), una técnica utilizada en Docker para optimizar la creación de imágenes. Con multi-stage build, puedes usar varias etapas dentro de un solo Dockerfile, lo que permite construir imágenes más pequeñas y eficientes al copiar solo los archivos necesarios para la imagen final, mientras descargas o construyes otras dependencias en etapas previas.

**Características principales:**
* Imágenes más pequeñas: Solo se incluye lo necesario en la imagen final, eliminando archivos intermedios, librerías o herramientas que solo son necesarias en la fase de construcción.
* Mejor organización: Puedes separar las fases de construcción y de ejecución en etapas diferentes, lo que mejora la legibilidad del Dockerfile.
Reducción de complejidad: Permite usar imágenes base diferentes en cada etapa, dependiendo de lo que necesites (por ejemplo, una imagen pesada para la compilación y otra ligera para la ejecución).

## Ejemplo Practico
```
# Etapa 1: Construcción
FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

# Etapa 2: Imagen final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]

```

## Como Ejecutar el proyecto

compile the docker file
```bash
$ docker build -t golang-simple-api .
```
```text
[+] Building 11.3s (18/18) FINISHED                                                                                                                                                   docker:default
 => [internal] load build definition from dockerfile                                                                                                                                            0.0s
 => => transferring dockerfile: 857B                                                                                                                                                            0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                                                2.0s
 => [internal] load metadata for docker.io/library/golang:1.21-alpine                                                                                                                           2.0s
 => [auth] library/golang:pull token for registry-1.docker.io                                                                                                                                   0.0s
 => [auth] library/alpine:pull token for registry-1.docker.io                                                                                                                                   0.0s
 => [internal] load .dockerignore                                                                                                                                                               0.0s
 => => transferring context: 2B                                                                                                                                                                 0.0s
 => [builder 1/7] FROM docker.io/library/golang:1.21-alpine@sha256:2414035b086e3c42b99654c8b26e6f5b1b1598080d65fd03c7f499552ff4dc94                                                             0.0s
 => [stage-1 1/3] FROM docker.io/library/alpine:latest@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d                                                                  0.0s
 => [internal] load build context                                                                                                                                                               0.0s
 => => transferring context: 37.74kB                                                                                                                                                            0.0s
 => CACHED [builder 2/7] WORKDIR /build                                                                                                                                                         0.0s
 => CACHED [builder 3/7] RUN apk add --no-cache git                                                                                                                                             0.0s
 => CACHED [builder 4/7] RUN git clone https://github.com/ErickAguilar95/golang-simple-api.git .                                                                                                0.0s
 => CACHED [builder 5/7] RUN go mod init golang-simple-api                                                                                                                                      0.0s
 => [builder 6/7] COPY . .                                                                                                                                                                      0.1s
 => [builder 7/7] RUN go build -o /build/app .                                                                                                                                                  9.0s
 => CACHED [stage-1 2/3] WORKDIR /root/                                                                                                                                                         0.0s
 => [stage-1 3/3] COPY --from=builder /build/app .                                                                                                                                              0.0s
 => exporting to image                                                                                                                                                                          0.0s
 => => exporting layers                                                                                                                                                                         0.0s
 => => writing image sha256:61095d6ed969996d407c26179cb87e7292c5cb0a82ea9fadb96ef5e616b8e960                                                                                                    0.0s
 => => naming to docker.io/library/golang-basic-api
```

run docker docker container
```bash
$ docker run -p 8000:8000 golang-simple-api
```
