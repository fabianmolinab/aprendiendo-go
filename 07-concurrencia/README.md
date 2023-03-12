# Concurrencia

La concurrencia es la capacidad de un programa para realizar varias tareas simultáneamente. En Go, la concurrencia se basa en el modelo de Goroutines, que son como hilos de ejecución ligeros y administrados por el propio lenguaje.

Las Goroutines se ejecutan en el mismo espacio de dirección que la Goroutine principal (también conocida como "main Goroutine") del programa. Cada Goroutine tiene su propia pila, pero comparte la memoria y el espacio de direcciones con otras Goroutines.

Además de las Goroutines, Go también proporciona canales (channels), que son una forma segura de comunicación entre Goroutines. Los canales son un mecanismo de sincronización que permite que las Goroutines se comuniquen y coordinen entre sí de manera efectiva.

