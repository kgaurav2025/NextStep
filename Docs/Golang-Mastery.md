# Go (Golang) Mastery Guide
## From Beginner to Expert & Solution Architect Level

---

## Table of Contents
1. [Beginner Level](#beginner-level)
2. [Intermediate Level](#intermediate-level)
3. [Advanced Level](#advanced-level)
4. [Expert Level](#expert-level)
5. [Solution Architect Level](#solution-architect-level)
6. [Practice Questions with Answers](#practice-questions-with-answers)

---

# Beginner Level

## 1. Introduction to Go

### What is Go?
Go (Golang) is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It was released in 2009.

### Key Features
- Simple and clean syntax
- Fast compilation
- Built-in concurrency (goroutines and channels)
- Garbage collection
- Strong standard library
- Cross-platform compilation

### Installation

```bash
# macOS
brew install go

# Verify installation
go version
```

### Hello World Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Run the program:**
```bash
go run main.go
```

---

## 2. Variables and Data Types

### Variable Declaration

```go
package main

import "fmt"

func main() {
    // Method 1: var keyword with type
    var name string = "John"
    
    // Method 2: var with type inference
    var age = 25
    
    // Method 3: Short declaration (inside functions only)
    city := "New York"
    
    // Multiple declarations
    var x, y, z int = 1, 2, 3
    
    // Zero values
    var defaultInt int       // 0
    var defaultString string // ""
    var defaultBool bool     // false
    
    fmt.Println(name, age, city)
    fmt.Println(x, y, z)
    fmt.Println(defaultInt, defaultString, defaultBool)
}
```

### Basic Data Types

```go
package main

import "fmt"

func main() {
    // Numeric Types
    var integer int = 42
    var integer8 int8 = 127           // -128 to 127
    var integer16 int16 = 32767
    var integer32 int32 = 2147483647
    var integer64 int64 = 9223372036854775807
    
    // Unsigned integers
    var uinteger uint = 42
    var uinteger8 uint8 = 255         // 0 to 255
    
    // Floating point
    var float32Val float32 = 3.14
    var float64Val float64 = 3.141592653589793
    
    // Boolean
    var isActive bool = true
    
    // String
    var message string = "Hello, Go!"
    
    // Byte (alias for uint8)
    var b byte = 'A'
    
    // Rune (alias for int32, represents a Unicode code point)
    var r rune = '世'
    
    fmt.Printf("Integer: %d\n", integer)
    fmt.Printf("Float: %f\n", float64Val)
    fmt.Printf("Boolean: %t\n", isActive)
    fmt.Printf("String: %s\n", message)
    fmt.Printf("Byte: %c\n", b)
    fmt.Printf("Rune: %c\n", r)
}
```

---

## 3. Constants

```go
package main

import "fmt"

// Package level constants
const Pi = 3.14159
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)

// iota - auto-incrementing constant
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
    Thursday       // 4
    Friday         // 5
    Saturday       // 6
)

// Bit flags with iota
const (
    Read    = 1 << iota // 1
    Write               // 2
    Execute             // 4
)

func main() {
    fmt.Println("Pi:", Pi)
    fmt.Println("Status OK:", StatusOK)
    fmt.Println("Monday:", Monday)
    fmt.Println("Read permission:", Read)
    fmt.Println("Write permission:", Write)
}
```

---

## 4. Control Flow

### If-Else Statements

```go
package main

import "fmt"

func main() {
    age := 25
    
    // Basic if-else
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
    
    // If with initialization statement
    if score := 85; score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }
}
```

### Switch Statements

```go
package main

import "fmt"

func main() {
    day := "Monday"
    
    // Basic switch
    switch day {
    case "Monday":
        fmt.Println("Start of work week")
    case "Friday":
        fmt.Println("TGIF!")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Midweek")
    }
    
    // Switch without expression (like if-else chain)
    hour := 14
    switch {
    case hour < 12:
        fmt.Println("Good morning")
    case hour < 17:
        fmt.Println("Good afternoon")
    default:
        fmt.Println("Good evening")
    }
    
    // Type switch
    var i interface{} = "hello"
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Loops

```go
package main

import "fmt"

func main() {
    // Basic for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // While-style loop
    count := 0
    for count < 5 {
        fmt.Println(count)
        count++
    }
    
    // Infinite loop
    // for {
    //     fmt.Println("Forever")
    // }
    
    // Range over slice
    fruits := []string{"apple", "banana", "cherry"}
    for index, fruit := range fruits {
        fmt.Printf("%d: %s\n", index, fruit)
    }
    
    // Range over map
    ages := map[string]int{"Alice": 25, "Bob": 30}
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
    
    // Break and Continue
    for i := 0; i < 10; i++ {
        if i == 3 {
            continue // Skip 3
        }
        if i == 7 {
            break // Stop at 7
        }
        fmt.Println(i)
    }
}
```

---

## 5. Functions

### Basic Functions

```go
package main

import "fmt"

// Simple function
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return // naked return
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    greet("Alice")
    
    result := add(5, 3)
    fmt.Println("5 + 3 =", result)
    
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", quotient)
    }
    
    area, perimeter := rectangle(5, 3)
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", area, perimeter)
    
    fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
}
```

### Anonymous Functions and Closures

```go
package main

import "fmt"

func main() {
    // Anonymous function
    func() {
        fmt.Println("Anonymous function")
    }()
    
    // Anonymous function with assignment
    square := func(x int) int {
        return x * x
    }
    fmt.Println("Square of 5:", square(5))
    
    // Closure
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }()
    
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
```

---

## 6. Arrays and Slices

### Arrays

```go
package main

import "fmt"

func main() {
    // Array declaration
    var arr1 [5]int
    fmt.Println("Zero value array:", arr1)
    
    // Array with initialization
    arr2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Initialized array:", arr2)
    
    // Array with inferred length
    arr3 := [...]string{"a", "b", "c"}
    fmt.Println("Inferred length:", arr3, "Length:", len(arr3))
    
    // Accessing elements
    arr2[0] = 10
    fmt.Println("Modified array:", arr2)
    
    // Array length
    fmt.Println("Length of arr2:", len(arr2))
    
    // Iterating over array
    for i, v := range arr2 {
        fmt.Printf("Index %d: %d\n", i, v)
    }
}
```

### Slices

```go
package main

import "fmt"

func main() {
    // Slice declaration
    var slice1 []int
    fmt.Println("Nil slice:", slice1, "Length:", len(slice1))
    
    // Slice with make
    slice2 := make([]int, 5)    // length 5, capacity 5
    slice3 := make([]int, 3, 5) // length 3, capacity 5
    fmt.Println("Make slice:", slice2, "Len:", len(slice2), "Cap:", cap(slice2))
    fmt.Println("Make slice with cap:", slice3, "Len:", len(slice3), "Cap:", cap(slice3))
    
    // Slice literal
    slice4 := []int{1, 2, 3, 4, 5}
    fmt.Println("Slice literal:", slice4)
    
    // Append to slice
    slice4 = append(slice4, 6, 7, 8)
    fmt.Println("After append:", slice4)
    
    // Slicing
    subSlice := slice4[2:5] // elements at index 2, 3, 4
    fmt.Println("Sub-slice [2:5]:", subSlice)
    
    // Copy slice
    destination := make([]int, len(slice4))
    copy(destination, slice4)
    fmt.Println("Copied slice:", destination)
    
    // Append slices
    slice5 := []int{100, 200}
    slice4 = append(slice4, slice5...)
    fmt.Println("After appending slice:", slice4)
}
```

---

## 7. Maps

```go
package main

import "fmt"

func main() {
    // Map declaration with make
    ages := make(map[string]int)
    ages["Alice"] = 25
    ages["Bob"] = 30
    fmt.Println("Ages:", ages)
    
    // Map literal
    scores := map[string]int{
        "Math":    95,
        "Science": 88,
        "English": 92,
    }
    fmt.Println("Scores:", scores)
    
    // Accessing values
    mathScore := scores["Math"]
    fmt.Println("Math score:", mathScore)
    
    // Check if key exists
    value, exists := scores["History"]
    if exists {
        fmt.Println("History score:", value)
    } else {
        fmt.Println("History score not found")
    }
    
    // Delete a key
    delete(scores, "English")
    fmt.Println("After delete:", scores)
    
    // Iterate over map
    for subject, score := range scores {
        fmt.Printf("%s: %d\n", subject, score)
    }
    
    // Length of map
    fmt.Println("Number of subjects:", len(scores))
}
```

---

## 8. Structs

```go
package main

import "fmt"

// Define a struct
type Person struct {
    Name    string
    Age     int
    Email   string
    Address Address // Nested struct
}

type Address struct {
    Street  string
    City    string
    Country string
}

// Method on struct
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s", p.Name)
}

// Pointer receiver (can modify the struct)
func (p *Person) Birthday() {
    p.Age++
}

func main() {
    // Create struct instances
    person1 := Person{
        Name:  "Alice",
        Age:   25,
        Email: "alice@example.com",
        Address: Address{
            Street:  "123 Main St",
            City:    "New York",
            Country: "USA",
        },
    }
    
    // Access fields
    fmt.Println("Name:", person1.Name)
    fmt.Println("City:", person1.Address.City)
    
    // Call methods
    fmt.Println(person1.Greet())
    
    // Modify with pointer receiver
    person1.Birthday()
    fmt.Println("After birthday, age:", person1.Age)
    
    // Anonymous struct
    employee := struct {
        ID   int
        Name string
    }{
        ID:   1,
        Name: "John",
    }
    fmt.Println("Employee:", employee)
    
    // Struct pointer
    person2 := &Person{Name: "Bob", Age: 30}
    fmt.Println("Person2 name:", person2.Name)
}
```

---

## 9. Pointers

```go
package main

import "fmt"

func main() {
    // Pointer basics
    x := 10
    var p *int = &x // p is a pointer to x
    
    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", &x)
    fmt.Println("Value of p:", p)
    fmt.Println("Value at p:", *p) // Dereferencing
    
    // Modify value through pointer
    *p = 20
    fmt.Println("New value of x:", x)
    
    // Pointer to struct
    type Point struct {
        X, Y int
    }
    
    point := &Point{X: 10, Y: 20}
    fmt.Println("Point:", point)
    fmt.Println("Point X:", point.X) // Automatic dereferencing
    
    // New function
    ptr := new(int) // Allocates memory, returns pointer
    *ptr = 100
    fmt.Println("Value at ptr:", *ptr)
}

// Pass by value vs pass by reference
func incrementValue(x int) {
    x++
}

func incrementPointer(x *int) {
    *x++
}
```

---

## 10. Error Handling

```go
package main

import (
    "errors"
    "fmt"
)

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

// Function returning error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Function with custom error
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Message: "age cannot be negative",
        }
    }
    if age > 150 {
        return &ValidationError{
            Field:   "age",
            Message: "age seems unrealistic",
        }
    }
    return nil
}

