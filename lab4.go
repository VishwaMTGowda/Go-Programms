package main 
import ( "fmt"  
"strings"  
)  
	func main() { 
		str := "the cat in the hat"
		 i :=  strings.Index(str, "hat") 
		fmt.Println(i)  

		var str2 string str2="The cat in the hat" 
		j:=strings.Index(str2,"in") 
		fmt.Println(j)  
}  