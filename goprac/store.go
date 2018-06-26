package main

import (
    "bufio"
    "fmt"
    "os"
  //  "io"
    "encoding/json"
    //"strconv"
    "reflect"
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
  document := query[pos + 1:]
  var jsonObject interface{}
  fmt.Println("document: " + document)
  json.Unmarshal([]byte(document), &jsonObject)
  page := jsonObject.(map[string]interface{})

  switch command {
  case  "add":
      fmt.Println("adding: " + document)

      fmt.Println("------------------")
      for key, value := range page {
        fmt.Println( key, " : ",value)
      }
      fmt.Println("------------------")

      s.Add(document)
  case "get":
    fmt.Println("getting: " + document)
    s.Get(document)
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
func IsMatchObjObj(page, document map[string]interface{}) (flag bool) {
  /*
  bool, for JSON booleans
  float64, for JSON numbers
  string, for JSON strings
  []interface{}, for JSON arrays
  map[string]interface{}, for JSON objects
  nil for JSON null

*/
defer  func() { if p := recover(); p != nil {
      flag = false
      return
  }
}()

  for k1, v1 := range document {
    for k2, v2 := range page {
    //  fmt.Println(k2)
    //  fmt.Println(reflect.TypeOf(v2).Kind())
      fmt.Println(k1," ",k2, " ", v1, " ", v2)
      if k1 == k2 {
        if reflect.TypeOf(v1).Kind() == reflect.Bool || reflect.TypeOf(v1).Kind() == reflect.Float64 || reflect.TypeOf(v1).Kind() == reflect.String{
          if reflect.TypeOf(v1).Kind() == reflect.TypeOf(v2).Kind() {
              flag = (v1 == v2)
            }
          }
        } else if reflect.TypeOf(v1).Kind() == reflect.Map{
          if (reflect.TypeOf(v2).Kind() == reflect.Map) {
            flag = IsMatchObjObj(v2, v1)
          } else if (reflect.TypeOf(v2).Kind() == reflect.Slice) {
            flag = IsMatchArrObj(v2, v1)
        } else if reflect.TypeOf(v1).Kind() == reflect.Slice{
          if (reflect.TypeOf(v2).Kind() == reflect.Map) {
            flag = IsMatchObjArr(v2, v1)
          } else if (reflect.TypeOf(v2).Kind() == reflect.Slice) {
            flag = IsMatchArrArr(v2, v1)
          }
        } else {
          flag = false
        }
      }
    }
  }
  return
}

//2.0
func IsMatchObjObj(page, document map[string]interface{}) (flag bool) {
  /*
  bool, for JSON booleans
  float64, for JSON numbers
  string, for JSON strings
  []interface{}, for JSON arrays
  map[string]interface{}, for JSON objects
  nil for JSON null

*/
defer  func() { if p := recover(); p != nil {
      flag = false
      return
  }
}()

  for docK, docV := range document {
    for pageK,pageV :=  range page {
      if (k1 == k2) {
        if reflect.TypeOf(v1).Kind() == reflect.Bool || reflect.TypeOf(v1).Kind() == reflect.Float64 || reflect.TypeOf(v1).Kind() == reflect.String{
          if reflect.TypeOf(v1).Kind() == reflect.TypeOf(v2).Kind() {
              flag = (v1 == v2)
            }
         } else if reflect.TypeOf(v1).Kind() == reflect.Map{
           if (reflect.TypeOf(v2).Kind() == reflect.Map) {
             flag = IsMatchObjObj(v2, v1)
           } else if (reflect.TypeOf(v2).Kind() == reflect.Slice) {
             flag = IsMatchArrObj(v2, v1)
         } else if reflect.TypeOf(v1).Kind() == reflect.Slice{
           if (reflect.TypeOf(v2).Kind() == reflect.Map) {
             flag = IsMatchObjArr(v2, v1)
           } else if (reflect.TypeOf(v2).Kind() == reflect.Slice) {
             flag = IsMatchArrArr(v2, v1)
           }
         }
       }
     } else {

     }
   }
 }

  return
}


func main() {
  datastore := &Store{[]string{}}
  queries := NewQueries()
  datastore.Process(queries)
  fmt.Println(datastore.storage)
  //fmt.Println(queries.data)


}
