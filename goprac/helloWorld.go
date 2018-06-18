package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  fmt.Println("Hello World")
  /*
  fmt.Println("echo 1")
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
  */
  /*
  fmt.Println("echo 2")
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }

  fmt.Println(s)
  */
  fmt.Println()
  fmt.Println(strings.Join(os.Args[1:], " "))
  //printing os.Args[0]
  fmt.Println("the command invoked: " + os.Args[0])

  //printing the indices of os.Args
  var s string = ""
  for i := 0; i < len(os.Args); i++ {
    s += strconv.Itoa(i) + ": " + os.Args[i] + "\n"
  }
  fmt.Println(s[:len(s)-1])
}
