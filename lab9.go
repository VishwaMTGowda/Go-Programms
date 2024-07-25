package main
import "fmt"

func linearSearch(key int, arr [] int)bool {
	for _, item:=range arr {
		
	   if item==key {
	     return true
} 
}
	return false
}
func main() {
	items:=[]int{95,78,56,84,25,35,15,26}
	fmt.Println(linearSearch(100,items))
}
