package main

import (
  "fmt"
  // By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
  "math/rand"
)

func main() {
  // By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
  fmt.Println("My favourite number is", rand.Intn(1000))
}
