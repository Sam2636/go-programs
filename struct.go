package main

import "fmt"

// Define a struct type
type Person struct {
    name string
    age  int
}

// Define a method on the Person struct
func (p Person) sayHello() {
    fmt.Printf("Hello, my name is %s and I'm %d years old\n", p.name, p.age)
}

func main() {
    // Create a new Person struct
    john := Person{name: "John", age: 30}

    // Call the sayHello method on the john object
    john.sayHello()
}
