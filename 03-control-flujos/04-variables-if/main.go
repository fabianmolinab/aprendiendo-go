package main

import "fmt"

func main () {
  if nombre, edad := "Alex", 26; nombre == "Alex" {

    fmt.Println("Hola, ", nombre)
  } else {

    fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
  }
}
