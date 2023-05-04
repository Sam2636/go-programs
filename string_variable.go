package main

import "fmt"

func main() {

	//Static Type Declaration in Go
   var x float64
   x = 20.0
   fmt.Println(x)
   fmt.Printf("x is of type %T\n", x)

   //Dynamic Type Declaration / Type Inference in Go
   var A float64 = 20.0

   B := 42 
   fmt.Println(A)
   fmt.Println(B)
   fmt.Printf("A is of type %T\n", A)
   fmt.Printf("B is of type %T\n", B)	

   //Mixed Variable Declaration in Go

   var a, b, c = 3, 4, "foo"  
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
}