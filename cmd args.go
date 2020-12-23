package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    argLength := len(os.Args[1:])
        // use fmt.Printf() to format string
    fmt.Printf("Arg length is %d", argLength) 
}