func main() {
    // Basic error handling
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
    
    // Custom error handling
    err = validateAge(-5)
    if err != nil {
        fmt.Println("Error:", err)
        
        // Type assertion for custom error
        if valErr, ok := err.(*ValidationError); ok {
            fmt.Println("Field:", valErr.Field)
            fmt.Println("Message:", valErr.Message)
        }
    }
    
    // Wrapping errors (Go 1.13+)
    err = fmt.Errorf("processing failed: %w", errors.New("original error"))
    fmt.Println("Wrapped error:", err)
    
    // Unwrapping errors
    unwrapped := errors.Unwrap(err)
    fmt.Println("Unwrapped:", unwrapped)
}
```

---

# Intermediate Level

## 11. Interfaces

```go
package main

import (
    "fmt"
    "math"
)

// Interface definition
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Circle implements Shape
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Rectangle implements Shape
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Function accepting interface
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Empty interface
func describe(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 10, Height: 5}
    
    // Polymorphism
    shapes := []Shape{circle, rectangle}
    for _, shape := range shapes {
        printShapeInfo(shape)
    }
    
    // Empty interface accepts any type
    describe(42)
    describe("hello")
    describe(circle)
    
    // Type assertion
    var i interface{} = "hello"
    s, ok := i.(string)
    if ok {
        fmt.Println("String value:", s)
    }
}
```

---

## 12. Concurrency - Goroutines

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello, %s! (%d)\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Start goroutine
    go sayHello("Alice")
    go sayHello("Bob")
    
    // Main goroutine continues
    fmt.Println("Main function")
    
    // Wait for goroutines (simple approach)
    time.Sleep(1 * time.Second)
    
    // Better approach: WaitGroup
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Duration(id*100) * time.Millisecond)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
}
```

---

## 13. Concurrency - Channels

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Unbuffered channel
    ch := make(chan string)
    
    go func() {
        ch <- "Hello from goroutine!"
    }()
    
    msg := <-ch
    fmt.Println(msg)
    
    // Buffered channel
    buffered := make(chan int, 3)
    buffered <- 1
    buffered <- 2
    buffered <- 3
    
    fmt.Println(<-buffered)
    fmt.Println(<-buffered)
    fmt.Println(<-buffered)
    
    // Channel with range
    numbers := make(chan int, 5)
    go func() {
        for i := 1; i <= 5; i++ {
            numbers <- i
        }
        close(numbers)
    }()
    
    for num := range numbers {
        fmt.Println("Received:", num)
    }
    
    // Select statement
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "from ch1"
    }()
    
    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "from ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received", msg2)
        case <-time.After(1 * time.Second):
            fmt.Println("Timeout")
        }
    }
}
```

---

## 14. Concurrency Patterns

### Worker Pool Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID   int
    Data string
}

type Result struct {
    JobID  int
    Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job.ID)
        time.Sleep(100 * time.Millisecond) // Simulate work
        results <- Result{
            JobID:  job.ID,
            Output: fmt.Sprintf("Processed: %s", job.Data),
        }
    }
}

func main() {
    numWorkers := 3
    numJobs := 10
    
    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)
    
    var wg sync.WaitGroup
    
    // Start workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }
    
    // Send jobs
    for i := 1; i <= numJobs; i++ {
        jobs <- Job{ID: i, Data: fmt.Sprintf("Data-%d", i)}
    }
    close(jobs)
    
    // Wait for workers and close results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Result: Job %d -> %s\n", result.JobID, result.Output)
    }
}
```

### Fan-Out/Fan-In Pattern

```go
package main

import (
    "fmt"
    "sync"
)

// Fan-out: distribute work to multiple goroutines
func fanOut(input <-chan int, numWorkers int) []<-chan int {
    outputs := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        outputs[i] = process(input)
    }
    return outputs
}

func process(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for n := range input {
            output <- n * n // Square the number
        }
    }()
    return output
}

// Fan-in: merge multiple channels into one
func fanIn(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    output := make(chan int)
    
    multiplex := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            output <- n
        }
    }
    
    wg.Add(len(channels))
    for _, c := range channels {
        go multiplex(c)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}

func main() {
    input := make(chan int)
    
    go func() {
        defer close(input)
        for i := 1; i <= 10; i++ {
            input <- i
        }
    }()
    
    // Fan-out to 3 workers
    outputs := fanOut(input, 3)
    
    // Fan-in results
    results := fanIn(outputs...)
    
    for result := range results {
        fmt.Println(result)
    }
}
```

---

## 15. Context Package

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func slowOperation(ctx context.Context) error {
    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Operation completed")
        return nil
    case <-ctx.Done():
        fmt.Println("Operation cancelled:", ctx.Err())
        return ctx.Err()
    }
}

func main() {
    // Context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    err := slowOperation(ctx)
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    // Context with cancel
    ctx2, cancel2 := context.WithCancel(context.Background())
    
    go func() {
        time.Sleep(1 * time.Second)
        cancel2()
    }()
    
    err = slowOperation(ctx2)
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    // Context with value
    type key string
    ctx3 := context.WithValue(context.Background(), key("userID"), 123)
    
    userID := ctx3.Value(key("userID")).(int)
    fmt.Println("User ID:", userID)
}
```

---

## 16. Testing

### Basic Testing

```go
// math.go
package mathutil

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

```go
// math_test.go
package mathutil

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Table-driven tests
func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 6},
        {"with zero", 5, 0, 0},
        {"negative numbers", -2, 3, -6},
        {"both negative", -2, -3, 6},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Multiply(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Multiply(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Test with error
func TestDivide(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("Expected error for division by zero")
    }
    
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if result != 5 {
        t.Errorf("Divide(10, 2) = %d; want 5", result)
    }
}

// Benchmark
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific test
go test -run TestAdd ./...

# Run benchmarks
go test -bench=. ./...

# Generate coverage report
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 17. Packages and Modules

### Creating a Module

```bash
# Initialize a new module
go mod init github.com/username/myproject

# Add dependencies
go get github.com/gin-gonic/gin

# Tidy up dependencies
go mod tidy

# Vendor dependencies
go mod vendor
```

### Package Structure

```
myproject/
├── go.mod
├── go.sum
├── main.go
├── internal/           # Private packages
│   └── database/
│       └── db.go
├── pkg/                # Public packages
│   └── utils/
│       └── helpers.go
└── cmd/                # Application entry points
    └── myapp/
        └── main.go
```

### Importing Packages

```go
package main

import (
    // Standard library
    "fmt"
    "net/http"
    
    // External packages
    "github.com/gin-gonic/gin"
    
    // Internal packages
    "github.com/username/myproject/internal/database"
    "github.com/username/myproject/pkg/utils"
)
```

---

## 18. JSON Handling

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Email   string   `json:"email,omitempty"`
    Active  bool     `json:"active"`
    Private string   `json:"-"` // Ignored in JSON
    Tags    []string `json:"tags"`
}

func main() {
    // Struct to JSON (Marshal)
    person := Person{
        Name:    "Alice",
        Age:     30,
        Email:   "alice@example.com",
        Active:  true,
        Private: "secret",
        Tags:    []string{"developer", "go"},
    }
    
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("JSON:", string(jsonData))
    
    // Pretty print JSON
    prettyJSON, _ := json.MarshalIndent(person, "", "  ")
    fmt.Println("Pretty JSON:\n", string(prettyJSON))
    
    // JSON to Struct (Unmarshal)
    jsonString := `{"name":"Bob","age":25,"active":false,"tags":["admin"]}`
    var person2 Person
    err = json.Unmarshal([]byte(jsonString), &person2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Parsed: %+v\n", person2)
    
    // Dynamic JSON with map
    var data map[string]interface{}
    json.Unmarshal([]byte(jsonString), &data)
    fmt.Println("Map:", data)
    fmt.Println("Name:", data["name"])
}
```

