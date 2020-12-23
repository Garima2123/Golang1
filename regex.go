package main 
  
import ( 
    "fmt"
    "regexp"
) 
  
func main() { 
  
    str := "geeksforgeeks"
  
    
    match1, err := regexp.MatchString("geeks", str) 
    fmt.Println("Match: ", match1, " Error: ", err) 
    str2 := "ComputerScience"
    match2, err := regexp.MatchString("geeks", str2) 
    fmt.Println("Match: ", match2, "Error: ", err) 
  
   
    match3, err := regexp.MatchString("geek(s", str2) 
    fmt.Println("Match: ", match3, "Error: ", err) 
} 
