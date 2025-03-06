In Golang, **functions** and **methods** are essential for structuring code, reusability, and encapsulation. Below is a detailed explanation with examples.

---

## **1. Functions in Golang**
A **function** in Go is a reusable block of code that performs a specific task. Functions can take parameters, return values, and support multiple return values.

### **Function Syntax**
```go
func functionName(parameterName dataType) returnType {
    // function body
    return value
}
```

### **Example: Basic Function**
```go
package main

import "fmt"

// Function that adds two numbers
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 5)
    fmt.Println("Sum:", result) // Output: Sum: 8
}
```

---

### **2. Function with Multiple Return Values**
Golang allows functions to return **multiple values**.

```go
func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Error: Division by zero"
    }
    return a / b, ""
}

func main() {
    result, err := divide(10, 2)
    if err != "" {
        fmt.Println(err)
    } else {
        fmt.Println("Result:", result) // Output: Result: 5
    }
}
```

---

### **3. Named Return Values**
Go allows naming return values, which makes the function more readable.

```go
func rectangleArea(length, width float64) (area float64) {
    area = length * width
    return // No need to explicitly return 'area'
}
```

---

### **4. Variadic Functions (Multiple Arguments)**
A variadic function accepts a variable number of arguments.

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4)) // Output: 10
}
```

---

### **5. Anonymous Functions**
Go supports anonymous (inline) functions.

```go
func main() {
    multiply := func(a, b int) int {
        return a * b
    }

    fmt.Println(multiply(4, 5)) // Output: 20
}
```

---

### **6. Higher-Order Functions (Passing Functions as Arguments)**
Go allows passing functions as arguments.

```go
func operate(a int, b int, op func(int, int) int) int {
    return op(a, b)
}

func add(a, b int) int {
    return a + b
}

func main() {
    result := operate(10, 5, add)
    fmt.Println("Result:", result) // Output: Result: 15
}
```

---

## **7. Methods in Golang**
A **method** is a function associated with a specific **struct** type. Methods allow encapsulating behavior within structs.

### **Method Syntax**
```go
func (receiver Type) methodName(parameters) returnType {
    // method body
}
```

### **Example: Method on a Struct**
```go
package main

import "fmt"

// Define a struct
type Rectangle struct {
    width, height float64
}

// Method to calculate area
func (r Rectangle) area() float64 {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 5, height: 3}
    fmt.Println("Area:", rect.area()) // Output: Area: 15
}
```

---

### **8. Pointer Receivers in Methods**
If you want the method to modify the original struct, use a **pointer receiver (`*Type`)**.

```go
type Counter struct {
    value int
}

// Method using pointer receiver
func (c *Counter) increment() {
    c.value++
}

func main() {
    c := Counter{value: 0}
    c.increment()
    fmt.Println("Counter:", c.value) // Output: Counter: 1
}
```

---

### **9. Method vs Function**
| Feature  | Function | Method |
|----------|---------|--------|
| Receiver | No | Yes (struct or pointer) |
| Scope | Independent | Tied to a struct |
| Use Case | General-purpose logic | Behavior of an object (struct) |

---

## **Summary**
- **Functions** are independent reusable blocks of code.
- **Methods** are functions associated with a struct (can use value or pointer receivers).
- **Variadic functions** accept multiple arguments (`...int`).
- **Higher-order functions** can pass functions as arguments.
- **Anonymous functions** allow inline logic.