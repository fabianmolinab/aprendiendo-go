package main

import (
	"errors"
	"fmt"
)

func division(dividiendo, divisor int) (int, error) {
  if divisor == 0 {
    return 0, errors.New("No es posible dividir entre 0")
  } else {
    return dividiendo / divisor, nil // Se pasa nil cuando no hay error
  }
}

func main() {
  var result, error = division(4, 2)

  if error == nil {
    fmt.Println("Resultado: ", result)
  } else {
    fmt.Println(error)
  }
}
