# Formá la palabra
**Formá la palabra** es un juego en el que tienes que ordenar un conjunto de letras para obtener una palabra.

Es un servidor HTTP hecho en **Go** ([fiber](https://gofiber.io)) con WebSockets para dar acabo una recreación de la versión hecha por **[ElSpreen](https://twitch.tv/elspreen)** en directo.

## Características
- [x] **Autenticación**
- [x] **Tiempo del juego** (personalizable en milisegundos, **segundos**, horas, días, etc.)
  - [x] Establecer tiempo del juego
  - [x] Pausar el tiempo del juego
  - [x] Añadir tiempo extra al juego
- [x] **Establecer la palabra**
- [x] **Pantallas de palabra correcta e incorrecta**
  - [x] Implementación servidor
  - [ ] Implementación cliente
- [ ] **Control por lado de cliente** (página con botones para realizar acciones del juego)

## Implementación
### Requisitos
- **[Go](https://go.dev/)**
- Clonar (descargar) este repositorio

### Ejecutar el servidor
```sh
go build cmd/flp/main.go
./main
# o (Windows)
main.exe
```