package main
import ("fmt")


/* We can declare array using 2 methods:
 1.With var keyword
 2.With := sign  */

func main() {
  var arr1 = [3]int{1,2,3}    //Intialising an array of size 3
  arr2 := [5]int{4,5,6,7,8}     //Intialising an array of size 5
  var arr3 = [...]int{1,2,3}      //Intialising an array without defining size

  fmt.Println("Length:",len(arr1),"\n","Array:",arr1)  //len function give the length of the array
  arr2[2]=10                 //We can change the values of the elements in an array
  fmt.Println(arr2)
  fmt.Println(arr3)

  fmt.Println("First Element=",arr1[0])   //Prints the first element of the arr1

  num := [5]int{1:10,2:40}  // 1:10 means: assign 10 to array index 1 (second element).
                           //   2:40 means: assign 40 to array index 2 (third element).
  fmt.Println(num)         //num= [0 10 40 0 0]

  var  balance = []int {1000, 2, 3, 17, 50}
  var avg float32
  avg = getAverage( balance, 5 ) ;       /* pass array as an argument */
  fmt.Printf( "Average value is: %f\n ", avg);
  twodimensionarrray()
}


 func getAverage(arr []int, size int) float32 {      //Example on how to pass an array to a function
	var i,sum int
	var avg float32  
 
	for i = 0; i < size;i++ {
	   sum += arr[i]
	}
	avg = float32(sum / size)
	return avg;
 }

func twodimensionarrray(){
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}   //Defining two dimensional array
	var i, j int

	for  i = 0; i < 5; i++ {             	/* output each array element's value */
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
		}
	}
}