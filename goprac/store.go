package main

import (
    "bufio"
    "fmt"
    "os"
  //  "io"
    "encoding/json"
    //"strconv"
    //"reflect"
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


  scanner := bufio.NewScanner(os.Stdin)
   for scanner.Scan() {
       queries.Add(scanner.Text())
   }
/*
  for {
      reader := bufio.NewReader(os.Stdin)
     //fmt.Print("Enter text: ")
     text, err := reader.ReadString('\n')
     if err != nil {
        if err == io.EOF {
            break
        }
    }

    // fmt.Println(text)
     queries.Add(text)
  }
  */
  return queries
}

type Store struct{
  storage []string
}

func (s *Store) Exec(query string, results *Store) {
  //parse query
  defer  func() { if p := recover(); p != nil {
        return
    }
  }()

  pos := Pos(query,' ')
  command  := query[0: pos]
  document := query[pos + 1:]
  var jsonObject interface{}
  fmt.Println("document: " + document)
  json.Unmarshal([]byte(document), &jsonObject)
  obj := jsonObject.(map[string]interface{})
  switch command {
  case  "add":
      fmt.Println("adding: " + document)

      fmt.Println("------------------")

      for key, value := range obj {
        fmt.Println( key, " : ",value)
      }

      //fmt.Println(obj["fruits"])
      //fmt.Println(obj["vegetables"])
      fmt.Println("------------------")
      s.Add(document)
  case "get":
    fmt.Println("getting: " + document)
    s.Get(obj, document)
  default:
    fmt.Println("command does not exist")
  }
}

func (s *Store) Add(document string) {
  s.storage = append (s.storage, document)
}

func (s *Store) Get(obj map[string]interface{}, document string) []string {
  res := []string{}
  var jsonObject interface{}
  json.Unmarshal([]byte(document), &jsonObject)
  docObj := jsonObject.(map[string]interface{})
  if IsMatchObj(obj, docObj) {
      //res = append (res, document)
  }
  return res
}

func (s *Store) Process(queries *Queue) []string {
  results := Store{[]string{}}
  fmt.Println("queries.data: ", queries.data)
  for _,query := range queries.data {
    s.Exec(query, &results)
    queries.Pop()
  }
  return results.storage
}
/*******************************************
helper functions
*********************************************/

/*
  return index of first occurence of <value> in the string
*/
func Pos(s string, value rune) int {
    for k, v := range s {
        if (v == value) {
            return k
        }
    }
    return -1
}

/*
converts bool to string
*/


func IsMatchObj(obj map[string]interface{}, document map[string]interface{}) (bool) {
  /*
  bool, for JSON booleans
  float64, for JSON numbers
  string, for JSON strings
  []interface{}, for JSON arrays
  map[string]interface{}, for JSON objects
  nil for JSON null
*/
    return true
}

func main() {
  datastore := &Store{[]string{}}
  queries := NewQueries()
  datastore.Process(queries)
  fmt.Println(datastore.storage)
  //fmt.Println(queries.data)


}
