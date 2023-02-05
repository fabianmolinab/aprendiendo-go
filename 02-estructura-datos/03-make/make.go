package main

import "fmt"

func main() {

  /*make es una función que se utiliza para crear slices, mapas y canales
    sintaxis --> make(type, length, capacity)
    type: es el tipo de datos que se quiere crear.
    lengh: es la longitud inicial
    capacity: es la capacidad del slice, que es opcional.
  */

  numeros := make([]int,3,3)

  numeros[0] = 100
  numeros[1] = 200
  numeros[2] = 300
  //numeros[3] = 400

  fmt.Println(numeros)
}
