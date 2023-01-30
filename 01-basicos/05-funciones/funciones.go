package main

import "fmt"
import "strconv"

func saludar (nombre string, edad int) string {

  saludo := "Hola soy " + nombre + " tengo " + strconv.Itoa(edad) + " y tengo años"

  return saludo
}

func main()  {

  encuentro := saludar("Fabian", 27)
  fmt.Println(encuentro)
}
