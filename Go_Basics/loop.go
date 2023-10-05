package main
import "fmt"

func main() {
  for i:=0; i < 50; i+=5 {
	if i==10{
		continue    //Used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
	}
	if i == 45{
	break  }       //Used to break/terminate the loop execution.
	fmt.Println(i)
  }

  //using gotostatement
  fmt.Printf("\n")
  k:=14
  LOOP: for k < 18 {
	if k == 15 {
	   /* skip the iteration */
	   k = k + 1
	   goto LOOP
	}
	fmt.Printf("value of k: %d\n", k)
	k++     
  }
   
  //Nested Loop
  fmt.Printf("\n")
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
  //The below loop uses range to iterate over an array and print both the indexes and the values at each iteration"
  fmt.Printf("\n")
  for idx, val := range fruits {
	fmt.Printf("%v\t%v\n", idx, val)
  }

  //To only show the value or the index, you can omit the other output using an underscore (_)
  //Here index is omitted
  fmt.Printf("\n")
  for _, val := range fruits {
	fmt.Printf("%v\n", val)
   }
}