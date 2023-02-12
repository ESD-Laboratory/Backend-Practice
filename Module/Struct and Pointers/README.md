# Struct

As go is not a pure object-oriented language, it does not have classes. However, it does have structs. Structs are similar to classes in other languages. 
They are a collection of fields. They can have methods attached to them. They can be passed around by value. They can be embedded in other structs. 
They can be compared with ==. They can be used as map keys. And they can be used to implement interfaces.

## Differences between structs and classes

Structs are values. Assigning one struct to another copies all the fields. (Two structs that contain the same fields may be unequal even if all the fields are equal, so don't use == to test for equality.)
Example:
```go
type Point struct{ X, Y int }

func main() {
    p := Point{1, 2}
    q := p
    q.X = 5
    fmt.Println(p, q) // "{1 2} {5 2}"
}
```

Structs are not comparable. If you want to test whether two structs contain the same data, you must write your own comparison function.
Example:
```go
type Point struct{ X, Y int }

func Equal(p, q Point) bool {
    return p.X == q.X && p.Y == q.Y
}
```

Structs cannot be inherited from or extended. To extend a struct, embed it in another struct.
Example:
```go
type Point struct{ X, Y int }

type Circle struct {
    Point
    Radius int
}

func main() {
    c := Circle{Point{8, 8}, 5}
    fmt.Println(c.X, c.Y, c.Radius)
}
```


Structs cannot have constructors. To initialize a struct, assign values to its fields.
Example:
```go
type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
```


## Differences between go structs and C structs

Go structs are not the same as C structs. In C, a struct is a collection of fields, but it is not a value. Assigning one struct to another copies the address, not the contents.
Example in C:
```c
struct Point { int x, y; };

void f() {
    Point p = {1, 2};
    Point q = p;
    q.x = 5;
    printf("%d %d\n", p.x, q.x); // "5 5"
}
```
In Go, a struct is a value. Assigning one struct to another copies all the fields.
Example in Go:
```go
type Point struct{ X, Y int }

func main() {
    p := Point{1, 2}
    q := p
    q.X = 5
    fmt.Println(p, q) // "{1 2} {5 2}"
}
```
## So how structs is important in backend development?

Structs are used to represent data. For example, the JSON package uses structs to represent JSON objects. The encoding/json package can decode JSON into a struct, and encode a struct into JSON.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Bob", 20}
    b, _ := json.Marshal(p)
    fmt.Println(string(b))
}
```

# Pointers
Pointers are used to pass references to values and records within your program. They are also used to share data across the program boundary. Pointers are also used to implement call-by-reference in functions.

## What is a pointer?
A pointer is a variable whose value is the address of another variable, i.e., direct address of the memory location. Like any variable or constant, you must declare a pointer before using it to store any variable address. The general form of a pointer variable declaration is:

```go
var var_name *var-type
```

Example:
```go
var ip *int        /* pointer to an integer */
var fp *float32    /* pointer to a float */
```

## How to use pointers?

The & operator is used to find the address of a variable. The * operator is used to access the value at the address specified by a pointer. The * operator is also used to declare a pointer variable.

Example:
```go
package main

import "fmt"

func main() {
    var a int = 20   /* actual variable declaration */
    var ip *int      /* pointer variable declaration */

    ip = &a  /* store address of a in pointer variable*/

    fmt.Printf("Address of a variable: %x\n", &a  )

    /* address stored in pointer variable */
    fmt.Printf("Address stored in ip variable: %x\n", ip )

    /* access the value using the pointer */
    fmt.Printf("Value of *ip variable: %d\n", *ip )
}
```

## What is a nil pointer?

A pointer that is not initialized is called a nil pointer. The nil pointer is a constant with a value of zero defined in several standard packages. For example, in package fmt, nil is a predefined name of type *File which is a valid pointer to the zero value of type File.

```go
package main

import "fmt"

func main() {
    var  ptr *int
    fmt.Printf("The value of ptr is : %x\n", ptr  )
}
```

## What is a pointer to a pointer?

A pointer to a pointer is a form of multiple indirection. It is a pointer that points to another pointer which points to the value that we want to access. The general form of a pointer to a pointer variable declaration is:

```go
var ptr **var-type
```

Example:
```go
package main

import "fmt"

func main() {
    var a int
    var ptr *int
    var pptr **int

    a = 3000

    /* take the address of var */
    ptr = &a

    /* take the address of ptr using address of operator & */
    pptr = &ptr

    /* take the value using pptr */
    fmt.Printf("Value of a = %d\n", a )
    fmt.Printf("Value available at *ptr = %d\n", *ptr )
    fmt.Printf("Value available at **pptr = %d\n", **pptr)
}
```

## What is a pointer to a struct?

A pointer to a struct is a form of multiple indirection. It is a pointer that points to another pointer which points to the value that we want to access. The general form of a pointer to a struct variable declaration is:

```go
var ptr *struct-name
```

Example:
```go
package main

import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main() {
    var Book1 Books        /* Declare Book1 of type Book */
    var Book2 Books        /* Declare Book2 of type Book */

    /* book 1 specification */
    Book1.title = "Go Programming"
    Book1.author = "ESD Laboratory"
    Book1.subject = "Go Programming Tutorial"
    Book1.book_id = 6495407

    /* book 2 specification */
    Book2.title = "UI/UX in Practice Vol1"
    Book2.author = "ESD Laboratory"
    Book2.subject = "UI/UX"
    Book2.book_id = 6495700

    /* print Book1 info */
    fmt.Printf( "Book 1 title : %s\n", Book1.title)

    /* print Book2 info */
    fmt.Printf( "Book 2 title : %s\n", Book2.title)
}
```

## What is a pointer to a function?

A pointer to a function is a form of multiple indirection. It is a pointer that points to another pointer which points to the value that we want to access. The general form of a pointer to a function variable declaration is:

```go
var ptr *func-name
```

Example:
```go
package main

import "fmt"

func getAverage(a []int, size int) int {
    var i, sum int
    var avg int

    for i = 0; i < size; i++ {
        sum += a[i]
    }
    avg = sum / size

    return avg
}

func main() {
    /* an int array with 5 elements */
    balance := []int{1000, 2, 3, 17, 50}
    var avg int

    /* pass pointer to the array as an argument */
    avg = getAverage( balance, 5 )

    /* output the returned value */
    fmt.Printf( "Average value is: %d ", avg )
}
```


