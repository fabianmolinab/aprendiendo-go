package main

import "fmt"

func main(){

  a := 20
  b := 10

  result := a + b
  fmt.Println("Suma:", result)

  // resta
  result = a - b
  fmt.Println("Resta:", result)

  //Multiplicación
  result = a * b 
  fmt.Println("Multiplicación:", result)

  // Division
  result = a / b
  fmt.Println("Division:", result)

  var div float64 = 3.0 / 2.0
  fmt.Println("Division float:", div)

  //Modulo
  result = a % b 
  fmt.Println("Modulo:", result)
}
