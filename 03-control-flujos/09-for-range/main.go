package main 

import "fmt"

func main() {

  nombres := [...]string{"fabian","nail","jeff"}

  //Recorre el arreglo
  for i := 0; i < len(nombres); i ++ {
    fmt.Println(nombres[i])
  }


 //Recorre el arreglo pero en este caso se coloca `_` para solo imprimir el elemento
  for _, elemento := range nombres {
    fmt.Println( elemento)
  }  
}
