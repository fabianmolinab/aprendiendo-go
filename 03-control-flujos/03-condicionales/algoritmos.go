//Calcular si un numero es par o impar

package main 

import "fmt"

func main() {

  var numero int

  fmt.Println("Este es un programa para cualcular si el numero es par o impar")
  fmt.Print("Ingrese el numero: ")
  fmt.Scanln(&numero)

  if numero%2 == 0 {
     fmt.Println(numero, "este numero es par")
  } else {
    fmt.Println(numero, "este numero es impar")
  }
}
