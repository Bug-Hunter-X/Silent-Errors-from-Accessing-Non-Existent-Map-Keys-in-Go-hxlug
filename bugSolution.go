```Go
package main

import (

    "fmt"
)

func main() {
    var m map[string]int
    m = map[string]int{}
    val, ok := m["a"]
    if ok {
        fmt.Println("Value found:", val)
    } else {
        fmt.Println("Key not found!")
    }
}
```