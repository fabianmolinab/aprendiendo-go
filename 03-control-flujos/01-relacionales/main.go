package main

import "fmt"

func main() {
  a := 4 
  b := 5

  //Condición
  var r bool = a == b

  //fmt.Printf se usa para imprimir el resultado de la comparación y %d & %t representa un decimal y un booleando respectivamente
  fmt.Printf("%d es igual que %d -> %t \n", a, b, r)

  //Diferente
  var d bool = a != b
  fmt.Printf("%d es diferente que %d -> %t \n", a, b, d)

  //Mayor que
  var m bool = a > b
  fmt.Printf("%d es mayor que %d -> %t \n", a, b, m)

  //Menor que 
  var n bool = a < b
  fmt.Printf("%d es menor que %d -> %t \n", a, b, n)

  //Mayor o igual que
  var i bool = a >= b
  fmt.Printf("%d es mayor o igual que %d -> %t \n", a, b, i)


  //Menor o igual que
  var e bool = a <= b
  fmt.Printf("%d es menor o igual que %d -> %t \n", a, b, e)
}
