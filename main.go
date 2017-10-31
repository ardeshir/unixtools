package main

import (
        "fmt"
        "os"
	"flag"
	"strings"
        // "path/filepath"
)

func main() {

	/*
        args := os.Args
        pwd, err := os.Getwd()

        if err == nil {
                fmt.Println(pwd)
        } else  {
                fmt.Println("Error: ", err)
        }

        if len(args)  == 1 {
                return
        }

        if args[1] != "-P" {
                return
        }

        fileInfo, err := os.Lstat(pwd)
        if fileInfo.Mode()&os.ModeSymlink != 0 {
                realpath, err := filepath.EvalSymlinks(pwd)
                if err == nil  {
                        fmt.Println(realpath)
                }
        } */
       //  USE FLAG
	minusA := flag.Bool("a", false, "a")
	minusS := flag.Bool("s", false, "s")

	flag.Parse()
	flags := flag.Args()
	if len(flags) == 0 {
		fmt.Println("Please provide an argument")
	}
	
	file := flags[0]
	foundIt := false

	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")
	
	for _, directory := range pathSlice {
		
		fullPath := directory + "/" + file
		fileInfo, err := os.Stat(fullPath)
		
             if err == nil {
                
                mode := fileInfo.Mode()
                
		if mode.IsRegular() {
			
			if mode&0111 != 0  {
				
				foundIt = true
				
				if *minusS == true {
					os.Exit(0)
				}
				if *minusA == true {
					fmt.Println(fullPath)
				} else {
					fmt.Println(fullPath)
					os.Exit(0)
				}
			}
		}
	    }
	}

	if foundIt == false {
		os.Exit(1)
	}
} // end of main




