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
   deleteFile() 
   // checkExistence() 
   
   
  // -----------------  footer ----------- //    
  if debugTrue() {
    u.V(version)
  }

}

// Function to create a file 
func createFile() {
    f, err := os.Create("testfile1.txt")
    u.ErrNil(err, "Unable to create file")
    defer f.Close()
    log.Printf("Created %s\n", f.Name())
}

// Function to delete a file 
func deleteFile() {
     // createFile()  // use the createFile first, then delete file
     checkExistence()
     err := os.Remove("testfile1.txt")
     u.ErrNil(err, "Unable to remove testfile1.txt")
     log.Println("Deleted testfile1.txt")
}

func checkExistence() {
    // createFile()
    fi, err := os.Stat("testfile1.txt")
    if err != nil {
        if os.IsNotExist(err) {
           log.Fatalln("File: testfile1.txt does not exist")
        } 
    } 
    
    log.Printf("Exists, last modified %v\n", fi.ModTime())
    // deleteFile()
}


// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {
    
     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }  
     return false 
}