package main
import ("fmt")

func binarySearch(num int,arr []int)bool {
	low:=0
	high:=len(arr)

	for low<=high {

		mid:=(low+high)/2

		if arr[mid]<num {
			low=mid+1
	    }else {
			high=mid-1
}

}
		if low==len(arr) || arr[low]!=num {
		return false
}
	return true
}


func main() {
		items:=[]int{1,2,9,20,31,45,63,70,100}
		fmt.Println(binarySearch(100,items))
}
