package main
import "fmt"

func swap(x, y int) (int, int) {         //Here first () denotes function arguments, second () denotes return datatypes
   return y, x
}

//Recursive Function Example
func testcount(x int) int {
	if x==11{
		return 0
	}
	fmt.Println(x)
	return testcount(x+1)
}


func main() {
	a, b := 45,23
   fmt.Println("Before Swapping:",a,b)
   c,d:=swap(a,b)                              //Calling Swap Function
   fmt.Println("After Swapping",c, d)

   testcount(1)      //Calling Recursive Function
}

