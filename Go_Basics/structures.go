package main
import  "fmt"

/* Syntax for creating Structure:-

   type struct_name struct {
     member1 datatype;
     member2 datatype;
     ...
   }
*/

type Student struct {
	name string;
	class int;
	rollno int;
	present bool;
}

func main(){
   var stud1 Student                   //declaring variable stud1 of type structure 'Student'
   var stud2 Student

   stud1.name="Hege"
   stud1.class= 12
   stud1.rollno= 18
   stud1.present= true
   fmt.Println("Name:", stud1.name)
   fmt.Println("Class:", stud1.class)
   fmt.Println("Rollno:", stud1.rollno)
   fmt.Println("Present:", stud1.present)

   stud2.name= "Cecilie"
   stud2.class=11
   stud2.rollno=48
   stud2.present=false
   PrintPerson(stud2)               //Passing Structure as Function Arguments

   var struct_pointer *Student         //Defining a Pointer
   struct_pointer = &stud1
   fmt.Println("Name:-",struct_pointer.name)    //Accessing the members of a structure using a pointer to that structure

}

func PrintPerson(stud Student){
	fmt.Println("Name:", stud.name)
   fmt.Println("Class:", stud.class)
   fmt.Println("Rollno:", stud.rollno)
   fmt.Println("Present:", stud.present)
}