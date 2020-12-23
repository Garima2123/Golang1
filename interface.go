package main
 
import (
    "fmt"
)
 
func checkType(i interface{}) {
    switch i.(type) {          // the switch uses the type of the interface
    case int:
        fmt.Println("Int")
    case string:
        fmt.Println("String")
    default:
        fmt.Println("Other")
    }
}
 
func main() {
    var i interface{} = "A string"
     
    checkType(i)   // String
}
