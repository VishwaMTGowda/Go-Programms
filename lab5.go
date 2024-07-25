package main 
import ("fmt") 

func main() {
    intSlice := [] int{1, 2, 3, 4, 10} 
    fmt.Println("slice: ", intSlice)

    last :=  intSlice[len(intSlice)-1] 
    fmt.Println("last element: ", last)

   first :=intSlice[0]
    fmt.Println("first element: ", first) 
    
    remove := intSlice[:len(intSlice)-1] 
    fmt.Println("remove last: ", remove) 

    // intSlice1 := [] int {1,2,3,4}
    // first1:=len(intSlice1)-1
    // fmt.Println("first",first1);

}
