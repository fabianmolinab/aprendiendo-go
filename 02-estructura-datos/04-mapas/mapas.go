package main

import "fmt"

func main(){

  /* Los mapas son un tipo de datos que permite almacenar pares de clave-valor
    mapa := make(map[tipoDeClave]tipoDeValor)
  */
  //Definiendo un mapa
  dias := make(map[int]string)

  //Agregar datos
  dias[1] = "Lunes"
  dias[2] = "Martes"
  dias[3] = "Miercoles"
  dias[4] = "Jueves"
  dias[5] = "Viernes"
  dias[6] = "Sabado"
  dias[7] = "Domingo"
  fmt.Println(dias)

  //Modificar datos
  dias[7] = "Sabado"
  fmt.Println(dias)

  //Eliminar datos
  delete(dias,1)
  fmt.Println(dias)

  //Saber la longitud del mapa
  fmt.Println(dias, len(dias))

  //Pasando de mapa a slice
  estudiantes := make(map[string][]int)

  estudiantes["Fabian"] = []int{13, 15, 16}

  //Acceder a un elemento del slice
  fmt.Println(estudiantes["Fabian"][2])
}
