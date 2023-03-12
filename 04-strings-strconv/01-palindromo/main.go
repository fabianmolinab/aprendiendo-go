package main

import (
       "strings" 
       "fmt"
       )


func reemplazar(palabra string) {
  
  /*
  Replace: Reemplaza todas las ocurrencias de una subcadena en una cadena dada con otra subcadena especificada
   Sintaxis: func(s, old, new string, n int) string

   s: Es la cadena destino
   old: es la cadena de caracteres que se quiere reemplazar
   new: es la nueva cadena de caracteres que reemplazar a la cadena `old`
   n: es el numero de veces que se quiere realizar el reemplazo.
  */

  palabra = strings.Replace(palabra, "z", "s",2)
  fmt.Println(palabra) // Lus asul
}

func esPalindromo(palabra string) {
  
  //ToLower: convierte cualquier string a minuscula
  palabra = strings.ToLower(palabra)


  //ToUpper: convierte cualquier string a mayuscula
  var palabraM string = strings.ToUpper(palabra) //LUZ AZUL

  fmt.Print(palabraM)
}

func main() {
  //esPalindromo("Luz azul")
  reemplazar("Luz azul")
}
