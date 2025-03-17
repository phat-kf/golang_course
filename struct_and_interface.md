## **Structs and Interfaces in Golang**  

In **Go**, structs and interfaces are essential building blocks for defining and organizing data, as well as achieving polymorphism.

---

# **1. Structs in Golang**  
A **struct** is a collection of fields that group related data together.  

### **1.1 Defining and Using a Struct**
```go
package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

func main() {
    // Initialize a struct
    p := Person{Name: "Alice", Age: 25}

    // Access fields
    fmt.Println(p.Name) // Output: Alice
    fmt.Println(p.Age)  // Output: 25

    // Modify a field
    p.Age = 30
    fmt.Println(p.Age)  // Output: 30
}
```

---

### **1.2 Struct Initialization Methods**
You can initialize a struct in different ways:

```go
p1 := Person{"Bob", 20}           // Positional initialization
p2 := Person{Name: "Charlie"}     // Partial initialization (Age = 0)
p3 := Person{}                    // Default values (empty string, 0)

fmt.Println(p1, p2, p3)
```

---

### **1.3 Struct with Methods**
Methods allow attaching functions to structs:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Method with a receiver
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

func main() {
    p := Person{Name: "David", Age: 40}
    p.Greet() // Output: Hello, my name is David
}
```

---

### **1.4 Struct with Pointer Receiver (Modify Fields)**
To modify the struct inside a method, use a pointer receiver (`*Person`):

```go
func (p *Person) Birthday() {
    p.Age++ // Modifies the original struct
}

func main() {
    p := Person{Name: "Emma", Age: 29}
    p.Birthday()
    fmt.Println(p.Age) // Output: 30
}
```

---

### **1.5 Embedding Structs (Inheritance-like Behavior)**
Go **does not** support inheritance but allows struct embedding:

```go
type Employee struct {
    Person  // Embedded struct
    Company string
}

func main() {
    e := Employee{
        Person:  Person{Name: "Frank", Age: 35},
        Company: "TechCorp",
    }
    
    fmt.Println(e.Name)  // Access embedded struct fields
    fmt.Println(e.Company)
}
```

---

# **2. Interfaces in Golang**  
An **interface** is a set of method signatures that a type must implement.

### **2.1 Defining and Implementing an Interface**
```go
package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak()
}

// Implement the interface in a struct
type Dog struct {
    Name string
}

func (d Dog) Speak() {
    fmt.Println(d.Name, "says Woof!")
}

func main() {
    d := Dog{Name: "Buddy"}
    d.Speak() // Output: Buddy says Woof!

    var s Speaker = d // Interface assignment
    s.Speak()
}
```

---

### **2.2 Interfaces with Multiple Implementations**
```go
type Cat struct {
    Name string
}

func (c Cat) Speak() {
    fmt.Println(c.Name, "says Meow!")
}

func main() {
    animals := []Speaker{Dog{"Rex"}, Cat{"Mittens"}}

    for _, a := range animals {
        a.Speak()
    }
}
```

---

### **2.3 Empty Interface (`interface{}`)**
An empty interface can hold **any** value:

```go
func PrintValue(v interface{}) {
    fmt.Println(v)
}

func main() {
    PrintValue(42)
    PrintValue("Hello")
    PrintValue([]int{1, 2, 3})
}
```

---

### **2.4 Type Assertion**
To extract the actual type from an `interface{}`:

```go
func Describe(i interface{}) {
    if v, ok := i.(string); ok {
        fmt.Println("String:", v)
    } else {
        fmt.Println("Not a string:", i)
    }
}

func main() {
    Describe("Go")
    Describe(100)
}
```

---

# **3. Differences Between Structs and Interfaces**
| Feature        | Structs                                      | Interfaces                                  |
|---------------|---------------------------------------------|---------------------------------------------|
| Data Storage  | Holds concrete data                         | Does not store data, only method signatures |
| Methods       | Can have methods                            | Defines method behavior                    |
| Usage         | Used to group related fields               | Used for polymorphism and abstraction      |
| Embedding     | Supports struct embedding                  | Cannot be embedded                         |

---

### **When to Use Structs vs. Interfaces**
✅ Use **structs** when you need to group related fields.  
✅ Use **interfaces** when you want to define behavior that multiple types can implement.  