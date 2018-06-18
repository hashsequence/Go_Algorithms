/*
%2
++++++++++++++++++++++++++++++
%2.3 variables
you can declare multiple variables in a tuple manner
var a, b, c = true, 3, 2

var name type = expression

if types is omitted then initial value determines type, otherwise
if we have the type then it is zero intialized

we always have either type or expression when intializing variables

%2.3.1 short variable declarations
this is the short variable declaration

:=

ex. freq := rand.Float64() * 3.0
t := 0.0

note: a short variable declaration does not neccesarily declare
all variables in the left hand side

ex.

in, err := os.Open(infile)
out, err := os.Open(outfile)

first statement declares in and err, and second statement declares out and
assigns a value to errors

note: you must declare at least one new variable using short decalartion
also a short declarations acts like an assignment only to variables declared
in the same lexical block

%2.3.2 pointers

var x int = 7
p := &x

p points to a reference to x

var p = f()

func f() *int {
  v := 1
  return &v
}

the above will work because since p recieves
the address, it will remain in existence

%2.7
universal blocks - lecical block for entire
source code
this has an error:

if f, err := os.Open(fname); err != nil { // compile error: unused: f
return err
}
f.ReadByte() // compile error: undefined f
f.Close() // compile error: undefined f

this doesn't:

if f, err := os.Open(fname); err != nil {
return err
} else {
// f and err are visible here too
f.ReadByte()
f.Close()
}

basically f, and err is visible within then
the same if, else block, so scope is within the if-else blocks
not just one brace of of the case statements, same for switches

%3
++++++++++++++++++++++++++++++++++++++++++++++++
%3.1 integers

The hig h-order bits that do not fit are
si lently dis carde d. If the original number isasig ned typ e, the result could be negat ive if the
lef tmost bit is a 1, as in the int8 example here:
var u uint8 = 255
fmt.Println(u, u+1, u*u) // "255 0 1"
var i int8 = 127
fmt.Println(i, i+1, i*i) // "127 -128 1"

go has AND NOT &^

The &^ op erator is bit cle ar (AND NOT): in the expression
z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1; otherwise it equals the
corresponding bit of x.

Arithmetically, a left
shif t x<<n is equivalent to multiplication by 2^n and a right shift x>>n is
equivalent to the floor of division by 2^n


left shifts fill vacated bits with zeroes but right shifts all the copies of the
sign bit to vacate bits

for type conversions you must convert them explicitly

%3.2 floating point numbers
%3.3 complex numbers
var x = complex64(3,4) //components is float32
var y = complex128(3,4) //components is float64

%3.4 booleans

%3.5 strings

substrings s[i:j]
[i,j)

concatenate with +

strings are immutable in go

Rune is a Type. It occupies 32bit and
is meant to represent a Unicode CodePoint


for example let say you want to process a string
literal with japanese character
use

r, size := utf8.DecodeRuneInString(s)

this returns the size and the encode string to r

Go's range loop decodes UTF8 strings naturally so it
indexes correctly since each character in utf-8 string could
be multibyte

index, value := range someSlice

or

index := range someSlice

%3.5.4 strings and byte slices
strings are immutable so use byte slices

ex.

s := "abc"
b := []byte(s)
s2 := string(b)

compiler might optmize and reference b to s, but usually
allocates memory for

%3.5.5 conversions between strings and numbers

integer to ascii
x := 123
y := fmt.sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x)) // "123 123"

strcov.Atoi converts from string to int8
strconv.ParseInt

%3.6 constants

the underyling type for constants are boleans, strings, or a number


as wirth variables, a sequence pf constqants can appear in one declaration, this would
be appropriate for a group of related values.

const (
e = 2.7182
pi = 3.1415
)

%3.6.1 the constant generator iota

used to create a sequence of related values without spelling
out each one explicitly, in a const declaration, the valuesof iota begins at zero and increments
by one for eah item in the sequence

type Weekday int8
const (
  Sunday Weekday = iota
  Monday
  Tuesday
  Wednesday
  Thurday
  Friday
  Saturday

)

this is basically the enumerations or enums in go func(
sunday is 0, monday is 1...

you can also do more complex expressions

type Flags uint

const (
  FlagUp Flags = 1 << iota
  FlagBroadcast
  FlagLoopback
  FlagPointToPoint
  FlagMulticast
)

const (
_ = 1 << (10 * iota)
KiB // 1024
MiB // 1048576
GiB // 1073741824
TiB // 1099511627776 (exceeds 1 << 32)
PiB // 1125899906842624
EiB // 1152921504606846976
ZiB // 1180591620717411303424 (exceeds 1 << 64)
YiB // 1208925819614629174706176
)

%3.6.2 untyped constants

the compiler represents these uncmmmited constants with much greater
numeric recision than values of basic types and arithmetic is more
more precise than machine arithmetic, 256 bits of precisoin
there are
untyped booleans
untyped integers
untyped Rune
untyped floating-point
untype complex
untyped string

ex. Yib and Zib


var f float64 = 3 + 0i // untyped complex -> float64
f=2 // untyped integer -> float64
f=1e123 // untyped floating-point -> float64
f = 'a' // untyped rune -> float64

The statements above are thu s equivalent to these:

var f float64 = float64(3 + 0i)
f=float64(2)
f=float64(1e123)
f=float64('a')

when convertinga constant from one type to another makesure that the target
type can hold the original value. note: if target type is float, it will round

%4 composite types
+++++++++++++++++++++++++++++++
%4.1 arrays

different ways to declared
var a [3]int8
var q [3]int = [3]int(1,2,3) //declaring and instantiating with array literals
q := [...]int(1,2,3) // elipses, length will be deterined by number of intialized
                      //elements

  also size of array is part of type so

  [3]int != [4]int

  q := [3]int{1, 2, 3}
  q=[4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

  initializing using key value pairs

  type Currency int
  const (
  USD Currency = iota
  EUR
  GBP
  RMB
  )
  symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: """}

  In this form, indices can appear in any order and som e may be omitted; as before,
  unspecified values take on the zero value for the element type. For instance,
r := [...]int{99: -1}

in go we can compare arrays if the array type is the same
to check if the elements are the same in same order

pointers and arrays:

func zero(ptr *[32]byte) {
  for i := range ptr {
    ptr[i] = 0
  }
}
or

func zero(ptr *[32]byte) {
*ptr = [32]byte{} //alocating an array of bytes to ptr
}

%4.2 slices
a slcies is written as []//
aray and slices are intimately connecte. a slice is a lightweight
data structure that fives acesss to a subsequence. a sliice
has three componenets: a pointer, a length, and a capacity.

the pointer points to the first element of the array that is reachable
through the slice , which is not neccessarily the array's first elements

ex.

months := [...]string{1 : "January", // ... //, 12 : "December"}

the slice operator s[i:j] where 0 <= i <= j <= cap(s)

revertsuve a slice of intes in place

func revevrse( s []int) {
  for i, j := 0, lens(s)-1, i < j; i, j = i+1, j+1 {
    s[i],s[j] = s[j],s[i]
  }
}

we cannot compare slices except Bytes have an in built Bytes.equals,
so make your own

difference between a slice an array,

he only time you're dealing with an array is when you create it with
 a size:

  names := [3]string{"leto", "paul", "teg"}
  //or
  names := [3]string{}
Everything else, is a slice:

  names := []string{"leto", "paul", "teg"}
  //or
  names := make([]string, 3)
  //or
  var names []string

  slcing beyond cap(s) causes a panic, but slicing beyond len(s) extends the
  slice, so the result may be longer

  fmt.Println(summer[:20]) // panic: out of range
endlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer) // "[June July August September October]"

A simple way to ro tat e a slice left by n elements is to app l y the reverse func tion three times,
firs t to the leading n elements, then to the remaining elements, and finally to the whole slice.
(To rot ate to the rig ht, make the third cal l firs t.)
s := []int{0, 1, 2, 3, 4, 5}
// Rotate s left by two positions.
reverse(s[:2])
reverse(s[2:])
reverse(s)
fmt.Println(s) // "[2 3 4 5 0 1]"

intializing a slice with make
make([]T, len, cap) // same as make([]T, cap)[:len]


%4.2.2 the append function

take a look at implementation of append for int8
func appendInt (x []int, y int) []int {
  var x []int
  var zlen := zlen + 1

  if zlen <= cap(x) {
    //there is room to grow. extend the slices
    z = x[:zlen]
  } else {
    //insufficient space. allocate a new arrays
    //double size, for amoritized linear complexity
    zcap := zlen
    if zcap < 2*len(x) {
      zcap = 2*len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z,x)  //builtin function
  }
  z[len(x)] = y
  return z

}

%4.2.2 in place slice techniques

take a look at this:

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
  i := 0
  for _, s := range strings {
    if s != "" {
      strings[i] = s
      i++
      }
    }
    return strings[:i]
}

the thing is that the input and return slice shares the
same underlying array so the original slice is modified

func remove(slice []int, i int) []int {
  copy(slice[i:], slice[i+1:])
  return slice[:len(slice)-1]
}

%4.3 maps

a map is a reference to a hashtable,
map[K]V where K is the K is the type of the key, and correspondently
for  V for value

we can use make to create a map:

ages := make(map[string]int) // mapping from strings to ints

or we can use a map literal to create a new map with some initial
key-value pairs

ages := map[string]int{
"alice": 31,
"charlie": 34,
}

or:

var x = map[string]int{}

deleting :

delete(map, key)

we cannot take the address of a map:

_=&ages["bob"] // compile error: cannot take address of map element:

looping through a map:

for name, age := range ages {
  fmt.Printf("%s\t%d\n", name, age)
}

order is not preserved so every iteration can be different

The zero value for a map typ e is nil, that is, a reference to no hash
table at all.

map dont have a equal sign so make your own:

func equal(x, y map[string]int) bool {
  if len(x) != len(y) {
  return false
  }
  for k, xv := range x {
    if yv, ok := y[k]; !ok || yv != xv {
      return false
    }
  }
  return true
}

when you access the value of a map you always
get a value, if the key is present then you get a value,
if not hen we get the zero value for that type

ex:
age, ok := ages["bob"]
if !ok { //"bob" is not a key in this map; age == 0. // }

you can also do this:
if age, ok := ages["bob"]; !ok { // ... // }

you can alsp make a map of a map:

  var graph = make(map[string]map[string]bool)
  func addEdge(from, to string) {
    edges := graph[from]
    if edges == nil {
      edges = make(map[string]bool)
      graph[from] = edges
      }
      edges[to] = true
    }
    func hasEdge(from, to string) bool {
      return graph[from][to]
  }

%4.4 structs

  ex.
  type Employee struct {
    ID int
    Name string
    Address string
    DoB time.Time
    Position string
    Salary int
    ManagerID int
  }
  var dilbert Employee
  accessing an element of dilbert:

  dilbert.Name

  the zero value for a struct is composed of them zero values
  of each of its field

  empty stuct is: struct{}

  %4.4.1 struct literals

  type Point struct{ X, Y int }
  p := Point{1, 2}

  two forms of struct literals
  first form, shown above, requires that a value be
  specified for every field, in the right order

  second form, in which a struct value is initialized by listing someSliceor all
  of the field names and their corresponding values

  ex.
  anim := gif.GIF{LoopCount: nframes}

  %4.4.2 struct embedding and anonymous fields

  Go's unusal struct embedding mechanism lets us use one named
  struct type as a n anonymous field of another struct types
  so that instead of doing x.d.e.f we can just do x.fmt.Printf("
  ", var)

  ex.
  lets  say we have to do this:

  type Point struct {
    X, Y int
  }
  type Circle struct {
    Center Point
    Radius int
  }
  type Wheel struct {
    Circle Circle
    Spokes int
  }

  then this happens:

  var w Wheel
  w.Circle.Center.X = 8
  w.Circle.Center.Y = 8
  w.Circle.Radius = 5
  w.Spokes = 20

  instead lets do this:

  type Circle struct {
  Point \\ notice we just have the struct type
  Radius int
  }
  type Wheel struct {
  Circle \\ notice we just have the struct type
  Spokes int
  }

  and now we can do this:

  var w Wheel
  w.X = 8 // equivalent to w.Circle.Point.X = 8
  w.Y = 8 // equivalent to w.Circle.Point.Y = 8
  w.Radius = 5 // equivalent to w.Circle.Radius = 5
  w.Spokes = 20

  note: the explict version still
  works since the name of the anonymous field is just the type of the struct

  however there is no shorthand for struct literals so:

  w=Wheel{8, 8, 5, 20} // compile error: unknown fields
  w=Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

  so we have to do this:

  w=Wheel{Circle{Point{8, 8}, 5}, 20}
  w=Wheel{
  Circle: Circle{
    Point: Point{X: 8, Y: 8},
    Radius: 5,
    },
    Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
  }


  %4.5 JSON

  a json object is a mpping from string to
  values

  converting a go data structure to a json
  is called marshaling with json.marshal

  data, err := json.Marshal(movies)
  if err != nil {
  log.Fatalf("JSON marshaling failed: %s", err)
  }
  fmt.Printf("%s\n", data)

  unmarshaling:

  var titles []struct{ Title string }
    if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
    }
  fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

  %4.6 Text and HTML Templates

  a templates is a string or file containing
  one or more portions enclosed in double
  braces {{...}}, called actions

  ex.
  const templ = `{{.TotalCount}} issues:
  {{range .Items}}----------------------------------------
  Number: {{.Number}}
  User: {{.User.Login}}
  Title: {{.Title | printf "%.64s"}}
  Age: {{.CreatedAt | daysAgo}} days
  {{end}}`

  %5 functions

  func name(parameter-list) (result-list) {
    body
  }

  a sequence of the samke type can be factored
  so that the type itself is written only once.
  ex.

  func f(i, j, k int, s, t string) { // ... // }
  func f(i int, j int, k int, s string, t string) { // ... // }

  there are four ways to declare a function
  with two parameters and one result, all of type int.
  the blank identifier can be used to emphasize that
  a parameter is unused


  func add(x int, y int) int { return x + y }
  func sub(x, y int) (z int) { z = x - y; return }
  func first(x int, _ int) int { return x }
  func zero(int, int) int { return 0 }

  the type of a function is sometimes called its signature

  %5.2 recursion

  //you should already know recursion

  %5.3 multiple return valuesof


  ex.
  links, err := fundLinks(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
      func main () {
      continue
    }
    for _,url := range os.Args[1:] {
    for _, link := range links {
      fmt.Println(link)
    }
  }
  }

  //findLinks performs an HTTP GET request for url, parses the
  //response as HTML, and extracts and returns the links.

  func findLinks(url string) ([]string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  if resp. StatusCode != http.StatusOK {
    resp.Body.Close()
    return nul, fmt.Errorf("getting %s: %s",url, resp.Status)
  }
  doc, err := html.Parse(res.Body)
  resp.Body.Close()
  if err != nil {
    return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
  }
  return visit(nil, doc), nil
  }

  Go's garbage collector recycles unsused memory, but do not assume it
  will release unused operating system resources like open files and
  network connections. They should be closed explicitly

  The result of calling a multi-valued function is a tuple of valuesof

  to ignore one of the values, assign it to a blank identifier

  links,_ := findLinks(url) // errors ignored


  a multi-valued call may appear as the sole argument when calling a function
  of multiple parameters:

  log.Println(findLinks(url))
  links, err := findLinks(url)
  log.Println(links, err)

  named results where the return values have names
  ex:

  func Size(rect image.Rectangle) (width, height int)
  func Split(path string) (dir, file string)
  func HourMinSec(t time.Time) (hour, minute, second int)

  you can do bare returns with named results


  //CountWorsAndImages does an HTTP GET request for the html
  //document url and returns the number of words and images in it.

  func CountWordsAndImages(url string) (word, images int, err error) {
    resp, err := http.GET(url)
    if err != nil {
      return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
      err = fmt.Errorf("parsing HTML: %s", err)
      return
    }
    words, images = countWordsAndImages(doc)
    return
  }
  func countWordsAndImages(n *html.Node) (words, images int) {// ....//}

  %5.4 errors

  the builtin type error is an interface type.
  an error is either nil or some non-neccesarily
  nil -> success
  non-nil -> something went wrong

  GO's approach sets it apart from many other language in which
  failures are reported using exceptions, not ordinary values

  %5.4.1 error handling strategies

  stategy #1: propgagate errors:


  always propogate the error, so a subroutine's error becomes
  the calling function's errors

  you can do this for example to construct new error messages:

  doc, err := html.Parse(resp.Body)
  resp.Body.Close()
  if err != nil {
  return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
  }

  if html.Parse failes then it can propogate and then
  you wil know where the error is

  strategy #2: retry operations

  //WaitForServer attempts to contact the server of a url.
  //It tries for one minute using exponentail back-off.
  //it reports an error if all attempts fail.

  func WaitForServer(url string) error {
    const timeout = 1 * time.minute
    deadline := time.Now().Before(deadline); tries++ {
      _, err := http.Head(url)
      if err == nil {
        return nil //success
      }
      log.Printf("Servernot responding (%s); retrying...", err)
      time.Sleep(time.Second << uint(tries)) // exponential backoff

      }
      return fmt.Errorf("server %s failed to respond after %s", url, timeout)
  }

  strategy #3: exit immediately

  if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1)
  }

  strategy #4: log the error and move on

  strategy #5 ignore errors

  %5.4.2 End of File (EOF)

  package io
  import "errors"
  // EOF is the error returned by Read when no more input is available.
  in := bufio.NewReader(os.Stdin)
  for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
      break // finished reading
    }
    if err != nil {
      return fmt.Errorf("read failed: %v", err)
    }
  // ...use r...

  %5.5 function values


  function are first-class values in GO, a function value may be called
  like any other function
  ex:

  func square(n int) int { return n*n}
  func negative(n int) int { return -n }
  func product(m, n int) int { return m * n }
  f := square
  fmt.Println(f(3)) // "9"
  f=negative
  fmt.Println(f(3)) // "-3"
  fmt.Printf("%T\n", f) // "func(int) int"
  f=product // compile error: can't assign f(int, int) int to f(int) int

  the zero value of a function type is nil. calling a nil function value causes panic:

  var f func(int) int
  f(3) // panic call of nil functions

  so do this:

  if f != nil {
    f(3)
  }

  but function values are not comparable



  %5.6 anonymous functinos

  a function literal is written like a function declaration, but without the name
  following the func keyword

  ex:

  func() (<type>) {
    //...//
    }

    // squares returns a function that returns
  // the next square number each time it is called.
  func squares() func() int {
      var x int
      return func() int {
        x++
        return x * x
      }
    }

  func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
  }

  the scope of ex survives since x exist inside squares as
  a hidden variable


  //this traverses the sorted map
  import (
    "fmt"
    "sort"
  )

  //map from strings to array of strings
  var prereqs = map[string][]string{
    "algorithms" : {"data structures"},
    "calculus" : {"linear algebra"},

    "compilers" : {
      "data structures",
      "formal languages",
      "computer organization",
    },
    "data structures" : {"discrete math"},
    "databases": {"data structures"},
    "discrete math": {"intro to programming"},
    "formal languages": {"discrete math"},
    "networks": {"operating systems"},
    "operating systems": {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
  }

  func main() {
    for i, course := range topoSort(prereqs) {
      fmt.Printf("%d: \t%s\n", i+1, course)
    }
  }

  func topoSort(m map[string][]string) []string{
    var order []string
    seen := make(map[string]bool)
    var visitAll func(items []string)
    visitAll = func(items []string) {
      for _, item := range items {
        if ! seen[item] {
          seen[item] = true
          fmt.Println("calling visitAll")
          visitAll(m[item])
          fmt.Println("appending ", item)
          order = append(order, item)
        }
      }
    }
    var keys []string
    for key := range m {
      keys = append(keys, key)
    }

    sort.Strings(keys)
    fmt.Println(keys)
    visitAll(keys)
    return order
  }

  package main

  import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "log"
    "os"
  )

  //BASIC WEB CRAWLER
  //Extract makes an HTTP GET request to the specified URL,
  //parse the response as HTML, and returns the links in the HTML document.
  func Extract(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
      return nil, err
    }
    if resp.StatusCode != http.StatusOK {
      resp.Body.Close()
      return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }

    var links []string
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
          for _, a := range n.Attr {
            if a.Key != "href" {
              continue
            }
            link, err := resp.Request.URL.Parse(a.Val)
            if err != nil {
              continue //ignore bad URLs
            }
            links = append(links, link.String())
          }
        }
    }
    forEachNode(doc, visitNode, nil)
    return links, nil
  }

  //breadFirst calls f for each item in the worklist
  //Any items returned by f are added to the worklist.
  //f is called at most once for each item

  func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
      items := worklist
      worklist = nil
      for _, item := range items {
        if !seen[item] {
          seen[item] = true
          worklist = append(worklist, f(item)...)
          //f(item)... vauses all the
          // items in the list returned by f to be appended to the worklist
        }
      }
    }
  }

  func crawl(url string) []string {
    fmt.Println(url)
    list, err := Extract(url)
    if err != nil {
      log.Print(err)
    }
    return list
  }

  func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
  	if pre != nil {
  		pre(n)
  	}
  	for c := n.FirstChild; c != nil; c = c.NextSibling {
  		forEachNode(c, pre, post)
  	}
  	if post != nil {
  		post(n)
  	}
  }

  func main() {
    // Crawl the web breadth-first,
  	// starting from the command-line arguments.
  	breadthFirst(crawl, os.Args[1:])
  }

  %5.6.1 Caveat: Capturing Iteration Variables
  consider a program that makes a set of directories and later
  remove them

  var rmdirs []func()
  for _, d := range tempDirs() {
    dir := d //this is neccessary
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
      os.RemoveAll(dir)
    })
  }

  for _, rmdir := range rmdirs {
    rmdir() //
  }

  //compare it to this

  var rmdirs []func()
  for _, dir := range tempDirs() {
      os.MkdirAll(dir, 0755)
      rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!
    })
  }

  //if you think about it in each iteratioin dir changes
  //evertime we append the func() { os.RemoveAll(dir)}
  //dir refers to the same reference which will be
  //the final iteration, and so will not removeAll ALL the dir only the finally

  %5.7 variadic functions

  a varadic function is one that can be called with varying numbers of arguments.

  ex.

  func(vals ...int) int {
    total := 0
    for _, val := range vals {
      total += val
    }
    return total
  }

    what it is reallly doing under the hood the caller is allocating an arrays
    then passes the slice of the entire array to the function.

    the ...int parameter behaves like a slice inside the function, but
    the function type of a variadic function is is distinct from a
    a fucntion with a slice type parameters

  %5.8 deferred function calls

  a defer statement is an ordinary funciton or method call prefixed by
  the keyword defer. the function and argument expressions are evaluated when
  the statement is executed, but the actual call is defered until the function that contains
  the defer statement has finished, whether normallly by executing a return statement or
  falling off the end, or abnormally, by panicking.


  so for example we have a function that opens a file, after the function is done
  the derfered statement will run to close any files that are opened

  ex:
  package ioutil

  func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
      return nil, err
    }
    defer f.Close()
    return ReadAll(f)
  }

  on entry and exit:

  func bigSlowOperation() {
    defer trace("bigSlowOperation")() // don't forget the extra parentheses
    // ...lots of work...
    time.Sleep(10 * time.Second) // simulate slow operation by sleeping
  }

  func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
  }

  $ go build gopl.io/ch5/trace
  $ ./trace
  2015/11/18 09:53:26 enter bigSlowOperation
  2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)

  if you take a look at the output, when we defer the trace statements
  The bigSlowOperation function below calls trace immediately, which
  does the ‘‘on entry’’ action then returns a function value that,
  when called, does the corresponding ‘‘on exit’’ action.

  %5.9 Panic

  go's type system catches many mistakes at compile time, but others, like an
  out of bounds array access or nil pointer dereference , requires checks
  at run-time, when go runtime detects these mistakes, it panics
  when there is a panic, the function execution stops, deferred function
  calls are executed, program crashes, and logs a stack trace

  you can also make your own panics:

  func MustCompile(expr string) *Regexp {
    re, err := Compile(expr)
    if err != nil {
      panic(err)
    }
    return re
  }

  panics are like exeptions in c++

  %5.10 Recover

  recovering from panics is like catching exceptions

  for ex:

  func Parse(input string) (s *Syntax, err error) {
    defer func() {
      if p := recover(); p != nil {
      err = fmt.Errorf("internal error: %v", p)
    }
    }()
    // ...parser...
  }





*/

package main

import (
  "fmt"
)

func main() {

}
