package main
import ("fmt")

/* 3 ways to declare slice:-
   1.Using slice_name :=[]dataype{values}
   2.Slicing from an array
   3.Using make() function  :- eg:- slice_name := make([]type, length, capacity) 
   */
func main(){
   myslice1 := []string{"Go", "Slices", "Are", "Powerful"}   // 1st method
   fmt.Println("Length=",len(myslice1))       // Returns the number of elements in the slice
   fmt.Println("Capacity=",cap(myslice1))     // Returns the capacity of the slice (the number of elements the slice can grow or shrink to)
   fmt.Println(myslice1,"\n")


   arr1 := [6]int{10, 11, 12, 13, 14,15}   //2nd method
   myslice := arr1[2:4]

   /* Here 'myslice' is a slice with length 2 which is made from arr1 of length 6.
   The slice starts from 2nd element of the array which has value 12 and slice can grow to the end of the array.So capacity of the slice is 4. */
   fmt.Printf("myslice = %v\n", myslice)    
   fmt.Printf("length = %d\n", len(myslice))  
   fmt.Printf("capacity = %d\n", cap(myslice)) //here capacity will be 4
   newslice := arr1[0:4]
   fmt.Printf("capacity = %d\n\n", cap(newslice))   //'newslice' started from element 0,then capacity would be 6.
   
 
   myslice2 := make([]int, 5, 10)        //3rd method        * 5 denotes length,10 denotes capacity 
   fmt.Println("Slice=",myslice2)               //Slice will be having default values
   fmt.Printf("length = %d\n", len(myslice2))  
   fmt.Printf("capacity = %d\n", cap(myslice2))


   /*Can access,change ,append elements in a  Slice */
   fmt.Println(newslice[2])   //Accessing element in Slice
   newslice[2]= 40           //Changing element


   oldslc := []int{1, 2, 3, 4, 5, 6}
   fmt.Printf("old = %v\n", oldslc)
   fmt.Printf("length = %d\n", len(oldslc))
   fmt.Printf("capacity = %d\n", cap(oldslc))
   newslc := append(oldslc, 20, 21)
   fmt.Printf("new = %v\n", newslc)
   fmt.Printf("length = %d\n", len(newslc))
   fmt.Printf("capacity = %d\n", cap(newslc)) 
   /*From above we can see that intially capacity of 'old' is 6,but after appending its capacity changes to 12  */
   
   /* Can append one slice to another 
    slice3 = append(slice1, slice2...) */
	slice3:= append(newslice,myslice...)
	fmt.Println("After appending two slices=",slice3)
	slice3 = append(slice3, 20, 21, 22, 23)          // Can change length of a slice by appending items
	fmt.Printf("Slice after modifying = %v\n", slice3)
    fmt.Printf("length = %d\n", len(slice3))
    fmt.Printf("capacity = %d\n", cap(slice3))
}