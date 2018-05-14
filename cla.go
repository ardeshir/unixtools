package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    
    if len(os.Args) == 1 {
        
        fmt.Println("Please give one or more floats.")
        os.Exit(1)
     }

     arg := os.Args
     min, _ := strconv.ParseFloat(arg[1], 64)
     max, _ := strconv.ParseFloat(arg[1], 64)

     for i := 2; i < len(arg); i++ {
        n,_ := strconv.ParseFloat(arg[i], 64)
        if n < min {
             min = n
        }
        if n > max {
             max = n
        }
     }
     
        fmt.Println("Min: ", min)
        fmt.Println("Max: ", max)
}



