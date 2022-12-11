# Implementación
Formá La Palabra es un servidor HTTP relacionado con el programa de televisión Formá La Palabra, interpretado por ElSpreen su canal de Twitch.

## Configuración
- `--help` Muestra un mensaje de ayuda
- `--public` Especifica una carpeta public/ (./public/)
- `--secret` Especifica una clave para la autenticación WebSocket (explicado en "API")
- `--address` Especifica la dirección de escucha (0.0.0.0:3000)

<hr/>

## API
- `*` - Por defecto, el servidor hace un mapeo de los ficheros dentro de la carpeta `public` especificada en su configuración para colocar los archivos del cliente.
- `/ws` - Crea una conexión WebSocket.
  - `?k=` - La clave secreta de autenticación especificada por la opción `--secret`

## WebSocket
Se realiza una conexión WebSocket a la ruta `/ws?k=` seguida de la clave secreta de autenticación.

### Control
El siguiente fragmento de código JSON contiene las propiedades de comandos de control al servidor.
```json
{
    "word": "Nueva palabra",
    "startTimer": true,
    "extraTime": 10,
    "setTime": 10,
    "pause": false,
    "wrong": false,
    "correct": false
}
```
> No deberías enviar todas las propiedades en un mismo mensaje pero el servidor es capaz de administrarlo igualmente.

- `word` string - Cambia la palabra del juego.
- `startTimer` bool - Si su valor es `true`, empieza el contador.
- `extraTime` int - Si su valor es diferente a 0, el valor se agrega como tiempo al juego (para restar, usar números negativos).
- `setTime` int - Si su valor es diferente a 0, se establece el tiempo del juego.
- `pause` bool - Si su valor es `true`: el contador está corriendo, este se pausa; si está pausado, se resume.
- `wrong` bool - Si su valor es `true`, emite la pantalla de error a todos los clientes.
- `correct` bool - Si su valor es `true`, emite la pantalla de afirmación a todos los clientes.

### Cliente
Todos los clientes conectados al servidor reciben un JSON con los siguientes eventos:

**Nueva palabra**
```json
{ "word": "nueva palabra" }
```
- `word` string - La nueva palabra

**Tick** (el contador avanza en 1)
```json
{ "time": 1 }
```
- `time` int - El nuevo tiempo del juego

**Pantalla**
```json
{ "screen": "correct" }
```
```json
{ "screen": "wrong" }
```
- `screen` "correct" o  "wrong" - La pantalla a mostrar
