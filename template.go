package main

import (
   "fmt"
   "flag"
   u "github.com/ardeshir/version"
)

var debug bool = false
var version string = "0.0.0"

func main() {



   minusO := flag.Bool("o", false, "o")
   minusC := flag.Bool("c", false, "c")
   minusK := flag.Int("k", 0, "an int")

   flag.Parse()

   fmt.Println("-o:", *minusO)
   fmt.Println("-c:", *minusC)
   fmt.Println("-k:", *minusK)

   for index, val := range flag.Args() {
      fmt.Println(index, ":" , val)
   }

 // -----------------  footer ----------- //    
  if debug == true {
    u.V(version)
  }

}

