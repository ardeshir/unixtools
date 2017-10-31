package main

import (
   "fmt"
   "flag"
   "github.com/ardeshir/simple"
)

var debug bool = false

func main() {

  fmt.Println(simple.Add(1,2))

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

  if debug == true {
  simple.Version()
  }

}

