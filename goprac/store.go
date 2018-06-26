package main

import (
    "bufio"
    "fmt"
    "os"
  //  "io"
    "encoding/json"
    //"strconv"
  //  "reflect"
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

func (s *Store) Exec(query string, results *[]string) {
  //parse query
  defer  func() { if p := recover(); p != nil {
        return
    }
  }()

  pos := Pos(query,' ')
  command  := query[0: pos]
  doc := query[pos + 1:]
  //fmt.Println("document: " + doc)

  switch command {
  case  "add":
      s.Add(doc)
  case "get":
    fmt.Println("getting: " + doc)
    s.Get(doc)
  default:
    fmt.Println("command does not exist")
  }
}

func (s *Store) Add(document string) {
  s.storage = append (s.storage, document)
}

func (s *Store) Get(document string) []string {
  defer  func() { if p := recover(); p != nil {
        return
    }
  }()

  res := []string{}
  var jsonObject interface{}
  json.Unmarshal([]byte(document), &jsonObject)
  docObj := jsonObject.(map[string]interface{})
  for _, value := range s.storage {
    var jsonObject interface{}
    json.Unmarshal([]byte(value), &jsonObject)
    page := jsonObject.(map[string]interface{})
    if IsMatchObjObj(page, docObj) {
        //res = append (res, document)
        fmt.Println(docObj, " is in ", page)
    } else {
        fmt.Println(docObj, " is not in ", page)
    }

  }


  return res
}

func (s *Store) Process(queries *Queue) []string {
  res := []string{}
  fmt.Println("queries.data: ", queries.data)
  for _,query := range queries.data {
    s.Exec(query, &res)
    queries.Pop()
  }
  return res
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

/*

IsMatchObjObj : check if the document is within the page when both parameters are map[string]interface{}

*/
func IsMatchObjObj(page, document interface{}) (flag bool) {
  /*
  bool, for JSON booleans
  float64, for JSON numbers
  string, for JSON strings
  []interface{}, for JSON arrays
  map[string]interface{}, for JSON objects
  nil for JSON null
*/
defer  func() { if p := recover(); p != nil {
      fmt.Errorf("Get paniced!!")
      flag = false
      return
  }
}()
  p_pg, _ := json.Marshal(page)
  p_doc, _ := json.Marshal(document)
  pg :=  fmt.Sprintf("%s",p_pg)
  doc := fmt.Sprintf("%s", p_doc)
  fmt.Println("page: ", pg)
  fmt.Println("document: ", doc)
  return
}

//2.0



func main() {
  datastore := &Store{[]string{}}
  queries := NewQueries()
  datastore.Process(queries)
  fmt.Println(datastore.storage)
  //fmt.Println(queries.data)


}
