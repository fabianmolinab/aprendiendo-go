package main

import "fmt"


func main() {
  //Slicen: es como un array que puede agregar mas elementos, son mas flexibles que un array y mas eficientes en memoria 

  //Crear un slicen
  numeros := []int{1,2,3}

  //len() se usa para determinar la longitud de un array,slice
  fmt.Println(numeros, len(numeros))

  //Modificar los elementos de un slice
  numeros[0] = 10
  fmt.Println(numeros[0])

  //Agregar datos
  numeros = append(numeros, 4)
  numeros = append(numeros, 5)
  fmt.Println(numeros, len(numeros))

  //Sub Slicen
  subSlicen := numeros[:2]
  numeros[0] = 100

  fmt.Println(subSlicen)
  fmt.Println(numeros)

}
