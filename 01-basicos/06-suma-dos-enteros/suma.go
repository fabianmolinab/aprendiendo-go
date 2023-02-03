// Suma dos numeros enteros
package main
import "fmt"

func main(){
  var number1 int
  var number2 int


  fmt.Print("Ingrese el primer numero: ")
  fmt.Scanf("%d", &number1)
  fmt.Print("Ingrese el segundo numero: ")
  fmt.Scanf("%d", &number2)

  var result = number1 + number2

  fmt.Println("El Resulatado es:", result)
}
