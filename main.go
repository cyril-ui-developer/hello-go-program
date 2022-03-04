package main

import "fmt"

func main() {
    msg := sayHello("Cyril")
    fmt.Println(msg)
   // println("Hello World test!")
}

func sayHello (name string) string {
  return fmt.Sprintf("H %s", name)
}

