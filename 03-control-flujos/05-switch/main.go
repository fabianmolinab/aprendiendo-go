package main 

import "fmt"

/*
 en Go funciona de la siguiente manera:
 1. Se evalúa una expresión y se compara con cada uno de los casos case.

 2. Si un caso es igual a la expresión, se ejecuta el código dentro de ese caso y se detiene la ejecución del switch.

 3. Si ningún caso coincide con la expresión, se ejecuta el bloque default (opcional).
*/
func main() {
  var day = "Monday"

  switch day {
  case "Monday":
    fmt.Println("Today is Monday")
  case "Tuesday":
    fmt.Println("Today is Tuesday")
  default:
    fmt.Println("Today is not Monday")
  }
}
