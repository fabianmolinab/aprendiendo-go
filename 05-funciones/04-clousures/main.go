package main

import "fmt"

/* Funciones Clousures: es una función que captura las variables definidas fuera de su ámbito léxico. 
Esto significa que la función puede acceder y modificar las variables de su entorno en el momento de su definición, así como también en el momento de su invocación.
*/

/* increment es una funcion que retorna otra función anonima que incremete una variable local cada vez que se llama.
*/

func increment() func() int {
    i := 0
    return func() int {
      i++
      return i
    }
}

func main() {
    inc := increment()

    fmt.Println(inc()) // 1
    fmt.Println(inc()) // 2
    fmt.Println(inc()) // 3
}
