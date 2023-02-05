package main

import "fmt"

//En los arrays son esticos no se agregar elementos pero si modificarlos
func main() {
  var numeros [5]int
  fmt.Println(numeros)

  //Modificando el array
  numeros[0] = 5
  numeros[1] = 10
  numeros[2] = 15
  fmt.Println(numeros)

  // Imprimir definiendo algun indice dentro del array
  fmt.Println(numeros[1])

  //Array con valores
  nombres := [3]string{"Fabian", "Juan", "Ragnar"}
  fmt.Println(nombres)
  colores := [...]string{"rojo", "verde", "azul"}

  //Modificamos el indice 2 
  colores[2] = "amarillo"
  fmt.Println(colores, len(colores))

  //Array con indices definidos
  monedas := [...]string{
    0: "Dolar",
    2: "Euro",
    4: "Peso",
    6: "Ruanes",
  }
  fmt.Println(monedas, len(monedas))

  //SubArreglo
  
  /*
  Del inicio hasta el indice 3
  subMoneda := monedas[:3]
  */

  /*Del indice 3 hasta el final
  subMoneda := monedas[3:]
  */

  //con indice definido
  subMoneda := monedas[3:7]

  fmt.Println(subMoneda)

}
