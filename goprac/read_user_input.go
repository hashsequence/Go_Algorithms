package main

import (
    "bufio"
    "fmt"
    "os"
)

type Queue struct {
  data []string
  size int
  back int
  front int
}

func NewQueue() *Queue {
  return &Queue{ []string{},0,  0,  0 }
}

func (q* Queue) Add(value string) {
  q.data = append(q.data, value)
  q.size = len(q.data)
  q.back = q.size - 1
}

func (q* Queue) Pop() (string, bool) {
  temp := q.data[0]
  if q.size > 0 {
    q.data = q.data[1:]
    q.front = 0
    q.size = len(q.data)
    return temp,true
  }
  return "", false
}

func (q* Queue) IsEmpty() bool {
  return q.size == 0
}

func NewQueries() *Queue {
  queries := NewQueue()
  for {
     reader := bufio.NewReader(os.Stdin)
     //fmt.Print("Enter text: ")
     text, _ := reader.ReadString('\n')
     if text == "END\n" {
       break
     }
    // fmt.Println(text)
     queries.Add(text)
  }
  return queries
}

func main() {
  queries := NewQueries()
  fmt.Println(queries.data)
}
