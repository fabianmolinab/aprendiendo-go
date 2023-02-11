package main

import "fmt"

func valorTotal (consumo float64, porcentajeDescuento float64)float64{

 var descuento = consumo * porcentajeDescuento 
 var montoDescuento = consumo - descuento
 var igv = montoDescuento * 0.19
 var valor = montoDescuento + igv
 return valor
}

func main(){
  /* App de restaurante
    Descuento de 10% hastas 100 de consumo
    Descuento de 20% hastas 200 de consumo
    Descuento de 30% hastas 300 de consumo
    igv 19%
  */

  var consumo float64

  fmt.Print("Ingrese consumo: ")
  fmt.Scanln(&consumo)

  if consumo <= 100 {

    valor := valorTotal(consumo, 0.10)
    fmt.Println("el valor a pagar es", valor)

  } else if consumo > 100 && consumo <= 200 {

    valor := valorTotal(consumo, 0.20)
    fmt.Println("el valor a pagar es", valor)

  } else if consumo > 200 {

    valor:= valorTotal(consumo, 0.30)
    fmt.Println("el valor a pagar es", valor)

  } else {
    fmt.Println("Error de consumo")
  }
}
