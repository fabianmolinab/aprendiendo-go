package main

import "fmt"
/*
  Los operadores logicos permiten combinar valores booleanos para producir un resultado logico.

    Los operadores lógicos más comunes en Go son:
    Conjunción (&&): devuelve verdadero si ambos valores son verdaderos.
    Disyunción (||): devuelve verdadero si al menos uno de los valores es verdadero.
    Negación (!): invierte el valor lógico.
*/

func main() {
  var a bool = true
  var b bool = false

  // Usando el operador de conjunción
  fmt.Println(a && b) // devuelve false

  // Usando el operador de disyunción
  fmt.Println(a || b) // devuelve true

  // Usando el operador de negación
  fmt.Println(!a) // devuelve  a = false
}
