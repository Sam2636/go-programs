package main

import "fmt"

func main() {
    // Declare a map with string keys and int values
    m := make(map[string]int)

    // Add some key-value pairs to the map
    m["one"] = 1
    m["two"] = 2

    // Access values in the map
    fmt.Println(m["one"])
    fmt.Println(m["two"])

    // Delete a key-value pair from the map
    delete(m, "two")

    // Check if a key exists in the map
    if val, ok := m["two"]; ok {
        fmt.Println("two is in the map with value", val)
    } else {
        fmt.Println("two is not in the map")
    }
}
