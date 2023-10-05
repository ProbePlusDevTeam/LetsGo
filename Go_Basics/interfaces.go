package main

import ("fmt" 
   "math")

type Shape interface {        /* define an interface */
   area() float64
}

type Circle struct {           /* define a circle */
   x,y,radius float64
}

func(circle Circle) area() float64 {           /* define a method for circle (implementation of Shape.area())*/
   return math.Pi * circle.radius * circle.radius
}

func getArea(shape Shape) float64 {        /* define a method for shape */
   return shape.area()
}

func main() {
   circle := Circle{x:0,y:0,radius:5}
   
   fmt.Printf("Circle area: %f\n",getArea(circle))
}