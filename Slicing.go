package main
import (
"fmt"
)
func main () {
a1 := [5]string{"Trichy", "is", "a", "Beautiful", "City"}
fmt.Printf("String : \n", a1)
Slice1 := a1[0:5]
fmt.Println("The slice is : \n", Slice1)
for i, v := range Slice1 {
fmt.Println("Index %d: ", i)
fmt.Println("Value %d: ", v)
fmt.Println()
}
}
