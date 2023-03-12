package utils

import (
  "fmt"
  "strconv"
)

func Deposito (saldo int) {
  fmt.Print("Esta en la seccion  de DEPOSITO \n")
  fmt.Print("Digite la cantidad que de desea ingresar \n")

  var cantidadDeposito string
  fmt.Scanln(&cantidadDeposito)

  convCantidadPrestamo, err := strconv.Atoi(cantidadDeposito)

  if err != nil {
    // Manejo del error
    fmt.Println("La cantidad de depósito no es un número válido")
  } else {
    // La conversión fue exitosa
     fmt.Printf("La cantidad de depósito es %d\n", convCantidadPrestamo)

     //Suma del deposito con el saldo
     nuevoSaldo := saldo + convCantidadPrestamo
     fmt.Print("Ahora el saldo de su cuenta es: ", nuevoSaldo)
  }
}
