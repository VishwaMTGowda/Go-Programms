package main
import (
	"fmt"

	)
func main() {
	var rows int
	var k int  = 0

	fmt.Println("Enter rows")
	fmt.Scanln(&rows)

	for i:= 1; i<= rows; i++ {
		k = 0

		for space:=1; space<= rows-i; space++ {
			fmt.Print(" ")
		}
		for {
			fmt.Print("*")
			k++
			if k == 2 *i -1 {
				break
			}
			
		}
		fmt.Println("")
	} 
	
}

