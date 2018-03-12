package main

import (
   // "fmt"
   "io"
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

    filename    := flag.String("file", defaultFile(), "Name of file to play with")
    text        := flag.String("text", "This will be printed", "something goes here")
    // newfilename := flag.String("newfile", defaultFile2(), "Name of file to play with")
    flag.Parse()
    

     // createFile(*filename)
      deleteFile(*filename) 
     // checkExistence(*filename) 
     // renameFile(*filename, "newfiletest.txt")
     // copyFile(*filename, *newfilename)
     // deleteFile(*newfilename)
     writeToFile(*filename, *text)
      
  if debugTrue() {
    u.V(version)
  }

}
func writeToFile(filename string, text string) {
    f, err := os.Create(filename)
    u.ErrNil(err, "Can't create new file")
    
    defer f.Close()
    
    if _, err := f.Write([]byte(text)); err != nil {
        log.Fatalln(err)
    }
    
    log.Println("Done\n")
}

// Function to copy a file
func copyFile(filename string, newfilename string) {
    of, err := os.Open(filename)
    u.ErrNil(err, "Can't Open old file")
    defer of.Close()
    
    nf, err2 := os.Create(newfilename)
    u.ErrNil(err2, "Can't create new file")
    defer nf.Close()
    
    bw, err3 := io.Copy(nf, of)
    u.ErrNil(err3, "Can't copy from old to new")
    log.Printf("Bytes written: %d\n", bw)
    
    if err4 := nf.Sync(); err4 != nil {
        log.Fatalln(err4)
    }
    log.Printf("Done copying from %s to %s\n", filename, newfilename)

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
    return "testfile.txt"
}

func defaultFile2() string {
    if os.Getenv("DEFAULT_FILE2") != "" {
        return os.Getenv("DEFAULT_FILE2")
    }
    return "newtestfile.txt"
}