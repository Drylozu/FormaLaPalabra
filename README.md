# Formá la palabra
**Formá la palabra** es un juego en el que tienes que ordenar un conjunto de letras para obtener una palabra.

Es un servidor HTTP hecho en **Go** ([fiber](https://gofiber.io)) con WebSockets para dar acabo una recreación de la versión hecha por **[ElSpreen](https://twitch.tv/elspreen)** en directo.

**[Lee la implementación de su API y configuración aquí](https://github.com/Drylozu/FormaLaPalabra/blob/main/IMPLEMENTATION.md)**.

## Características
- [x] **Autenticación**
- [x] **Tiempo del juego** (personalizable en milisegundos, **segundos**, horas, días, etc.)
  - [x] Establecer tiempo del juego
  - [x] Pausar el tiempo del juego
  - [x] Añadir tiempo extra al juego
- [x] **Establecer la palabra**
- [x] **Pantallas de palabra correcta e incorrecta**
  - [x] Implementación servidor
  - [x] Implementación cliente
- [ ] **Control por lado de cliente** (página con botones para realizar acciones del juego)

## Uso
- Clonar (descargar) este repositorio

### Docker
- Instalar **[Docker](https://www.docker.com/)**.
```sh
docker build -t forma-la-palabra .

docker run -dp 3000:3000 forma-la-palabra flp --secret "texto secreto para autenticación"
```

### En tu máquina
- Instalar **[Go](https://go.dev/)**.
```sh
# descargar dependencias y crear un ejecutable
go mod download
go build cmd/flp/main.go

# ejecutar el servidor
./main
# o (Windows)
main.exe
```