## MONTH 1: Go Language Foundations

### Week 1: Core Language Basics

**Why Go?** Cloud-native default language, excellent for high-performance backends, strong backend credibility.

#### Topics to Master:

**Variables & Types:**
- [x] Variable declaration (var, :=, const)

  **Explanation:** Go provides multiple ways to declare variables. `var` is the traditional way, `:=` is short declaration (type inferred), and `const` is for constants.
  ```go
  // Using var (explicit type)
  var name string = "John"
  var age int = 25
  
  // Using var (type inferred)
  var city = "New York"
  
  // Short declaration (most common inside functions)
  count := 10
  message := "Hello"
  
  // Constants (cannot be changed)
  const Pi = 3.14159
  const MaxUsers = 100
  ```

- [x] Basic types (int, int64, float64, string, bool, byte, rune)

  **Explanation:** Go has strict typing with specific types for different purposes.
  ```go
  var i int = 42           // Platform-dependent integer (32 or 64 bit)
  var i64 int64 = 42       // Always 64-bit integer
  var f float64 = 3.14     // 64-bit floating point
  var s string = "hello"   // UTF-8 string
  var b bool = true        // Boolean
  var by byte = 'A'        // Alias for uint8 (ASCII character)
  var r rune = '世'        // Alias for int32 (Unicode code point)
  ```

- [x] Type conversions (explicit only, no implicit)

  **Explanation:** Go requires explicit type conversion - no automatic type coercion.
  ```go
  var i int = 42
  var f float64 = float64(i)  // int to float64
  var s string = strconv.Itoa(i)  // int to string
  
  // This will NOT compile:
  // var f float64 = i  // Error: cannot use i (type int) as type float64
  ```

- [x] Zero values (every type has one - CRITICAL concept)

  **Explanation:** Variables declared without initialization get their "zero value" automatically.
  ```go
  var i int       // 0
  var f float64   // 0.0
  var s string    // "" (empty string)
  var b bool      // false
  var p *int      // nil
  var sl []int    // nil
  var m map[string]int  // nil
  ```

**Data Structures (YOU WILL USE DAILY):**
- [x] Arrays (fixed size: `[5]int`)

  **Explanation:** Arrays have a fixed size defined at compile time. Size is part of the type.
  ```go
  // Declaration
  var arr [5]int                    // [0, 0, 0, 0, 0]
  arr2 := [3]string{"a", "b", "c"}  // Initialized
  arr3 := [...]int{1, 2, 3}         // Size inferred from elements
  
  // Access
  arr[0] = 10
  fmt.Println(arr[0])  // 10
  fmt.Println(len(arr))  // 5
  ```

- [x] Slices (dynamic: `[]int`) - append, copy, slicing

  **Explanation:** Slices are dynamic, flexible views into arrays. Most commonly used collection.
  ```go
  // Creation
  nums := []int{1, 2, 3}        // Slice literal
  nums2 := make([]int, 5)       // make(type, length)
  nums3 := make([]int, 0, 10)   // make(type, length, capacity)
  
  // Append (may allocate new array if capacity exceeded)
  nums = append(nums, 4, 5)     // [1, 2, 3, 4, 5]
  
  // Slicing (creates view, not copy!)
  subset := nums[1:3]           // [2, 3] (index 1 to 2)
  
  // Copy (creates independent copy)
  copied := make([]int, len(nums))
  copy(copied, nums)
  ```

- [x] Maps (`map[string]int`) - create, read, delete, check existence

  **Explanation:** Maps are key-value pairs (like objects/dictionaries in other languages).
  ```go
  // Creation
  m := make(map[string]int)     // Empty map
  m2 := map[string]int{         // Map literal
      "apple":  5,
      "banana": 3,
  }
  
  // Set value
  m["key"] = 42
  
  // Get value (returns zero value if not found)
  val := m["key"]
  
  // Check existence (IMPORTANT!)
  val, exists := m["key"]
  if exists {
      fmt.Println("Found:", val)
  }
  
  // Delete
  delete(m, "key")
  ```

- [x] Slice vs Array (know the difference!)

  **Explanation:** Arrays are fixed-size values; slices are dynamic references.
  ```go
  // Array: fixed size, copied when passed
  arr := [3]int{1, 2, 3}
  arr2 := arr        // Creates a COPY
  arr2[0] = 99       // arr is unchanged
  
  // Slice: dynamic, reference to underlying array
  sl := []int{1, 2, 3}
  sl2 := sl          // Same underlying array!
  sl2[0] = 99        // sl[0] is also 99!
  
  // Function behavior
  func modify(arr [3]int) { arr[0] = 0 }  // Receives copy
  func modify(sl []int) { sl[0] = 0 }     // Modifies original!
  ```

- [x] Nil slices vs empty slices

  **Explanation:** Nil slice has no underlying array; empty slice points to empty array.
  ```go
  var nilSlice []int           // nil slice (len=0, cap=0, nil)
  emptySlice := []int{}        // empty slice (len=0, cap=0, not nil)
  emptySlice2 := make([]int, 0) // empty slice
  
  fmt.Println(nilSlice == nil)   // true
  fmt.Println(emptySlice == nil) // false
  
  // Both work the same for most operations!
  nilSlice = append(nilSlice, 1)   // Works fine
  emptySlice = append(emptySlice, 1) // Works fine
  
  // JSON difference:
  // nil slice -> null
  // empty slice -> []
  ```

**Control Flow:**
- [x] if/else (with init statement)

  **Explanation:** Go's if can include an initialization statement before the condition.
  ```go
  // Basic if/else
  if x > 10 {
      fmt.Println("big")
  } else if x > 5 {
      fmt.Println("medium")
  } else {
      fmt.Println("small")
  }
  
  // With init statement (variable scoped to if block)
  if val, err := doSomething(); err != nil {
      fmt.Println("Error:", err)
  } else {
      fmt.Println("Success:", val)
  }
  // val and err not accessible here
  ```

