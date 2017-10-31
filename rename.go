package main

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"
)

func main() {

	minusOv := flag.Bool("overwrite", false, "overwrite")

	flag.Parse()
	flags := flag.Args()

	if len(flags) < 2 {
		fmt.Println("please provide two arguments!")
		os.Exit(1)
	}

	source := flags[0]
	dest := flags[1]
	fInfo, err := os.Stat(source)

	if err == nil {
		mode := fInfo.Mode()
		if mode.IsRegular() == false {
			fmt.Println("Sorry, we only support regular files as source!")
			os.Exit(1)
		}
	} else {
		fmt.Println("Error reading:", source)
		os.Exit(1)
	}

	newDest := dest
	destInfo, err := os.Stat(dest)
	if err == nil {
		mode := destInfo.Mode()
		if mode.IsDir() {
			justName := filepath.Base(source)
			newDest = dest + "/" + justName
		}
	}

	dest = newDest
	destInfo, err = os.Stat(dest)

	if err == nil {
		if *minusOv == false {
			fmt.Println("Destination file already exists!")
			os.Exit(1)
		}
	}

	err = os.Rename(source, dest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