---

## 19. HTTP Server

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Message struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Go HTTP Server!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    response := Message{
        Status:  "success",
        Message: "Hello from API!",
    }
    
    json.NewEncoder(w).Encode(response)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var data map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "received": data,
    })
}

// Middleware example
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/api", apiHandler)
    mux.HandleFunc("/post", postHandler)
    
    // Apply middleware
    handler := loggingMiddleware(mux)
    
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
```

---

## 20. HTTP Client

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

func main() {
    // Create HTTP client with timeout
    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    // GET request
    resp, err := client.Get("https://api.github.com/users/golang")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Println("Status:", resp.Status)
    fmt.Println("Body:", string(body[:200])+"...")
    
    // POST request with JSON
    data := map[string]string{
        "name":  "Test",
        "email": "test@example.com",
    }
    jsonData, _ := json.Marshal(data)
    
    resp2, err := client.Post(
        "https://httpbin.org/post",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp2.Body.Close()
    
    // Custom request with headers
    req, _ := http.NewRequest("GET", "https://api.github.com/users/golang", nil)
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "Bearer your-token-here")
    
    resp3, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp3.Body.Close()
    
    fmt.Println("Custom request status:", resp3.Status)
}
```

---

# Advanced Level

## 21. Generics (Go 1.18+)

```go
package main

import (
    "fmt"
    "golang.org/x/exp/constraints"
)

// Generic function
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Generic function with multiple type parameters
func Map[T any, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}

// Generic struct
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}

// Custom type constraint
type Number interface {
    int | int32 | int64 | float32 | float64
}

func Sum[T Number](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

func main() {
    // Using generic Min
    fmt.Println(Min(3, 5))       // 3
    fmt.Println(Min("a", "b"))   // a
    fmt.Println(Min(3.14, 2.71)) // 2.71
    
    // Using generic Map
    numbers := []int{1, 2, 3, 4, 5}
    doubled := Map(numbers, func(n int) int { return n * 2 })
    fmt.Println(doubled) // [2 4 6 8 10]
    
    strings := Map(numbers, func(n int) string { 
        return fmt.Sprintf("Number: %d", n) 
    })
    fmt.Println(strings)
    
    // Using generic Stack
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)
    
    if val, ok := intStack.Pop(); ok {
        fmt.Println("Popped:", val) // 3
    }
    
    // Using Sum with constraint
    ints := []int{1, 2, 3, 4, 5}
    floats := []float64{1.1, 2.2, 3.3}
    fmt.Println("Sum ints:", Sum(ints))     // 15
    fmt.Println("Sum floats:", Sum(floats)) // 6.6
}
```

---

## 22. Reflection

```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string `json:"name" validate:"required"`
    Age     int    `json:"age" validate:"min=0,max=150"`
    Email   string `json:"email" validate:"email"`
    private string // unexported field
}

func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}

func (p *Person) SetName(name string) {
    p.Name = name
}

func inspectType(i interface{}) {
    t := reflect.TypeOf(i)
    v := reflect.ValueOf(i)
    
    fmt.Println("Type:", t)
    fmt.Println("Kind:", t.Kind())
    fmt.Println("Value:", v)
    
    // Handle pointer
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
        v = v.Elem()
    }
    
    // Inspect struct fields
    if t.Kind() == reflect.Struct {
        fmt.Println("\nFields:")
        for i := 0; i < t.NumField(); i++ {
            field := t.Field(i)
            value := v.Field(i)
            
            // Check if field is exported
            if field.PkgPath == "" {
                fmt.Printf("  %s (%s): %v\n", 
                    field.Name, field.Type, value.Interface())
                
                // Get struct tags
                if tag := field.Tag.Get("json"); tag != "" {
                    fmt.Printf("    json tag: %s\n", tag)
                }
                if tag := field.Tag.Get("validate"); tag != "" {
                    fmt.Printf("    validate tag: %s\n", tag)
                }
            }
        }
        
        // Inspect methods
        fmt.Println("\nMethods:")
        for i := 0; i < t.NumMethod(); i++ {
            method := t.Method(i)
            fmt.Printf("  %s\n", method.Name)
        }
    }
}

// Dynamic function call
func callMethod(obj interface{}, methodName string, args ...interface{}) []reflect.Value {
    v := reflect.ValueOf(obj)
    method := v.MethodByName(methodName)
    
    if !method.IsValid() {
        panic(fmt.Sprintf("Method %s not found", methodName))
    }
    
    in := make([]reflect.Value, len(args))
    for i, arg := range args {
        in[i] = reflect.ValueOf(arg)
    }
    
    return method.Call(in)
}

// Create struct dynamically
func createStruct() {
    // Define fields
    fields := []reflect.StructField{
        {
            Name: "Name",
            Type: reflect.TypeOf(""),
            Tag:  reflect.StructTag(`json:"name"`),
        },
        {
            Name: "Age",
            Type: reflect.TypeOf(0),
            Tag:  reflect.StructTag(`json:"age"`),
        },
    }
    
    // Create struct type
    structType := reflect.StructOf(fields)
    
    // Create instance
    instance := reflect.New(structType).Elem()
    instance.Field(0).SetString("John")
    instance.Field(1).SetInt(30)
    
    fmt.Printf("Dynamic struct: %+v\n", instance.Interface())
}

func main() {
    person := Person{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }
    
    inspectType(person)
    
    // Dynamic method call
    results := callMethod(person, "Greet")
    fmt.Println("\nDynamic call result:", results[0].String())
    
    // Create dynamic struct
    fmt.Println()
    createStruct()
}
```

---

## 23. Unsafe Package

```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    // Get size of types
    var i int64
    var f float64
    var s string
    
    fmt.Println("Size of int64:", unsafe.Sizeof(i))     // 8
    fmt.Println("Size of float64:", unsafe.Sizeof(f))   // 8
    fmt.Println("Size of string:", unsafe.Sizeof(s))    // 16 (pointer + length)
    
    // Pointer arithmetic (BE CAREFUL!)
    arr := [5]int{10, 20, 30, 40, 50}
    ptr := &arr[0]
    
    // Get element at index 2 using pointer arithmetic
    ptr2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 2*unsafe.Sizeof(arr[0])))
    fmt.Println("Element at index 2:", *ptr2) // 30
    
    // Convert between pointer types
    type MyInt int
    var x int = 42
    myIntPtr := (*MyInt)(unsafe.Pointer(&x))
    fmt.Println("Converted:", *myIntPtr)
    
    // Struct alignment and padding
    type Example struct {
        a bool  // 1 byte
        b int64 // 8 bytes
        c bool  // 1 byte
    }
    
    var ex Example
    fmt.Println("Size of Example:", unsafe.Sizeof(ex))         // 24 (with padding)
    fmt.Println("Align of Example:", unsafe.Alignof(ex))       // 8
    fmt.Println("Offset of a:", unsafe.Offsetof(ex.a))         // 0
    fmt.Println("Offset of b:", unsafe.Offsetof(ex.b))         // 8
    fmt.Println("Offset of c:", unsafe.Offsetof(ex.c))         // 16
    
    // Optimized struct layout
    type OptimizedExample struct {
        b int64 // 8 bytes
        a bool  // 1 byte
        c bool  // 1 byte
        // 6 bytes padding
    }
    
    var optEx OptimizedExample
    fmt.Println("Size of OptimizedExample:", unsafe.Sizeof(optEx)) // 16
}
```

---

## 24. Build Tags and Cross Compilation

### Build Tags

```go
// +build linux darwin
// +build amd64

// OR in Go 1.17+:
//go:build (linux || darwin) && amd64

package main
```

### Platform-Specific Files

```
myapp/
├── main.go
├── file_linux.go    // Only compiled on Linux
├── file_darwin.go   // Only compiled on macOS
└── file_windows.go  // Only compiled on Windows
```

### Cross Compilation

```bash
# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=darwin GOARCH=amd64 go build -o myapp-macos
GOOS=windows GOARCH=amd64 go build -o myapp.exe

# Common GOOS/GOARCH combinations
# linux/amd64, linux/arm64
# darwin/amd64, darwin/arm64
# windows/amd64, windows/386
# freebsd/amd64
```

### Build Flags

```bash
# Strip debug information (smaller binary)
go build -ldflags="-s -w" -o myapp

# Inject version at build time
go build -ldflags="-X main.Version=1.0.0 -X main.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o myapp

# Static build (no external dependencies)
CGO_ENABLED=0 go build -a -installsuffix cgo -o myapp
```

---

## 25. Profiling and Performance

### pprof

```go
package main

import (
    "log"
    "net/http"
    _ "net/http/pprof" // Import for side effects
)

func main() {
    // Start pprof server
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Your application code
    // ...
}
```

### Accessing Profiles

```bash
# CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Memory profile
go tool pprof http://localhost:6060/debug/pprof/heap

# Goroutine profile
go tool pprof http://localhost:6060/debug/pprof/goroutine

# Block profile
go tool pprof http://localhost:6060/debug/pprof/block

# Generate profile from test
go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=.
go tool pprof cpu.prof
```

### Benchmark Testing

