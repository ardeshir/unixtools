package main

import (
   // "fmt"
   // "flag"
   "os"
   "log"
   u "github.com/ardeshir/version"
)

var ( 
 debug bool = false
 version string = "0.0.0"
 )

func main() {


   // --  FLAGS -----------------------//
   /*
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
   */

   // ========================================//
   //  Create file with os.Create()
   createFile()


  // -----------------  footer ----------- //    
  if debugTrue() {
    u.V(version)
  }

}



func createFile() {
    f, err := os.Create("testfile1.txt")
    u.ErrNil(err, "Unable to create file")
    defer f.Close()
    log.Printf("Created %s\n", f.Name())
}




// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {
    
     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }  
     return false 
}