- [x] for loops (Go's only loop - 3 forms)

  **Explanation:** `for` is the only loop in Go, but it has multiple forms.
  ```go
  // Form 1: Traditional for loop
  for i := 0; i < 10; i++ {
      fmt.Println(i)
  }
  
  // Form 2: While-style loop
  count := 0
  for count < 10 {
      count++
  }
  
  // Form 3: Infinite loop
  for {
      // runs forever until break
      if done {
          break
      }
  }
  ```

- [x] switch (no break needed, case fallthrough)

  **Explanation:** Go's switch doesn't fall through by default (unlike C/Java).
  ```go
  // Basic switch
  switch day {
  case "Monday":
      fmt.Println("Start of week")
  case "Friday":
      fmt.Println("Almost weekend")
  default:
      fmt.Println("Regular day")
  }
  
  // Multiple values per case
  switch day {
  case "Saturday", "Sunday":
      fmt.Println("Weekend!")
  }
  
  // Explicit fallthrough (rare)
  switch num {
  case 1:
      fmt.Println("one")
      fallthrough  // continues to next case
  case 2:
      fmt.Println("two")
  }
  
  // Switch without expression (like if-else chain)
  switch {
  case x > 10:
      fmt.Println("big")
  case x > 5:
      fmt.Println("medium")
  }
  ```

- [x] range (iterating slices, maps, strings)

  **Explanation:** `range` iterates over slices, maps, strings, and channels.
  ```go
  // Slice: index, value
  nums := []int{10, 20, 30}
  for i, v := range nums {
      fmt.Printf("index %d: %d\n", i, v)
  }
  for _, v := range nums { }  // ignore index
  for i := range nums { }     // index only
  
  // Map: key, value (random order!)
  m := map[string]int{"a": 1, "b": 2}
  for k, v := range m {
      fmt.Printf("%s: %d\n", k, v)
  }
  
  // String: index, rune (Unicode code point)
  for i, r := range "Hello世界" {
      fmt.Printf("%d: %c\n", i, r)
  }
  ```

**Functions:**
- [x] Multiple return values

  **Explanation:** Go functions can return multiple values - commonly used for value+error.
  ```go
  func divide(a, b float64) (float64, error) {
      if b == 0 {
          return 0, errors.New("division by zero")
      }
      return a / b, nil
  }
  
  // Calling
  result, err := divide(10, 2)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println(result)  // 5
  ```

- [x] Named return values

  **Explanation:** Return values can be named and used as variables, enabling "naked return".
  ```go
  func split(sum int) (x, y int) {
      x = sum * 4 / 9
      y = sum - x
      return  // naked return - returns x and y
  }
  
  // Better for documentation:
  func getUser(id int) (user User, found bool) {
      // return values are initialized to zero values
      // makes it clear what the function returns
  }
  ```

- [x] Variadic functions (`...`)

  **Explanation:** Functions that accept variable number of arguments.
  ```go
  func sum(nums ...int) int {
      total := 0
      for _, n := range nums {
          total += n
      }
      return total
  }
  
  // Calling
  sum(1, 2, 3)      // 6
  sum(1, 2, 3, 4, 5) // 15
  
  // Spread slice into variadic
  nums := []int{1, 2, 3}
  sum(nums...)      // 6
  ```

- [x] Functions as values (closures)

  **Explanation:** Functions are first-class citizens; they can be stored in variables and passed around.
  ```go
  // Function as variable
  add := func(a, b int) int {
      return a + b
  }
  result := add(2, 3)  // 5
  
  // Closure - captures outer variables
  func counter() func() int {
      count := 0
      return func() int {
          count++  // captures 'count' from outer scope
          return count
      }
  }
  
  c := counter()
  fmt.Println(c())  // 1
  fmt.Println(c())  // 2
  fmt.Println(c())  // 3
  ```

- [x] Anonymous functions

  **Explanation:** Functions without a name, often used inline.
  ```go
  // Immediately invoked
  func() {
      fmt.Println("I run immediately")
  }()
  
  // As goroutine
  go func(msg string) {
      fmt.Println(msg)
  }("hello")
  
  // As callback
  process(data, func(result int) {
      fmt.Println("Got result:", result)
  })
  ```

**Pointers:**
- [x] & (address of) and * (dereference)

  **Explanation:** `&` gets the memory address; `*` accesses the value at an address.
  ```go
  x := 10
  p := &x      // p is a pointer to x (type *int)
  
  fmt.Println(p)   // 0xc0000140a8 (memory address)
  fmt.Println(*p)  // 10 (dereferenced value)
  
  *p = 20      // modify x through pointer
  fmt.Println(x)   // 20
  ```

- [x] When to use pointers vs values

  **Explanation:** Use pointers when you need to modify the original or for large structs.
  ```go
  // Use POINTERS when:
  // 1. You need to modify the original
  func double(n *int) {
      *n *= 2
  }
  
  // 2. Struct is large (avoid copying)
  func processLargeStruct(s *LargeStruct) { }
  
  // 3. Need to represent "no value" (nil)
  func findUser(id int) *User {
      return nil  // user not found
  }
  
  // Use VALUES when:
  // 1. Small types (int, bool, small structs)
  // 2. You want immutability
  // 3. Thread safety (each goroutine gets its own copy)
  ```

- [x] Pointer receivers vs value receivers

  **Explanation:** Methods can have pointer or value receivers, affecting mutation and performance.
  ```go
  type Counter struct {
      value int
  }
  
  // Value receiver - gets COPY of struct
  func (c Counter) GetValue() int {
      return c.value
  }
  
  // Pointer receiver - can MODIFY struct
  func (c *Counter) Increment() {
      c.value++  // modifies the original
  }
  
  c := Counter{value: 0}
  c.Increment()  // Go auto-dereferences, even though c is not pointer
  fmt.Println(c.GetValue())  // 1
  
  // Rule of thumb: If ANY method needs pointer receiver,
  // make ALL methods use pointer receiver for consistency
  ```

- [x] nil pointers

  **Explanation:** Pointers can be nil; always check before dereferencing.
  ```go
  var p *int  // nil pointer
  
  // Dereferencing nil causes panic!
  // fmt.Println(*p)  // panic: runtime error
  
  // Always check for nil
  if p != nil {
      fmt.Println(*p)
  }
  
  // Methods can be called on nil receivers (useful pattern)
  type Node struct {
      value int
      next  *Node
  }
  
  func (n *Node) String() string {
      if n == nil {
          return "<nil>"
      }
      return fmt.Sprintf("%d", n.value)
  }
  ```

#### Quick Self-Test (Week 1):
```go
// Can you explain what each line does?
nums := []int{1, 2, 3}
nums = append(nums, 4)
m := make(map[string]int)
m["key"] = 1
if val, ok := m["key"]; ok { fmt.Println(val) }
```

---

### Week 2: Go Patterns & Idioms

#### Structs & Methods:
- [x] Defining structs

  **Explanation:** Structs are custom types that group related data together.
  ```go
  type Person struct {
      Name    string
      Age     int
      Email   string
      Active  bool
  }
  
  // Creating instances
  p1 := Person{Name: "John", Age: 30}  // Named fields (preferred)
  p2 := Person{"Jane", 25, "jane@email.com", true}  // Positional (order matters)
  var p3 Person  // Zero value: {"", 0, "", false}
  ```

- [x] Struct literals (named vs positional)

  **Explanation:** Named fields are clearer and safer when struct fields change.
  ```go
  // Named (recommended) - order doesn't matter, missing fields get zero value
  user := User{
      Email: "john@example.com",
      Name:  "John",
  }
  
  // Positional (fragile) - must provide ALL fields in order
  user := User{"john@example.com", "John", 25}
  
  // Pointer to struct
  userPtr := &User{Name: "John"}
  ```

- [x] Embedded structs (composition)

  **Explanation:** Go uses composition over inheritance. Embed structs to inherit fields/methods.
  ```go
  type Address struct {
      Street string
      City   string
  }
  
  type Person struct {
      Name    string
      Address         // Embedded (anonymous field)
  }
  
  p := Person{
      Name: "John",
      Address: Address{
          Street: "123 Main St",
          City:   "NYC",
      },
  }
  
  // Access embedded fields directly (promoted)
  fmt.Println(p.City)         // "NYC" - promoted field
  fmt.Println(p.Address.City) // "NYC" - also works
  ```

- [x] Methods (value receiver vs pointer receiver)

  **Explanation:** Methods are functions with a receiver argument.
  ```go
  type Rectangle struct {
      Width, Height float64
  }
  
  // Value receiver - does NOT modify original
  func (r Rectangle) Area() float64 {
      return r.Width * r.Height
  }
  
  // Pointer receiver - CAN modify original
  func (r *Rectangle) Scale(factor float64) {
      r.Width *= factor
      r.Height *= factor  // modifies the original
  }
  
  rect := Rectangle{10, 5}
  fmt.Println(rect.Area())  // 50
  rect.Scale(2)
  fmt.Println(rect.Area())  // 200
  ```

- [x] Constructor patterns (NewXxx functions)

  **Explanation:** Go uses factory functions (convention: NewTypeName) instead of constructors.
  ```go
  type Server struct {
      host string
      port int
  }
  
  // Constructor function
  func NewServer(host string, port int) *Server {
      if port == 0 {
          port = 8080  // default
      }
      return &Server{
          host: host,
          port: port,
      }
  }
  
  // With validation
  func NewServer(host string, port int) (*Server, error) {
      if host == "" {
          return nil, errors.New("host is required")
      }
      return &Server{host: host, port: port}, nil
  }
  
  server := NewServer("localhost", 8080)
  ```

#### Interfaces:
- [ ] Implicit implementation (no `implements` keyword)

  **Explanation:** Types implement interfaces automatically by having the required methods.
  ```go
  type Speaker interface {
      Speak() string
  }
  
  type Dog struct{ Name string }
  type Cat struct{ Name string }
  
  // Dog implements Speaker (no explicit declaration!)
  func (d Dog) Speak() string {
      return "Woof!"
  }
  
  // Cat implements Speaker
  func (c Cat) Speak() string {
      return "Meow!"
  }
  
  // Now both can be used as Speaker
  func MakeSpeak(s Speaker) {
      fmt.Println(s.Speak())
  }
  
  MakeSpeak(Dog{Name: "Rex"})   // Woof!
  MakeSpeak(Cat{Name: "Whiskers"}) // Meow!
  ```

- [x] Empty interface (`interface{}` / `any`)

  **Explanation:** Empty interface accepts any type (like Object in Java or any in TypeScript).
  ```go
  // interface{} and any are equivalent (any is alias since Go 1.18)
  func printAnything(v any) {
      fmt.Println(v)
  }
  
  printAnything(42)
  printAnything("hello")
  printAnything([]int{1, 2, 3})
  
  // Commonly used in:
  var data map[string]any  // JSON-like structure
  data = map[string]any{
      "name": "John",
      "age":  30,
      "tags": []string{"go", "developer"},
  }
  ```

- [x] Type assertions (`value.(Type)`)

  **Explanation:** Extract the concrete type from an interface value.
  ```go
  var i interface{} = "hello"
  
  // Type assertion (panics if wrong type!)
  s := i.(string)
  fmt.Println(s)  // "hello"
  
  // Safe type assertion with ok check
  s, ok := i.(string)
  if ok {
      fmt.Println(s)
  }
  
  n, ok := i.(int)  // ok is false, n is 0
  if !ok {
      fmt.Println("not an int")
  }
  ```

- [x] Type switches

  **Explanation:** Switch on the type of an interface value.
  ```go
  func describe(i interface{}) {
      switch v := i.(type) {
      case int:
          fmt.Printf("int: %d\n", v)
      case string:
          fmt.Printf("string: %s\n", v)
      case bool:
          fmt.Printf("bool: %t\n", v)
      case []int:
          fmt.Printf("slice of ints: %v\n", v)
      default:
          fmt.Printf("unknown type: %T\n", v)
      }
  }
  
  describe(42)        // int: 42
  describe("hello")   // string: hello
  describe(true)      // bool: true
  ```

- [ ] Common interfaces (io.Reader, io.Writer, error, Stringer)

  **Explanation:** Standard library interfaces you'll use constantly.
  ```go
  // io.Reader - anything that can be read from
  type Reader interface {
      Read(p []byte) (n int, err error)
  }
  // Examples: files, network connections, strings.Reader
  
  // io.Writer - anything that can be written to
  type Writer interface {
      Write(p []byte) (n int, err error)
  }
  // Examples: files, http.ResponseWriter, bytes.Buffer
  
  // error - the error interface
  type error interface {
      Error() string
  }
  
  // fmt.Stringer - custom string representation
  type Stringer interface {
      String() string
  }
  
  type Person struct{ Name string }
  func (p Person) String() string {
      return fmt.Sprintf("Person(%s)", p.Name)
  }
  ```

#### Error Handling:
- [ ] The `error` interface

  **Explanation:** `error` is a built-in interface with a single method.
  ```go
  type error interface {
      Error() string
  }
  
  // Functions return error as last value
  func doSomething() (string, error) {
      return "", errors.New("something went wrong")
  }
  
  // Always check errors!
  result, err := doSomething()
  if err != nil {
      // Handle error
      log.Fatal(err)
  }
  // Use result
  ```

- [ ] Creating errors (`errors.New`, `fmt.Errorf`)

  **Explanation:** Multiple ways to create error values.
  ```go
  import "errors"
  
  // Simple error
  err := errors.New("file not found")
  
  // Formatted error with context
  err := fmt.Errorf("failed to open file %s: permission denied", filename)
  
  // Custom error type
  type ValidationError struct {
      Field   string
      Message string
  }
  
  func (e ValidationError) Error() string {
      return fmt.Sprintf("%s: %s", e.Field, e.Message)
  }
  
  return ValidationError{Field: "email", Message: "invalid format"}
  ```

- [ ] Error wrapping (`%w` and `errors.Is`, `errors.As`)

  **Explanation:** Wrap errors to add context while preserving the original error.
  ```go
  // Wrap error with context using %w
  if err != nil {
      return fmt.Errorf("failed to process user %d: %w", userID, err)
  }
  
  // Check if error is a specific error (unwraps automatically)
  if errors.Is(err, os.ErrNotExist) {
      fmt.Println("File doesn't exist")
  }
  
  // Extract specific error type
  var pathErr *os.PathError
  if errors.As(err, &pathErr) {
      fmt.Println("Path:", pathErr.Path)
  }
  
  // Unwrap to get original error
  original := errors.Unwrap(err)
  ```

- [ ] Sentinel errors (`var ErrNotFound = errors.New(...)`)

  **Explanation:** Predefined errors for common conditions that callers can check.
  ```go
  // Define sentinel errors (package-level)
  var (
      ErrNotFound     = errors.New("not found")
      ErrUnauthorized = errors.New("unauthorized")
      ErrInvalidInput = errors.New("invalid input")
  )
  
  func GetUser(id int) (*User, error) {
      user := db.Find(id)
      if user == nil {
          return nil, ErrNotFound
      }
      return user, nil
  }
  
  // Caller can check for specific errors
  user, err := GetUser(123)
  if errors.Is(err, ErrNotFound) {
      // Handle not found case
  }
  ```

- [ ] When to panic vs return error

  **Explanation:** Return errors for expected failures; panic only for programming bugs.
  ```go
  // USE ERROR RETURN for:
  // - File not found, network timeout, invalid input
  // - Anything the caller should handle
  func ReadConfig(path string) (Config, error) {
      data, err := os.ReadFile(path)
      if err != nil {
          return Config{}, err  // Let caller decide
      }
      return parseConfig(data)
  }
  
  // USE PANIC for:
  // - Programming errors (nil pointer you didn't expect)
  // - Impossible situations
  // - During initialization if app can't continue
  func MustCompileRegex(pattern string) *regexp.Regexp {
      re, err := regexp.Compile(pattern)
      if err != nil {
          panic("invalid regex: " + err.Error())
      }
      return re
  }
  
  // Convention: functions starting with "Must" panic on error
  ```

#### Important Patterns:
- [ ] defer (cleanup, runs LIFO)

  **Explanation:** `defer` schedules a function to run when the surrounding function returns.
  ```go
  func readFile(path string) error {
      f, err := os.Open(path)
      if err != nil {
          return err
      }
      defer f.Close()  // Guaranteed to run when function returns
      
      // Read file...
      return nil
  }
  
  // LIFO order (Last In, First Out)
  func demo() {
      defer fmt.Println("1")
      defer fmt.Println("2")
      defer fmt.Println("3")
  }
  // Output: 3, 2, 1
  
  // Deferred function arguments evaluated immediately
  x := 10
  defer fmt.Println(x)  // Will print 10, not 20
  x = 20
  ```

- [ ] panic and recover (rare, for truly exceptional cases)

  **Explanation:** panic stops execution; recover catches panics in deferred functions.
  ```go
  // Panic - stops normal execution
  func divide(a, b int) int {
      if b == 0 {
          panic("division by zero")  // Program crashes
      }
      return a / b
  }
  
  // Recover - catch panics (only works in defer)
  func safeDivide(a, b int) (result int, err error) {
      defer func() {
          if r := recover(); r != nil {
              err = fmt.Errorf("recovered: %v", r)
          }
      }()
      
      result = divide(a, b)
      return result, nil
  }
  
  result, err := safeDivide(10, 0)
  // err: "recovered: division by zero"
  ```

- [ ] init() functions (runs before main)

  **Explanation:** Special function that runs automatically before main().
  ```go
  package main
  
  var config Config
  
  func init() {
      // Runs before main()
      // Used for: setup, loading config, registering drivers
      config = loadConfig()
      fmt.Println("Initialized!")
  }
  
  func main() {
      fmt.Println("Main running")
      // config is already loaded
  }
  
  // Multiple init() functions allowed in same file
  // Order: package imports → package init() → main()
  ```

- [ ] Packages and modules (go mod init, go mod tidy)

  **Explanation:** Modules are how Go manages dependencies; packages organize code.
  ```bash
  # Initialize a new module
  go mod init github.com/username/myproject
  
  # Download dependencies and clean up go.mod
  go mod tidy
  
  # Download dependencies
  go mod download
  ```
  
  ```go
  // Package declaration (every file needs one)
  package main  // Executable package
  package utils // Library package
  
  // Importing packages
  import (
      "fmt"                              // Standard library
      "github.com/gin-gonic/gin"         // External package
      "myproject/internal/handler"        // Local package
  )
  ```

- [ ] Public vs private (capitalization)

  **Explanation:** Go uses capitalization for visibility, not keywords.
  ```go
  package user
  
  // Exported (public) - Capital letter
  type User struct {
      Name  string  // exported field
      Email string  // exported field
      age   int     // unexported (private to package)
  }
  
  // Exported function
  func NewUser(name string) *User {
      return &User{Name: name}
  }
  
  // Unexported function (private to package)
  func validateEmail(email string) bool {
      return strings.Contains(email, "@")
  }
  
  // From another package:
  u := user.NewUser("John")  // ✓ Works
  u.Name                      // ✓ Works
  u.age                       // ✗ Error: unexported field
  user.validateEmail()        // ✗ Error: unexported function
  ```

#### Hands-on Project (Week 1-2):
```
Project 1: CLI Task Manager
- Add, list, complete, delete tasks
- Store in JSON file
- Use flag package for CLI args
- Proper error handling throughout
- Demonstrates: structs, slices, maps, file I/O, JSON

Project 2: Concurrent File Processor
- Read multiple files concurrently
- Count words/lines in each
- Aggregate results
- Demonstrates: goroutines, channels, sync.WaitGroup
```

#### Key Differences from JavaScript/TypeScript:
| Concept | TypeScript | Go |
|---------|------------|----|
| Error handling | try/catch | Multiple return values |
| Null safety | undefined/null | nil + explicit checks |
| OOP | Classes | Structs + Interfaces |
| Async | Promises/async-await | Goroutines + Channels |
| Generics | Built-in | Added in Go 1.18 |
| Zero values | undefined | Type-specific (0, "", nil) |
| Collections | Array, Object | Slice, Map |

#### Resources:
- **Official:** https://go.dev/tour (A Tour of Go) - DO THIS FIRST
- **Practice:** https://gobyexample.com (Concise examples)
- **Book:** "Learning Go" by Jon Bodner
- **Course:** "Go: The Complete Developer's Guide" - Udemy
- **Exercises:** https://exercism.io/tracks/go

---

### Week 3: Concurrency Deep Dive

#### Goroutines:
- [ ] Starting goroutines (`go func()`)

  **Explanation:** Goroutines are lightweight threads managed by Go runtime.
  ```go
  // Start a goroutine with 'go' keyword
  go doSomething()
  
  // Anonymous function as goroutine
  go func() {
      fmt.Println("Running in goroutine")
  }()
  
  // With parameters (capture by value!)
  for i := 0; i < 5; i++ {
      go func(n int) {  // Pass i as parameter
          fmt.Println(n)
      }(i)  // i is passed here
  }
  
  // Main must wait for goroutines to finish
  time.Sleep(time.Second)  // Bad way
  // Better: use sync.WaitGroup or channels
  ```

- [ ] Goroutine lifecycle

  **Explanation:** Goroutines run until their function returns or the program exits.
  ```go
  func main() {
      go longRunning()  // Starts goroutine
      
      // If main() exits, all goroutines are killed immediately!
      // Goroutines don't have IDs or handles - no way to "stop" them directly
      
      time.Sleep(time.Second)  // Wait for goroutine
  }
  
  func longRunning() {
      for {
          // This goroutine runs until:
          // 1. Function returns (break/return)
          // 2. Program exits
          // 3. panic (unless recovered)
      }
  }
  ```

- [ ] Race conditions (and how to detect: `go run -race`)

  **Explanation:** Race conditions occur when goroutines access shared data without synchronization.
  ```go
  // RACE CONDITION - DO NOT DO THIS
  counter := 0
  for i := 0; i < 1000; i++ {
      go func() {
          counter++  // Multiple goroutines read/write same variable!
      }()
  }
  // counter will NOT be 1000 - undefined behavior
  
  // Detect races at runtime
  // go run -race main.go
  // go test -race ./...
  
  // Fix with mutex
  var mu sync.Mutex
  counter := 0
  for i := 0; i < 1000; i++ {
      go func() {
          mu.Lock()
          counter++
          mu.Unlock()
      }()
  }
  ```

- [ ] Goroutine leaks (and how to avoid)

  **Explanation:** Leaked goroutines are goroutines that never terminate, wasting resources.
  ```go
  // LEAK: Goroutine blocked forever
  func leak() {
      ch := make(chan int)
      go func() {
          val := <-ch  // Blocked forever - nothing sends to ch!
          fmt.Println(val)
      }()
      // ch is never sent to, goroutine leaks
  }
  
  // Fix: Use context for cancellation
  func noLeak(ctx context.Context) {
      ch := make(chan int)
      go func() {
          select {
          case val := <-ch:
              fmt.Println(val)
          case <-ctx.Done():
              return  // Exit when cancelled
          }
      }()
  }
  
  // Fix: Always close channels or use buffered channels
  func noLeak2() {
      ch := make(chan int, 1)  // Buffered - sender doesn't block
      go func() {
          ch <- 42  // Won't block even if no receiver
      }()
  }
  ```

#### Channels:
- [ ] Unbuffered channels (synchronous)

  **Explanation:** Sender blocks until receiver is ready (and vice versa).
  ```go
  ch := make(chan int)  // Unbuffered channel
  
  go func() {
      ch <- 42  // Blocks until someone receives
  }()
  
  val := <-ch  // Blocks until someone sends
  fmt.Println(val)  // 42
  
  // Synchronization point - both goroutines meet at channel operation
  ```

- [ ] Buffered channels (async up to capacity)

  **Explanation:** Sender doesn't block until buffer is full.
  ```go
  ch := make(chan int, 3)  // Buffer size of 3
  
  ch <- 1  // Doesn't block (buffer: [1])
  ch <- 2  // Doesn't block (buffer: [1, 2])
  ch <- 3  // Doesn't block (buffer: [1, 2, 3])
  // ch <- 4  // Would block! Buffer full
  
  fmt.Println(<-ch)  // 1 (buffer: [2, 3])
  fmt.Println(<-ch)  // 2 (buffer: [3])
  
  fmt.Println(len(ch))  // 1 - items in buffer
  fmt.Println(cap(ch))  // 3 - buffer capacity
  ```

- [ ] Sending and receiving

  **Explanation:** The `<-` operator is used for both sending and receiving.
  ```go
  ch := make(chan string, 1)
  
  // Send (arrow points INTO channel)
  ch <- "hello"
  
  // Receive (arrow points OUT of channel)
  msg := <-ch
  
  // Receive and discard
  <-ch
  
  // Receive with ok (check if channel closed)
  val, ok := <-ch
  if !ok {
      fmt.Println("Channel closed")
  }
  ```

- [ ] Closing channels

  **Explanation:** Close signals that no more values will be sent.
  ```go
  ch := make(chan int)
  
  go func() {
      for i := 0; i < 5; i++ {
          ch <- i
      }
      close(ch)  // Signal: no more values coming
  }()
  
  // Receiving from closed channel returns zero value immediately
  val := <-ch  // 0 if closed
  
  // Check if closed
  val, ok := <-ch
  if !ok {
      fmt.Println("Closed!")
  }
  
  // Only sender should close - never close from receiver
  // Closing already closed channel = panic!
  // Sending to closed channel = panic!
  ```

- [ ] Range over channels

  **Explanation:** Range receives values until channel is closed.
  ```go
  ch := make(chan int)
  
  go func() {
      for i := 0; i < 5; i++ {
          ch <- i
      }
      close(ch)  // Must close for range to terminate!
  }()
  
  // Range automatically breaks when channel closes
  for val := range ch {
      fmt.Println(val)  // 0, 1, 2, 3, 4
  }
  
  // Without close(), range would block forever!
  ```

- [ ] Directional channels (`chan<-`, `<-chan`)

  **Explanation:** Restrict channel usage to send-only or receive-only.
  ```go
  // Bidirectional (default)
  ch := make(chan int)
  
  // Send-only: can only send
  func send(ch chan<- int) {
      ch <- 42
      // <-ch  // Error: cannot receive from send-only channel
  }
  
  // Receive-only: can only receive
  func receive(ch <-chan int) {
      val := <-ch
      // ch <- 1  // Error: cannot send to receive-only channel
  }
  
  // Commonly used in function signatures for safety
  func producer() <-chan int {  // Returns receive-only
      ch := make(chan int)
      go func() {
          ch <- 42
          close(ch)
      }()
      return ch
  }
  ```

#### Select Statement:
- [ ] Multiplexing channels

  **Explanation:** Select waits on multiple channel operations simultaneously.
  ```go
  ch1 := make(chan string)
  ch2 := make(chan string)
  
  go func() { ch1 <- "from 1" }()
  go func() { ch2 <- "from 2" }()
  
  // Wait for whichever is ready first
  select {
  case msg1 := <-ch1:
      fmt.Println(msg1)
  case msg2 := <-ch2:
      fmt.Println(msg2)
  }
  
  // If multiple ready, one is chosen randomly
  // Loop to receive from multiple channels
  for i := 0; i < 2; i++ {
      select {
      case msg := <-ch1:
          fmt.Println("ch1:", msg)
      case msg := <-ch2:
          fmt.Println("ch2:", msg)
      }
  }
  ```

- [ ] Default case (non-blocking)

  **Explanation:** Default executes immediately if no other case is ready.
  ```go
  ch := make(chan int)
  
  // Non-blocking receive
  select {
  case val := <-ch:
      fmt.Println("Received:", val)
  default:
      fmt.Println("No value ready")  // Executes immediately
  }
  
  // Non-blocking send
  select {
  case ch <- 42:
      fmt.Println("Sent")
  default:
      fmt.Println("Channel not ready")
  }
  
  // Polling pattern
  for {
      select {
      case val := <-ch:
          process(val)
      default:
          // Do other work
          time.Sleep(10 * time.Millisecond)
      }
  }
  ```

- [ ] Timeouts with `time.After`

  **Explanation:** time.After returns a channel that receives after duration.
  ```go
  ch := make(chan string)
  
  select {
  case result := <-ch:
      fmt.Println(result)
  case <-time.After(3 * time.Second):
      fmt.Println("Timeout!")
  }
  
  // Request with timeout pattern
  func fetchWithTimeout(url string, timeout time.Duration) (string, error) {
      ch := make(chan string, 1)
      
      go func() {
          result := fetch(url)
          ch <- result
      }()
      
      select {
      case result := <-ch:
          return result, nil
      case <-time.After(timeout):
          return "", errors.New("timeout")
      }
  }
  ```

#### Sync Package:
- [ ] sync.Mutex (protect shared state)

  **Explanation:** Mutex provides mutual exclusion for shared resources.
  ```go
  type SafeCounter struct {
      mu    sync.Mutex
      count int
  }
  
  func (c *SafeCounter) Increment() {
      c.mu.Lock()
      defer c.mu.Unlock()  // Always unlock with defer
      c.count++
  }
  
  func (c *SafeCounter) Value() int {
      c.mu.Lock()
      defer c.mu.Unlock()
      return c.count
  }
  
  // Safe for concurrent use
  counter := &SafeCounter{}
  for i := 0; i < 1000; i++ {
      go counter.Increment()
  }
  ```

- [ ] sync.RWMutex (multiple readers, single writer)

  **Explanation:** Allows multiple readers OR one writer (more efficient for read-heavy workloads).
  ```go
  type Cache struct {
      mu    sync.RWMutex
      items map[string]string
  }
  
  // Multiple goroutines can read simultaneously
  func (c *Cache) Get(key string) string {
      c.mu.RLock()          // Read lock
      defer c.mu.RUnlock()
      return c.items[key]
  }
  
  // Only one goroutine can write
  func (c *Cache) Set(key, value string) {
      c.mu.Lock()           // Write lock (exclusive)
      defer c.mu.Unlock()
      c.items[key] = value
  }
  ```

- [ ] sync.WaitGroup (wait for goroutines)

  **Explanation:** Wait for a collection of goroutines to finish.
  ```go
  var wg sync.WaitGroup
  
  for i := 0; i < 5; i++ {
      wg.Add(1)  // Increment counter BEFORE starting goroutine
      go func(n int) {
          defer wg.Done()  // Decrement when done
          fmt.Println(n)
      }(i)
  }
  
  wg.Wait()  // Block until counter is 0
  fmt.Println("All goroutines finished")
  
  // Common pattern:
  // 1. Add(n) before starting n goroutines
  // 2. Done() at end of each goroutine
  // 3. Wait() in main goroutine
  ```

- [ ] sync.Once (one-time initialization)

  **Explanation:** Ensure code runs exactly once, even with multiple goroutines.
  ```go
  var (
      once     sync.Once
      instance *Database
  )
  
  func GetDatabase() *Database {
      once.Do(func() {
          // This runs exactly once, even if called concurrently
          instance = &Database{}
          instance.Connect()
      })
      return instance
  }
  
  // Safe to call from multiple goroutines
  // First call initializes, others wait and return cached
  for i := 0; i < 10; i++ {
      go func() {
          db := GetDatabase()  // All get same instance
          db.Query()
      }()
  }
  ```

#### Context Package (CRITICAL FOR BACKEND):
- [ ] context.Background() and context.TODO()

  **Explanation:** Root contexts - starting points for context chains.
  ```go
  // Background - the root context (never cancelled)
  ctx := context.Background()
  
  // TODO - placeholder when unsure which context to use
  ctx := context.TODO()  // Same as Background, but signals "fix later"
  
  // Usage: Background at entry points (main, HTTP handlers)
  func main() {
      ctx := context.Background()
      runServer(ctx)
  }
  
  func HandleRequest(w http.ResponseWriter, r *http.Request) {
      ctx := r.Context()  // Get context from request
      processRequest(ctx)
  }
  ```

- [ ] context.WithCancel

  **Explanation:** Create a context that can be manually cancelled.
  ```go
  ctx, cancel := context.WithCancel(context.Background())
  
  go func() {
      for {
          select {
          case <-ctx.Done():
              fmt.Println("Cancelled!")
              return
          default:
              // Do work
          }
      }
  }()
  
  // Later: cancel the goroutine
  cancel()  // Signals cancellation
  
  // Always call cancel to release resources
  defer cancel()
  ```

- [ ] context.WithTimeout

  **Explanation:** Context that automatically cancels after a duration.
  ```go
  // Cancel after 5 seconds
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()  // Always call cancel!
  
  result, err := doSlowOperation(ctx)
  if err == context.DeadlineExceeded {
      fmt.Println("Operation timed out")
  }
  
  // Check timeout in function
  func doSlowOperation(ctx context.Context) error {
      select {
      case <-time.After(10 * time.Second):
          return nil  // Completed
      case <-ctx.Done():
          return ctx.Err()  // Timeout or cancelled
      }
  }
  ```

- [ ] context.WithDeadline

  **Explanation:** Context that cancels at a specific time.
  ```go
  // Cancel at specific time
  deadline := time.Now().Add(30 * time.Second)
  ctx, cancel := context.WithDeadline(context.Background(), deadline)
  defer cancel()
  
  // Check remaining time
  if deadline, ok := ctx.Deadline(); ok {
      remaining := time.Until(deadline)
      fmt.Printf("Time remaining: %v\n", remaining)
  }
  ```

- [ ] context.WithValue (use sparingly)

  **Explanation:** Pass request-scoped values through context (use carefully).
  ```go
  type key string
  const userIDKey key = "userID"
  
  // Add value to context
  ctx := context.WithValue(context.Background(), userIDKey, "user-123")
  
  // Retrieve value
  func getUser(ctx context.Context) string {
      userID, ok := ctx.Value(userIDKey).(string)
      if !ok {
          return ""
      }
      return userID
  }
  
  // Best practices:
  // - Use custom types for keys (avoid collisions)
  // - Only for request-scoped data (userID, traceID)
  // - Don't pass optional parameters this way
  // - Don't pass dependencies (use function args)
  ```

- [ ] Passing context through function calls

  **Explanation:** Context should be the first parameter, named `ctx`.
  ```go
  // Convention: context is first parameter
  func GetUser(ctx context.Context, id string) (*User, error) {
      // Pass context to called functions
      result, err := db.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
      if err != nil {
          return nil, err
      }
      return parseUser(result), nil
  }
  
  // HTTP handler example
  func handler(w http.ResponseWriter, r *http.Request) {
      ctx := r.Context()  // Get from request
      
      ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
      defer cancel()
      
      user, err := GetUser(ctx, "123")
      if err != nil {
          http.Error(w, err.Error(), 500)
          return
      }
      json.NewEncoder(w).Encode(user)
  }
  ```

#### Common Concurrency Patterns:
- [ ] Worker pools

  **Explanation:** Fixed number of workers processing jobs from a shared queue.
  ```go
  func workerPool() {
      jobs := make(chan int, 100)
      results := make(chan int, 100)
      
      // Start 3 workers
      for w := 1; w <= 3; w++ {
          go worker(w, jobs, results)
      }
      
      // Send jobs
      for j := 1; j <= 9; j++ {
          jobs <- j
      }
      close(jobs)
      
      // Collect results
      for r := 1; r <= 9; r++ {
          <-results
      }
  }
  
  func worker(id int, jobs <-chan int, results chan<- int) {
      for j := range jobs {
          fmt.Printf("Worker %d processing job %d\n", id, j)
          time.Sleep(time.Second)
          results <- j * 2
      }
  }
  ```

- [ ] Fan-out, fan-in

  **Explanation:** Fan-out: multiple goroutines read from same channel. Fan-in: merge multiple channels into one.
  ```go
  // Fan-out: distribute work
  func fanOut(input <-chan int, n int) []<-chan int {
      outputs := make([]<-chan int, n)
      for i := 0; i < n; i++ {
          outputs[i] = worker(input)
      }
      return outputs
  }
  
  // Fan-in: merge results
  func fanIn(channels ...<-chan int) <-chan int {
      out := make(chan int)
      var wg sync.WaitGroup
      
      for _, ch := range channels {
          wg.Add(1)
          go func(c <-chan int) {
              defer wg.Done()
              for v := range c {
                  out <- v
              }
          }(ch)
      }
      
      go func() {
          wg.Wait()
          close(out)
      }()
      
      return out
  }
  ```

- [ ] Pipeline

  **Explanation:** Chain stages connected by channels, each stage processes and forwards data.
  ```go
  // Stage 1: Generate numbers
  func generate(nums ...int) <-chan int {
      out := make(chan int)
      go func() {
          for _, n := range nums {
              out <- n
          }
          close(out)
      }()
      return out
  }
  
  // Stage 2: Square numbers
  func square(in <-chan int) <-chan int {
      out := make(chan int)
      go func() {
          for n := range in {
              out <- n * n
          }
          close(out)
      }()
      return out
  }
  
  // Usage: chain stages together
  func main() {
      ch := generate(2, 3, 4)
      out := square(ch)
      
      for result := range out {
          fmt.Println(result)  // 4, 9, 16
      }
  }
  ```

- [ ] Semaphore with buffered channel

  **Explanation:** Limit concurrent operations using a buffered channel.
  ```go
  // Semaphore - limit to 3 concurrent operations
  sem := make(chan struct{}, 3)
  
  for _, url := range urls {
      go func(url string) {
          sem <- struct{}{}        // Acquire (blocks if 3 already running)
          defer func() { <-sem }() // Release
          
          fetch(url)
      }(url)
  }
  
  // Or using x/sync/semaphore package
  import "golang.org/x/sync/semaphore"
  
  sem := semaphore.NewWeighted(3)
  
  for _, url := range urls {
      if err := sem.Acquire(ctx, 1); err != nil {
          return err
      }
      go func(url string) {
          defer sem.Release(1)
          fetch(url)
      }(url)
  }
  ```

#### Hands-on Project:
```
Project 3: Concurrent URL Checker
- Check multiple URLs concurrently
- Timeout after 5 seconds per URL
- Cancel all requests if user interrupts (Ctrl+C)
- Report results as they complete
- Demonstrates: goroutines, channels, context, select
```

---

### Week 4: Go Intermediate + Gin Intro

#### Standard Library Essentials:
- [ ] `io` (readers, writers, utilities)

  **Explanation:** The io package provides basic interfaces for I/O operations.
  ```go
  import "io"
  
  // io.Reader - read data from source
  func readAll(r io.Reader) ([]byte, error) {
      return io.ReadAll(r)  // Read until EOF
  }
  
  // io.Writer - write data to destination
  func writeData(w io.Writer, data []byte) error {
      _, err := w.Write(data)
      return err
  }
  
  // Copy between reader and writer
  io.Copy(dst, src)  // Copy all data from src to dst
  
  // Combine readers
  r := io.MultiReader(r1, r2, r3)  // Read from r1, then r2, then r3
  
  // Limit reading
  limitedReader := io.LimitReader(r, 1024)  // Read max 1024 bytes
  ```

- [ ] `os` (files, env vars, args, ReadFile/WriteFile)

  **Explanation:** OS package provides platform-independent OS functionality.
  ```go
  import "os"
  
  // File operations
  data, err := os.ReadFile("config.json")  // Read entire file
  err = os.WriteFile("output.txt", data, 0644)  // Write entire file
  
  // Open file for more control
  f, err := os.Open("file.txt")        // Read-only
  f, err := os.Create("file.txt")      // Create/truncate, write-only
  f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
  defer f.Close()
  
  // Environment variables
  home := os.Getenv("HOME")            // Get env var
  os.Setenv("MY_VAR", "value")         // Set env var
  
  // Command-line arguments
  args := os.Args           // ["program", "arg1", "arg2"]
  programName := os.Args[0]
  
  // Working directory
  pwd, _ := os.Getwd()      // Current directory
  os.Chdir("/tmp")          // Change directory
  
  // Exit program
  os.Exit(1)                // Exit with status code
  ```

- [ ] `strings` and `strconv`

  **Explanation:** String manipulation and type conversion utilities.
  ```go
  import (
      "strings"
      "strconv"
  )
  
  // strings package
  strings.Contains("hello", "ell")      // true
  strings.HasPrefix("hello", "he")      // true
  strings.HasSuffix("hello", "lo")      // true
  strings.Split("a,b,c", ",")           // ["a", "b", "c"]
  strings.Join([]string{"a", "b"}, "-") // "a-b"
  strings.ToUpper("hello")              // "HELLO"
  strings.ToLower("HELLO")              // "hello"
  strings.TrimSpace("  hi  ")           // "hi"
  strings.Replace("hello", "l", "L", -1) // "heLLo"
  
  // strconv package - convert between strings and numbers
  i, err := strconv.Atoi("42")          // string to int
  s := strconv.Itoa(42)                  // int to string
  f, err := strconv.ParseFloat("3.14", 64)  // string to float64
  b, err := strconv.ParseBool("true")   // string to bool
  s := strconv.FormatFloat(3.14, 'f', 2, 64)  // float to string
  ```

- [ ] `time` (durations, parsing, formatting)

  **Explanation:** Time manipulation, formatting, and duration handling.
  ```go
  import "time"
  
  // Current time
  now := time.Now()
  
  // Duration
  d := 5 * time.Second
  d := time.Minute + 30*time.Second
  time.Sleep(d)
  
  // Time arithmetic
  later := now.Add(24 * time.Hour)    // Add duration
  diff := later.Sub(now)              // Difference (duration)
  
  // Parsing time (reference: Mon Jan 2 15:04:05 MST 2006)
  t, err := time.Parse("2006-01-02", "2024-01-15")
  t, err := time.Parse(time.RFC3339, "2024-01-15T10:30:00Z")
  
  // Formatting time
  s := now.Format("2006-01-02 15:04:05")  // "2024-01-15 10:30:00"
  s := now.Format(time.RFC3339)           // "2024-01-15T10:30:00Z"
  
  // Time comparisons
  t1.Before(t2)
  t1.After(t2)
  t1.Equal(t2)
  
  // Timer and Ticker
  timer := time.NewTimer(5 * time.Second)
  <-timer.C  // Blocks for 5 seconds
  
  ticker := time.NewTicker(1 * time.Second)
  for range ticker.C {
      // Runs every second
  }
  ```

- [ ] `encoding/json` (Marshal, Unmarshal, struct tags)

  **Explanation:** JSON encoding and decoding with struct tags for field mapping.
  ```go
  import "encoding/json"
  
  type User struct {
      ID        int    `json:"id"`
      Name      string `json:"name"`
      Email     string `json:"email,omitempty"`  // Omit if empty
      Password  string `json:"-"`                 // Never include
      CreatedAt time.Time `json:"created_at"`
  }
  
  // Struct to JSON (Marshal)
  user := User{ID: 1, Name: "John"}
  jsonBytes, err := json.Marshal(user)
  // {"id":1,"name":"John","created_at":"..."}
  
  // Pretty print
  jsonBytes, err := json.MarshalIndent(user, "", "  ")
  
  // JSON to Struct (Unmarshal)
  var user User
  err := json.Unmarshal(jsonBytes, &user)
  
  // Streaming (for large data or HTTP)
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&user)
  
  encoder := json.NewEncoder(w)
  err := encoder.Encode(user)
  
  // Dynamic JSON (unknown structure)
  var data map[string]interface{}
  json.Unmarshal(jsonBytes, &data)
  ```

- [ ] `net/http` (basic server, client)

  **Explanation:** HTTP client and server implementation.
  ```go
  import "net/http"
  
  // Simple HTTP Server
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, World!")
  })
  http.ListenAndServe(":8080", nil)
  
  // HTTP Client - GET
  resp, err := http.Get("https://api.example.com/users")
  if err != nil {
      log.Fatal(err)
  }
  defer resp.Body.Close()
  body, _ := io.ReadAll(resp.Body)
  
  // HTTP Client - POST
  data := []byte(`{"name": "John"}`)
  resp, err := http.Post(
      "https://api.example.com/users",
      "application/json",
      bytes.NewBuffer(data),
  )
  
  // Custom client with timeout
  client := &http.Client{
      Timeout: 10 * time.Second,
  }
  req, _ := http.NewRequest("GET", url, nil)
  req.Header.Set("Authorization", "Bearer token")
  resp, err := client.Do(req)
  ```

- [ ] `log` and `log/slog` (structured logging)

  **Explanation:** Standard logging and structured logging (Go 1.21+).
  ```go
  import (
      "log"
      "log/slog"
  )
  
  // Basic log package
  log.Println("Info message")
  log.Printf("User %s logged in", username)
  log.Fatal("Fatal error")  // Logs and calls os.Exit(1)
  log.Panic("Panic!")       // Logs and panics
  
  // Custom logger
  logger := log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
  logger.Println("Custom logger")
  
  // Structured logging with slog (Go 1.21+)
  slog.Info("User logged in",
      "user_id", 123,
      "ip", "192.168.1.1",
  )
  // Output: 2024/01/15 10:30:00 INFO User logged in user_id=123 ip=192.168.1.1
  
  // JSON handler
  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  logger.Info("Request", "method", "GET", "path", "/api/users")
  // {"time":"...","level":"INFO","msg":"Request","method":"GET","path":"/api/users"}
  ```

#### Testing:
- [ ] Writing tests (`_test.go` files)

  **Explanation:** Tests live in files ending with `_test.go` in the same package.
  ```go
  // math.go
  package math
  
  func Add(a, b int) int {
      return a + b
  }
  
  // math_test.go
  package math
  
  import "testing"
  
  func TestAdd(t *testing.T) {
      result := Add(2, 3)
      if result != 5 {
          t.Errorf("Add(2, 3) = %d; want 5", result)
      }
  }
  
  // Run tests
  // go test           # Run tests in current package
  // go test ./...     # Run all tests recursively
  // go test -v        # Verbose output
  ```

- [ ] Table-driven tests

  **Explanation:** Run same test logic with multiple inputs/expected outputs.
  ```go
  func TestAdd(t *testing.T) {
      tests := []struct {
          name     string
          a, b     int
          expected int
      }{
          {"positive numbers", 2, 3, 5},
          {"negative numbers", -1, -2, -3},
          {"zero", 0, 0, 0},
          {"mixed", -1, 2, 1},
      }
      
      for _, tt := range tests {
          t.Run(tt.name, func(t *testing.T) {
              result := Add(tt.a, tt.b)
              if result != tt.expected {
                  t.Errorf("Add(%d, %d) = %d; want %d",
                      tt.a, tt.b, result, tt.expected)
              }
          })
      }
  }
  ```

- [ ] Subtests (`t.Run`)

  **Explanation:** Create nested tests for better organization and selective running.
  ```go
  func TestUser(t *testing.T) {
      t.Run("Create", func(t *testing.T) {
          // Test user creation
      })
      
      t.Run("Update", func(t *testing.T) {
          // Test user update
      })
      
      t.Run("Delete", func(t *testing.T) {
          // Test user deletion
      })
  }
  
  // Run specific subtest
  // go test -run TestUser/Create
  // go test -run TestUser/Update
  ```

- [ ] Test helpers

  **Explanation:** Helper functions reduce test code duplication.
  ```go
  func TestSomething(t *testing.T) {
      user := createTestUser(t)  // Helper
      // ... test code
  }
  
  func createTestUser(t *testing.T) *User {
      t.Helper()  // Marks as helper - errors report caller's line
      
      user, err := NewUser("test@example.com")
      if err != nil {
          t.Fatalf("failed to create user: %v", err)
      }
      return user
  }
  
  // Cleanup
  func TestWithDatabase(t *testing.T) {
      db := setupTestDB(t)
      t.Cleanup(func() {
          db.Close()  // Runs after test completes
      })
      // ... test code
  }
  ```

- [ ] Mocking (interfaces for testability)

  **Explanation:** Use interfaces to inject mock implementations for testing.
  ```go
  // Define interface for dependencies
  type UserRepository interface {
      GetUser(id string) (*User, error)
      SaveUser(user *User) error
  }
  
  // Production implementation
  type PostgresUserRepo struct { db *sql.DB }
  func (r *PostgresUserRepo) GetUser(id string) (*User, error) { ... }
  
  // Mock for testing
  type MockUserRepo struct {
      users map[string]*User
  }
  
  func (m *MockUserRepo) GetUser(id string) (*User, error) {
      if user, ok := m.users[id]; ok {
          return user, nil
      }
      return nil, ErrNotFound
  }
  
  // Service uses interface - can inject either implementation
  type UserService struct {
      repo UserRepository
  }
  
  // Test with mock
  func TestUserService(t *testing.T) {
      mock := &MockUserRepo{users: map[string]*User{
          "1": {ID: "1", Name: "Test"},
      }}
      service := &UserService{repo: mock}
      
      user, err := service.GetUser("1")
      // ...
  }
  ```

- [ ] Benchmarks (`func BenchmarkXxx`)

  **Explanation:** Measure performance of code with benchmarks.
  ```go
  func BenchmarkAdd(b *testing.B) {
      // b.N is set by the testing framework
      for i := 0; i < b.N; i++ {
          Add(2, 3)
      }
  }
  
  // Run benchmarks
  // go test -bench=.
  // go test -bench=BenchmarkAdd
  // go test -bench=. -benchmem  # Include memory stats
  
  // Output:
  // BenchmarkAdd-8   1000000000   0.5 ns/op   0 B/op   0 allocs/op
  
  // Reset timer for setup
  func BenchmarkProcess(b *testing.B) {
      data := loadTestData()  // Setup
      b.ResetTimer()          // Don't count setup time
      
      for i := 0; i < b.N; i++ {
          Process(data)
      }
  }
  ```

- [ ] `go test -cover` for coverage

  **Explanation:** Measure test coverage of your code.
  ```bash
  # Show coverage percentage
  go test -cover
  # coverage: 75.0% of statements
  
  # Generate coverage profile
  go test -coverprofile=coverage.out
  
  # View coverage in browser
  go tool cover -html=coverage.out
  
  # Coverage by function
  go tool cover -func=coverage.out
  
  # Coverage for all packages
  go test -cover ./...
  ```

#### Generics (Go 1.18+):
- [ ] Type parameters

  **Explanation:** Write functions and types that work with any type.
  ```go
  // Generic function
  func Min[T int | float64](a, b T) T {
      if a < b {
          return a
      }
      return b
  }
  
  // Usage - type inferred
  Min(1, 2)       // int
  Min(1.5, 2.5)   // float64
  
  // Generic type
  type Stack[T any] struct {
      items []T
  }
  
  func (s *Stack[T]) Push(item T) {
      s.items = append(s.items, item)
  }
  
  func (s *Stack[T]) Pop() T {
      item := s.items[len(s.items)-1]
      s.items = s.items[:len(s.items)-1]
      return item
  }
  
  // Usage
  intStack := &Stack[int]{}
  intStack.Push(1)
  ```

- [ ] Type constraints

  **Explanation:** Limit which types can be used with generics.
  ```go
  // Built-in constraints
  import "golang.org/x/exp/constraints"
  
  func Sum[T constraints.Integer | constraints.Float](nums []T) T {
      var sum T
      for _, n := range nums {
          sum += n
      }
      return sum
  }
  
  // Define custom constraint
  type Number interface {
      int | int32 | int64 | float32 | float64
  }
  
  func Double[T Number](n T) T {
      return n * 2
  }
  
  // comparable constraint - types that support == and !=
  func Contains[T comparable](slice []T, item T) bool {
      for _, v := range slice {
          if v == item {
              return true
          }
      }
      return false
  }
  
  // any constraint - accepts any type (alias for interface{})
  func Print[T any](v T) {
      fmt.Println(v)
  }
  ```

- [ ] Common use cases (slices, maps utilities)

  **Explanation:** Generics are great for collection utilities.
  ```go
  // Map function
  func Map[T, U any](slice []T, fn func(T) U) []U {
      result := make([]U, len(slice))
      for i, v := range slice {
          result[i] = fn(v)
      }
      return result
  }
  
  // Filter function
  func Filter[T any](slice []T, fn func(T) bool) []T {
      var result []T
      for _, v := range slice {
          if fn(v) {
              result = append(result, v)
          }
      }
      return result
  }
  
  // Usage
  nums := []int{1, 2, 3, 4, 5}
  doubled := Map(nums, func(n int) int { return n * 2 })
  evens := Filter(nums, func(n int) bool { return n%2 == 0 })
  
  // Standard library: golang.org/x/exp/slices and maps
  import "golang.org/x/exp/slices"
  
  slices.Contains(nums, 3)        // true
  slices.Sort(nums)               // Sort in place
  slices.Reverse(nums)            // Reverse in place
  ```

#### Project Structure:
- [ ] Standard layout (`cmd/`, `internal/`, `pkg/`)

  **Explanation:** Recommended directory structure for Go projects.
  ```
  myproject/
  ├── cmd/                    # Main applications
  │   ├── server/
  │   │   └── main.go         # go run ./cmd/server
  │   └── cli/
  │       └── main.go         # go run ./cmd/cli
  │
  ├── internal/               # Private code (cannot be imported)
  │   ├── handler/            # HTTP handlers
  │   ├── service/            # Business logic
  │   └── repository/         # Database access
  │
  ├── pkg/                    # Public library code (can be imported)
  │   └── utils/
  │
  ├── api/                    # API definitions (OpenAPI, proto)
  ├── configs/                # Config files
  ├── scripts/                # Build/deploy scripts
  ├── go.mod
  ├── go.sum
  └── Makefile
  ```

- [ ] When to use `internal/`

  **Explanation:** `internal/` makes code truly private to your module.
  ```go
  // internal/ is special: Go compiler enforces visibility
  
  myproject/
  ├── internal/
  │   └── secret/
  │       └── secret.go       # Package myproject/internal/secret
  └── cmd/
      └── app/
          └── main.go
  
  // From main.go (SAME module):
  import "myproject/internal/secret"  // ✓ Works
  
  // From ANOTHER module:
  import "myproject/internal/secret"  // ✗ Compile error!
  
  // Use internal/ for:
  // - Implementation details you might change
  // - Code not ready for public API
  // - Preventing external dependencies on unstable code
  ```

- [ ] Organizing by feature vs layer

  **Explanation:** Two common approaches to organizing code.
  ```
  // BY LAYER (traditional)
  internal/
  ├── handlers/        # All HTTP handlers
  │   ├── user.go
  │   └── order.go
  ├── services/        # All business logic
  │   ├── user.go
  │   └── order.go
  └── repositories/    # All database code
      ├── user.go
      └── order.go
  
  // BY FEATURE (domain-driven)
  internal/
  ├── user/            # Everything about users
  │   ├── handler.go
  │   ├── service.go
  │   └── repository.go
  └── order/           # Everything about orders
      ├── handler.go
      ├── service.go
      └── repository.go
  
  // Feature-based is often better for larger projects:
  // - Related code is together
  // - Easier to understand one feature
  // - Clear ownership boundaries
  ```

#### Gin Framework Basics:
> **Note on Gin:** You are using Gin because it is the most popular Go web framework, uses the standard `net/http` library (unlike Fiber), and has a massive ecosystem of middleware. It strikes the perfect balance between performance and productivity.

- [ ] Routing and handlers (`gin.Context`)

  **Explanation:** Gin provides fast routing and a powerful context object.
  ```go
  import "github.com/gin-gonic/gin"
  
  func main() {
      r := gin.Default()  // Logger + Recovery middleware
      
      // Route handlers
      r.GET("/users", getUsers)
      r.GET("/users/:id", getUser)      // Path parameter
      r.POST("/users", createUser)
      r.PUT("/users/:id", updateUser)
      r.DELETE("/users/:id", deleteUser)
      
      // Route groups
      api := r.Group("/api/v1")
      {
          api.GET("/users", getUsers)
          api.POST("/users", createUser)
      }
      
      r.Run(":8080")
  }
  
  func getUser(c *gin.Context) {
      id := c.Param("id")                    // Path param
      name := c.Query("name")                // Query string ?name=x
      page := c.DefaultQuery("page", "1")    // With default
      
      c.JSON(http.StatusOK, gin.H{
          "id": id,
          "name": name,
      })
  }
  ```

- [ ] Middleware (logging, recovery, CORS)

  **Explanation:** Middleware runs before/after handlers for cross-cutting concerns.
  ```go
  // Built-in middleware
  r := gin.New()  // No middleware
  r.Use(gin.Logger())      // Request logging
  r.Use(gin.Recovery())    // Panic recovery
  
  // Custom middleware
  func AuthMiddleware() gin.HandlerFunc {
      return func(c *gin.Context) {
          token := c.GetHeader("Authorization")
          if token == "" {
              c.JSON(401, gin.H{"error": "unauthorized"})
              c.Abort()  // Stop chain
              return
          }
          
          // Set value for later handlers
          c.Set("user_id", "123")
          c.Next()  // Continue to next handler
      }
  }
  
  // Apply middleware
  r.Use(AuthMiddleware())  // All routes
  
  // Group-specific middleware
  admin := r.Group("/admin", AuthMiddleware())
  
  // CORS
  import "github.com/gin-contrib/cors"
  r.Use(cors.Default())  // Allow all origins
  ```

- [ ] Request parsing (`ShouldBindJSON`, `ShouldBindQuery`)

  **Explanation:** Bind request data to Go structs automatically.
  ```go
  type CreateUserRequest struct {
      Name  string `json:"name" binding:"required"`
      Email string `json:"email" binding:"required,email"`
      Age   int    `json:"age" binding:"gte=0,lte=120"`
  }
  
  func createUser(c *gin.Context) {
      var req CreateUserRequest
      
      // Bind JSON body
      if err := c.ShouldBindJSON(&req); err != nil {
          c.JSON(400, gin.H{"error": err.Error()})
          return
      }
      
      // ShouldBind* variants (return error, don't abort):
      c.ShouldBindJSON(&req)    // JSON body
      c.ShouldBindQuery(&req)   // Query params
      c.ShouldBindUri(&req)     // Path params
      c.ShouldBind(&req)        // Auto-detect
      
      // Bind* variants abort on error:
      c.BindJSON(&req)          // Aborts with 400 on error
  }
  
  type Pagination struct {
      Page  int `form:"page" binding:"gte=1"`
      Limit int `form:"limit" binding:"gte=1,lte=100"`
  }
  
  func listUsers(c *gin.Context) {
      var pagination Pagination
      c.ShouldBindQuery(&pagination)
  }
  ```

- [ ] Response handling (`c.JSON`)

  **Explanation:** Send various response types.
  ```go
  func handler(c *gin.Context) {
      // JSON response
      c.JSON(200, gin.H{
          "message": "success",
          "data":    user,
      })
      
      // With struct
      c.JSON(200, user)
      
      // String response
      c.String(200, "Hello %s", name)
      
      // HTML response
      c.HTML(200, "template.html", data)
      
      // File response
      c.File("path/to/file.pdf")
      
      // Redirect
      c.Redirect(302, "/new-url")
      
      // Error responses
      c.JSON(400, gin.H{"error": "bad request"})
      c.JSON(404, gin.H{"error": "not found"})
      c.JSON(500, gin.H{"error": "internal error"})
      
      // Abort (stop handler chain)
      c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
  }
  ```

- [ ] Validation (built-in `go-playground/validator`)

  **Explanation:** Gin uses go-playground/validator for request validation.
  ```go
  type User struct {
      Name     string `binding:"required,min=2,max=50"`
      Email    string `binding:"required,email"`
      Age      int    `binding:"gte=0,lte=150"`
      Password string `binding:"required,min=8"`
      Phone    string `binding:"omitempty,e164"`  // Optional but validated if present
      Website  string `binding:"omitempty,url"`
  }
  
  // Common validation tags:
  // required      - Field must be present
  // email         - Valid email format
  // url           - Valid URL
  // min=n, max=n  - String length or number value
  // gte=n, lte=n  - Greater/less than or equal
  // len=n         - Exact length
  // oneof=a b c   - One of listed values
  // omitempty     - Skip validation if empty
  
  // Custom validation
  import "github.com/go-playground/validator/v10"
  
  func main() {
      v := validator.New()
      v.RegisterValidation("customtag", func(fl validator.FieldLevel) bool {
          return fl.Field().String() != "invalid"
      })
  }
  ```

- [ ] Configuration management (Viper)

  **Explanation:** Viper is the standard for configuration management.
  ```go
  import "github.com/spf13/viper"
  
  type Config struct {
      Server struct {
          Port int    `mapstructure:"port"`
          Host string `mapstructure:"host"`
      } `mapstructure:"server"`
      Database struct {
          URL string `mapstructure:"url"`
      } `mapstructure:"database"`
  }
  
  func LoadConfig() (*Config, error) {
      viper.SetConfigName("config")       // config.yaml
      viper.SetConfigType("yaml")
      viper.AddConfigPath(".")
      viper.AddConfigPath("./configs")
      
      // Environment variable overrides
      viper.AutomaticEnv()
      viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
      // SERVER_PORT env var → server.port config
      
      if err := viper.ReadInConfig(); err != nil {
          return nil, err
      }
      
      var config Config
      if err := viper.Unmarshal(&config); err != nil {
          return nil, err
      }
      
      return &config, nil
  }
  
  // config.yaml
  // server:
  //   port: 8080
  //   host: localhost
  // database:
  //   url: postgres://localhost/db
  ```

#### Hands-on Project:
```
Build: User Management API (Gin)
- POST /users (create with validation)
- GET /users/:id
- PUT /users/:id
- DELETE /users/:id
- Proper error handling
- Request logging middleware
- Input validation
- Structured logging (zerolog/zap)
```

#### Project Structure:
```
user-service/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── config/              # Configuration
│   ├── handler/             # HTTP handlers
│   ├── middleware/          # Gin middleware
│   ├── model/               # Domain models
│   ├── repository/          # Database access
│   └── service/             # Business logic
├── pkg/                     # Reusable packages
├── go.mod
├── go.sum
└── Makefile
```

#### Resources:
- **Gin Docs:** https://gin-gonic.com/docs/
- **Project Structure:** https://github.com/golang-standards/project-layout

---

## MONTH 2: Go Backend Deep Dive