```go
package main

import (
    "testing"
)

func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := ""
        for j := 0; j < 100; j++ {
            s += "a"
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        for j := 0; j < 100; j++ {
            sb.WriteString("a")
        }
        _ = sb.String()
    }
}
```

---

## 26. Memory Management

### Escape Analysis

```bash
# See what escapes to heap
go build -gcflags="-m" main.go
go build -gcflags="-m -m" main.go  # More verbose
```

### Memory Optimization Tips

```go
package main

import (
    "sync"
)

// Use sync.Pool for frequently allocated objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

func processWithPool() {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    
    // Use buf...
}

// Pre-allocate slices when size is known
func efficientSlice() {
    // Bad: multiple reallocations
    var bad []int
    for i := 0; i < 1000; i++ {
        bad = append(bad, i)
    }
    
    // Good: single allocation
    good := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        good = append(good, i)
    }
}

// Avoid unnecessary string-to-byte conversions
func efficientStringBytes(s string) {
    // This creates a copy
    b := []byte(s)
    _ = b
    
    // For read-only access, use unsafe (carefully)
    // import "unsafe"
    // b := unsafe.Slice(unsafe.StringData(s), len(s))
}
```

---

# Expert Level

## 27. Custom Runtime Behavior

### init Functions

```go
package main

import "fmt"

var globalVar int

func init() {
    // Called before main
    globalVar = 42
    fmt.Println("Init 1 called")
}

func init() {
    // Multiple init functions are allowed
    fmt.Println("Init 2 called")
}

func main() {
    fmt.Println("Main called, globalVar:", globalVar)
}
```

### Controlling Goroutine Scheduling

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    // Set number of OS threads
    runtime.GOMAXPROCS(4)
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
    
    // Number of goroutines
    fmt.Println("NumGoroutine:", runtime.NumGoroutine())
    
    // Yield the processor
    var wg sync.WaitGroup
    wg.Add(2)
    
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine 1:", i)
            runtime.Gosched() // Yield to other goroutines
        }
    }()
    
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine 2:", i)
            runtime.Gosched()
        }
    }()
    
    wg.Wait()
    
    // Force garbage collection
    runtime.GC()
    
    // Get memory statistics
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc: %d MB\n", m.Alloc/1024/1024)
    fmt.Printf("TotalAlloc: %d MB\n", m.TotalAlloc/1024/1024)
    fmt.Printf("NumGC: %d\n", m.NumGC)
}
```

---

## 28. CGO (Calling C from Go)

```go
package main

