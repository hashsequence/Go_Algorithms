/*
%8 goroutines and channel
++++++++++++++++++++++++++++++++++++++++++++++
%8.1 goroutines

each concurrently executing activity is called a goroutinge
goroutines are like threads except the difference is quantitive

when a program starts only the main gorouting main() is called

new goroutines are created by the go statements

ex.

f() // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait

when the main function returns all goroutines are aruptly terminated
and the program exits. other than by returning or exiting
the program, there is no programmatic way for one goroutine to stop another,
but there are ways for goroutines to request that it stop itself

ex. sample concurrent program:

func main() {
  go spinner(100 * time.Millisecond)
  const n = 45
  fibN := fib(n) // slow
  fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)


func spinner(delay time.Duration) {
  for {
    for _, r := range `-\|/` {
      fmt.Printf("\r%c", r)
      time.Sleep(delay)
    }
  }
}
func fib(x int) int {
  if x < 2 {
  return x
  }
  return fib(x-1) + fib(x-2)
}

%8.2 concurrent clock















*/

package main

import (
  "io"
  "log"
  "net"
  "time"
)

func main() {
  listerner, err := net.Listen("tcp", "localhost:8000")
  if err != nil {
    log.Fatal(err)
  }
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Print(err) // e.g., connection aborted
      continue
    }
    handleConn(conn) //handle one connection at a time
  }
}

func handleConn(c .net.Conn) {
  defer c.Close()
  for {
    _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
    if err != nil {
      return // eg. client disconnected
    }
    time.Sleep(1*time.Second)
  }
}
