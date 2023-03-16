package main

import "fmt"

//Función variádica -
func sumar(nombre string, numeros ...int) (string, int) {

  //Se formatea la cadena de texto con el nombre como el parametro de la funcion
  mensaje := fmt.Sprintf("La suma de %s es: ", nombre)

  var total int

  //Recorre el slice 'numeros' utilizando un 'for' y se suma todos los numeros que contiene al valor actual de 'total'
  for _, num := range numeros {
    total += num
  }

  //Retornar multiples valores
  return mensaje, total
}

func main() {

  //Se llama la funcion sumar pasando los argumento ordenados y almacenandolos en variables
  mensaje, result := sumar("Alex", 10, 20, 40, 70, 60)
  fmt.Println(mensaje, result)
}