/*
#include <stdio.h>
#include <stdlib.h>

void sayHello(const char* name) {
    printf("Hello, %s from C!\n", name);
}

int add(int a, int b) {
    return a + b;
}
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    // Call C function
    name := C.CString("Go Developer")
    defer C.free(unsafe.Pointer(name))
    
    C.sayHello(name)
    
    // Call C function with return value
    result := C.add(5, 3)
    fmt.Println("5 + 3 =", int(result))
}
```

### Building with CGO

```bash
# Enable CGO
CGO_ENABLED=1 go build

# Specify C compiler
CC=gcc go build

# Cross-compile with CGO (requires cross-compiler)
CC=x86_64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build
```

---

## 29. Plugin System

```go
// plugin.go - Compile as plugin
package main

import "fmt"

var Message = "Hello from plugin!"

func Greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func Add(a, b int) int {
    return a + b
}
```

```bash
# Build plugin
go build -buildmode=plugin -o myplugin.so plugin.go
```

```go
// main.go - Load and use plugin
package main

import (
    "fmt"
    "plugin"
)

func main() {
    // Load plugin
    p, err := plugin.Open("myplugin.so")
    if err != nil {
        panic(err)
    }
    
    // Look up exported variable
    msgSymbol, err := p.Lookup("Message")
    if err != nil {
        panic(err)
    }
    msg := *msgSymbol.(*string)
    fmt.Println(msg)
    
    // Look up exported function
    greetSymbol, err := p.Lookup("Greet")
    if err != nil {
        panic(err)
    }
    greet := greetSymbol.(func(string) string)
    fmt.Println(greet("Plugin User"))
    
    // Look up another function
    addSymbol, err := p.Lookup("Add")
    if err != nil {
        panic(err)
    }
    add := addSymbol.(func(int, int) int)
    fmt.Println("5 + 3 =", add(5, 3))
}
```

---

## 30. Assembly in Go

```go
// add_amd64.s
// func Add(a, b int64) int64
TEXT ·Add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    ADDQ b+8(FP), AX
    MOVQ AX, ret+16(FP)
    RET
```

```go
// add.go
package main

func Add(a, b int64) int64

func main() {
    result := Add(5, 3)
    println("5 + 3 =", result)
}
```

---

## 31. Distributed Systems Patterns

### Service Discovery

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceRegistry struct {
    client *clientv3.Client
    lease  clientv3.Lease
}

func NewServiceRegistry(endpoints []string) (*ServiceRegistry, error) {
    client, err := clientv3.New(clientv3.Config{
        Endpoints:   endpoints,
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        return nil, err
    }
    
    return &ServiceRegistry{
        client: client,
        lease:  clientv3.NewLease(client),
    }, nil
}

func (sr *ServiceRegistry) Register(ctx context.Context, serviceName, address string, ttl int64) error {
    // Create lease
    leaseResp, err := sr.lease.Grant(ctx, ttl)
    if err != nil {
        return err
    }
    
    // Put with lease
    key := fmt.Sprintf("/services/%s/%s", serviceName, address)
    _, err = sr.client.Put(ctx, key, address, clientv3.WithLease(leaseResp.ID))
    if err != nil {
        return err
    }
    
    // Keep alive
    keepAliveChan, err := sr.lease.KeepAlive(ctx, leaseResp.ID)
    if err != nil {
        return err
    }
    
    go func() {
        for range keepAliveChan {
            // Keep alive received
        }
    }()
    
    return nil
}

func (sr *ServiceRegistry) Discover(ctx context.Context, serviceName string) ([]string, error) {
    key := fmt.Sprintf("/services/%s/", serviceName)
    resp, err := sr.client.Get(ctx, key, clientv3.WithPrefix())
    if err != nil {
        return nil, err
    }
    
    addresses := make([]string, 0, len(resp.Kvs))
    for _, kv := range resp.Kvs {
        addresses = append(addresses, string(kv.Value))
    }
    
    return addresses, nil
}
```

### Circuit Breaker Pattern

```go
package main

import (
    "errors"
    "sync"
    "time"
)

type State int

const (
    StateClosed State = iota
    StateOpen
    StateHalfOpen
)

type CircuitBreaker struct {
    mu              sync.Mutex
    state           State
    failures        int
    successes       int
    maxFailures     int
    resetTimeout    time.Duration
    halfOpenMax     int
    lastFailureTime time.Time
}

func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:        StateClosed,
        maxFailures:  maxFailures,
        resetTimeout: resetTimeout,
        halfOpenMax:  3,
    }
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
    cb.mu.Lock()
    
    // Check if we should transition from Open to HalfOpen
    if cb.state == StateOpen {
        if time.Since(cb.lastFailureTime) > cb.resetTimeout {
            cb.state = StateHalfOpen
            cb.successes = 0
        } else {
            cb.mu.Unlock()
            return errors.New("circuit breaker is open")
        }
    }
    
    cb.mu.Unlock()
    
    // Execute the function
    err := fn()
    
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if err != nil {
        cb.failures++
        cb.lastFailureTime = time.Now()
        
        if cb.state == StateHalfOpen || cb.failures >= cb.maxFailures {
            cb.state = StateOpen
        }
        return err
    }
    
    // Success
    if cb.state == StateHalfOpen {
        cb.successes++
        if cb.successes >= cb.halfOpenMax {
            cb.state = StateClosed
            cb.failures = 0
        }
    } else {
        cb.failures = 0
    }
    
    return nil
}

func (cb *CircuitBreaker) State() State {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    return cb.state
}
```

---

## 32. Rate Limiting

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Token Bucket Rate Limiter
type TokenBucket struct {
    mu         sync.Mutex
    tokens     float64
    maxTokens  float64
    refillRate float64 // tokens per second
    lastRefill time.Time
}

func NewTokenBucket(maxTokens, refillRate float64) *TokenBucket {
    return &TokenBucket{
        tokens:     maxTokens,
        maxTokens:  maxTokens,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()
    
    // Refill tokens
    now := time.Now()
    elapsed := now.Sub(tb.lastRefill).Seconds()
    tb.tokens = min(tb.maxTokens, tb.tokens+elapsed*tb.refillRate)
    tb.lastRefill = now
    
    if tb.tokens >= 1 {
        tb.tokens--
        return true
    }
    return false
}

// Sliding Window Rate Limiter
type SlidingWindow struct {
    mu         sync.Mutex
    timestamps []time.Time
    windowSize time.Duration
    maxRequests int
}

func NewSlidingWindow(windowSize time.Duration, maxRequests int) *SlidingWindow {
    return &SlidingWindow{
        timestamps:  make([]time.Time, 0),
        windowSize:  windowSize,
        maxRequests: maxRequests,
    }
}

func (sw *SlidingWindow) Allow() bool {
    sw.mu.Lock()
    defer sw.mu.Unlock()
    
    now := time.Now()
    windowStart := now.Add(-sw.windowSize)
    
    // Remove old timestamps
    valid := make([]time.Time, 0)
    for _, ts := range sw.timestamps {
        if ts.After(windowStart) {
            valid = append(valid, ts)
        }
    }
    sw.timestamps = valid
    
    if len(sw.timestamps) < sw.maxRequests {
        sw.timestamps = append(sw.timestamps, now)
        return true
    }
    return false
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func main() {
    // Token bucket example
    limiter := NewTokenBucket(10, 2) // 10 tokens max, 2 tokens/sec refill
    
    for i := 0; i < 15; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Rate limited\n", i+1)
        }
    }
    
    // Wait and try again
    time.Sleep(2 * time.Second)
    if limiter.Allow() {
        fmt.Println("After wait: Allowed")
    }
}
```

---

# Solution Architect Level

## 33. Microservices Architecture

### Project Structure for Microservices

```
my-microservices/
├── api-gateway/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes/
│   └── Dockerfile
├── user-service/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/
│   │   ├── repository/
│   │   ├── service/
│   │   └── handlers/
│   ├── proto/
│   │   └── user.proto
│   └── Dockerfile
├── order-service/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/
│   │   ├── repository/
│   │   ├── service/
│   │   └── handlers/
│   └── Dockerfile
├── shared/
│   ├── config/
│   ├── logger/
│   ├── middleware/
│   └── utils/
├── docker-compose.yml
├── Makefile
└── README.md
```

### Clean Architecture Example

```go
// domain/user.go
package domain

import (
    "context"
    "time"
)

type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id string) (*User, error)
    GetByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, offset, limit int) ([]*User, error)
}

type UserService interface {
    Register(ctx context.Context, email, name, password string) (*User, error)
    GetUser(ctx context.Context, id string) (*User, error)
    UpdateUser(ctx context.Context, user *User) error
    DeleteUser(ctx context.Context, id string) error
    ListUsers(ctx context.Context, offset, limit int) ([]*User, error)
}
```

```go
// service/user_service.go
package service

import (
    "context"
    "errors"
    "time"
    
    "github.com/google/uuid"
    "myapp/internal/domain"
)

type userService struct {
    repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
    return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, email, name, password string) (*domain.User, error) {
    // Check if user exists
    existing, _ := s.repo.GetByEmail(ctx, email)
    if existing != nil {
        return nil, errors.New("user already exists")
    }
    
    user := &domain.User{
        ID:        uuid.New().String(),
        Email:     email,
        Name:      name,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    
    if err := s.repo.Create(ctx, user); err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
    user.UpdatedAt = time.Now()
    return s.repo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context, offset, limit int) ([]*domain.User, error) {
    return s.repo.List(ctx, offset, limit)
}
```

---

## 34. Event-Driven Architecture

### Event Bus

```go
package eventbus

import (
    "context"
    "encoding/json"
    "sync"
)

type Event struct {
    Type      string          `json:"type"`
    Payload   json.RawMessage `json:"payload"`
    Timestamp int64           `json:"timestamp"`
}

type Handler func(ctx context.Context, event Event) error

type EventBus struct {
    mu       sync.RWMutex
    handlers map[string][]Handler
}

func New() *EventBus {
    return &EventBus{
        handlers: make(map[string][]Handler),
    }
}

func (eb *EventBus) Subscribe(eventType string, handler Handler) {
    eb.mu.Lock()
    defer eb.mu.Unlock()
    eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *EventBus) Publish(ctx context.Context, event Event) error {
    eb.mu.RLock()
    handlers := eb.handlers[event.Type]
    eb.mu.RUnlock()
    
    for _, handler := range handlers {
        if err := handler(ctx, event); err != nil {
            return err
        }
    }
    return nil
}

func (eb *EventBus) PublishAsync(ctx context.Context, event Event) {
    eb.mu.RLock()
    handlers := eb.handlers[event.Type]
    eb.mu.RUnlock()
    
    for _, handler := range handlers {
        go handler(ctx, event)
    }
}
```

### Kafka Integration

```go
package messaging

import (
    "context"
    "encoding/json"
    
    "github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
    writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
    return &KafkaProducer{
        writer: &kafka.Writer{
            Addr:     kafka.TCP(brokers...),
            Topic:    topic,
            Balancer: &kafka.LeastBytes{},
        },
    }
}

func (p *KafkaProducer) Publish(ctx context.Context, key string, value interface{}) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    
    return p.writer.WriteMessages(ctx, kafka.Message{
        Key:   []byte(key),
        Value: data,
    })
}

func (p *KafkaProducer) Close() error {
    return p.writer.Close()
}

type KafkaConsumer struct {
    reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
    return &KafkaConsumer{
        reader: kafka.NewReader(kafka.ReaderConfig{
            Brokers:  brokers,
            Topic:    topic,
            GroupID:  groupID,
            MinBytes: 10e3, // 10KB
            MaxBytes: 10e6, // 10MB
        }),
    }
}

func (c *KafkaConsumer) Consume(ctx context.Context, handler func(key, value []byte) error) error {
    for {
        msg, err := c.reader.ReadMessage(ctx)
        if err != nil {
            return err
        }
        
        if err := handler(msg.Key, msg.Value); err != nil {
            // Log error but continue processing
            continue
        }
    }
}

func (c *KafkaConsumer) Close() error {
    return c.reader.Close()
}
```

---

## 35. gRPC Services

### Proto Definition

```protobuf
// user.proto
syntax = "proto3";

package user;

option go_package = "github.com/myapp/proto/user";

service UserService {
    rpc CreateUser(CreateUserRequest) returns (UserResponse);
    rpc GetUser(GetUserRequest) returns (UserResponse);
    rpc UpdateUser(UpdateUserRequest) returns (UserResponse);
    rpc DeleteUser(DeleteUserRequest) returns (Empty);
    rpc ListUsers(ListUsersRequest) returns (ListUsersResponse);
    rpc StreamUsers(StreamUsersRequest) returns (stream UserResponse);
}

message User {
    string id = 1;
    string email = 2;
    string name = 3;
    int64 created_at = 4;
    int64 updated_at = 5;
}

message CreateUserRequest {
    string email = 1;
    string name = 2;
    string password = 3;
}

message GetUserRequest {
    string id = 1;
}

message UpdateUserRequest {
    string id = 1;
    string email = 2;
    string name = 3;
}

message DeleteUserRequest {
    string id = 1;
}

message ListUsersRequest {
    int32 offset = 1;
    int32 limit = 2;
}

message ListUsersResponse {
    repeated User users = 1;
    int32 total = 2;
}

message StreamUsersRequest {
    int32 batch_size = 1;
}

message UserResponse {
    User user = 1;
}

message Empty {}
```

### gRPC Server

```go
package main

import (
    "context"
    "log"
    "net"
    
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    
    pb "github.com/myapp/proto/user"
)

type userServer struct {
    pb.UnimplementedUserServiceServer
}

func (s *userServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
    // Implementation
    user := &pb.User{
        Id:    "generated-id",
        Email: req.Email,
        Name:  req.Name,
    }
    return &pb.UserResponse{User: user}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
    // Implementation
    return &pb.UserResponse{}, nil
}

func (s *userServer) StreamUsers(req *pb.StreamUsersRequest, stream pb.UserService_StreamUsersServer) error {
    // Streaming implementation
    users := []*pb.User{
        {Id: "1", Name: "User 1"},
        {Id: "2", Name: "User 2"},
    }
    
    for _, user := range users {
        if err := stream.Send(&pb.UserResponse{User: user}); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &userServer{})
    
    // Enable reflection for tools like grpcurl
    reflection.Register(s)
    
    log.Println("gRPC server starting on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

---

## 36. Database Patterns

### Repository Pattern with PostgreSQL

```go
package repository

import (
    "context"
    "database/sql"
    "time"
    
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    
    "myapp/internal/domain"
)

type PostgresUserRepository struct {
    db *sqlx.DB
}

func NewPostgresUserRepository(dsn string) (*PostgresUserRepository, error) {
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    return &PostgresUserRepository{db: db}, nil
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *domain.User) error {
    query := `
        INSERT INTO users (id, email, name, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `
    _, err := r.db.ExecContext(ctx, query, 
        user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt)
    return err
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
    var user domain.User
    query := `SELECT id, email, name, created_at, updated_at FROM users WHERE id = $1`
    err := r.db.GetContext(ctx, &user, query, id)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &user, err
}

func (r *PostgresUserRepository) List(ctx context.Context, offset, limit int) ([]*domain.User, error) {
    var users []*domain.User
    query := `
        SELECT id, email, name, created_at, updated_at 
        FROM users 
        ORDER BY created_at DESC 
        LIMIT $1 OFFSET $2
    `
    err := r.db.SelectContext(ctx, &users, query, limit, offset)
    return users, err
}

// Transaction support
func (r *PostgresUserRepository) WithTransaction(ctx context.Context, fn func(*sqlx.Tx) error) error {
    tx, err := r.db.BeginTxx(ctx, nil)
    if err != nil {
        return err
    }
    
    if err := fn(tx); err != nil {
        tx.Rollback()
        return err
    }
    
    return tx.Commit()
}
```

### CQRS Pattern

```go
package cqrs

import (
    "context"
)

// Commands
type CreateUserCommand struct {
    Email    string
    Name     string
    Password string
}

type UpdateUserCommand struct {
    ID    string
    Email string
    Name  string
}

// Queries
type GetUserQuery struct {
    ID string
}

type ListUsersQuery struct {
    Offset int
    Limit  int
}

// Command Handler
type CommandHandler interface {
    HandleCreateUser(ctx context.Context, cmd CreateUserCommand) (string, error)
    HandleUpdateUser(ctx context.Context, cmd UpdateUserCommand) error
}

// Query Handler
type QueryHandler interface {
    HandleGetUser(ctx context.Context, query GetUserQuery) (*UserDTO, error)
    HandleListUsers(ctx context.Context, query ListUsersQuery) ([]*UserDTO, error)
}

// DTO for queries
type UserDTO struct {
    ID        string `json:"id"`
    Email     string `json:"email"`
    Name      string `json:"name"`
    CreatedAt string `json:"created_at"`
}

// Implementation
type userCommandHandler struct {
    writeRepo WriteRepository
    eventBus  EventBus
}

type userQueryHandler struct {
    readRepo ReadRepository
}
```

---

## 37. Observability

### Structured Logging with Zap

```go
package logger

import (
    "context"
    
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

type Logger struct {
    *zap.SugaredLogger
}

func New(environment string) (*Logger, error) {
    var config zap.Config
    
    if environment == "production" {
        config = zap.NewProductionConfig()
        config.EncoderConfig.TimeKey = "timestamp"
        config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    } else {
        config = zap.NewDevelopmentConfig()
        config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    }
    
    logger, err := config.Build()
    if err != nil {
        return nil, err
    }
    
    return &Logger{logger.Sugar()}, nil
}

func (l *Logger) WithRequestID(requestID string) *Logger {
    return &Logger{l.With("request_id", requestID)}
}

func (l *Logger) WithUserID(userID string) *Logger {
    return &Logger{l.With("user_id", userID)}
}

// Context-aware logging
type contextKey string

const loggerKey contextKey = "logger"

func WithContext(ctx context.Context, logger *Logger) context.Context {
    return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *Logger {
    if logger, ok := ctx.Value(loggerKey).(*Logger); ok {
        return logger
    }
    // Return default logger
    l, _ := New("development")
    return l
}
```

### Distributed Tracing with OpenTelemetry

```go
package tracing

import (
    "context"
    
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
    "go.opentelemetry.io/otel/trace"
)

func InitTracer(serviceName, jaegerEndpoint string) (*sdktrace.TracerProvider, error) {
    exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(
        jaeger.WithEndpoint(jaegerEndpoint),
    ))
    if err != nil {
        return nil, err
    }
    
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
        )),
    )
    
    otel.SetTracerProvider(tp)
    return tp, nil
}

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
    tracer := otel.Tracer("myapp")
    return tracer.Start(ctx, name)
}

// Usage example
func ProcessRequest(ctx context.Context) error {
    ctx, span := StartSpan(ctx, "ProcessRequest")
    defer span.End()
    
    // Add attributes
    span.SetAttributes(
        attribute.String("user.id", "123"),
        attribute.Int("items.count", 5),
    )
    
    // Call other services
    if err := callDatabase(ctx); err != nil {
        span.RecordError(err)
        return err
    }
    
    return nil
}

func callDatabase(ctx context.Context) error {
    ctx, span := StartSpan(ctx, "DatabaseQuery")
    defer span.End()
    
    // Database operation
    return nil
}
```

### Metrics with Prometheus

```go
package metrics

import (
    "net/http"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestsTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )
    
    requestDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Duration of HTTP requests",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "endpoint"},
    )
    
    activeConnections = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "active_connections",
            Help: "Number of active connections",
        },
    )
    
    requestSize = promauto.NewSummary(
        prometheus.SummaryOpts{
            Name:       "http_request_size_bytes",
            Help:       "Size of HTTP requests",
            Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
        },
    )
)

func RecordRequest(method, endpoint, status string, duration float64) {
    requestsTotal.WithLabelValues(method, endpoint, status).Inc()
    requestDuration.WithLabelValues(method, endpoint).Observe(duration)
}

func IncrementConnections() {
    activeConnections.Inc()
}

func DecrementConnections() {
    activeConnections.Dec()
}

func Handler() http.Handler {
    return promhttp.Handler()
}
```

---

## 38. Security Best Practices

### JWT Authentication

```go
package auth

import (
    "errors"
    "time"
    
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID string   `json:"user_id"`
    Email  string   `json:"email"`
    Roles  []string `json:"roles"`
    jwt.RegisteredClaims
}

type JWTManager struct {
    secretKey     []byte
    accessExpiry  time.Duration
    refreshExpiry time.Duration
}

func NewJWTManager(secretKey string, accessExpiry, refreshExpiry time.Duration) *JWTManager {
    return &JWTManager{
        secretKey:     []byte(secretKey),
        accessExpiry:  accessExpiry,
        refreshExpiry: refreshExpiry,
    }
}

func (m *JWTManager) GenerateAccessToken(userID, email string, roles []string) (string, error) {
    claims := &Claims{
        UserID: userID,
        Email:  email,
        Roles:  roles,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.accessExpiry)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "myapp",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(m.secretKey)
}

func (m *JWTManager) GenerateRefreshToken(userID string) (string, error) {
    claims := &jwt.RegisteredClaims{
        Subject:   userID,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.refreshExpiry)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
        Issuer:    "myapp",
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(m.secretKey)
}

func (m *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return m.secretKey, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, errors.New("invalid token")
}
```

### Password Hashing

```go
package auth

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

### Input Validation

```go
package validation

import (
    "regexp"
    "unicode"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateEmail(email string) bool {
    return emailRegex.MatchString(email)
}

func ValidatePassword(password string) []string {
    var errors []string
    
    if len(password) < 8 {
        errors = append(errors, "Password must be at least 8 characters")
    }
    
    var hasUpper, hasLower, hasNumber, hasSpecial bool
    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsNumber(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }
    
    if !hasUpper {
        errors = append(errors, "Password must contain at least one uppercase letter")
    }
    if !hasLower {
        errors = append(errors, "Password must contain at least one lowercase letter")
    }
    if !hasNumber {
        errors = append(errors, "Password must contain at least one number")
    }
    if !hasSpecial {
        errors = append(errors, "Password must contain at least one special character")
    }
    
    return errors
}
```

---

# Practice Questions with Answers

## Beginner Level Questions

### Q1: What is the zero value for different types in Go?

**Answer:**
- `int`, `float64`: `0`
- `string`: `""` (empty string)
- `bool`: `false`
- `pointer`: `nil`
- `slice`, `map`, `channel`: `nil`
- `struct`: Zero value of each field

```go
var i int       // 0
var s string    // ""
var b bool      // false
var p *int      // nil
var sl []int    // nil
var m map[string]int // nil
```

---

### Q2: What is the difference between `var` and `:=`?

**Answer:**
- `var` can be used anywhere (package or function level)
- `:=` (short declaration) can only be used inside functions
- `var` can declare without initialization
- `:=` must have initial value

```go
// Package level - only var works
var globalVar = 10

func main() {
    var a int      // Declaration without initialization
    var b = 20     // Declaration with type inference
    c := 30        // Short declaration (only in functions)
}
```

---

### Q3: How do you create a slice with a specific capacity?

**Answer:**
```go
// Using make: make([]T, length, capacity)
slice := make([]int, 0, 10)  // length=0, capacity=10

// Check length and capacity
fmt.Println(len(slice))  // 0
fmt.Println(cap(slice))  // 10
```

---

### Q4: What is a defer statement and when does it execute?

**Answer:**
`defer` schedules a function call to be executed after the surrounding function returns. Deferred calls are executed in LIFO (Last In, First Out) order.

```go
func main() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    fmt.Println("Main function")
}
// Output:
// Main function
// Second defer
// First defer
```

Common use cases:
- Closing files
- Unlocking mutexes
- Cleanup operations

---

### Q5: What is the difference between arrays and slices?

**Answer:**

| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed at compile time | Dynamic |
| Declaration | `[5]int` | `[]int` |
| Value type | Yes (copied on assignment) | Reference type |
| Can use `append` | No | Yes |

```go
// Array
arr := [3]int{1, 2, 3}
arr2 := arr  // Creates a copy

// Slice
slice := []int{1, 2, 3}
slice2 := slice  // Both point to same underlying array
```

---

## Intermediate Level Questions

### Q6: Explain the difference between buffered and unbuffered channels.

**Answer:**

**Unbuffered Channel:**
- Has no capacity
- Send blocks until receive is ready
- Provides synchronization

```go
ch := make(chan int)  // Unbuffered
go func() {
    ch <- 42  // Blocks until received
}()
value := <-ch  // Receives and unblocks sender
```

**Buffered Channel:**
- Has specific capacity
- Send blocks only when buffer is full
- Receive blocks only when buffer is empty

```go
ch := make(chan int, 3)  // Buffered with capacity 3
ch <- 1  // Doesn't block
ch <- 2  // Doesn't block
ch <- 3  // Doesn't block
ch <- 4  // Blocks! Buffer is full
```

---

### Q7: What is a goroutine leak and how do you prevent it?

**Answer:**
A goroutine leak occurs when a goroutine is started but never terminated, consuming memory indefinitely.

**Common causes:**
1. Blocked channel operations
2. Infinite loops without exit conditions
3. Waiting on resources that never become available

**Prevention:**
```go
// Use context for cancellation
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return  // Exit when context is cancelled
        default:
            // Do work
        }
    }
}

// Use done channel
func worker(done <-chan struct{}) {
    for {
        select {
        case <-done:
            return
        default:
            // Do work
        }
    }
}

// Close channels when done
func producer(ch chan<- int) {
    defer close(ch)
    for i := 0; i < 10; i++ {
        ch <- i
    }
}
```

---

### Q8: Explain interface satisfaction in Go.

**Answer:**
In Go, interface satisfaction is implicit. A type satisfies an interface by implementing all its methods - no explicit declaration needed.

```go
type Writer interface {
    Write([]byte) (int, error)
}

// File satisfies Writer by implementing Write
type File struct{}

func (f File) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}

// No explicit "implements Writer" declaration needed
var w Writer = File{}  // Works because File has Write method
```

Empty interface `interface{}` is satisfied by all types.

---

### Q9: What is a race condition and how do you detect it?

**Answer:**
A race condition occurs when multiple goroutines access shared data concurrently, and at least one access is a write.

**Detection:**
```bash
go run -race main.go
go test -race ./...
```

**Prevention:**
```go
// Method 1: Mutex
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

// Method 2: Channels
func incrementWithChannel(ch chan int) {
    val := <-ch
    ch <- val + 1
}

// Method 3: Atomic operations
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}
```

---

### Q10: Explain the select statement.

**Answer:**
`select` allows a goroutine to wait on multiple channel operations. It blocks until one of its cases can proceed.

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case ch3 <- value:
    fmt.Println("Sent to ch3")
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No channel ready")  // Non-blocking
}
```

**Key behaviors:**
- If multiple cases are ready, one is chosen randomly
- `default` makes select non-blocking
- Can use `time.After` for timeouts

---

## Advanced Level Questions

### Q11: Explain escape analysis in Go.

**Answer:**
Escape analysis is a compile-time analysis that determines whether a variable can be safely allocated on the stack or must be allocated on the heap.

**Variables escape to heap when:**
1. They're returned from a function
2. Stored in a global variable
3. Sent to a channel
4. Stored in an interface
5. Size is too large or unknown at compile time

```bash
# See escape analysis
go build -gcflags="-m" main.go
```

```go
// Escapes to heap (returned)
func newInt() *int {
    x := 42
    return &x  // x escapes to heap
}

// Stays on stack
func sumInt() int {
    x := 42
    return x  // x stays on stack
}
```

---

### Q12: What is the difference between value and pointer receivers?

**Answer:**

| Feature | Value Receiver | Pointer Receiver |
|---------|---------------|------------------|
| Modifies original | No | Yes |
| Copies data | Yes | No (copies pointer) |
| Can be nil | No | Yes |
| Use when | Small, immutable data | Large data, need mutation |

```go
type Counter struct {
    value int
}

// Value receiver - doesn't modify original
func (c Counter) IncrementValue() {
    c.value++  // Only modifies copy
}

// Pointer receiver - modifies original
func (c *Counter) IncrementPointer() {
    c.value++  // Modifies original
}

func main() {
    c := Counter{value: 0}
    c.IncrementValue()
    fmt.Println(c.value)  // 0
    
    c.IncrementPointer()
    fmt.Println(c.value)  // 1
}
```

**Best Practice:** Be consistent - if one method needs pointer receiver, use pointer receivers for all methods.

---

### Q13: Explain context package and its use cases.

**Answer:**
The `context` package provides:
1. Cancellation signals
2. Deadlines/timeouts
3. Request-scoped values

```go
// Cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Deadline
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Values (use sparingly)
ctx := context.WithValue(context.Background(), "userID", "123")
userID := ctx.Value("userID").(string)

// Check if done
select {
case <-ctx.Done():
    return ctx.Err()  // context.Canceled or context.DeadlineExceeded
default:
    // Continue work
}
```

---

### Q14: What are generics and how do you use type constraints?

**Answer:**
Generics (Go 1.18+) allow writing functions and types that work with any type satisfying specified constraints.

```go
// Built-in constraints from golang.org/x/exp/constraints
// comparable, ordered, signed, unsigned, float, complex, integer

// Custom constraint
type Number interface {
    int | int32 | int64 | float32 | float64
}

// Generic function
func Sum[T Number](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// Generic type
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

// Constraint with methods
type Stringer interface {
    String() string
}

func PrintAll[T Stringer](items []T) {
    for _, item := range items {
        fmt.Println(item.String())
    }
}
```

---

### Q15: How do you implement graceful shutdown?

**Answer:**
```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    server := &http.Server{
        Addr: ":8080",
    }
    
    // Channel to receive shutdown signal
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGTERM)
    
    // Start server in goroutine
    go func() {
        if err := server.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("Server error: %v", err)
        }
    }()
    
    log.Println("Server started on :8080")
    
    // Wait for shutdown signal
    <-done
    log.Println("Shutting down server...")
    
    // Create context with timeout for shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Gracefully shutdown
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server shutdown failed: %v", err)
    }
    
    log.Println("Server stopped gracefully")
}
```

---

## Expert Level Questions

### Q16: Explain the Go scheduler (GMP model).

**Answer:**
Go uses a M:N scheduler with three entities:

**G (Goroutine):** Lightweight thread managed by Go runtime
**M (Machine):** OS thread
**P (Processor):** Logical processor, holds context for executing Go code

**Key concepts:**
- Each P has a local run queue of goroutines
- There's also a global run queue
- GOMAXPROCS sets number of Ps
- Work stealing: idle P steals work from busy P's queue

**Scheduling triggers:**
- `go` statement creates new G
- Garbage collection
- System calls
- Channel operations
- `runtime.Gosched()`

```go
// Check and set GOMAXPROCS
runtime.GOMAXPROCS(4)
numProcs := runtime.GOMAXPROCS(0)

// Get number of goroutines
numGoroutines := runtime.NumGoroutine()
```

---

### Q17: How do you implement custom memory allocation?

**Answer:**
Use `sync.Pool` for object reuse:

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        buf := make([]byte, 4096)
        return &buf
    },
}

func getBuffer() *[]byte {
    return bufferPool.Get().(*[]byte)
}

func putBuffer(buf *[]byte) {
    // Reset buffer before returning
    *buf = (*buf)[:0]
    bufferPool.Put(buf)
}

// Usage
func processData() {
    buf := getBuffer()
    defer putBuffer(buf)
    
    // Use buffer...
}
```

For arena-style allocation (Go 1.20+):
```go
import "arena"

func processWithArena() {
    a := arena.NewArena()
    defer a.Free()
    
    // Allocate from arena
    data := arena.MakeSlice[byte](a, 0, 1024)
}
```

---

### Q18: Explain the difference between sync.Mutex and sync.RWMutex.

**Answer:**

**sync.Mutex:** Exclusive lock for both reads and writes
```go
var mu sync.Mutex
var data int

func write(val int) {
    mu.Lock()
    data = val
    mu.Unlock()
}

func read() int {
    mu.Lock()
    defer mu.Unlock()
    return data
}
```

**sync.RWMutex:** Allows multiple readers OR one writer
```go
var rwmu sync.RWMutex
var data int

func write(val int) {
    rwmu.Lock()       // Exclusive lock
    data = val
    rwmu.Unlock()
}

func read() int {
    rwmu.RLock()      // Shared lock
    defer rwmu.RUnlock()
    return data
}
```

**Use RWMutex when:**
- Reads are much more frequent than writes
- Read operations don't modify state

---

### Q19: How do you implement dependency injection in Go?

**Answer:**
```go
// 1. Define interfaces
type UserRepository interface {
    GetByID(ctx context.Context, id string) (*User, error)
    Create(ctx context.Context, user *User) error
}

type EmailService interface {
    Send(ctx context.Context, to, subject, body string) error
}

// 2. Create concrete implementations
type PostgresUserRepository struct {
    db *sql.DB
}

type SMTPEmailService struct {
    host string
    port int
}

// 3. Inject dependencies through constructor
type UserService struct {
    repo  UserRepository
    email EmailService
}

func NewUserService(repo UserRepository, email EmailService) *UserService {
    return &UserService{
        repo:  repo,
        email: email,
    }
}

// 4. Wire everything together (can use wire, dig, fx)
func main() {
    // Create dependencies
    db := setupDatabase()
    repo := NewPostgresUserRepository(db)
    email := NewSMTPEmailService("smtp.example.com", 587)
    
    // Inject dependencies
    userService := NewUserService(repo, email)
    
    // Use service
    userService.CreateUser(ctx, user)
}
```

---

### Q20: How do you optimize Go code for high performance?

**Answer:**

**1. Reduce allocations:**
```go
// Pre-allocate slices
slice := make([]int, 0, expectedSize)

// Use sync.Pool
var pool = sync.Pool{New: func() interface{} { return new(Buffer) }}

// Reuse strings.Builder
var sb strings.Builder
sb.Reset()
```

**2. Avoid interface{} when possible:**
```go
// Use generics instead
func Sum[T Number](nums []T) T { ... }
```

**3. Use buffered I/O:**
```go
buffered := bufio.NewReader(file)
```

**4. Optimize struct layout:**
```go
// Bad (24 bytes with padding)
type Bad struct {
    a bool   // 1 byte + 7 padding
    b int64  // 8 bytes
    c bool   // 1 byte + 7 padding
}

// Good (16 bytes)
type Good struct {
    b int64  // 8 bytes
    a bool   // 1 byte
    c bool   // 1 byte + 6 padding
}
```

**5. Use appropriate data structures:**
```go
// Map for O(1) lookups
// Slice for iteration
// sync.Map for concurrent access
```

**6. Profile and benchmark:**
```bash
go test -bench=. -benchmem
go test -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

---

## Solution Architect Level Questions

### Q21: Design a rate-limited API gateway in Go.

**Answer:**
```go
package gateway

import (
    "net/http"
    "sync"
    "time"
    
    "golang.org/x/time/rate"
)

type APIGateway struct {
    limiters map[string]*rate.Limiter
    mu       sync.RWMutex
    rate     rate.Limit
    burst    int
}

func NewAPIGateway(r rate.Limit, burst int) *APIGateway {
    return &APIGateway{
        limiters: make(map[string]*rate.Limiter),
        rate:     r,
        burst:    burst,
    }
}

func (g *APIGateway) getLimiter(key string) *rate.Limiter {
    g.mu.RLock()
    limiter, exists := g.limiters[key]
    g.mu.RUnlock()
    
    if exists {
        return limiter
    }
    
    g.mu.Lock()
    defer g.mu.Unlock()
    
    limiter = rate.NewLimiter(g.rate, g.burst)
    g.limiters[key] = limiter
    return limiter
}

func (g *APIGateway) RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        clientIP := r.RemoteAddr
        limiter := g.getLimiter(clientIP)
        
        if !limiter.Allow() {
            w.Header().Set("Retry-After", "1")
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

// Cleanup old limiters periodically
func (g *APIGateway) StartCleanup(interval time.Duration) {
    go func() {
        for range time.Tick(interval) {
            g.mu.Lock()
            // Clear all limiters (simple approach)
            g.limiters = make(map[string]*rate.Limiter)
            g.mu.Unlock()
        }
    }()
}
```

---

### Q22: Design a distributed cache with Go.

**Answer:**
```go
package cache

import (
    "context"
    "encoding/json"
    "sync"
    "time"
    
    "github.com/redis/go-redis/v9"
)

type Cache interface {
    Get(ctx context.Context, key string) ([]byte, error)
    Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
    Delete(ctx context.Context, key string) error
}

// Local cache layer
type LocalCache struct {
    items map[string]cacheItem
    mu    sync.RWMutex
}

type cacheItem struct {
    value     []byte
    expiresAt time.Time
}

func NewLocalCache() *LocalCache {
    c := &LocalCache{items: make(map[string]cacheItem)}
    go c.cleanup()
    return c
}

func (c *LocalCache) Get(ctx context.Context, key string) ([]byte, error) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, ok := c.items[key]
    if !ok || time.Now().After(item.expiresAt) {
        return nil, ErrNotFound
    }
    return item.value, nil
}

// Redis cache layer
type RedisCache struct {
    client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
    return &RedisCache{
        client: redis.NewClient(&redis.Options{Addr: addr}),
    }
}

func (c *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
    return c.client.Get(ctx, key).Bytes()
}

// Multi-level cache
type MultiLevelCache struct {
    local  Cache
    remote Cache
}

func NewMultiLevelCache(local, remote Cache) *MultiLevelCache {
    return &MultiLevelCache{local: local, remote: remote}
}

func (c *MultiLevelCache) Get(ctx context.Context, key string) ([]byte, error) {
    // Try local first
    if val, err := c.local.Get(ctx, key); err == nil {
        return val, nil
    }
    
    // Try remote
    val, err := c.remote.Get(ctx, key)
    if err != nil {
        return nil, err
    }
    
    // Populate local cache
    c.local.Set(ctx, key, val, 1*time.Minute)
    return val, nil
}
```

---

### Q23: Design a pub/sub system in Go.

**Answer:**
```go
package pubsub

import (
    "context"
    "sync"
)

type Message struct {
    Topic   string
    Payload []byte
}

type Subscriber struct {
    id      string
    channel chan Message
    topics  map[string]bool
    mu      sync.RWMutex
}

type PubSub struct {
    subscribers map[string][]*Subscriber
    mu          sync.RWMutex
}

func NewPubSub() *PubSub {
    return &PubSub{
        subscribers: make(map[string][]*Subscriber),
    }
}

func (ps *PubSub) Subscribe(topics ...string) *Subscriber {
    sub := &Subscriber{
        id:      generateID(),
        channel: make(chan Message, 100),
        topics:  make(map[string]bool),
    }
    
    ps.mu.Lock()
    defer ps.mu.Unlock()
    
    for _, topic := range topics {
        ps.subscribers[topic] = append(ps.subscribers[topic], sub)
        sub.topics[topic] = true
    }
    
    return sub
}

func (ps *PubSub) Publish(ctx context.Context, topic string, payload []byte) error {
    ps.mu.RLock()
    subs := ps.subscribers[topic]
    ps.mu.RUnlock()
    
    msg := Message{Topic: topic, Payload: payload}
    
    for _, sub := range subs {
        select {
        case sub.channel <- msg:
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Channel full, skip or log
        }
    }
    
    return nil
}

func (ps *PubSub) Unsubscribe(sub *Subscriber, topics ...string) {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    
    for _, topic := range topics {
        subs := ps.subscribers[topic]
        for i, s := range subs {
            if s.id == sub.id {
                ps.subscribers[topic] = append(subs[:i], subs[i+1:]...)
                break
            }
        }
        delete(sub.topics, topic)
    }
}

func (s *Subscriber) Messages() <-chan Message {
    return s.channel
}

func (s *Subscriber) Close() {
    close(s.channel)
}
```

---

### Q24: How would you design a microservices architecture in Go?

**Answer:**

**Architecture Components:**

1. **API Gateway**
   - Rate limiting
   - Authentication
   - Request routing
   - Load balancing

2. **Service Discovery**
   - Consul/etcd for service registration
   - Health checks
   - Dynamic routing

3. **Services**
   - Clean architecture
   - gRPC for internal communication
   - REST for external APIs

4. **Message Queue**
   - Kafka/RabbitMQ for async communication
   - Event sourcing
   - Saga pattern for transactions

5. **Observability**
   - Prometheus for metrics
   - Jaeger for tracing
   - ELK for logging

**Project Structure:**
```
microservices/
├── api-gateway/
├── user-service/
├── order-service/
├── payment-service/
├── notification-service/
├── shared/
│   ├── proto/
│   ├── config/
│   ├── middleware/
│   └── utils/
├── infrastructure/
│   ├── kubernetes/
│   └── terraform/
└── docker-compose.yml
```

---

### Q25: Design a high-availability system with Go.

**Answer:**

**Key Patterns:**

1. **Leader Election:**
```go
func (n *Node) runElection(ctx context.Context) {
    ticker := time.NewTicker(n.electionTimeout)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            if n.isLeader() {
                n.sendHeartbeats(ctx)
            } else if n.heartbeatMissed() {
                n.startElection(ctx)
            }
        case <-ctx.Done():
            return
        }
    }
}
```

2. **Health Checks:**
```go
type HealthChecker struct {
    checks map[string]Check
}

type Check func(ctx context.Context) error

func (h *HealthChecker) Run(ctx context.Context) HealthStatus {
    results := make(map[string]error)
    var wg sync.WaitGroup
    
    for name, check := range h.checks {
        wg.Add(1)
        go func(n string, c Check) {
            defer wg.Done()
            results[n] = c(ctx)
        }(name, check)
    }
    
    wg.Wait()
    return h.aggregate(results)
}
```

3. **Circuit Breaker with Fallback:**
```go
func (cb *CircuitBreaker) Execute(primary, fallback func() error) error {
    if cb.IsOpen() {
        return fallback()
    }
    
    err := primary()
    if err != nil {
        cb.RecordFailure()
        return fallback()
    }
    
    cb.RecordSuccess()
    return nil
}
```

4. **Retry with Exponential Backoff:**
```go
func RetryWithBackoff(ctx context.Context, fn func() error, maxRetries int) error {
    var err error
    for i := 0; i < maxRetries; i++ {
        err = fn()
        if err == nil {
            return nil
        }
        
        backoff := time.Duration(math.Pow(2, float64(i))) * 100 * time.Millisecond
        select {
        case <-time.After(backoff):
        case <-ctx.Done():
            return ctx.Err()
        }
    }
    return err
}
```

---

## Summary

This comprehensive guide covers Go from beginner to solution architect level:

- **Beginner:** Syntax, data types, control flow, functions, basic data structures
- **Intermediate:** Interfaces, concurrency, testing, packages, HTTP
- **Advanced:** Generics, reflection, profiling, memory management
- **Expert:** Runtime internals, CGO, plugins, distributed patterns
- **Solution Architect:** Microservices, event-driven architecture, high availability

**Next Steps:**
1. Practice each concept with hands-on coding
2. Build small projects progressing to larger ones
3. Contribute to open-source Go projects
4. Read the Go source code for deeper understanding
5. Stay updated with Go releases and proposals

---

*Document Version: 1.0*
*Last Updated: January 2026*
