# `🚀 Go Image Optimizer: Benchmark de Concurrencia`

Este proyecto es una guía práctica para entender cómo funcionan las goroutines y los channels en Go aplicados a un caso de uso del mundo real: la optimización masiva de imágenes.

Muchas veces vemos ejemplos de concurrencia con `time.Sleep`, pero aquí ponemos a prueba los núcleos del CPU procesando imágenes reales (decodificación y compresión JPEG).

## `🛠️ ¿Qué hace este programa?`

El script compara dos formas de procesar imágenes en una carpeta:
- Modo Secuencial: Procesa una imagen a la vez (un solo núcleo trabajando).
- Worker Pool (Concurrente): Levanta un grupo de workers basado en la cantidad de CPUs de tu máquina para procesar múltiples imágenes en paralelo.

## `📈 ¿Por qué usar un Worker Pool?`

En el backend, no queremos saturar los recursos ni tampoco desperdiciarlos. Usar un Worker Pool nos permite:

    Limitar el uso de memoria: No lanzamos 10,000 procesos de golpe.

    Maximizar la CPU: Ponemos a trabajar todos los núcleos disponibles.

    Reducir tiempos: Podemos ver una mejora en el tiempo de procesamiento

## `🚀 Cómo ejecutarlo`
Dentro de la carpeta `/images` vas a ver una serie de imagenes. Podes fijarte vos mismo cuanto pesa cada una. Una vez estés preparado, podes ejecutar el script con el siguiente comando:
```shell
go run main.go
```
Luego de ejecutarlo, se habrá creado una carpeta `/output`. Esta carpeta dentro contiene las mismas imagenes pero ahora pesando un 20% menos. Además, vas a tener un log de salida que te va a indicar el tiempo de procesamiento secuencial, y el tiempo de procesamiento con múltiples workers.

