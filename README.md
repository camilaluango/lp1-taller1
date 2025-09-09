# Taller de Programación Concurrente en Go (Goroutines, Locks y Canales)

Este repositorio contiene **9 problemas** progresivos para que un estudiante desarrolle intuición y práctica en el uso de **Programación Concurrente** con Go.

## Cómo desarrollar el taller

1. Prepara el ambiente de desarrollo como se indicó en el Taller #0.


2. Ingresa a `https://github.com/UR-CC/lp1-taller1` y realiza un **fork** con tu cuenta.

3. Clona tu **fork** del repositorio `lp1-taller1`:

    ```bash
    git clone https://github.com/TU-USUARIO/lp1-taller1.git
    cd lp1-taller1
    ```

4. Para ejecutar cada problema utiliza:

    ```bash
    go run ./problema1
    go run ./problema2
    # ...
    go run ./problema9
    ```

5. Para cada problema, busca los comentarios `// TODO:` y sigue las indicaciones.
 
6. Para detectar condiciones de carrera, puedes ejecutar el detector de carreras:

    ```bash
    go run -race ./problema3
    ```

---

## Contenido

### Problema 1 – Goroutines y WaitGroup (Hola desde varios hilos)

Comprender la diferencia entre ejecución secuencial y concurrente, lanzar goroutines y esperar su finalización con `sync.WaitGroup`.

### Problema 2 – Temporización y `time.Sleep` (tareas simuladas)

Simular tareas que “toman tiempo” y observar beneficios de concurrencia midiendo duración total; comparar secuencial vs concurrente.

### Problema 3 – Contador compartido (condición de carrera, `Mutex`/`atomic`)

Provocar una condición de carrera incrementando un contador desde múltiples goroutines y luego **arreglarla** usando `sync.Mutex` o `sync/atomic`.

### Problema 4 – Productor–Consumidor con canales

Usar un canal para coordinar un productor y múltiples consumidores, evitando *busy waiting* y practicando cierre de canales.

### Problema 5 – Lectores–Escritores con `RWMutex`

Uermitir múltiples lecturas concurrentes pero escritura exclusiva sobre una “base de datos” (mapa) compartida. Practicar `RLock/RUnlock`.

### Problema 6 – Deadlock (y cómo prevenirlo)

Reproducir un interbloqueo entre dos mutex para entender su causa y luego **prevenirlo** imponiendo un orden global de adquisición de recursos.

### Problema 7 – Pool de workers (trabajos y resultados)

Crear un *pool* de goroutines “workers” que procesen trabajos de una cola (canal) y devuelvan resultados. Comparar vs lanzar goroutines ad‑hoc.

### Problema 8 – “Futuros” con canales (recolectar resultados asíncronos)

Construir funciones que devuelven un canal como “futuro” del resultado. Lanzar varias tareas y recolectar sus resultados.

### Problema 9 – Filósofos Comensales (problema clásico)

Modelar el problema clásico con mutex/canales; comprender cómo evitar deadlock (p. ej., orden global de recursos o limitar concurrencia).

---

## Recomendaciones

- Para cada problema, empieza por leer todoos comentarios `// TODO:`, antes de empezar a modificar el código.
- Itera: primero haz que funcione, luego mide y mejora.
- Usa `-race` para revelar errores sutiles.
- Agrega *logs* con `fmt.Printf` para entender el orden de ejecución.
