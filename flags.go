package main

import (
   "fmt"
   "flag"
   "os"
   "strconv"
   u "github.com/ardeshir/version"
)

var ( 
    debug bool = false
    version string = "0.0.1"
    )

func main() {

   name := flag.String("name", defaultName() , "name of the gopher string flag")
   age  := flag.Int("age", defaultAge(), "Age of gopher int flag")
   
   debug := flag.Bool("d", defaultDebug(), "The d bool flag")
   minC := flag.Bool("c", true, "The c bool flag")
   minK := flag.Int("k", 3, "The k int flag")

   flag.Parse()

   fmt.Println("\n------- Old Flags ------\n")
   fmt.Println("-d:", *debug)
   fmt.Println("-c:", *minC)
   fmt.Println("-k:", *minK)
   fmt.Println("\n------- Values of old flags ------")
   fmt.Printf("\n -c: %t\n -d: %t\n -k: %d\n", *minC, *debug, *minK)
   
   fmt.Println("\n------- New Flags ------\n")
   fmt.Println("name: ", *name)
   fmt.Println("age: ", *age)

   for index, val := range flag.Args() {
      fmt.Println(index, ":" , val)
   }

  // -----------------  footer ----------- //    
  if *debug   {
     u.V(version)
  }

}

func defaultName() string {
    if os.Getenv("DEFAULT_NAME") != "" {
        return os.Getenv("DEFAULT_NAME")
    }
    return "Programmer"
}

func defaultAge() int {
    if os.Getenv("DEFAULT_AGE") != "" {
        age, err := strconv.Atoi(os.Getenv("DEFAULT_AGE"))
        if err == nil {
            return age
        }
    }
    return 27
}

func defaultDebug() bool {
    
    if os.Getenv("DEFAULT_DEBUG") != "" {
        debug, err := strconv.ParseBool(os.Getenv("DEFAULT_DEBUG") )
        
        if err == nil {
            return debug
        }
        fmt.Printf("DEBUG value: %v\n", debug)
    }
    
    
    fmt.Printf("DEBUG value: %v\n", debug)
    return false
}
