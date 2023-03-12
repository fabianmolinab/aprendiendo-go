//Este es el algoritmo del banco

package main

import (
	"banco/utils"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//TODO: - Modularizar la aplicación

func main() {
  
  var nombre string

  //Preguntarle el nombre
  fmt.Print("Ingrese su nombre por favor: ")
  fmt.Scanln(&nombre)

  fmt.Println("Buenas", nombre)
  
  //Pregunta para iniciar la cuenta en el banco
  fmt.Println("Desea ud tener una cuenta con nuestro banco? s/n")

  //Se almacena la respuesta
  var respuesta string
  fmt.Scanln(&respuesta)

  //Condicional que registra la respuesta sobre la cuenta de banco
  if (respuesta  == "s" || respuesta == "S" || respuesta == "Si" || respuesta == "SI") {

    //Llama la función que genera el numero de cuenta
    var numeroCuentaGen int = numeroCuenta()
    fmt.Println("Su numero cuenta es: ", numeroCuentaGen)

    fmt.Println("Procesando.........")
    time.Sleep(3 * time.Second)

    //Se llama la función que imprime las opciones
    preguntaAccion()

    //Registra la respuesta del usuario
    var respuestaAccion string
    fmt.Scanln(&respuestaAccion)

    var saldo int = 0
    acciones(respuestaAccion, saldo)

  } else {
    time.Sleep(3* time.Second)
    fmt.Println("Gracias por su tiempo, hasta la proxima :D")
  }
}

func numeroCuenta() int {
  //Genera el numero aletorio cada vez que se reinicia la función
  rand.Seed(time.Now().UnixNano())

  //Generador de numero aletorio de 0 a 10000
  aletorioNum := rand.Intn(10001)

  return aletorioNum
}

func preguntaAccion() {
  fmt.Println(`Marque el numero con la acción que desea realizar:
  1. Pedir Prestamo
  2. Ingresar dinero a la cuenta
  3. Retirar dinero
  0. Salir 
  `) 
}

func acciones(respuesta string, saldo int) {
  
  switch respuesta {
  case "1":
    //Caso 1 de pedir prestamo
    fmt.Print("Ingreso a la opción de pedir prestamo \n")
    fmt.Print("Ingrese el monto que desea pedir el prestamo: \n")

    var peticionPrestamo string
    fmt.Scan(&peticionPrestamo)

    //strconv Convierte el string del monto a int
    convPeticionPrestamo, _ := strconv.Atoi(peticionPrestamo)

    //Llama a la función "Prestamo" que da la respuesta al prestamo
    utils.Prestamo(convPeticionPrestamo)

  case "2":
    //Caso 2 deposito
    utils.Deposito(saldo)

  case "3":

  default:
  }
}
