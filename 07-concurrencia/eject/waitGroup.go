package eject

import(
  "sync"
  "fmt"
)

/* 
  WaitGroup: La espera de grupo es una estructura que permite esperar a que se completen todas las goroutines antes de finalizar el programa. Para usar WaitGroup, debemos agregar goroutines al grupo antes de ejecutarlas y luego llamar a la función Done() cuando terminen. Finalmente, usamos la función Wait() para esperar a que todas las goroutines en el grupo terminen.
*/

func WaitGroup() {
    var wg sync.WaitGroup
    wg.Add(2) // agregamos dos goroutines al grupo
    go sayHello(&wg)
    go sayGoodbye(&wg)
    wg.Wait() // esperamos a que ambas goroutines terminen
}

func sayHello(wg *sync.WaitGroup) {
    defer wg.Done() // indicamos que la goroutine ha terminado
    fmt.Println("Hola")
}

func sayGoodbye(wg *sync.WaitGroup) {
    defer wg.Done() // indicamos que la goroutine ha terminado
    fmt.Println("Adiós")
}
