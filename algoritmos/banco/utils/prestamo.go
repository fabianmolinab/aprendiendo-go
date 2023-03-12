package utils

func Prestamo( monto int ) {
  switch {
    case monto >= 1000000:
      println("Su prestamo fue aprobado")
    default: 
      println("Su prestamo no fue aprobado")
  }
}

