package main

import(
    "bufio"
    "fmt"
    "os"
    )

func main() {
    var f  *os.File
    f = os.Stdin
    defer f.Close()

    scanr := bufio.NewScanner(f)
    for scanr.Scan() {
        fmt.Println(" > ", scanr.Text() )
    }
}


