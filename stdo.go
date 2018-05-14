package main

import (
       //  "fmt"
        "os"
	// "flag"
	// "strings"
	"io"
        // "path/filepath"
)

func main() {

        myString := ""
        args := os.Args
        // pwd, err := os.Getwd()


        if len(args)  == 1 {
                myString = "Please give me one arg"
        } else {
        	myString = args[1]
        }


       io.WriteString(os.Stdout, myString)
       io.WriteString(os.Stdout, "\n")

} // end of main


