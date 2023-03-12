package eject

import "fmt"

/*Channels
 Los canales son estructuras que permiten la comunicación entre goroutines. Un canal puede enviar y recibir valores de un tipo de dato específico. Para crear un canal en Go, usamos la función make() y especificamos el tipo de dato que queremos que maneje.
 */

func Channels() {
    ch := make(chan int) // creación de un canal de enteros
    go sendValues(ch) //goroutines
    receiveValues(ch)
}

func sendValues(ch chan<- int) { // canal de entrada
    ch <- 10
    ch <- 20
    ch <- 30
    close(ch)
}

func receiveValues(ch <-chan int) { // canal de salida
    for v := range ch {
        fmt.Println(v)
    }
}
