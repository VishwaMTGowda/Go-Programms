mpackage main
import "fmt"
func main() {
	var n int
	fmt.Print("enter the integer number:")
	fmt.Scanln(&n)
	i:=1

for {
if i>10 {
	break
}
	fmt.Println(n,"X",i,"=",n*i)
	i++
}
}
