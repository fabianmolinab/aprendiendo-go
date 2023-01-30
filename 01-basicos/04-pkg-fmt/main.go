package main
import "fmt"

func main() {
  
  var cumple = 27

  // Print: Imprime sin saltos de linea
  fmt.Print(cumple)
  fmt.Print(cumple, "latimosamente esa es mi edad")
  
  // Println: Imprime en una sola linea
  fmt.Println(cumple)

  var nombre string = "Fabian"
  var apellido string = "Molina"
  var edad int= 28 

  // Imprime un mensaje formateado
  fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)
  fmt.Printf("Hola, %v y su apellido es %v \n", nombre, apellido)

  // Guarda el mensaje formateado en la variable
  mensaje := fmt.Sprintf("Hola, %v y su edad es %v", nombre, edad)
  fmt.Println(mensaje)

  //Con el carater %T nos imprime el tipo de dato de la variable
  fmt.Printf("edad: %T \n", edad) // edad: int

  //Scanln: permite almacenar una variable ingresada por consola 
  fmt.Println("Ingrese otro nombre: ")
  fmt.Scanln(&nombre)

  fmt.Println("otro nombre: ", nombre)

}
