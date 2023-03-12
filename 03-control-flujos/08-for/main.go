
package main

import "fmt"

func main() {

  /*Sintaxis: for inicializacion; condicion; post {codigo a ejecutar}

    inicializacion: es un sentencia opcional que se ejecuta al comienzo del bucle
    condicion: es la expresion booleana que se evalua en cada iteración del bucle
    post: es una sentencia opcional que se ejecuta al final de cada iteración.
    codigo a ejecutar: es el bloque de codigo que se ejecuta si la condicion es verdadera
  */

  for i := 0; i < 10; i++ {
    //Codigo se ejecutara 10 veces
  }

  //Break: Permite salir del bucle actual y continuar con la siguiente instrucción 
  for i := 0; i < 10; i++ {

    if i == 5 { //Cuando llegue a 5 se detiene
      break
    }
    fmt.Println(i)
  }

  //Continue: Permite saltar a la siguiente iteración del bucle sin ejecutar las instrucciones restantes.

  for i := 0; i < 10; i++ {
    if i%2 == 0 {
      continue
    }
    fmt.Println(i)
  }

}
