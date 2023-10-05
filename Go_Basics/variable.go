package main 
import "fmt"

/* 2 Methods to declare variable
   1.Using var keyword
   2.Using := sign

/* 3 Methods to print the Output:-
   Using Print,Printf ,Println function
 */  

func main(){
	var myfirstname string="Probeplus"
	var mylastname ="Team"
	k:=2
	fmt.Println(k)
	var  x,y bool
	fmt.Println(x,y)
	var mynum int=50
	h:="Hi!!"
	var a,b=6,"Hello!"
	var(
		c int 
		d string="HELLO!"
	)
	fmt.Print(myfirstname,mylastname)    //Prints its arguments with their default format without adding any space.
	fmt.Print(myfirstname," ",mylastname,"\n")
	fmt.Print(mynum,a,"\n")    //Print() inserts a space between the arguments if neither are strings
	fmt.Println(mynum,b)     //In Println() function a whitespace is added between the arguments, and a newline is added at the end
    fmt.Printf("c has value: %v and type: %T\n",c,c)     /* %v for value & %T for type*/
    fmt.Printf("d has value: %v and type: %T\n", d,d)
	fmt.Printf("%v\n",h)      //Prints the value in the default format
	fmt.Printf("%#v\n",h)     //Prints the value in Go-syntax format
	fmt.Printf("%v%%\n",a)    //prints % symbol

	/*For converting a type of a variable,we can use
	  type_name(expression)    */
	var sum int = 17
    var count int = 5
    var mean float32
	mean = float32(sum)/float32(count)
	fmt.Println("Mean=",mean)
}