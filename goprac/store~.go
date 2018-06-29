package main

import (
    "bufio"
    "fmt"
    "os"
  //  "io"
    "encoding/json"
  //  "strings"
  //  "strconv"
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
  doc := query[pos + 1:]
  //fmt.Println("document: " + doc)

  switch command {
  case  "add":
      s.Add(doc)
  case "get":
//    fmt.Println("getting: " + doc)
   s.Get(doc, results)
  case "delete":
    s.Delete(doc)
  default:
    fmt.Println("command does not exist")
  }
}

func (s *Store) Add(document string) {
  fmt.Println("-------------------------------------------")
  fmt.Println("ADD | Document: ", document, "\n")
  s.storage = append (s.storage, document)
  fmt.Println("-------------------------------------------")
}

func (s *Store) Delete(document string) {
    fmt.Println("\n-------------------------------------------")
    fmt.Println("DELETE | Document: ", document, "\n")

    var jsonObject2 interface{}
    json.Unmarshal([]byte(document), &jsonObject2)
    doc := jsonObject2.(map[string]interface{})

  for _, page := range s.storage {

    var jsonObject interface{}
    json.Unmarshal([]byte(page), &jsonObject)
    pg := jsonObject.(map[string]interface{})

    if CheckIfPageContainsDoc(pg, doc) {
      fmt.Println("DELETE |  ", document, " matches the page ", page, " so deleting it\n")
    } else {
        fmt.Println("DELETE |  ", document, "doesnt matches the page ", page, " \n")
    }
  }
    fmt.Println("-------------------------------------------")
}

func (s *Store) Get(document string, results *[]string) {
  defer  func() { if p := recover(); p != nil {
        return
    }
  }()
    fmt.Println("-------------------------------------------")

    var jsonObject2 interface{}
    json.Unmarshal([]byte(document), &jsonObject2)
    doc := jsonObject2.(map[string]interface{})

  for _, page := range s.storage {

      var jsonObject interface{}
      json.Unmarshal([]byte(page), &jsonObject)
      pg := jsonObject.(map[string]interface{})


   fmt.Println("GET | DOCUMENT: ", document)
   fmt.Println("GET | PAGE: ", page)
    if CheckIfPageContainsDoc(pg, doc) {
        *results = append (*results, page)
        fmt.Println(document, " is in ", page)
    } else {
        fmt.Println(document, " is not in ", page)
    }

  }
    fmt.Println("\nGET | FINAL RESULTS | doc: ", document, " | results: ", *results, "\n")
    fmt.Println("-------------------------------------------")

}

