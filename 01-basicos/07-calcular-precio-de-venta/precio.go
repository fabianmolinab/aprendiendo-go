package main
import "fmt"

 func impuesto(venta float64) {
  var porcentaje float64 = 0.18
  valorIGV := venta * porcentaje

  valorTotal := valorIGV + venta
  fmt.Println("El impuesto a su compra es:", valorIGV)
  fmt.Println("El valor a pagar es:", valorTotal)
}

func main() {
  var venta float64
  fmt.Print("Ingrese el valor de la venta: ")
  fmt.Scanf("%f", &venta)
  impuesto(venta)
}
