package main

import "fmt"

// Una función es una función sin un nombre en especifico.
func main()  {
  
  sum := func(a,b int)int {
    return a + b 
  }

  result := sum(3, 4)
  fmt.Println(result)
}