func (s *Store) Process(queries *Queue) []string {
//  fmt.Println("queries.data: ", queries.data)
  final_res := []string{}
  for _,query := range queries.data {
    res := []string{}
    s.Exec(query, &res)
    final_res = append(final_res, res...)
  }
  return final_res
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
/*****************************************************************
CheckIfPageContainsDoc : check if the document is within the page of the storage


******************************************************************/









func CheckIfPageContainsDoc(pg, doc map[string]interface{}) (flag bool) {
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
  //pg_byte, _ := json.Marshal(page)
//  doc_byte, _ := json.Marshal(document)
  //var jsonObject interface{}
//  var jsonObject2 interface{}
  //json.Unmarshal(pg_byte, &jsonObject)
  //json.Unmarshal(doc_byte, &jsonObject2)
//  json.Unmarshal([]byte(page), &jsonObject)
//  json.Unmarshal([]byte(document), &jsonObject2)
//  pg := jsonObject.(map[string]interface{})
//  doc := jsonObject2.(map[string]interface{})
//  pg_str :=  fmt.Sprintf("%s",pg_byte)
  //doc_str := fmt.Sprintf("%s", doc_byte)
  //pg_str, _ = strconv.Unquote(pg_str)
  //doc_str, _ = strconv.Unquote(doc_str)
//fmt.Println("GET | JsonObject: ", jsonObject, "\n")
fmt.Println("GET | document: ", doc, "\n")
fmt.Println("GET | page: ", pg, "\n")

  flag = false

  OuterLoop:
  for doc_key, doc_value := range doc {
   if _, ok:= pg[doc_key]; ok && reflect.TypeOf(pg[doc_key]).Kind() == reflect.TypeOf(doc_value).Kind(){
     flag = true
     if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Bool || reflect.TypeOf(pg[doc_key]).Kind() == reflect.Float64 || reflect.TypeOf(pg[doc_key]).Kind() == reflect.String {
       fmt.Println("CHECKIFPAGECONTAINSDOC | comparing ", doc_value, " and ", pg[doc_key])
         if !reflect.DeepEqual(pg[doc_key], doc_value) {
           flag = false
           break
         }
          fmt.Println("", doc_value, " and ", pg[doc_key], " are equal")
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Map {
      // if reflect.TypeOf(doc_value).Kind() == reflect.Map {
         fmt.Println("CHECKIFPAGECONTAINSDOC |comparing ", doc_value.(map[string]interface{}), " and ", pg[doc_key].(map[string]interface{}))
         /*
            if reflect.DeepEqual(pg[doc_key].(map[string]interface{}),doc_value.(map[string]interface{})) {
              flag = false
              return
            }
            */

         if !IsSubObj(pg[doc_key].(map[string]interface{}), doc_value.(map[string]interface{})) {
           flag = false
           break
         }

       //}
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Slice {
       //if reflect.TypeOf(doc_value).Kind() == reflect.Slice {
          fmt.Println("CHECKIFPAGECONTAINSDOC | comparing ", doc_value.([]interface{}), " and ", pg[doc_key].([]interface{}))
          /*
         if !reflect.DeepEqual(pg[doc_key].([]interface{}),doc_value.([]interface{})) {
           flag = false
           return
         }
         */
         if !ArrHasSameValues(pg[doc_key].([]interface{}), doc_value.([]interface{})) {
           flag = false
           break
         }

      // }
     }
   } else {
     for _, pg_value := range pg {
       if reflect.TypeOf(pg_value).Kind() == reflect.Map{
         flag = false
         //pg_byte, _ := json.Marshal(pg_value)
         //sub_page :=  fmt.Sprintf("%s",pg_byte)
      //   _, _ = strconv.Unquote(sub_page)
        //fmt.Println("CHECKIFPAGECONTAINSDOC | error: ",err)
         fmt.Println("CHECKIFPAGECONTAINSDOC | MAP | looking in the subpage: ", pg_value , " for ", doc )
         if !CheckIfPageContainsDoc(pg_value.(map[string]interface{}), doc) {
           flag = false
         } else {
           flag = true
           break OuterLoop
         }

       } else if reflect.TypeOf(pg_value).Kind() == reflect.Slice {
         flag = false
        // _, _ = strconv.Unquote(sub_page)
        for _, sub_pg_value := range pg_value.([]interface{}) {
        //  sub_pg_byte, _ := json.Marshal(sub_pg_value)
          //sub_page :=  fmt.Sprintf("%s",sub_pg_byte)
          fmt.Println("CHECKIFPAGECONTAINSDOC | SLICE | looking in the subpage: ", sub_pg_value , " for: ", doc )
          if reflect.TypeOf(sub_pg_value).Kind() == reflect.Map {
            if !CheckIfPageContainsDoc(sub_pg_value.(map[string]interface{}), doc) {
              flag = false
            } else {
              flag = true
              break OuterLoop
            }
          } else if reflect.TypeOf(sub_pg_value).Kind() == reflect.Slice {
            fmt.Println("CHECKIFPAGECONTAINSFDOC | SLICE | checking if doc is in sub page array: ", sub_pg_value)
            if !CheckIfArrContainsDoc(sub_pg_value.([]interface{}), doc) {
              flag = false
            } else {
              flag = true
              break OuterLoop
            }
          }
        }
       }
     }
   }
  }
  return
}

/***************************************************
Compare Functions
*****************************************************/
/*
IsSubObj : checks if the document is a sub object of the page
*/
func IsSubObj (pg, doc map[string]interface{}) (flag bool) {
  defer  func() { if p := recover(); p != nil {
        fmt.Errorf("Get paniced!!")
        flag = false
        return
    }
  }()
  flag = false
 //check if page contain all the key value pairs of the doc
  for doc_key, doc_value := range doc {
   if _, ok:= pg[doc_key]; ok && reflect.TypeOf(pg[doc_key]).Kind() == reflect.TypeOf(doc_value).Kind() {
       fmt.Println("ISSUBOBJ | in MATCH| comparing ", doc, " and ", pg)
     flag = true
     if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Bool || reflect.TypeOf(pg[doc_key]).Kind() == reflect.Float64 || reflect.TypeOf(pg[doc_key]).Kind() == reflect.String {
       if pg[doc_key] != doc_value {
         flag = false
        break
      }
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Map {
        if !IsSubObj(pg[doc_key].(map[string]interface{}), doc_value.(map[string]interface{})) {
        flag = false
        break
      }
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Slice {
         if !ArrHasSameValues(pg[doc_key].([]interface{}), doc_value.([]interface{})) {
           flag = false
           break
         }
     } else {
       flag = false
     }
   } else {
          fmt.Println("ISSOBOB | in NOT MATCH| comparing ", doc, " and ", pg)
     flag =  false
   }
 }


 //check if the doc contains the page, note: uncommenting this part turns the function into a equals operation
/*
 for pg_key, pg_value := range pg {
  if _, ok:= doc[pg_key]; ok && reflect.TypeOf(doc[pg_key]).Kind() == reflect.TypeOf(pg_value).Kind() {
      fmt.Println("IsSubObj in MATCH| comparing ", pg, " and ", doc)
    if reflect.TypeOf(doc[pg_key]).Kind() == reflect.Bool || reflect.TypeOf(doc[pg_key]).Kind() == reflect.Float64 || reflect.TypeOf(doc[pg_key]).Kind() == reflect.String {
      if doc[pg_key] != pg_value {
        flag = false
       break
     }
    } else if reflect.TypeOf(doc[pg_key]).Kind() == reflect.Map {
       if !IsSubObj(doc[pg_key].(map[string]interface{}), pg_value.(map[string]interface{})) {
       flag = false
       break
     }
    } else if reflect.TypeOf(doc[pg_key]).Kind() == reflect.Slice {
        if !ArrHasSameValues(doc[pg_key].([]interface{}), pg_value.([]interface{})) {
          flag = false
          break
        }
    } else {
      flag = false
    }
  } else {
         fmt.Println("IsSubObj in NOT MATCH| comparing ", pg, " and ", doc)
    flag =  false
  }
 }
*/
 return
}

/*
checks if the two arrays have the same values
if I want to compute arrays with the same values and the same count for each element I can use a map to map the elements to the count and then
check if the two maps are equivalent
*/
func ArrHasSameValues(pg, doc []interface{}) (flag bool) {
  defer  func() { if p := recover(); p != nil {
        fmt.Errorf("Get paniced!!")
        flag = false
        return
    }
  }()
  fmt.Println("ARRHASSAMEVALUE | comparing", pg, " and ", doc)
  flag = true
  for _, doc_value := range doc {
    if !Contains(pg, doc_value) {
      flag = false
      break
    }
  }
  return
}

func Contains(s []interface{}, e interface{}) bool {
  defer  func() bool { if p := recover(); p != nil {
        fmt.Errorf("Get paniced!!")
        return false
    }
    return true
  }()
    for _, a := range s {
      switch reflect.ValueOf(e).Kind() {
      case reflect.Map:
        if reflect.TypeOf(a).Kind() == reflect.Map {
          if IsSubObj(a.(map[string]interface{}), e.(map[string]interface{})) {
            return true
          }
        }
      case reflect.Slice:
        if reflect.TypeOf(a).Kind() == reflect.Slice {
          if ArrHasSameValues(a.([]interface{}), e.([]interface{})) {
            return true
          }
        }
      default:
        if reflect.DeepEqual(a,e) {
          return true
        }
      }
    }
    return false
}

/*
IsDocInArr : checks for the doc in the array
*/

func CheckIfArrContainsDoc(pg []interface{}, doc map[string]interface{}) (flag bool) {
  defer  func() bool { if p := recover(); p != nil {
        fmt.Errorf("Get paniced!!")
          return false
      }
      return flag
    }()
  fmt.Println("CHECKIFARRCONTAINSDOC | checking the array: ",pg, " for: ", doc)
  for _, pg_value := range pg {
    if reflect.TypeOf(pg_value).Kind() == reflect.Map {
      //pg_byte, _ := json.Marshal(pg_value)
    //  page :=  fmt.Sprintf("%s",pg_byte)
    //  doc_byte, _ := json.Marshal(doc)
    //  document :=  fmt.Sprintf("%s",doc_byte)
      if !CheckIfPageContainsDoc(pg_value.(map[string]interface{}), doc) {
        flag = false
      } else {
        flag = true
        break
      }
    } else if reflect.TypeOf(pg_value).Kind() == reflect.Slice {
      if !CheckIfArrContainsDoc(pg_value.([]interface{}), doc) {
        flag = false
      } else {
        flag = true
        break
      }
    }
  }
  return
}
func main() {
  datastore := &Store{[]string{}}
  queries := NewQueries()
  res := datastore.Process(queries)
  fmt.Println("\n")
  for _, value := range datastore.storage {
      fmt.Println("STORAGE: ", value)
  }
  fmt.Println("\n")
  for _, value := range queries.data {
      fmt.Println("QUERIES: ", value)
  }

  fmt.Println("FINAL RESULTS | ", res)



}
