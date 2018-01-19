package main

import (
   // "fmt"
   "flag"
   "os"
   "log"
   u "github.com/ardeshir/version"
)

var ( 
 debug bool = false
 version string = "0.0.0"
 )

func main() {

    filename := flag.String("file", defaultFile(), "Name of file to play with")
    flag.Parse()
    

     createFile(*filename)
    // deleteFile(*filename) 
    // checkExistence(*filename) 
     renameFile(*filename, "newfiletest.txt")
     deleteFile("newfiletest.txt")
  
  if debugTrue() {
    u.V(version)
  }

}
// Function to rename a file
func renameFile ( filename string, newname string) {
        checkExistence(filename)
        err := os.Rename(filename, newname)
        u.ErrNil(err, "File was corrupted")
        
    fi, err2 := os.Stat(newname)
    if err2 != nil {
        if os.IsNotExist(err2) {
           log.Fatalf("File: %s does not exist", fi.Name)
        } 
    } 
    
    log.Printf("Exists, last modified %v\n", fi.ModTime())
    // deleteFile(newname)
}


// Function to create a file 
func createFile( filename string) {
    f, err := os.Create(filename)
    u.ErrNil(err, "Unable to create file")
    defer f.Close()
    log.Printf("Created %s\n", f.Name())
}

// Function to delete a file 
func deleteFile( filename string) {
     // createFile()  // use the createFile first, then delete file
     checkExistence(filename)
     err := os.Remove(filename)
     u.ErrNil(err, "Unable to remove testfile1.txt")
     log.Printf("Deleted %s", filename)
}

func checkExistence( filename string) {
    // createFile()
    fi, err := os.Stat(filename)
    if err != nil {
        if os.IsNotExist(err) {
           log.Fatalf("File: %s does not exist", fi.Name)
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

// Function to check env variable DEFAULT_FILE to get
func defaultFile() string {
    if os.Getenv("DEFAULT_FILE") != "" {
        return os.Getenv("DEFAULT_FILE")
    }
    return "testfile1.txt"
}