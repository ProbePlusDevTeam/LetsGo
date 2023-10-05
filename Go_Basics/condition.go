package main
import ("fmt")

func main() {
  //Using if..else
  time := 20
  if time < 18 {
    fmt.Println("Good day.")
  } else {                                   //The brackets in the else statement should be like  } else {
    fmt.Println("Good evening.")             
  }
  //using if..elseif..else 
  timing := 2
  if timing < 10 {
    fmt.Println("Good morning.")
  } else if timing < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }

  //Expression Switch
  //Using Single case switch 
  day := 4
  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
	fmt.Println("Not a weekday")
  }
   //Expression Switch
   //Multi Case Switch
  switch day{
   case 1,3,5:
     fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
     fmt.Println("Weekend")
  }

  //Type Switch
  var x interface{}
   x=2
   switch i := x.(type) {
      case nil:	  
         fmt.Printf("type of x :%T",i)                
      case int:	  
         fmt.Printf("x is int")                       
      case float64:
         fmt.Printf("x is float64")           
      case func(int) float64:
         fmt.Printf("x is func(int)")                      
      case bool, string:
         fmt.Printf("x is bool or string")       
      default:
         fmt.Printf("don't know the type")     
   }   
}
