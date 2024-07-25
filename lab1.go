


package main 
import "fmt" 

func main() { 
	fmt.Println("Enter number:") 
	var n int 
	fmt.Scanln(&n) 
	if n%2==0 { 
	fmt.Println(n,"is Even number") 
	}else{ 
	fmt.Println(n,"is odd number") 
} 
} 