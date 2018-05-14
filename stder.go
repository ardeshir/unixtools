package main

import (
       //  "fmt"
        "os"
	// "flag"
	// "strings"
	"io"
        // "path/filepath"
)


// redirect the file descriptor of standard error(2) to the standard output (1)
//  go run stder.go > /tmp/output 2>&1 

func main() {

        myString := ""
        args := os.Args
        // pwd, err := os.Getwd()


        if len(args)  == 1 {
                myString = "Please give me one arg"
        } else {
        	myString = args[1]
        }
       
       io.WriteString(os.Stdout, "This is Standard output\n")
       io.WriteString(os.Stderr, myString)
       io.WriteString(os.Stderr, "\n")

} // end of main


